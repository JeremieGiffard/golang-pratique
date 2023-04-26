package cmd

import (
	"github.com/spf13/cobra"
)

var vehiculepurchase = &cobra.Command{
	Use:   "vehiculepurchase <type of vehicule> <name of vehicule1> <name of vehicule2> ",
	Short: "vehiculepurchase <type of vehicule> <name of vehicule1> <name of vehicule2> ",
	Long:  "example : vehiculepurchase car toyota renault",
	Run: func(cmd *cobra.Command, args []string) {
		if NeedsLicense(args[0]) {
			println("If you have a licence, " + ChooseVehicle(args[1], args[2]))
		} else {
			println(ChooseVehicle(args[1], args[2]))
		}

	},
}

func NeedsLicense(kind string) bool {
	var doNeedLicence bool
	if kind == "car" || kind == "truck" {
		doNeedLicence = true
	}
	return doNeedLicence
}
func ChooseVehicle(option1 string, option2 string) string {
	var choiceMessage string
	if option1 < option2 {
		choiceMessage = option1 + " is clearly the better choice."
	} else {
		choiceMessage = option2 + " is clearly the better choice."
	}
	return choiceMessage
}
func CalculateResellPrice(originalPrice float64, age float64) {

}

func init() {
	rootCmd.AddCommand(vehiculepurchase)

}
