package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var techPalace = &cobra.Command{
	Use:   "techPalace [--name=<random name>]",
	Short: "Print techPalace --name=<random name>",
	Long:  `All software has versions. techPalace's`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		WelcomeMessage(name)
	},
}

func init() {
	rootCmd.AddCommand(techPalace)
	techPalace.PersistentFlags().String("name", "", "Name of visitor")
}

func WelcomeMessage(name string) {
	fmt.Printf("Welcome to the Tech Palace, %s", name)
}
