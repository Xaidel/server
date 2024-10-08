package main

import (
	"csprobe/server/common"
	"csprobe/server/inits"
	"csprobe/server/models"
	"fmt"
)

func init() {
	common.LoadEnv()
	inits.ConnectDB()
}

func main() {
	err := inits.DATABASE.AutoMigrate(&models.User{}, &models.Curriculum{}, &models.Course{})
	if err != nil {
		fmt.Println("Error Migrating!!")
		return
	}
	fmt.Println("Models Successfully Migrated!!")
}
