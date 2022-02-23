package utils

import "github.com/HDYS-TTBYS/go-todo-api/config"

func IsSecure() bool {
	if config.GetConfig().GO_TODO_ENV == "dev" {
		return false
	} else {
		return true
	}
}
