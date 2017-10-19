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

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		bearerToken := os.Getenv("BEARER_TOKEN")
		if len(bearerToken) == 0 || len(org) == 0 || len(name) == 0 {
			log.Fatal().Msg("BEARER_TOKEN environment var must be set. --org and --name flags must also be set")
		}
		robots.BearerToken = bearerToken

		robot, err := robots.Create(name, org)
		if jsonMode {
			ToJsonPrint(robot)
			return
		}

		fmt.Println("Created robot: ", name, "under", org)
		fmt.Println("Name:", robot.Name)
		fmt.Println("Token:", robot.Token)
		if err != nil {
			log.Fatal().Msgf("error creating new robot: %s", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
}
