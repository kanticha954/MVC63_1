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

func GetPatientStatus() (list []Patient_status, err error) {
	if err = configure.DB.Select("p.hnid,firstname,lastname, s.covid_status").
		Table("covid.patient as p , covid.patient_covid_status as s").
		Where("s.covid_status = 'Positive' AND p.hnid = s.hnid").
		Group("p.hnid").
		Find(&list).Error; err != nil {
		return

	}
	return

}

func CountPatient() (count []Count_Patient, err error) {
	if err = configure.DB.Select("hnid, COUNT(hnid) as count_patient").
		Table("covid.patient_covid_status").
		Where("covid_status = 'Positive'").
		Find(&count).Error; err != nil {
		return

	}
	return

}

func CountPerHospital() (count []Count_hopital, err error) {
	if err = configure.DB.Select("h.hid, title, COUNT(h.hid) as count_hospital").
		Table("covid.patient_covid_status as s, covid.hospital as h, covid.patient as p").
		Where("covid_status = 'Positive' AND p.hnid = s.hnid AND h.hid = p.hid").
		Group("h.hid").
		Find(&count).Error; err != nil {
		return

	}
	return

}
