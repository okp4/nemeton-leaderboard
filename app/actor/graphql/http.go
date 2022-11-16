package graphql

import (
	"context"
	"net/http"
	"sync"

	"github.com/rs/zerolog/log"
)

type server struct {
	srv *http.Server
	wg  *sync.WaitGroup
}

func makeHTTPServer(listenAddr string, router *http.ServeMux) *server {
	return &server{
		srv: &http.Server{
			Addr:    listenAddr,
			Handler: router,
		},
		wg: &sync.WaitGroup{},
	}
}

func (s *server) start() {
	s.wg.Add(1)

	go func() {
		defer s.wg.Done()

		if err := s.srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("HTTP server stopped unexpectedly")
		}
	}()
}

func (s *server) stop(ctx context.Context) error {
	if err := s.srv.Shutdown(ctx); err != nil {
		return err
	}

	s.wg.Wait()
	return nil
}
