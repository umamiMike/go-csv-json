package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Header struct {
	Type  string `json:"type"`
	Name  string `json: "name"`
	Value string `json:"value"`
}
type Config struct {
	Host     string   `json:"host"`
	Endpoint string   `json:"endpoint"`
	Csvfile  string   `json:"csvfile"`
	Headers  []Header `json:"headers"`
}

func LoadConfig(c string) (Config, error) {
	var config Config
	cfile, err := os.Open(c)
	fc, _ := ioutil.ReadFile(c)
	if isJSON(string(fc)) == false {
		fmt.Println("the config file is not in JSON format")
	}
	defer cfile.Close()
	if err != nil {
	}
	jsonParser := json.NewDecoder(cfile)
	err = jsonParser.Decode(&config)
	return config, err
}

func Readme() string {
	fmtString := `
*********************************************************************************************
json lib! encode any csv file as json  ******************************************************
*********************************************************************************************

*********************************************************************************************
*********************************************************************************************
`
	return fmtString
}
