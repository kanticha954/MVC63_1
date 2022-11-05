package model

import (
	"MVC63_1/configure"

	_ "github.com/go-sql-driver/mysql"
)

// (input)(sent input)
func GetHospital() (hospital []Hospital, err error) {
	if err = configure.DB.Select("hid,title").
		Table("covid.hospital").
		Find(&hospital).Error; err != nil {
		return

	}
	return

}

func GetPatient() (patient []Patient, err error) {
	if err = configure.DB.Select("hnid,firstname,lastname").
		Table("covid.patient").
		Find(&patient).Error; err != nil {
		return

	}
	return

}
