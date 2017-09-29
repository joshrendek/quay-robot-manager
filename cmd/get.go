package cmd

import (
	"os"

	"fmt"

	"github.com/joshrendek/quay-robot-manager/robots"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		bearerToken := os.Getenv("BEARER_TOKEN")
		if len(bearerToken) == 0 || len(org) == 0 || len(name) == 0 {
			log.Fatal().Msg("BEARER_TOKEN environment var must be set. --org and --name flags must also be set")
		}
		robots.BearerToken = bearerToken

		robot, err := robots.Get(name, org)
		if err != nil {
			log.Fatal().Msgf("error creating new robot: %s", err)
		}
		if jsonMode {
			ToJsonPrint(robot)
			return
		}
		fmt.Println("Name:", robot.Name)
		fmt.Println("Token:", robot.Token)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
