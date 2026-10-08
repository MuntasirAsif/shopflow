package main

import (
	"fmt"
	"github.com/muntasirasif/shopflow/config"
	"github.com/muntasirasif/shopflow/database"
)

func main(){
	config.LoadEnv()
	db, err := database.Connect()
	if err != nil {
		fmt.Println("Unable to connect to database:", err)
		return
	}
	defer db.Close()
	fmt.Println("Connected to database")
}