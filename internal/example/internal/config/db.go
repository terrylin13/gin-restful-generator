package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		// Host:     os.Getenv("DB_HOST"),
		// Port:     convertToInt(os.Getenv("DB_PORT")),
		// User:     os.Getenv("DB_USER"),
		// Password: os.Getenv("DB_PASSWORD"),
		// DBName:   os.Getenv("DB_NAME"),
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "123456",
		DBName:   "test",
	}
	return &dbConfig
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

func GetDB() (*gorm.DB, error) {
	dbConfig := BuildDBConfig()
	// return gorm.Open("mysql", DbURL(dbConfig))
	return gorm.Open(mysql.Open(DbURL(dbConfig)), &gorm.Config{})
}

// func convertToInt(s string) int {
// 	i, err := strconv.Atoi(s)
// 	if err != nil {
// 		// handle error
// 		fmt.Println(err)
// 		os.Exit(2)
// 	}
// 	return i
// }
