package src

type PayloadReport struct {
	Data []*Person `json:"data"`
}

type Person struct {
	Age            NullInt   `json:"Age"`
	ConfirmDate    string    `json:"omitempty"`
	District       string    `json:"District,omitempty"`
	Gender         string    `json:"Gender,omitempty"`
	GenderEn       string    `json:"GenderEn,omitempty"`
	Nation         string    `json:"Nation,omitempty"`
	NationEn       string    `json:"NationEn,omitempty"`
	Province       NanString `json:"Province,omitempty"`
	ProvinceEn     NanString `json:"ProvinceEn,omitempty"`
	ProvinceID     int       `json:"ProvinceId"`
	StatQuarantine int       `json:"StatQuarantine"`
	No             *string   `json:"No,omitempty"`
}

type Summary struct {
	Province map[string]int `json:"Province"`
	AgeGroup map[string]int `json:"AgeGroup"`
}

func NewSummary() Summary {
	return Summary{
		Province: make(map[string]int),
		AgeGroup: make(map[string]int),
	}
}
