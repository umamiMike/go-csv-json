package main

import (
	"encoding/json"
	"os"
)

type Header struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type Config struct {
	Host     string   `json:"host"`
	Endpoint string   `json:"endpoint"`
	Csvfile  string   `json:"csvfile"`
	Headers  []Header `json:"headers"`
}

var configfile = "./conf.json" //REMOVE ME

func LoadConfig(c string) (Config, error) {
	var config Config
	cfile, err := os.Open(c)
	defer cfile.Close()
	if err != nil {
	}
	jsonParser := json.NewDecoder(cfile)
	err = jsonParser.Decode(&config)
	return config, err
}

func printReadme() string {
	fmtString := `
This is a tool to run batch commands based on csv document
The first row must contain the names of the encoded post commands
lets say you have a post to make to an endpoint to change the middle name of
a series of clients.  The enpoint might look like  http:/hostname/update-client
and the form data might be client_id=###&firstname=string&middlename=string&lastname=string

your csv file should look something like:

client_id,firstname,middlename,lastname
3456,michael,wayne,wilding

the tool will run through every line of the csv document
and make a post with the data on each row
so the example from the csv above would look like
/update-client?client_id=3465&firstname=michael&middlename=wayne&lastname=wilding

The other thing to note is the config file, where you fill in the requisite info
including a list of any headers




*************** Copy  into a json file (EX: config.json) ***********************
` + configfiletoecho + ` ********************************************************************************
Then run ajaxFromCsv /path/to/file.csv
`
	return fmtString
}

var configfiletoecho = `{
        "host" : "http://example.com",
        "endpoint" : "/endpoint",
        "csvfile": "/path/to/file.csv",
        "headers" : [
                {  
                        "type" : "Cookie",
                        "value" : "PHPSESSID="
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