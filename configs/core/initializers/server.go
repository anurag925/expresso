package initializers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

type Server interface {
	Provider() *HttpServer
	Start(port int)
}

type HttpServer struct {
	*echo.Echo
}

func NewHttpServer(echo *echo.Echo) *HttpServer {
	return &HttpServer{echo}
}

type EchoServer struct {
	server *HttpServer
}

var _ Server = (*EchoServer)(nil)

func NewEchoServer(server *HttpServer) *EchoServer {
	return &EchoServer{server}
}

func (s *EchoServer) Provider() *HttpServer {
	return s.server
}

func (s *EchoServer) Start(port int) {
	// Start Server()
	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      s.server,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := s.server.StartServer(&srv); err != nil && err != http.ErrServerClosed {
			panic("shutting down the Server: " + err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the Server() with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		panic("unable to close application " + err.Error())
	}
}

// // Start the scheduler service to queue periodic tasks
// go func() {
// 	if err := init.app.Tasks().StartScheduler(); err != nil {
// 		panic("scheduler shutdown: " + err.Error())
// 	}
// }()
