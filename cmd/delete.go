package cmd

import (
	"fmt"

	"os"

	"github.com/joshrendek/quay-robot-manager/robots"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		bearerToken := os.Getenv("BEARER_TOKEN")
		if len(bearerToken) == 0 || len(org) == 0 || len(name) == 0 {
			log.Fatal().Msg("BEARER_TOKEN environment var must be set. --org and --name flags must also be set")
		}
		robots.BearerToken = bearerToken
		if jsonMode {
			log.Fatal().Msg("delete doesnt have a json output mode")
		}
		fmt.Println("Deleting robot: ", name, "under", org)
		err := robots.Delete(name, org)
		if err != nil {
			log.Fatal().Msgf("error creating new robot: %s", err)
		}
		robots, err := robots.All(org)
		if err != nil {
			log.Fatal().Msgf("error listing robots: %s", err)
		}
		fmt.Println("** Robots **")
		for _, r := range robots {
			fmt.Println("Name:", r.Name)
		}
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
