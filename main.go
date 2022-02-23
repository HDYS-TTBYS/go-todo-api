package main

import (
	"github.com/HDYS-TTBYS/go-todo-api/config"
	"github.com/HDYS-TTBYS/go-todo-api/database"
	firebasecon "github.com/HDYS-TTBYS/go-todo-api/firebaseCon"
	"github.com/HDYS-TTBYS/go-todo-api/server"
	_ "gorm.io/driver/postgres"
)

func main() {
	config.Init()
	database.Init(false)
	firebasecon.Init()
	if err := server.Init(); err != nil {
		panic(err)
	}
}
