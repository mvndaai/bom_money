package bommath_test

import (
	"fmt"

	bommath "github.com/mvndaai/bom_math"
)

func ExampleJudgeDailyWages() {
	approximations := []struct {
		name           string
		centsCostBarly int
	}{}

	for _, a := range approximations {
		fmt.Printf("%s - %v", a.name, bommath.JudgeDailyWages(a.centsCostBarly))
	}
}
