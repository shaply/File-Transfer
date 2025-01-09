package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/shaply/File-Transfer/server/service/guest"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Server struct {
	addr string
	db   *mongo.Database
}

func NewServer(addr string, db *mongo.Database) *Server {
	return &Server{addr: addr, db: db}
}

func (s *Server) Start() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	guestHandler := guest.NewGuestHandler()
	guestHandler.RegisterRoutes(subrouter)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})

	return http.ListenAndServe(s.addr, corsHandler.Handler(router))
}
