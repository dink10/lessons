package apiserver

import (
	"fmt"
	"github.com/zn11ch/SimpleWebApp/config"
	"github.com/zn11ch/SimpleWebApp/internal/store"
	"net/http"
)

type ApiServer struct {
	host  string
	port  int
	store *store.Store
}

func New(cfg *config.Config) *ApiServer {
	return &ApiServer{host: cfg.Host, port: cfg.Port}
}

func (s *ApiServer) Start() error {
	s.configureStore()
	s.configureHandlers()
	return http.ListenAndServe(fmt.Sprintf("%s:%d", s.host, s.port), nil)
}

func (s *ApiServer) configureStore() error {
	config := store.NewConfig("host=localhost dbname=students_db sslmode=disable user=postgres password=mysecretpassword")
	st := store.New(config)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}
func (s *ApiServer) configureHandlers() {
	http.HandleFunc("/", s.index)
	http.HandleFunc("/add", s.add)
	http.HandleFunc("/edit/", s.edit)
	http.HandleFunc("/view/", s.view)
}
