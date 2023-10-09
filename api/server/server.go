package api

import (
	"encoding/json"
	"net/http"
	storage "rentalreviewspt/database"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}
func (s *Server) Start() error {
	http.HandleFunc("/user", s.handleGetUserByID)
	http.HandleFunc("/rent", s.handleGetUserByID)
	return http.ListenAndServe(s.listenAddr, nil)
}
func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request) {

	user := s.store.Get(10)
	json.NewEncoder(w).Encode(user)
}
