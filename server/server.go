package server

import (
	"net/http"

	"github.com/HDYS-TTBYS/go-todo-api/config"
	"github.com/HDYS-TTBYS/go-todo-api/controllers"
)

// Init initialize server
func Init() error {
	c := config.GetConfig()
	ro := NewRouter(
		controllers.NewHealthController(),
		controllers.NewCsrfController(),
		controllers.NewAuthController(),
	)
	r, err := ro.StartServer(c)
	if err != nil {
		return err
	}

	http.ListenAndServe(":"+c.LISTEN_PORT, r)
	return nil
}
