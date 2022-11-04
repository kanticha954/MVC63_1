package configure

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// read data
func BuildDBConfig() *DBConfig {
	get := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		DBName:   "covid",
		Password: "kanticha954",
	}
	return &get

}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
