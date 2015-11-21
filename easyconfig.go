package config

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"	
)

type DatabaseUser struct {
	Name string
	Password string
}

type DatabaseDefinition struct {
	Name string
	Protocol string
	Address string
}

type DatabaseConfig struct {
	ReadUser DatabaseUser
	WriteUser DatabaseUser
	Dev DatabaseDefinition
	Prod DatabaseDefinition
}

type Configuration struct {
	Database DatabaseConfig
}

var (
	cwd,_ = os.Getwd()
)

func Get(configObject interface{}, configFile string) Configuration {
	if (configObject != interface{}) {
		fmt.Print("\n\nCONFIG ALREADY SET!!! %+v",configObject)
		return configObject
	}
	
	fmt.Print("reading config file")
	
	file, err := ioutil.ReadFile(cwd + configFilePath)
	
	if err != nil {
		fmt.Print("\n\nerror opening config file: %v", err)
		os.Exit(1)
	}
	
	err = json.Unmarshal(file, &configObject)
	if err != nil {
		fmt.Print("\n\n%+v", err)
	}
	return configObject
}
