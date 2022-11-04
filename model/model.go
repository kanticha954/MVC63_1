package model

//get datatype
type Patient struct {
	HNID      string `json:"HNID"`
	FirstName string `json:"Firstname"`
	LastName  string `json:"Lastname"`
	Hid_FK    string `json:"HID"`
}

type Hospital struct {
	Hid   string `json:"HID"`
	Title string `json:"Title"`
}
