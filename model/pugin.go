package model

import (
	"MVC63_1/configure"

	_ "github.com/go-sql-driver/mysql"
)

// (input)(sent input)
func GetPatient() (patient []Patient, err error) {
	if err = configure.DB.Select("hnid,firstname,lastname,hid").
		Table("covid.patient").
		Find(&patient).Error; err != nil {
		return

	}
	return

}

func GetHospital() (hospital []Hospital, err error) {
	if err = configure.DB.Select("hid,title").
		Table("covid.hospital").
		Find(&hospital).Error; err != nil {
		return

	}
	return

}

func CountPatient() (count []Count_Patient, err error) {
	if err = configure.DB.Select("COUNT(hnid) as count_patient").
		Table("covid.patient_covid_status").
		Where("covid_status = 'Positive'").
		Find(&count).Error; err != nil {
		return

	}
	return

}
