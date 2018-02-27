package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

//flags
var Csv string
var Source string

var Run = &cobra.Command{
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("You need to supply two arguments, [config file] [csv file to process]")
		}
		if len(args) < 2 {
			return fmt.Errorf("You need to supply a path to a csv file")
		}
		if len(args) < 3 {
			return fmt.Errorf("You need to supply a rate in Milliseconds denoting how fast you want to hit the endpoint")
		}
		return nil
	},
	Use:          "run /path/to/config.json /path/to/file.csv",
	Short:        "run http post to all rows in csv file",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		configtouse, err := LoadConfig(args[0])
		rate, err := strconv.Atoi(args[2])
		if err != nil {
			return fmt.Errorf("could not convert string to int for rate")
		}
		csvfile := args[1]
		fmt.Println(csvfile)
		runningConfig = configtouse
		if err != nil {
			cmd.Println(err)
		}
		processCsv(configtouse, csvfile, rate)
		return nil
	},
}

func createConfig() *cobra.Command {

	myCommand := &cobra.Command{
		Args:  cobra.MinimumNArgs(1),
		Use:   "create-config /path/to/file",
		Short: "generates boilerplate config file for you",
		Long:  "Will create the boilerplate config file for you, supply a filename to generate",
		RunE: func(cmd *cobra.Command, args []string) error {
			f, err := os.OpenFile(args[0], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
			defer f.Close()
			if _, err = f.WriteString(configfiletoecho); err != nil {
				panic(err)
			}
			return nil
		},
	}
	return myCommand
}
