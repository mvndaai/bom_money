package bommath

import (
	"fmt"
)

const (
	Leah = 1 << iota
	Shiblum
	Shiblon
	Senine
	Seon
	Shum

	Antion = 3 * Shiblon
	Limnah = Senine + Seon + Shum

	Senum = Senine
	Amnor = Seon
	Ezrom = Shum
	Onti  = Limnah
)

func PrintBinary() {
	fmt.Printf("Leah\t%d\n", Leah)
	fmt.Printf("Shiblum\t%d\n", Shiblum)
	fmt.Printf("Shiblon\t%d\n", Shiblon)
	fmt.Printf("Senine\t%d\n", Senine)
	fmt.Printf("Antion\t%d\n", Antion)
	fmt.Printf("Seon\t%d\n", Seon)
	fmt.Printf("Shum\t%d\n", Shum)
	fmt.Printf("Limnah\t%d\n", Limnah)
}

func BinaryPrintDollarValuesArranged(costBarley int) {

	each := []struct {
		name        string
		multipliler float64
	}{
		{name: "silver leah", multipliler: Leah},
		{name: "silver shiblum", multipliler: Shiblum},
		{name: "silver shiblon", multipliler: Shiblon},
		{name: "silver senum, gold senine", multipliler: Senine},
		{name: "gold antion", multipliler: Antion},
		{name: "silver amnor, gold seon", multipliler: Seon},
		{name: "silver ezrom, gold shum", multipliler: Shum},
		{name: "silver onti, gold limnah", multipliler: Limnah},
	}

	base := float64(costBarley) / 8
	for _, v := range each {
		fmt.Printf("$%.2f - %s\n", base*v.multipliler/100, v.name)
	}
}
