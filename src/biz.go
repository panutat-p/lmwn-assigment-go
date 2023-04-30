package src

func GenerateSummary(report []*Person) Summary {
	var (
		summary = NewSummary()
	)

	for _, e := range report {
		key := string(e.ProvinceEn)
		if _, ok := summary.Province[key]; !ok {
			summary.Province[key] = e.ProvinceID
		}
		group := DetermineAge(int(e.Age))
		summary.AgeGroup[group] += 1
	}

	return summary
}

func GenerateFilteredSummary(report []*Person, provinces Set) Summary {
	var (
		summary = NewSummary()
	)

	for _, e := range report {
		key := string(e.ProvinceEn)
		if provinces.Has(key) {
			if _, ok := summary.Province[key]; !ok {
				summary.Province[key] = e.ProvinceID
			}
			group := DetermineAge(int(e.Age))
			summary.AgeGroup[group] += 1
		}
	}

	return summary
}

func DetermineAge(age int) string {
	if age == -1 {
		return "N/A"
	}

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
