package commands

import (
	"log"

	"github.com/spf13/cobra"
)

func getSearchName(jokeTerm string) {
	log.Printf("You searched for: %v", jokeTerm)
}

var RandomCommand = &cobra.Command{
	Use:   "search",
	Short: "Get person by name",
	Long:  "This command return a person name",
	RunE: func(cmd *cobra.Command, args []string) error {
		randomterm, _ := cmd.Flags().GetString("term")
		if randomterm != "" {
			getSearchName(randomterm)
		} else {
			log.Print("No flag")
		}
		return nil
	},
}
