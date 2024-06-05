package ApiHandler

import (
	"log"
	"net/http"
)

const defaultListenAddr = ":8000"

type ServerOpts struct {
	Port string
}

type Server interface {
	Start() error
}

func NewServer(port string) Server {
	if len(port) == 0 {
		port = defaultListenAddr
	}
	return &ServerOpts{
		Port: port,
	}
}

func (s *ServerOpts) Start() error {
	log.Println("Starting Server at Port: ", s.Port)
	err := http.ListenAndServe(s.Port, nil)
	if err != nil {
		log.Fatalf("Failed to start the server")
		return err
	}

	return nil
}
