package model

//get datatype
type Patient struct {
	Hnid      string `json:"hnid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Hid       string `json:"hid"`
}

type Hospital struct {
	Hid   string `json:"hid"`
	Title string `json:"title"`
}

type Patient_status struct {
	Hnid        string `json:"hnid"`
	CovidStatus string `json:"covid_status"`
}
