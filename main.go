package main

import (
	"github.com/HDYS-TTBYS/go-todo-api/config"
	"github.com/HDYS-TTBYS/go-todo-api/database"
	"github.com/HDYS-TTBYS/go-todo-api/server"
	_ "gorm.io/driver/postgres"
)

func main() {
	config.Init()
	database.Init(false)
	if err := server.Init(); err != nil {
		panic(err)
	}
}
