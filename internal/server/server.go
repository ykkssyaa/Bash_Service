package server

import (
	"github.com/ykkssyaa/Bash_Service/internal/service"
	"net/http"
	"time"
)

type HttpServer struct {
	services *service.Services
	//logger   *logger.Logger
}

func NewHttpServer(addr string, services *service.Services) *http.Server {
	srv := &HttpServer{services: services}

	r := NewRouter(srv)

	return &http.Server{
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Addr:           addr,
	}
}
