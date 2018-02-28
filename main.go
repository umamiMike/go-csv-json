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
	"github.com/davecgh/go-spew/spew"
	"github.com/recursionpharma/go-csv-map"
	"github.com/spf13/cobra"
	"gopkg.in/cheggaaa/pb.v2"

	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var runningConfig Config

func main() {
	rootCmd := &cobra.Command{
		Use:   "posty",
		Short: "A tool for posting to an endpoint from rows in a CSV",
		// Long:  printReadme(),
	}
	rootCmd.AddCommand(Run)
	// Run.PersistentFlags().StringP("rate", "r", "100", "Number is milliseconds")

	rootCmd.AddCommand(createConfig())
	// rootCmd.PersistentFlags().StringP("config", "c", "configfile", "config file (default is $HOME/.cobra.yaml)")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func processCsv(confi Config, csvfile string, rate int) {
	spew.Dump("config %v", confi)
	fmt.Println("Rate in Milliseconds, if recieving too many 502 errors restart with bigger number", rate)
	c := make(chan int)
	f, err := os.Open(csvfile)
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

	recordCount := len(records)
	bar := pb.New(recordCount)
	bar.SetWidth(100)
	bar.Start()
	go incrementBar(bar, c)
	for _, v := range records { //v is the map we are going to parse into the values, can be accessed like ->  v["header"]
		data := url.Values{}
		for i, j := range v {
			buildUrlData(&data, i, j)
		}
		time.Sleep(time.Millisecond * time.Duration(rate)) //rate limits based on argument
		go performCall(data, c)
	}
}

func incrementBar(bar *pb.ProgressBar, c chan int) {
	for {
		inc := <-c
		if inc == 1 {
			bar.Increment()
		}
	}
}
func buildUrlData(d *url.Values, k string, v string) {
	d.Set(k, v)
}

func makeRequest(data url.Values) *http.Request {
	req, _ := http.NewRequest("POST", runningConfig.Host+runningConfig.Endpoint, strings.NewReader(data.Encode()))
	for _, header := range runningConfig.Headers {
		if header.Type == "Cookie" {
			cookie := http.Cookie{Name: header.Name, Value: header.Value}
			req.AddCookie(&cookie)
		} else {
			req.Header.Set(header.Type, header.Value)
		}

	}
	return req
}

func performCall(data url.Values, c chan int) {
	client := http.Client{}
	req := makeRequest(data)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("error")
	}
	if resp.StatusCode != http.StatusOK {
		c <- -1
		performCall(data, c)
	}
	bodybytes, _ := ioutil.ReadAll(resp.Body)
	var gr goodResponse
	json.Unmarshal(bodybytes, &gr)
	//purpose built TODO: Refactor
	if gr.Result != true {
		c <- -1
		performCall(data, c)
	}
	defer resp.Body.Close()
	c <- 1
}

type goodResponse struct {
	Result bool   `json:"result"`
	Reason string `json: "reason"`
}
