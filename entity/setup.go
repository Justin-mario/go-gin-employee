package entity

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	username := "root"       //user account name
	password := "HUNters007" //password
	host := "127.0.0.1"      //Database address, can be IP or domain name
	port := 3306             //Database port
	Dbname := "employee_db"  //data storage name
	timeout := "10s"         //Connection timeout, 10 seconds

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	database, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic("Failed To Connect To DataBase")
	}

	database.AutoMigrate(&Employee{})

	DB = database

}
