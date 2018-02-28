package main

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

func printReadme() string {
	fmtString := `
*********************************************************************************************
Posty!   Hit that Endpoint!   ***************************************************************
*********************************************************************************************
This is a tool to run batch commands based on csv document

The first row must contain the names of the encoded post commands
your csv file should look something like:

first row ->             id,firstname,middlename,lastname
each successive row ->  3456,michael,wayne,wilding

use the 'posty create-config ' command to generate the boilerplate json file
Fill in the appropriate data, including the hostname, origin, endpoint, and any cookies.

The enpoint might look like     http://hostname/update-client
The form data might be          client_id=###&firstname=string&middlename=string&lastname=string

the tool will run through every line of the csv document
and make a post with the data on each row
so the example from the csv above would look like
http://hostname/update-client/update-client?id=3465&firstname=michael&middlename=wayne&lastname=wilding

*********************************************************************************************
*********************************************************************************************
`
	return fmtString
}

var configfiletoecho = `{
        "host" : "http://example.com",
        "endpoint" : "/endpoint",
        "headers" : [
                {  
												"type" : "Cookie",
                        "name" : "PHPSESSID",
                        "value" : ""
                },
                {
                        "type" : "Content-Type",
                        "value" : "application/x-www-form-urlencoded; charset=UTF-8"
                },
                {  
                        "type" : "Origin",
                        "value" : "http://example.com"
                }
        ]
}
`
