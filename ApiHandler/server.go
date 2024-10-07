package ApiHandler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const defaultListenAddr = ":8000"

type ServerOpts struct {
	Port string
}

type Server interface {
	Start(router *mux.Router) error
}

func NewServer(port string) Server {
	if len(port) == 0 {
		port = defaultListenAddr
	}
	return &ServerOpts{
		Port: port,
	}
}

func (s *ServerOpts) Start(router *mux.Router) error {
	log.Println("Starting Server at Port: ", s.Port)
	err := http.ListenAndServe(s.Port, router)
	if err != nil {
		log.Fatalf("Failed to start the server")
		return err
	}

	return nil
}
