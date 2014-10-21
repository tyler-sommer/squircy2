package squircy

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Network       string
	Nick          string
	Username      string
	Password      string
	Channel       string
	OwnerNick     string
	OwnerHost     string
	RedisHost     string
	RedisDatabase int
}

func NewConfiguration(fname string) (config *Configuration) {
	config = &Configuration{}
	
	file, err := os.Open(fname)
	if err != nil {
		panic("Could not open configuration: " + err.Error())
	}

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(config); err != nil {
		panic("Could not decode configuration: " + err.Error())
	}
	
	return
}