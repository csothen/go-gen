package generator

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

const (
	DEFAULT_APPLICATION_LOCATION = "cmd/"
	DEFAULT_CONTROLLER_LOCATION  = "internal/"
	DEFAULT_SERVICE_LOCATION     = "internal/"
	DEFAULT_REPOSITORY_LOCATION  = "internal/"
	DEFAULT_MODEL_LOCATION       = "pkg/"
)

var GeneratorConfig Configuration = GetConfig()

type Configuration struct {
	APPLICATION_LOCATION string
	CONTROLLER_LOCATION  string
	SERVICE_LOCATION     string
	REPOSITORY_LOCATION  string
	MODEL_LOCATION       string
}

func GetConfig() Configuration {
	configuration := Configuration{}

	fileName := fmt.Sprintf("./generation_config.json")

	if err := gonfig.GetConf(fileName, &configuration); err != nil {
		useDefault(&configuration)
	}

	return configuration
}

func useDefault(conf *Configuration) {
	conf.APPLICATION_LOCATION = DEFAULT_APPLICATION_LOCATION
	conf.CONTROLLER_LOCATION = DEFAULT_CONTROLLER_LOCATION
	conf.SERVICE_LOCATION = DEFAULT_SERVICE_LOCATION
	conf.REPOSITORY_LOCATION = DEFAULT_REPOSITORY_LOCATION
	conf.MODEL_LOCATION = DEFAULT_MODEL_LOCATION
}
