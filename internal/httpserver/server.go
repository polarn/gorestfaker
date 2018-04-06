package httpserver

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	port int
}

func (this *Server) Start() error {
	http.HandleFunc("/", DecodeHttpHandler)
	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", this.port), nil)
}

func NewServer(port int) *Server {
	return &Server{
		port: port,
	}
}

func DecodeHttpHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] [%s] [%s %s] %s\n", r.RemoteAddr, r.Host, r.Method, r.RequestURI, r.UserAgent())
}
