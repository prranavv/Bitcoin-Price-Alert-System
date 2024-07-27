package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtkey = []byte("secret_key")

// handleRegisterUser creates an User in the user table
func (h *Handler) handleRegisterUser(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = h.Db.InsertingIntoUser(credentials)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// NewHandler returns a pointer to a Handler
func NewHandler(d *DB) *Handler {
	return &Handler{Db: d}
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var credential Credentials
	err := json.NewDecoder(r.Body).Decode(&credential)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	expectedpassword, err := h.Db.GettingFromUser(credential.Username)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(expectedpassword)
	if expectedpassword != credential.Password {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(32)
	expirationtime := time.Now().Add(time.Minute * 5)
	claims := &Claims{
		Username: credential.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationtime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationtime,
	})
	fmt.Println("Logged in")
}

// handleCreateAlert is a function that creates an alert for the user.
func (h *Handler) handleCreateAlert(w http.ResponseWriter, r *http.Request) {
	var req RequestPostData
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("Error Occured at Decoding ")
		return
	}
	err = h.Db.AddingToAlert(req.Price, "Created")
	if err != nil {
		fmt.Println(err)
	}
}

// handleDeleteAlert is a function deletes an alert that is given in the request
func (h *Handler) handleDeleteAlert(w http.ResponseWriter, r *http.Request) {
	var req RequestDeleteData
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("Error Occured at Decoding ")
		return
	}
	err = h.Db.UpdatingFromAlert(req.AlertID, "Deleted")
	if err != nil {
		fmt.Println(err)
	}
}

// handleListAlerts is a function lists all alerts.
func (h *Handler) handleListAlerts(w http.ResponseWriter, r *http.Request) {
	alerts, err := h.Db.GettingFromAlert()
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	i := 0
	for _, alert := range alerts {
		i++
		fmt.Printf("%d. AlertID-%d Price-%d Status-%s \n", i, alert.AlertID, alert.Price, alert.Status)
	}

}
