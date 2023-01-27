package my_micro_serv

import (
	"appDB/package/api"
	"appDB/package/storage"
	"github.com/gorilla/mux"
)

type Server struct {
	db storage.InstanceDB
	api *api.API
}

func New(db storage.InstanceDB)*Server {
	return &Server{db : db, api : api.NewAPI(db)}
}

func (s *Server) Router()*mux.Router{
	return s.api.Router()
}