package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "golang-pratique",
	Short: "This is the first command",
	Long: `A longer description 
	for the first command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("if needed : golang-pratique -h")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

}
func initConfig() {

}
