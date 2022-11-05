package model

//get datatype
type Patient struct {
	Hnid      string `json:"hnid"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Hid       string `json:"hid"`
}

type Count_Patient struct {
	CountPatient int    `json:"count_patient"`
	Hnid         string `json:"hnid"`
}

type Hospital struct {
	Hid   string `json:"hid"`
	Title string `json:"title"`
}

type Patient_status struct {
	Hnid        string `json:"hnid"`
	CovidStatus string `json:"covid_status"`
}
