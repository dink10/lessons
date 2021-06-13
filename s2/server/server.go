package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	host    string
	port    int
	storage map[string]Student
}

type Student struct {
	Name string
}

func NewServer(host string, port int) *Server {
	return &Server{
		host:    host,
		port:    port,
		storage: make(map[string]Student),
	}
}

func (s *Server) Run() error {
	address := fmt.Sprintf("%s:%d", s.host, s.port) // s.host + ":" s.port

	s.initRouter()

	return http.ListenAndServe(address, nil)
}

func (s *Server) initRouter() {
	// localhost:3000 == localhost:3000/
	// localhost:3000/root/1 	localhost:3000/root/a 	localhost:3000/root/$
	http.HandleFunc("/", root) // => root(ResponseWriter, Request)
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
