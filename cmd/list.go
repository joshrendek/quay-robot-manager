package cmd

import (
	"os"

	"fmt"

	"github.com/joshrendek/quay-robot-manager/robots"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		bearerToken := os.Getenv("BEARER_TOKEN")
		if len(bearerToken) == 0 || len(org) == 0 {
			log.Fatal().Msg("BEARER_TOKEN environment var must be set. --org flag must also be set")
		}
		robots.BearerToken = bearerToken
		robots, err := robots.All(org)
		if err != nil {
			log.Fatal().Msgf("error listing robots: %s", err)
		}

		if jsonMode {
			ToJsonPrint(robots)
			return
		}

		fmt.Println("** Robots **")
		for _, r := range robots {
			fmt.Println("Name:", r.Name)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
