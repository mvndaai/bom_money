package main

import (
	"fmt"

	bommath "github.com/mvndaai/bom_math"
)

func main() {
	bommath.PrintBinary()
	printApproximations()
}

func printApproximations() {
	approximations := []struct {
		name           string
		centsCostBarly int
	}{
		{name: "5lb Baking Flour", centsCostBarly: 1300},                           //https://www.amazon.com/Bobs-Red-Mill-Unbleached-5-pound/dp/B00V911X4G/ref=sr_1_8?crid=2JFXTXDK5JGUR
		{name: "congress person dalily wage", centsCostBarly: 174_000_00 / 52 / 7}, // https://www.congressionalinstitute.org/2019/02/21/how-much-do-members-of-congress-get-paid-2/
		{name: "$1", centsCostBarly: 100},
		{name: "Silver leah (smallest at) $1", centsCostBarly: 800},
	}

	fmt.Println("Approximations:")
	for _, a := range approximations {
		fmt.Printf("** %s - Barly Cost: $%.2f **\n", a.name, float64(a.centsCostBarly)/100)
		// bommath.PrintDollarValuesArranged(a.centsCostBarly)
		// fmt.Println()
		bommath.BinaryPrintDollarValuesArranged(a.centsCostBarly)
		fmt.Println()
	}
}
