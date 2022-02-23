package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	firebasecon "github.com/HDYS-TTBYS/go-todo-api/firebaseCon"
	"github.com/HDYS-TTBYS/go-todo-api/utils"
)

type AuthController interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}
type authController struct{}

func NewAuthController() AuthController {
	return &authController{}
}

type authResponse struct {
	Message string `json:"message"`
}

func newAuthResponse(message string) *authResponse {
	return &authResponse{message}
}

func getIDTokenFromBody(r *http.Request) (string, error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	var parsedBody struct {
		IDToken string `json:"idToken"`
	}
	err = json.Unmarshal(b, &parsedBody)
	return parsedBody.IDToken, err
}

func (hc *authController) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	client, err := firebasecon.GetFirebaseApp().Auth(r.Context())
	if err != nil {
		responseJson(w, http.StatusBadRequest, newAuthResponse(err.Error()))
		return
	}
	idToken, err := getIDTokenFromBody(r)
	if err != nil {
		responseJson(w, http.StatusBadRequest, newAuthResponse("Invalid ID token"))
		return
	}

	decoded, err := client.VerifyIDToken(r.Context(), idToken)
	if err != nil {
		responseJson(w, http.StatusUnauthorized, newAuthResponse("Invalid ID token"))
		return
	}
	// サインインが5分より古い場合は、エラーを返します。
	if time.Now().Unix()-decoded.Claims["auth_time"].(int64) > 5*60 {
		responseJson(w, http.StatusUnauthorized, newAuthResponse("Recent sign-in required"))
		return
	}

	expiresIn := time.Hour * 24 * 5
	cookie, err := client.SessionCookie(r.Context(), idToken, expiresIn)
	if err != nil {
		responseJson(w, http.StatusInternalServerError, newAuthResponse("Failed to create a session cookie"))
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    cookie,
		Path:     "/",
		MaxAge:   int(expiresIn.Seconds()),
		HttpOnly: true,
		Secure:   utils.IsSecure(),
	})
	responseJson(w, http.StatusOK, newAuthResponse("success"))
}

func (hc *authController) Logout(w http.ResponseWriter, r *http.Request) {
	client, err := firebasecon.GetFirebaseApp().Auth(r.Context())
	if err != nil {
		responseJson(w, http.StatusBadRequest, newAuthResponse(err.Error()))
		return
	}
	cookie, err := r.Cookie("session")
	if err != nil {
		responseJson(w, http.StatusUnauthorized, newAuthResponse("No ID token"))
		return
	}

	decoded, err := client.VerifySessionCookie(r.Context(), cookie.Value)
	if err != nil {
		responseJson(w, http.StatusUnauthorized, newAuthResponse("Invalid ID token"))
		return
	}
	if err := client.RevokeRefreshTokens(r.Context(), decoded.UID); err != nil {
		responseJson(w, http.StatusInternalServerError, newAuthResponse("Failed to revoke refresh token"))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: 0,
	})
	responseJson(w, http.StatusOK, newAuthResponse("success"))
}
