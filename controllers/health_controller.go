package controllers

import (
	"net/http"
)

type HealthController interface {
	Get(w http.ResponseWriter, r *http.Request)
}
type healthController struct{}

func NewHealthController() HealthController {
	return &healthController{}
}

type healthResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func newRealthResponse(status int, message string) *healthResponse {
	return &healthResponse{status, message}
}

func (hc *healthController) Get(w http.ResponseWriter, r *http.Request) {
	responseJson(w, r, http.StatusOK, newRealthResponse(200, "OK"))
}
