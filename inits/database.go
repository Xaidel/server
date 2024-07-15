package inits

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"fmt"
	"csprobe/server/common"
)

var DATABASE *gorm.DB

func ConnectDB(){
	common.LoadEnv()
	DB_USERNAME := common.GetEnv("DB_USERNAME")
	DB_PASSWORD := common.GetEnv("DB_PASSWORD")
	DB_HOST := common.GetEnv("DB_HOST")
	DB_PORT := common.GetEnv("DB_PORT")
	DB_DATABASE := common.GetEnv("DB_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_DATABASE) 
	var err error
	DATABASE, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if DATABASE != nil{
		fmt.Println("Database connected succesfully")
		fmt.Println(DATABASE)
	}
}
