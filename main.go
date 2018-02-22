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
	"github.com/spf13/cobra"
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
	cmd := &cobra.Command{
		Use:   "run",
		Short: "process",
		RunE: func(cmd *cobra.Command, args []string) error {
			processCsv()
			return nil
		},
		SilenceUsage: true,
	}
	cmd.AddCommand(printTimeCmd())
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		fmt.Println(printReadme())
	}
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

func printTimeCmd() *cobra.Command {
	return &cobra.Command{
		Use: "curtime",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
