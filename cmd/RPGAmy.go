package cmd

import "github.com/spf13/cobra"

var RPGAmy = &cobra.Command{
	Use:   "blabla",
	Short: "blablabla",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func CanFastAttack(knightIsAwake bool) bool {
	var fastAttack bool
	if !knightIsAwake {
		fastAttack = true
	}
	return fastAttack
}

func CanSpy(knightIsAwake bool, archerIsAwake bool, prisonerIsAwake bool) bool {
	var spy bool
	if knightIsAwake || archerIsAwake || prisonerIsAwake {
		spy = true
	}
	return spy
}

func CanSignalPrisoner(archerIsAwake bool, prisonerIsAwake bool) bool {
	var canSignal bool
	if !archerIsAwake && prisonerIsAwake {
		canSignal = true
	}
	return canSignal
}

func CanFreePrisoner(knightIsAwake bool, archerIsAwake bool, prisonerIsAwake bool, petDogIsPresent bool) bool {
	var freePrisoner bool
	if petDogIsPresent && !archerIsAwake {
		freePrisoner = true
	} else if !petDogIsPresent && prisonerIsAwake && !archerIsAwake && !knightIsAwake {
		freePrisoner = true
	}
	return freePrisoner
}
