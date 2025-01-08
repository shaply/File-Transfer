package api

import "go.mongodb.org/mongo-driver/v2/mongo"

type Server struct {
	addr string
	db   *mongo.Database
}

func NewServer(addr string, db *mongo.Database) *Server {
	return &Server{addr: addr, db: db}
}

func (s *Server) Start() error {
	return nil
}
