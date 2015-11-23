package easyConfig

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"io/ioutil"	
)

var cwd,_ = os.Getwd()

func New(configObject interface{}, configFilePath string) interface{} {

	fmt.Println("reading config file: ", configFilePath)
	
	file, err := ioutil.ReadFile(path.Join(cwd, configFilePath))
	
	if err != nil {
		fmt.Println("\n\nerror opening config file: %+v", err)
		os.Exit(1)
	}
	
	err = json.Unmarshal(file, &configObject)
	if err != nil {
		fmt.Println("\n\n%+v", err)
	}
	return configObject
}
