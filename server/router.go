package server

import (
	"time"

	"github.com/HDYS-TTBYS/go-todo-api/config"
	"github.com/HDYS-TTBYS/go-todo-api/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	"github.com/go-chi/httprate"
	"github.com/gorilla/csrf"
)

type Router interface {
	StartServer(c *config.Config) (*chi.Mux, error)
}

type router struct {
	healthController controllers.HealthController
	csrfController   controllers.CsrfController
}

func NewRouter(hc controllers.HealthController, cc controllers.CsrfController) Router {
	return &router{hc, cc}
}

//Serverを起動する
func (ro router) StartServer(c *config.Config) (*chi.Mux, error) {
	r := chi.NewRouter()
	//panic
	r.Use(middleware.Recoverer)
	//logging
	l := httplog.NewLogger("go-todo-api", httplog.Options{
		JSON: true,
	})
	r.Use(httplog.RequestLogger(l))
	//timeout
	r.Use(middleware.Timeout(time.Second * 60))
	//Httpレート制限
	r.Use(httprate.LimitByIP(100, 1*time.Minute))
	//Cors
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{c.FRONTEND_URL},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{""},
		//これを追加すると Cookie が取得できる
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	//application/json; charset=UTF-8 のみ許可
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.ContentCharset("UTF-8"))
	//gorilla csrf
	var secure bool
	if c.GO_TODO_ENV == "dev" {
		secure = false
	} else {
		secure = true
	}
	csrfMiddleware := csrf.Protect(
		[]byte(c.CSRF_KEY),
		csrf.TrustedOrigins([]string{c.FRONTEND_URL}),
		csrf.Secure(secure),
		csrf.Path("/"),
	)
	r.Use(csrfMiddleware)

	r.Route("/api/v1", func(api chi.Router) {
		api.Get("/", ro.healthController.Get)
		api.Get("/csrf", ro.csrfController.Get)
	})

	return r, nil
}
