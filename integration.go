package connect

import (
	"io"
	"net"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/segmentio/connect/internal/api"
)

// Integration is the iterface defining a Segment connect integration
type Integration interface {
	Init() error

	// Process is a function that takes the incoming body reader
	// from the current request.
	// It is client's responsibility to close the body when ready,
	// usually this can be achieved using `defer r.Close()`.
	Process(r io.ReadCloser) error
}

func Run(i Integration) {
	listenAddress := os.Getenv("LISTEN_ADDRESS")
	if listenAddress == "" {
		listenAddress = ":3000"
	}

	// Create listener
	listener, err := net.Listen("tcp", listenAddress)
	if err != nil {
		logrus.Error(err)
		return
	}

	if os.Getenv("DEBUG") != "" {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if err := i.Init(); err != nil {
		logrus.Fatal(err)
	}

	// Create HTTP server
	server := api.NewHttpServer(i.Process)

	// Run listen and serve
	logrus.Infof("Server started at %v", listener.Addr())
	logrus.Fatal(http.Serve(listener, server))
}
