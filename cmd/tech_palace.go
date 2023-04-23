package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var techPalace = &cobra.Command{
	Use:   "techPalace [--name=<random name>] [--stars=<Integer>]",
	Short: "Print techPalace --name=<random name> [--stars=<Integer>]",
	Long:  `All software has versions. techPalace's`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		stars, _ := cmd.Flags().GetInt("stars")
		//
		WelcomeMessage(name, stars)

	},
}

func init() {
	rootCmd.AddCommand(techPalace)
	techPalace.PersistentFlags().String("name", "", "Name of visitor")
	techPalace.PersistentFlags().Int("stars", 0, "number of stars")
}

func WelcomeMessage(name string, stars int) {
	Premium := AddBorder(stars)
	fmt.Printf("%s\nWelcome to the Tech Palace, %s\n%s", Premium, name, Premium)
}
func AddBorder(stars int) string {
	return strings.Repeat("*", stars)
}
