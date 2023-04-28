package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Params struct {
	option1, option2 string
	licence          bool
}

var vehiculepurchase = &cobra.Command{
	Use:   "vehiculepurchase <type of vehicule> <name of vehicule1> <name of vehicule2> ",
	Short: "vehiculepurchase <type of vehicule> <name of vehicule1> <name of vehicule2> ",
	Long:  "example : vehiculepurchase car toyota renault",
	Run: func(cmd *cobra.Command, args []string) {

		ChooseVehicle(Params{option1: args[1], option2: args[2], licence: NeedsLicense(args[0])})

	},
}

func NeedsLicense(kind string) bool {
	var doNeedLicence bool
	if kind == "car" || kind == "truck" {
		doNeedLicence = true
	}
	return doNeedLicence
}
func ChooseVehicle(p Params) {
	var choiceMessage string
	var licence string
	if p.licence {
		licence = "if you have a licence, "
	}
	if p.option1 < p.option2 {
		choiceMessage = p.option1
	} else {
		choiceMessage = p.option2
	}
	fmt.Printf("%s%s is clearly the better choice.", licence, choiceMessage)
}
func CalculateResellPrice(originalPrice float64, age float64) {
	var calculatedPrice float64
	if age < 3 {
		calculatedPrice = originalPrice * 0.8
	} else if age >= 3 && age < 10 {
		calculatedPrice = originalPrice * 0.3
	} else {
		calculatedPrice = originalPrice * 0.5
	}
	fmt.Printf("%f estimated price.", calculatedPrice)

}

func init() {
	rootCmd.AddCommand(vehiculepurchase)

}
