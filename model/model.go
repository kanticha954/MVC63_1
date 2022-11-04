package model

//get datatype
type Patient struct {
	PatientId     string `json:"hnid"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	HospitalIdRef string `json:"hid"`
}

type Hospital struct {
	Hid   string `json:"HID"`
	Title string `json:"Title"`
}
