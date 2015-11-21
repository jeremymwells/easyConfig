package easyConfig

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"	
)

var cwd,_ = os.Getwd()

func New(configObject interface{}, configFilePath string) interface{} {

	fmt.Print("reading config file: %s", configFilePath)
	
	file, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", cwd, configFilePath))
	
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
