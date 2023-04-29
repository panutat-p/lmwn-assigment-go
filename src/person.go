package src

type Response struct {
	Data []Person `json:"data"`
}

type Person struct {
	ConfirmDate    string      `json:"confirm_date"`
	No             interface{} `json:"no"`
	Age            int         `json:"age"`
	Gender         string      `json:"gender"`
	GenderEn       string      `json:"gender_en"`
	Nation         interface{} `json:"nation"`
	NationEn       string      `json:"nation_en"`
	Province       string      `json:"province"`
	ProvinceID     int         `json:"province_id"`
	District       interface{} `json:"district"`
	ProvinceEn     string      `json:"province_en"`
	StatQuarantine int         `json:"stat_quarantine"`
}
