/**
* TODO:  Fix the way I am setting  cookies in the header. Currently If I add one I would overwrite it with the next.
* I am sure there are a few I havent discovered yet, but one I know about is adding cookies to the header.
*
* TODO: it will probably be easier to supply the csv file as an argument rather than in the config
 */
package main

import (
	"encoding/json"
	"fmt"
	//	"github.com/davecgh/go-spew/spew"
	"github.com/recursionpharma/go-csv-map"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
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

func LoadConfig() (Config, error) {
	var config Config
	cfile, err := os.Open(os.Args[1])
	defer cfile.Close()
	if err != nil {

	}
	jsonParser := json.NewDecoder(cfile)
	err = jsonParser.Decode(&config)
	return config, err
}
func main() {
	if len(os.Args) < 2 {
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
{
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
********************************************************************************

Then run ajaxFromCsv /path/to/file.csv


`
		fmt.Println(fmtString)
		return
	}

	processCsv()
}

func processCsv() {
	confi, err := LoadConfig()
	if err != nil {
		fmt.Println("error loading config", err)
	}
	f, err := os.Open(confi.Csvfile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	reader := csvmap.NewReader(f)
	reader.Columns, err = reader.ReadHeader()
	if err != nil {
		fmt.Println(" error with ReadHeader", err)
		os.Exit(1)
	}
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, v := range records { //v is the map we are going to parse into the values
		data := url.Values{}
		for i, j := range v {
			buildData(&data, i, j)
		}
		performCall(&data)
	}
}
func buildData(d *url.Values, k string, v string) {
	d.Set(k, v)
}

func makeRequest(data *url.Values) *http.Request {
	conf, err := LoadConfig()
	if err != nil {
		fmt.Println("error loading config", err)
	}
	req, _ := http.NewRequest("POST", conf.Host+conf.Endpoint, strings.NewReader(data.Encode()))
	for _, header := range conf.Headers {
		req.Header.Set(header.Type, header.Value)
	}
	return req
}
func performCall(data *url.Values) {
	req := makeRequest(data)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error")
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status, "\n")
	fmt.Println("response Headers:", resp.Header, "\n")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body), "\n")
}
