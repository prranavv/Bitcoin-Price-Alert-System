package main

import "github.com/go-chi/chi"

func routes(h *Handler) *chi.Mux {
	mux := chi.NewRouter()

	mux.Group(func(mux chi.Router) {
		mux.Post("/login", h.handleLogin)
		mux.Post("/register", h.handleRegisterUser)
	})

	//Authenticate
	mux.Group(func(mux chi.Router) {
		mux.Use(autheticate)
		mux.Post("/alerts/create", h.handleCreateAlert)
		mux.Post("/alerts/delete", h.handleDeleteAlert)
		mux.Get("/alerts/list", h.handleListAlerts)
	})
	return mux
}
