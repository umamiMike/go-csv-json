/**
* TODO:  Fix the way I am setting  cookies in the header. Currently If I add one I would overwrite it with the next.
* I am sure there are a few I havent discovered yet, but one I know about is adding cookies to the header.
*
* TODO: it will probably be easier to supply the csv file as an argument rather than in the config
 */
package main

import (
	"fmt"
	// "github.com/davecgh/go-spew/spew"
	"github.com/recursionpharma/go-csv-map"
	"github.com/spf13/cobra"
	"gopkg.in/cheggaaa/pb.v2"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	cmd := &cobra.Command{
		Use:          "run",
		Short:        "Hi All, Use me",
		SilenceUsage: true,
	}
	cmd.AddCommand(printTimeCmd())
	cmd.AddCommand(createConfig())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func processCsv() {
	confi, err := LoadConfig()
	c := make(chan int)
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

	recordCount := len(records)
	bar := pb.StartNew(recordCount)
	for _, v := range records { //v is the map we are going to parse into the values
		data := url.Values{}
		for i, j := range v {
			buildUrlData(&data, i, j)
		}

		time.Sleep(time.Millisecond * 45)
		go performCall(data, c)
		bar.Increment()
	}
}

func buildUrlData(d *url.Values, k string, v string) {
	d.Set(k, v)
}

func makeRequest(data url.Values) *http.Request {
	conf, err := LoadConfig()
	if err != nil {
		fmt.Print("error loading config", err)
	}
	req, _ := http.NewRequest("POST", conf.Host+conf.Endpoint, strings.NewReader(data.Encode()))
	for _, header := range conf.Headers {
		req.Header.Set(header.Type, header.Value)
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
	if resp.StatusCode != 200 {
		fmt.Print("response Status Code:"+resp.Status, " for record: ", data.Get("id"))
		time.Sleep(1 * time.Second)
		performCall(data, c)
	}
	defer resp.Body.Close()
	c <- 1
}

func createConfig() *cobra.Command {

	return &cobra.Command{
		Use:   "create-config",
		Short: "generates boilerplate config file for you",
		Long:  "Will create the boilerplate config file for you",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) >= 0 {
				f, err := os.OpenFile(args[0], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
				defer f.Close()
				if _, err = f.WriteString(configfiletoecho); err != nil {
					panic(err)
				}
			}
			return nil
		},
	}
}
