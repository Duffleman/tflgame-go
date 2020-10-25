package main

import (
	"database/sql"
	"io/ioutil"
	"os"

	"tflgame/server/app"
	"tflgame/server/db"
	"tflgame/server/lib/config"
	"tflgame/server/lib/tfl"
	"tflgame/server/rpc"

	ksuid "github.com/cuvva/ksuid-go"
	_ "github.com/lib/pq" // required for the PGSQL driver to be loaded
	"github.com/sirupsen/logrus"
)

func main() {
	env := config.EnvironmentName(os.Getenv("env"))

	ksuid.SetEnvironment(env)

	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: false,
	}

	pgDb, err := sql.Open("postgres", "postgresql://postgres@localhost/tflgame?sslmode=disable")
	if err != nil {
		panic(err)
	}

	db := db.New(pgDb)

	privateKey, _ := ioutil.ReadFile("./ec_private.pem")
	publicKey, _ := ioutil.ReadFile("./ec_public.pem")

	signingKey, err := config.NewSigningKey(privateKey, publicKey)
	if err != nil {
		panic(err)
	}

	internalKey := "test"

	tflClient := tfl.NewClient("https://api.tfl.gov.uk")

	app := app.New(db, signingKey, tflClient, logger)

	r := rpc.New(app, logger, internalKey)

	// user endpoints
	r.Route("/authenticate", r.Authenticate, rpc.AuthenticateSchema, rpc.UnsafeNoAuth)
	r.Route("/create_user", r.CreateUser, rpc.CreateUserSchema, rpc.UnsafeNoAuth)
	r.Route("/change_handle", r.ChangeHandle, rpc.ChangeHandleSchema, rpc.JWTAuth)
	r.Route("/release_handle", r.ReleaseHandle, rpc.ReleaseHandleSchema, rpc.JWTAuth)
	r.Route("/change_pin", r.ChangePin, rpc.ChangePinSchema, rpc.JWTAuth)
	r.Route("/list_events", r.ListEvents, rpc.ListEventsSchema, rpc.JWTAuth)

	// server endpoints
	r.Route("/sync_tfl_data", r.SyncTFLData, nil, rpc.InternalOnlyAuth)

	// game endpoints
	r.Route("/test_game_options", r.TestGameOptions, rpc.TestGameOptionsSchema, rpc.UnsafeNoAuth)
	r.Route("/get_game_options", r.GetGameOptions, nil, rpc.UnsafeNoAuth)
	r.Route("/create_game", r.CreateGame, rpc.CreateGameSchema, rpc.JWTAuth)
	r.Route("/submit_answer", r.SubmitAnswer, rpc.SubmitAnswerSchema, rpc.JWTAuth)
	r.Route("/get_current_game", r.GetCurrentGame, rpc.GetCurrentGameSchema, rpc.JWTAuth)
	r.Route("/get_hint", r.GetHint, rpc.GetHintSchema, rpc.JWTAuth)
	r.Route("/get_game_state", r.GetGameState, rpc.GetGameStateSchema, rpc.JWTAuth)
	r.Route("/explain_score", r.ExplainScore, rpc.ExplainScoreSchema, rpc.JWTAuth)
	r.Route("/get_leaderboard", r.GetLeaderboard, nil, rpc.UnsafeNoAuth)

	addr := config.GetEnv("addr").(string)
	r.Serve(addr)
}
