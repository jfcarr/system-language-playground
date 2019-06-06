package main

import (
	"fmt"

	"github.com/xrash/smetrics"
)

type DefaultValues struct {
	JaroWinkler   JaroWinklerDefaults
	WagnerFischer WagnerFischerDefaults
}

type JaroWinklerDefaults struct {
	BoostThreshold float64
	PrefixSize     int
}

type WagnerFischerDefaults struct {
	InsertionCost    int
	DeletionCost     int
	SubstitutionCost int
}

var defaults DefaultValues

func getJaroWinkler(string1 string, string2 string) float64 {
	result := smetrics.JaroWinkler(string1, string2, defaults.JaroWinkler.BoostThreshold, defaults.JaroWinkler.PrefixSize)

	return result
}

func getWagnerFischer(string1 string, string2 string) int {
	result := smetrics.WagnerFischer(string1, string2, defaults.WagnerFischer.InsertionCost, defaults.WagnerFischer.DeletionCost, defaults.WagnerFischer.SubstitutionCost)

	return result
}

func testStrings(string1 string, string2 string) {
	jaroWinklerRating := getJaroWinkler(string1, string2)
	wagnerFischerValue := getWagnerFischer(string1, string2)

	fmt.Printf("%s <=> %s  :  JaroWinkler is %.3f, WagnerFischer is %d\n", string1, string2, jaroWinklerRating, wagnerFischerValue)
}

func initDefaults() {
	defaults.JaroWinkler.BoostThreshold = 0.7
	defaults.JaroWinkler.PrefixSize = 4

	defaults.WagnerFischer.InsertionCost = 1
	defaults.WagnerFischer.DeletionCost = 1
	defaults.WagnerFischer.SubstitutionCost = 2
}

func main() {
	initDefaults()

	// Exact
	testStrings("Allan Parsons Project", "Allan Parsons Project")

	// Similar
	testStrings("The Allan Parsons Project", "Allan Parsons Project")
	testStrings("Allan Parsons Project", "Alan Parsons Project")
	testStrings("Allan Parsons Project", "Allen Parsons Project")
	testStrings("Sammy Davis Jr", "Sammy Davis Jr.")
	testStrings("Sammy Davis Jr", "Sammy Davis Junior")
	testStrings("The Captain and Tennille", "The Captain & Tennille")

	// Dissimilar
	testStrings("Jack White", "Sammy Davis Jr.")
}
