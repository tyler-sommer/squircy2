package script

import (
	"errors"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/robertkrimen/otto"
	"github.com/veonik/squircy2/event"
)

const maxExecutionTime = 5 // in seconds
var Halt = errors.New("Execution limit exceeded")
var UnknownScriptType = errors.New("Unknown script type")

type scriptDriver interface {
	Handle(e event.Event, fnName string)
	RunUnsafe(code string) (interface{}, error)
	String() string
}

type javascriptDriver struct {
	VM *jsVm
	log.FieldLogger
}

func (d javascriptDriver) Handle(e event.Event, fnName string) {
	d.VM.Interrupt = make(chan func(), 1)
	data, err := d.VM.ToValue(e.Data)
	if err != nil {
		d.Debugln("An error occurred while creating event data", err)
		return
	}
	_, err = d.VM.Call(fnName, otto.NullValue(), data)
	if err != nil {
		d.Debugln("An error occurred while executing the Javascript handler", err)
	}
}

func (d javascriptDriver) RunUnsafe(unsafe string) (val interface{}, err error) {
	start := time.Now()
	quit := make(chan struct{})
	defer func() {
		if e := recover(); e != nil {
			if e == Halt {
				duration := time.Since(start)
				d.Infoln("Some Javascript code took too long! Stopping after: ", duration)
			}
			err = e.(error)
		}
		close(quit)
	}()

	d.VM.Interrupt = make(chan func(), 1)

	go func() {
		time.Sleep(maxExecutionTime * time.Second)
		select {
		case <-quit:
			return
		default:
			d.VM.Interrupt <- func() {
				panic(Halt)
			}
		}
	}()

	v, err := d.VM.Run(unsafe)
	val, _ = v.Export()

	return
}

func (d javascriptDriver) String() string {
	return "js"
}
