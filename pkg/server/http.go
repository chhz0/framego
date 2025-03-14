package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chhz0/gokit/pkg/server/engines"
)

type HttpConfig struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	TLS          *TLSConfig
}

type TLSConfig struct {
	Cert string
	Key  string
}

type httpServer struct {
	cfg    *HttpConfig
	server *http.Server
}

// Listen implements Server.
func (hs *httpServer) ListenAndServe() error {
	if hs.cfg.TLS != nil {
		if err := hs.server.ListenAndServeTLS(hs.cfg.TLS.Cert, hs.cfg.TLS.Key); err != nil && err != http.ErrServerClosed {
			return err
		}
	} else {
		if err := hs.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
	}

	return hs.wait()
}

// Shutdown implements Server.
func (hs *httpServer) Shutdown(ctx context.Context) error {
	return hs.server.Shutdown(ctx)
}

func (hs *httpServer) wait() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(quit)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return hs.server.Shutdown(ctx)
}

func NewHttp(cfg *HttpConfig, engine engines.Handler) Server {
	return &httpServer{
		cfg: cfg,
		server: &http.Server{
			Addr:         cfg.Addr,
			Handler:      engine.Handler(),
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
		},
	}
}
