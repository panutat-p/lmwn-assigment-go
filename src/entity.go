package src

type Response struct {
	Data []*Person `json:"data"`
}

type Person struct {
	Age            int    `json:"Age,omitempty"`
	ConfirmDate    string `json:"ConfirmDate,omitempty"`
	District       string `json:"District,omitempty"`
	Gender         string `json:"Gender,omitempty"`
	GenderEn       string `json:"GenderEn,omitempty"`
	Nation         string `json:"Nation,omitempty"`
	NationEn       string `json:"NationEn,omitempty"`
	Province       string `json:"Province,omitempty"`
	ProvinceEn     string `json:"ProvinceEn,omitempty"`
	ProvinceID     int    `json:"ProvinceId"`
	StatQuarantine int    `json:"StatQuarantine"`
	No             string `json:"No,omitempty"`
}

type Summary struct {
	Province map[string]int
	AgeGroup map[string]int
}
