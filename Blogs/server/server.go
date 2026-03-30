package server

import (
	"blogs/router"
	"log"
)

type Server struct {
}

var App = &Server{}

func (*Server) Start(ip, port string) {
	r := router.Router()
	addr := ip + ":" + port
	log.Printf("Server starting on %s\n", addr)
	if err := r.Run(addr); err != nil {
		log.Println(err)
	}
}
