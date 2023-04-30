package src

func GenerateSummary(report []*Person, provinces Set) Summary {
	var (
		summary = NewSummary()
	)

	for _, e := range report {
		if provinces.Has(e.ProvinceEn) {
			if _, ok := summary.Province[e.ProvinceEn]; !ok {
				summary.Province[e.ProvinceEn] = e.ProvinceID
			}
			key := DetermineAge(e.Age)
			summary.AgeGroup[key] += 1
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
