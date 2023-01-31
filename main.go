package main

import (
	"BankDatabase/api"
	"BankDatabase/model"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting server...")
	time.Sleep(3 * time.Second)
	model.InitialMigration()
	api.StartServer()
}
