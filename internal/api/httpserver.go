package api

import (
	"io"
	"net/http"

	"github.com/gohttp/app"
	"github.com/gohttp/logger"
	"github.com/gohttp/response"
)

type ProcessFunc func(r io.ReadCloser) error

// Api structure
type Server struct {
	server http.Handler
	*app.App

	processFunc ProcessFunc
}

// New creates a new Server, initializes it and returns it.
func NewHttpServer(p ProcessFunc) *Server {
	api := &Server{processFunc: p, App: app.New()}
	api.Use(logger.New())
	api.Post("/listen", api.requestHandler)
	return api
}

// requestHandler is an HTTP Handler function that listens on /set
// endpoint. Uses the `SetMessage` struct in the `message` package
// to unmarshal the message.
func (s *Server) requestHandler(w http.ResponseWriter, r *http.Request) {
	if err := s.processFunc(r.Body); err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	response.OK(w)
}
