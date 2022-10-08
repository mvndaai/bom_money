// https://www.churchofjesuschrist.org/study/scriptures/bofm/alma/11?lang=eng

package bommath

import (
	"fmt"
	"strings"
)

var values = map[string]map[string]func(barlyCostCents int) int{}

const (
	MaterialGold   = "gold"
	MaterialSilver = "silver"
)

// 4 Now these are the names of the different pieces of their gold, and of their silver, according to their value. And the names are given by the Nephites, for they did not reckon after the amanner of the Jews who were at Jerusalem; neither did they measure after the manner of the Jews; but they altered their reckoning and their measure, according to the minds and the circumstances of the people, in every generation, until the reign of the judges, they having been bestablished by king Mosiah.
// 5 Now the reckoning is thus—a senine of gold, a seon of gold, a shum of gold, and a limnah of gold.
const (
	GoldSenine = "senine"
	GoldSeon   = "seon"
	GoldShum   = "shum"
	GoldLimnah = "limnah"

	// 19 Now an antion of gold is equal to three shiblons.
	GoldAntion = "antion"
)

// 6 A senum of silver, an amnor of silver, an ezrom of silver, and an onti of silver.
const (
	SilverSenum = "senum"
	SilverAmnor = "amnor"
	SilverEzrom = "ezrom"
	SilverOnti  = "onti"

	// 14 Now this is the value of the lesser numbers of their reckoning—
	// 15 A shiblon is half of a senum; therefore, a shiblon for half a measure of barley.
	// 16 And a shiblum is a half of a shiblon.
	// 17 And a leah is the half of a shiblum.
	SilverShiblon = "shiblon"
	SilverShiblum = "shiblum"
	SilverLeah    = "leah"
)

func init() {
	values[MaterialGold] = map[string]func(barlyCostCents int) int{
		// 7 A senum of silver was equal to a senine of gold, and either for a measure of barley, and also for a measure of every kind of grain.
		GoldSenine: func(costBarly int) int { return costBarly },
		// 8 Now the amount of a seon of gold was twice the value of a senine.
		GoldSeon: func(costBarly int) int { return 2 * costBarly },
		// 9 And a shum of gold was twice the value of a seon.
		GoldShum: func(costBarly int) int { return 4 * costBarly },
		// 10 And a limnah of gold was the value of them all.
		GoldLimnah: func(costBarly int) int { return 7 * costBarly },
	}

	values[MaterialSilver] = map[string]func(barlyCostCents int) int{
		// 7 A senum of silver was equal to a senine of gold, and either for a measure of barley, and also for a measure of every kind of grain.
		SilverSenum: func(costBarly int) int { return costBarly },
		// 11 And an amnor of silver was as great as two senums.
		SilverAmnor: func(costBarly int) int { return 2 * costBarly },
		// 12 And an ezrom of silver was as great as four senums.
		SilverEzrom: func(costBarly int) int { return 4 * costBarly },
		// 13 And an onti was as great as them all.
		SilverOnti: func(costBarly int) int { return 7 * costBarly },
	}

	// 14 Now this is the value of the lesser numbers of their reckoning—
	// 15 A shiblon is half of a senum; therefore, a shiblon for half a measure of barley.
	values[MaterialSilver][SilverShiblon] = func(costBarly int) int { return costBarly / 2 }
	// 16 And a shiblum is a half of a shiblon.
	values[MaterialSilver][SilverShiblum] = func(costBarly int) int { return costBarly / 4 }
	// 17 And a leah is the half of a shiblum.
	values[MaterialSilver][SilverLeah] = func(costBarly int) int { return costBarly / 8 }
	// 18 Now this is their number, according to their reckoning.
	// 19 Now an antion of gold is equal to three shiblons.
	values[MaterialGold][GoldAntion] = func(costBarly int) int { return (costBarly / 2) * 3 }
}

// 3 And the judge received for his wages according to his time—a senine of gold for a day, or a senum of silver, which is equal to a senine of gold; and this is according to the law which was given.
func JudgeDailyWages(costBarley int) int {
	return values[MaterialGold][GoldSenine](costBarley)
}

func GetCentValues(costBarley int) map[string]map[string]int {
	out := map[string]map[string]int{}
	for material, m := range values {
		out[material] = map[string]int{}
		for name, f := range m {
			out[material][name] = f(costBarley)
		}
	}
	return out
}

func GetDollarValues(costBarley int) map[string]map[string]string {
	centValues := GetCentValues(costBarley)

	out := map[string]map[string]string{}
	for material, m := range centValues {
		out[material] = map[string]string{}
		for name, cents := range m {
			out[material][name] = fmt.Sprintf("$%.2f", float64(cents)/100)
		}
	}
	return out
}

func PrintDollarValues(costBarley int) {
	dollarValues := GetDollarValues(costBarley)

	for material, m := range dollarValues {
		for name, dollars := range m {
			// out[material][name] = fmt.Sprintf("$%f", float64(cents)/100)
			fmt.Printf("%s %s = %s\n", material, name, dollars)
		}
	}
}

func PrintDollarValuesArranged(costBarley int) {
	dollarValues := GetDollarValues(costBarley)

	type coin struct {
		name     string
		material string
	}

	list := [][]coin{
		{{name: SilverLeah, material: MaterialSilver}},
		{{name: SilverShiblum, material: MaterialSilver}},
		{{name: SilverShiblon, material: MaterialSilver}},
		{{name: GoldAntion, material: MaterialGold}},
		{{name: SilverSenum, material: MaterialSilver}, {name: GoldSenine, material: MaterialGold}},
		{{name: SilverAmnor, material: MaterialSilver}, {name: GoldSeon, material: MaterialGold}},
		{{name: SilverEzrom, material: MaterialSilver}, {name: GoldShum, material: MaterialGold}},
		{{name: SilverOnti, material: MaterialSilver}, {name: GoldLimnah, material: MaterialGold}},
	}

	for _, c := range list {
		var names []string
		for _, v := range c {
			names = append(names, fmt.Sprint(v.material, " ", v.name))
		}
		joined := strings.Join(names, ", ")
		fmt.Printf("%s - %s\n", dollarValues[c[0].material][c[0].name], joined)
	}
}
