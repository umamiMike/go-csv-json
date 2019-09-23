package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/smartystreets/csv"
	"github.com/umamimike/go-csv-json/config"
	"github.com/umamimike/go-csv-json/utils"
	"github.com/umamimike/go-csv-json/validate"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	infile := os.Args[1]
	delimiter := os.Args[2]
	processCsv(infile, delimiter)
}

func processCsv(csvfile string, delimiter string) {

	f, err := os.Open(csvfile)
	if err != nil {
		log.Panic(err)
	}
	type keyval []interface{}
	reader := bufio.NewReader(f)
	heads, _, _ := reader.ReadLine()
	headSlice := strings.Split(string(heads), delimiter)

	scanner := csv.NewScanner(f)
	var parsedCsv keyval
	for scanner.Scan() {
		recordMap := make(map[string]string)
		if err := scanner.Error(); err != nil {
			log.Panic(err)
		} else {
			for key, record := range scanner.Record() {

				fmt.Println("the key in the scanner record range is: ", key, ", and the record is: ", record)
				fmt.Println("the key in the headslice is", headSlice[key])
				recordMap[headSlice[key]] = record
			}
		}

		parsedCsv = append(parsedCsv, recordMap)
	}
	j, _ := json.MarshalIndent(parsedCsv, "", "    ")
	ioutil.WriteFile("output.json", j, 0644)

}
