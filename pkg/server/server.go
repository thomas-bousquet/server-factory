package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/thomas-bousquet/server-factory/internal/http/handler"
	"github.com/thomas-bousquet/server-factory/pkg/server/config"
	"github.com/thomas-bousquet/server-factory/pkg/server/logger"
	"go.uber.org/zap"
)

type Routes map[string]func(resp http.ResponseWriter, req *http.Request)

type Server struct {
	Logger *zap.Logger
	Config config.Config
	Router *mux.Router
}

func NewServer() Server {
	router := mux.NewRouter()

	config, err := config.NewConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	logger, err := logger.NewLogger(config)

	if err != nil {
		log.Fatal(err.Error())
	}

	server := Server{Logger: logger, Router: router, Config: config}

	server.registerDefaultRoutes()
	return server
}

func (s *Server) RegisterRoute(pattern string, handler func(http.ResponseWriter, *http.Request)) {

}

func (s *Server) registerDefaultRoutes() {

	healthCheckPath, healthCheckHandler := handler.HealthCheck()
	serverRoutes := map[string]func(resp http.ResponseWriter, req *http.Request){
		healthCheckPath: healthCheckHandler,
	}

	for path, handler := range serverRoutes {
		s.Router.HandleFunc(path, handler)
	}
}

func (s Server) Serve() {

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", s.Config.AppPort),
		Handler:      s.Router,
		ReadTimeout:  s.Config.ReadTimeout,
		WriteTimeout: s.Config.WriteTimeout,
		IdleTimeout:  s.Config.IdleTimeout,
	}

	startError := make(chan error)
	go func() {
		startError <- srv.ListenAndServe()
	}()

	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-startError:
		s.Logger.Fatal(err.Error())
	case signal := <-quitSignal:
		err := s.gracefulShutdown(signal.String(), &srv)

		if err != nil {
			s.Logger.Fatal(err.Error())
		}
	}
}

func (s Server) gracefulShutdown(sig string, srv *http.Server) error {
	s.Logger.Info("os signal received", zap.String("signal", sig))

	ctx, cancel := context.WithTimeout(context.TODO(), s.Config.GracefullShutdownTimeout)
	defer cancel()
	return srv.Shutdown(ctx)
}
