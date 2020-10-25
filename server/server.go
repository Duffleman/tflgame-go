package server

import (
	"database/sql"
	"io/ioutil"

	"tflgame/server/app"
	"tflgame/server/db"
	"tflgame/server/lib/config"
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

	PostgresURI string `json:"postgres_uri"`

	PrivateKeyFile string `json:"private_key_file"`
	PublicKeyFile  string `json:"public_key_file"`

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

		PrivateKeyFile: "./ec_private.pem",
		PublicKeyFile:  "./ec_public.pem",

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

	privateKey, err := ioutil.ReadFile(cfg.PrivateKeyFile)
	if err != nil {
		return err
	}

	publicKey, err := ioutil.ReadFile(cfg.PublicKeyFile)
	if err != nil {
		return err
	}

	signingKey, err := config.NewSigningKey(privateKey, publicKey)
	if err != nil {
		panic(err)
	}

	tflClient := tfl.NewClient(cfg.TFLURL, cfg.TFLKey)

	app := app.New(db, signingKey, tflClient, logger)

	r := rpc.New(app, logger, cfg.InternalKeys)

	return r.Run(cfg.Server)
}
