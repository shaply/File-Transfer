package guest

import (
	"net/http"

	"github.com/gorilla/mux"
)

type GuestHandler struct {
	clients map[string]*http.Client
}

func NewGuestHandler() *GuestHandler {
	return &GuestHandler{
		clients: make(map[string]*http.Client),
	}
}

func (h *GuestHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/guest", func(http.ResponseWriter, *http.Request) {}).Methods("POST")
}
