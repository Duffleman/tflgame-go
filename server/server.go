package server

import (
	"database/sql"
	"fmt"
	"time"

	"tflgame/server/app"
	"tflgame/server/db"
	"tflgame/server/lib/config"
	"tflgame/server/lib/limiter"
	"tflgame/server/lib/tfl"
	"tflgame/server/rpc"

	"github.com/cuvva/ksuid-go"
	_ "github.com/lib/pq" // required for the PGSQL driver to be loaded
	"github.com/sirupsen/logrus"
)

type Config struct {
	Server config.Server `json:"server"`
	Env    string        `json:"env"`

	InternalKeys []string `json:"internal_keys"`

	PostgresURI string       `json:"postgres_uri"`
	Limit       config.Redis `json:"limit"`

	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`

	TFLURL string `json:"tfl_url"`
	TFLKey string `json:"tfl_key"`
}

func DefaultConfig() Config {
	return Config{
		Env: "local",
		Server: config.Server{
			Addr:     "127.0.0.1:3000",
			Graceful: 5,
		},

		InternalKeys: []string{"test"},

		PostgresURI: "postgresql://postgres@localhost/tflgame?sslmode=disable",
		Limit: config.Redis{
			URI: "redis://localhost/1",
		},

		PrivateKey: "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIATvKM7wNazPtLUWvJLwX06XCjwe5VKQvi+b6qd9v6JmoAoGCCqGSM49\nAwEHoUQDQgAEeNY4hNPKbvOurrmSW25dgYNbSn/Oa6NdLDd884awxKlzAdC+o5T0\np5LK4LZYbmSlKsb1a5kftBR4waloCoalug==\n-----END EC PRIVATE KEY-----",
		PublicKey:  "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEeNY4hNPKbvOurrmSW25dgYNbSn/O\na6NdLDd884awxKlzAdC+o5T0p5LK4LZYbmSlKsb1a5kftBR4waloCoalug==\n-----END PUBLIC KEY-----",

		TFLURL: "https://api.tfl.gov.uk",
		TFLKey: "",
	}
}

func Run(cfg Config) error {
	ksuid.SetEnvironment(cfg.Env)

	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: false,
	}

	pgDb, err := sql.Open("postgres", cfg.PostgresURI)
	if err != nil {
		panic(err)
	}

	db := db.New(pgDb)

	signingKey, err := config.NewSigningKey([]byte(cfg.PrivateKey), []byte(cfg.PublicKey))
	if err != nil {
		panic(err)
	}

	rateRedis, err := cfg.Limit.Connect()
	if err != nil {
		return fmt.Errorf("redis: %w", err)
	}

	rateLimit := &rpc.Limiter{
		IP: limiter.Tiered{
			&limiter.Sliding{1 * time.Minute, 3, rateRedis, "tflgame/1min/ip"},
			&limiter.Sliding{1 * time.Hour, 60, rateRedis, "tflgame/1hr/ip"},
		},
	}

	tflClient := tfl.NewClient(cfg.TFLURL, cfg.TFLKey)

	app := app.New(db, signingKey, tflClient, logger)

	r := rpc.New(app, logger, cfg.InternalKeys, rateLimit)

	return r.Run(cfg.Server)
}
