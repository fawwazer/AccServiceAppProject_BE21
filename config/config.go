package config

import (
	"Alterra/ALTA_BE21_Rizal/AccServiceProject_BE21/users"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() *gorm.DB {
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "", "localhost", 3306, "db_applogin")
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println("terjadi sebuah kesalahan", err.Error())
		return nil
	}
	return db
}

func Migrate(connection *gorm.DB) error {
	err := connection.AutoMigrate(&users.Users{})
	return err
}
