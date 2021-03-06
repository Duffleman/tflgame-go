package config

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis"
)

// Redis configures a connection to a Redis database.
type Redis struct {
	URI string `json:"uri"`
}

// Options returns a configured redis.Options structure.
func (r Redis) Options() (*redis.Options, error) {
	return redis.ParseURL(r.URI)
}

// Connect returns a connected redis.Client instance.
func (r Redis) Connect() (*redis.Client, error) {
	opts, err := r.Options()
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)

	if err := client.Ping().Err(); err != nil {
		return nil, err
	}

	return client, nil
}

// DefaultGraceful is the graceful shutdown timeout applied when no
// configuration value is given.
const DefaultGraceful = 5

// Server configures the binding and security of an HTTP server.
type Server struct {
	Addr string `json:"addr"`

	// Graceful enables graceful shutdown and is the time in seconds to wait
	// for all outstanding requests to terminate before forceably killing the
	// server. When no value is given, DefaultGraceful is used. Graceful
	// shutdown is disabled when less than zero.
	Graceful int `json:"graceful"`
}

// ListenAndServe configures a HTTP server and begins listening for clients.
func (cfg *Server) ListenAndServe(srv *http.Server) error {
	// only set listen address if none is already configured
	if srv.Addr == "" {
		srv.Addr = cfg.Addr
	}

	if cfg.Graceful == 0 {
		cfg.Graceful = DefaultGraceful
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM)

	errs := make(chan error, 1)

	go func() {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			errs <- err
		}
	}()

	select {
	case err := <-errs:
		return err

	case <-stop:
		if cfg.Graceful > 0 {
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Graceful)*time.Second)
			defer cancel()

			return srv.Shutdown(ctx)
		}

		return nil
	}
}
