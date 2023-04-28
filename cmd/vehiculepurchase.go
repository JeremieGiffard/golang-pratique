package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

type Params struct {
	option1, option2 string
	licence          bool
}

var vehiculepurchase = &cobra.Command{
	Use:   "vehiculepurchase <type of vehicule> <name of vehicule1> <name of vehicule2> ",
	Short: "vehiculepurchase <type of vehicule> <name of vehicule1> <name of vehicule2> \n Or vehiculepurchase <price> <age> --sell=true",
	Long:  "example : vehiculepurchase car toyota renault \n example with Flag sell : vehiculepurchase 100 2 --sell=true",
	Run: func(cmd *cobra.Command, args []string) {
		sell, _ := cmd.Flags().GetBool("sell")
		if sell {
			CalculateResellPrice(args[0], args[1])
		} else {
			ChooseVehicle(Params{option1: args[1], option2: args[2], licence: NeedsLicense(args[0])})
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
func CalculateResellPrice(originalPrice string, age string) {
	floatOriginalPrice, _ := strconv.ParseFloat(originalPrice, 64)
	FloatAge, _ := strconv.ParseFloat(age, 64)
	var calculatedPrice float64
	if FloatAge < 3 {
		calculatedPrice = floatOriginalPrice * 0.8
	} else if FloatAge >= 3 && FloatAge < 10 {
		calculatedPrice = floatOriginalPrice * 0.7
	} else {
		calculatedPrice = floatOriginalPrice * 0.5
	}
	fmt.Printf("%geuro estimated price.", calculatedPrice)

}

func init() {
	rootCmd.AddCommand(vehiculepurchase)
	vehiculepurchase.PersistentFlags().Bool("sell", false, "sell my vehicle")

}
