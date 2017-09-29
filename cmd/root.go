package cmd

import (
	"fmt"
	"os"

	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var org string
var name string
var jsonMode bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "quay-robot-manager",
	Short: "A brief description of your application",
	Long:  ``,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ToJsonPrint(in interface{}) {
	jsonBytes, err := json.MarshalIndent(in, "", "\t")
	if err != nil {
		log.Fatal().Msgf("error marshalling to json: %s", err)
	}
	fmt.Println(string(jsonBytes))
}

func init() {
	RootCmd.PersistentFlags().StringVar(&name, "name", "", "name of the robot")
	RootCmd.PersistentFlags().StringVar(&org, "org", "", "Organization name to use")
	RootCmd.PersistentFlags().BoolVar(&jsonMode, "json", false, "Output mode: json")

	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
