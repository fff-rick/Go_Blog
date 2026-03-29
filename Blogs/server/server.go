package server

import (
	"blogs/router"
	"log"
	"net/http"
)

type Server struct {
}

var App = &Server{}

func (*Server) Start(ip, port string) {
	server := http.Server{
		Addr: ip + ":" + port,
	}
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
