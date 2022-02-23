package controllers

import (
	"net/http"

	"github.com/gorilla/csrf"
)

type CsrfController interface {
	Get(w http.ResponseWriter, r *http.Request)
}
type csrfController struct{}

func NewCsrfController() CsrfController {
	return &csrfController{}
}

type csrfResponse struct {
	X_CSRF_Token string `json:"X-CSRF-Token"`
}

func newCsrfResponse(token string) *csrfResponse {
	return &csrfResponse{token}
}

func (hc *csrfController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-CSRF-Token", csrf.Token(r))
	responseJson(w, r, http.StatusOK, newCsrfResponse(csrf.Token(r)))
}
