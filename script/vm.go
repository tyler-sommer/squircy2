package script

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/robertkrimen/otto"
	"github.com/veonik/squircy2/data"
	"github.com/veonik/squircy2/event"
)

// timer represents a function call to be performed after a delay.
type timer struct {
	t      *time.Timer
	dur    time.Duration
	repeat bool
	call   otto.FunctionCall
}

// jsVM is a small wrapper around Otto VM, facilitating the event loop
// and setTimeout/setInterval functionality.
type jsVm struct {
	*otto.Otto

	registry map[*timer]*timer
	ready    chan *timer
	quit     chan struct{}
}

func newJavascriptVm(m *ScriptManager) *jsVm {
	vm := &jsVm{otto.New(), make(map[*timer]*timer), make(chan *timer), make(chan struct{})}
	must := func(err error) {
		if err != nil {
			m.logger.Warnln("Error binding vm value", err)
		}
	}
	getFnName := func(fn otto.Value) (name string) {
		if fn.IsFunction() {
			name = fmt.Sprintf("__Handler%x", sha1.Sum([]byte(fn.String())))
		} else {
			name = fn.String()
		}

		return
	}
	newTimer := func(call otto.FunctionCall, repeat bool) (*timer, otto.Value, error) {
		delay, _ := call.Argument(1).ToInteger()
		if delay <= 0 {
			delay = 1
		}

		res := &timer{nil, time.Duration(delay) * time.Millisecond, repeat, call}
		vm.registry[res] = res

		res.t = time.AfterFunc(res.dur, func() {
			vm.ready <- res
		})

		val, err := vm.ToValue(res)
		if err != nil {
			return nil, otto.UndefinedValue(), err
		}

		return res, val, nil
	}
	clearTimer := func(call otto.FunctionCall) otto.Value {
		ti, _ := call.Argument(0).Export()
		if ti, ok := ti.(*timer); ok {
			ti.t.Stop()
			delete(vm.registry, ti)
		}
		return otto.UndefinedValue()
	}

	must(vm.Set("setTimeout", func(call otto.FunctionCall) otto.Value {
		_, v, err := newTimer(call, false)
		if err != nil {
			return otto.UndefinedValue()
		}
		return v
	}))
	must(vm.Set("setInterval", func(call otto.FunctionCall) otto.Value {
		_, v, err := newTimer(call, true)
		if err != nil {
			return otto.UndefinedValue()
		}
		return v
	}))
	must(vm.Set("clearTimeout", clearTimer))
	must(vm.Set("clearInterval", clearTimer))
	must(vm.Set("async", func(call otto.FunctionCall) otto.Value {
		if len(call.ArgumentList) != 2 {
			panic(vm.MakeCustomError("ArgumentError", "expected 2 arguments"))
		}
		async := call.Argument(0)
		if !async.IsFunction() {
			panic(vm.MakeCustomError("ArgumentError", "expected argument 1 to be a function"))
		}
		await := call.Argument(1)
		if !await.IsFunction() {
			panic(vm.MakeCustomError("ArgumentError", "expected argument 2 to be a function"))
		}
		go func(cvm *otto.Otto) {
			c := fmt.Sprintf("(%s)()", async.String())
			// Execute the async function on a copy of the VM
			result, err := cvm.Eval(c)
			if err != nil {
				if _, err := await.Call(await, nil, vm.MakeCustomError("AsyncError", err.Error())); err != nil {
					m.logger.Warnln("Error running async callback", err)
				}
				return
			}
			// Call the await function with the result on the original VM
			if _, err := await.Call(await, result); err != nil {
				m.logger.Warnln("Error running async callback", err)
			}
		}(vm.Copy())
		return otto.TrueValue()
	}))
	must(vm.Set("Http", &m.httpHelper))
	if v, err := vm.Object("Http"); err == nil {
		must(v.Set("Send", func(call otto.FunctionCall) otto.Value {
			o := call.Argument(0).Object()
			if o == nil {
				fmt.Printf("Expected argument 0 to be object, got: %v\n", o)
				return otto.UndefinedValue()
			}
			urlVal, err := o.Get("url")
			if err != nil {
				fmt.Printf("'url' is a required option\n")
				return otto.UndefinedValue()
			}
			url, err := urlVal.ToString()
			if err != nil {
				fmt.Printf("Could not coerce 'url' into string\n")
				return otto.UndefinedValue()
			}
			typVal, _ := o.Get("type")
			typ, _ := typVal.ToString()
			successCb, _ := o.Get("success")
			datVal, _ := o.Get("data")
			dat, _ := datVal.ToString()
			headerVal, _ := o.Get("headers")
			headers := make([]string, 0)
			if headerVal.IsString() {
				h, _ := headerVal.ToString()
				headers = append(headers, h)
			} else if headerVal.IsObject() {
				h := headerVal.Object()
				for _, k := range h.Keys() {
					v, _ := h.Get(k)
					headers = append(headers, fmt.Sprintf("%s: %s", k, v))
				}
			}
			go func() {
				var res string
				switch typ {
				case "post":
					res = m.httpHelper.Post(url, dat, headers...)
				default:
					res = m.httpHelper.Get(url, headers...)
				}
				if successCb.IsFunction() {
					successCb.Call(successCb, res)
				}
			}()

			return otto.UndefinedValue()
		}))
	} else {
		must(err)
	}
	must(vm.Set("Config", &m.configHelper))
	must(vm.Set("Irc", &m.ircHelper))
	must(vm.Set("Os", &m.osHelper))
	if v, err := vm.Object("({})"); err == nil {
		must(v.Set("ReadAll", func(call otto.FunctionCall) otto.Value {
			res, err := m.fileHelper.ReadAll(call.Argument(0).String())
			if err != nil {
				e, _ := otto.ToValue(err.Error())
				panic(e)
			}
			ret, _ := vm.ToValue(res)
			return ret
		}))
		must(vm.Set("File", v))
	} else {
		must(err)
	}

	must(vm.Set("Math", &m.mathHelper))
	if v, err := vm.Object("Math"); err == nil {
		must(v.Set("random", (&m.mathHelper).Rand))
		must(v.Set("round", (&m.mathHelper).Round))
		must(v.Set("ceil", (&m.mathHelper).Ceil))
		must(v.Set("floor", (&m.mathHelper).Floor))
	}

	must(vm.Set("bind", func(call otto.FunctionCall) otto.Value {
		eventType := call.Argument(0).String()
		fn := call.Argument(1)
		fnName := getFnName(fn)
		if fn.IsFunction() {
			must(m.driver.VM.Set(fnName, func(call otto.FunctionCall) otto.Value {
				if _, err := fn.Call(call.This, call.ArgumentList); err != nil {
					m.logger.Warnln("Error invoking callback", err)
				}
				return otto.UndefinedValue()
			}))
		}
		m.scriptHelper.Bind(Javascript, event.EventType(eventType), fnName)
		val, _ := otto.ToValue(fnName)
		return val
	}))

	must(vm.Set("unbind", func(call otto.FunctionCall) otto.Value {
		eventType := call.Argument(0).String()
		fnName := getFnName(call.Argument(1))
		m.scriptHelper.Unbind(Javascript, event.EventType(eventType), fnName)
		return otto.UndefinedValue()
	}))
	must(vm.Set("trigger", func(call otto.FunctionCall) otto.Value {
		eventType := call.Argument(0).String()
		dat, _ := call.Argument(1).Export()
		if dat == nil {
			dat = make(map[string]interface{}, 0)
		}
		m.scriptHelper.Trigger(event.EventType(eventType), dat.(map[string]interface{}))
		return otto.UndefinedValue()
	}))
	must(vm.Set("use", func(call otto.FunctionCall) otto.Value {
		coll := call.Argument(0).String()

		db := data.NewGenericRepository(m.database, coll)
		obj, _ := vm.Object("({})")
		must(obj.Set("Save", func(call otto.FunctionCall) otto.Value {
			exp, _ := call.Argument(0).Export()
			var model data.GenericModel
			switch t := exp.(type) {
			case data.GenericModel:
				model = t

			case map[string]interface{}:
				model = data.GenericModel(t)
			}
			switch t := model["ID"].(type) {
			case int64:
				model["ID"] = int(t)

			case int:
				model["ID"] = t

			case float64:
				model["ID"] = int(t)
			}
			db.Save(model)

			id, _ := vm.ToValue(model["ID"])

			return id
		}))
		must(obj.Set("Delete", func(call otto.FunctionCall) otto.Value {
			i, _ := call.Argument(0).ToInteger()
			db.Delete(int(i))

			res, _ := vm.ToValue(true)
			return res
		}))
		must(obj.Set("Fetch", func(call otto.FunctionCall) otto.Value {
			i, _ := call.Argument(0).ToInteger()
			val := db.Fetch(int(i))
			v, err := vm.ToValue(val)

			if err != nil {
				panic(err)
			}

			return v
		}))
		must(obj.Set("FetchAll", func(call otto.FunctionCall) otto.Value {
			vals := db.FetchAll()
			v, err := vm.ToValue(vals)

			if err != nil {
				m.logger.Debugln("An error occurred: ", err)
			}

			return v
		}))
		must(obj.Set("Index", func(call otto.FunctionCall) otto.Value {
			exp, _ := call.Argument(0).Export()
			cols := make([]string, 0)
			for _, val := range exp.([]interface{}) {
				cols = append(cols, val.(string))
			}
			db.Index(cols)

			return otto.UndefinedValue()
		}))
		must(obj.Set("Query", func(call otto.FunctionCall) otto.Value {
			qry, _ := call.Argument(0).Export()
			vals := db.Query(qry)
			v, err := vm.ToValue(vals)

			if err != nil {
				m.logger.Debugln("An error occurred: ", err)
			}

			return v
		}))

		return obj.Value()
	}))

	go vm.Loop()

	return vm
}

// Loop kicks off the VM's event loop.
func (vm *jsVm) Loop() {
	for {
		select {
		case ti := <-vm.ready:
			var args []interface{}
			if len(ti.call.ArgumentList) > 2 {
				tmp := ti.call.ArgumentList[2:]
				// args[1] will end up being "this" when the function is invoked,
				// so we need to offset each actual argument by one
				args = make([]interface{}, 2+len(tmp))
				for i, value := range tmp {
					args[i+2] = value
				}
			} else {
				args = make([]interface{}, 1)
			}
			args[0] = ti.call.ArgumentList[0]
			// Since we are calling "Function.call.call", we pass the function
			// to be called in as the first argument, which means args[1] will
			// end up being the function's "this" binding, and any further values
			// in args will be passed to the function as arguments.
			_, err := vm.Call("Function.call.call", nil, args...)
			if err == nil && ti.repeat {
				ti.t.Reset(ti.dur)
			} else {
				delete(vm.registry, ti)
			}
		case <-vm.quit:
			return
		}
	}
}
