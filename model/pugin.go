package model

import (
	"MVC63_1/configure"

	_ "github.com/go-sql-driver/mysql"
)

// (input)(sent input)
func GetHospital() (hospital Hospital, err error) {
	if err = configure.DB.Select("Title").
		Table("covid.hospital").
		Find(&hospital).Error; err != nil {
		return

	}
	return

}
