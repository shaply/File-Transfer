package types

import "github.com/gorilla/mux"

type Handler interface {
	RegisterRoutes(ruoter *mux.Router)
}
