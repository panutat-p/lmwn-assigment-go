package src

import "fmt"

func GenerateSummary(report []*Person, provinces Set) Summary {
	var (
		summary = NewSummary()
	)

	for _, e := range report {
		fmt.Printf("%+v\n", e)
		key := string(e.ProvinceEn)
		if provinces.Has(key) {
			if _, ok := summary.Province[key]; !ok {
				summary.Province[key] = e.ProvinceID
			}
			group := DetermineAge(e.Age)
			summary.AgeGroup[group] += 1
		}
	}

	return summary
}

func DetermineAge(v *int) string {
	if v == nil {
		return "N/A"
	}

	age := *v
	switch {
	case age >= 0 && age <= 30:
		return "0-30"
	case age <= 60:
		return "0-60"
	case age > 60:
		return "61+"
	default:
		return "N/A"
	}
}
