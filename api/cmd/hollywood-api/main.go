package main

import (
	"context"
	"fmt"
	"github.com/hxcuber/hollywood-api/api/internal/api/router"
	"github.com/hxcuber/hollywood-api/api/internal/controller/actorController"
	"github.com/hxcuber/hollywood-api/api/internal/controller/movieController"
	"github.com/hxcuber/hollywood-api/api/internal/controller/systemController"
	"github.com/hxcuber/hollywood-api/api/internal/repository"
	"github.com/hxcuber/hollywood-api/api/pkg/db/pg"
	"github.com/hxcuber/hollywood-api/api/pkg/httpserv"
	"log"
	"os"
	"strings"
)

const (
	host     = "localhost"
	port     = 5432
	username = "hxcuber"
	password = "hxcuber"
	dbname   = "hollywood"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Exiting...")
}

func run(ctx context.Context) error {
	log.Println("Starting app initialization")

	dbOpenConns := 4
	dbIdleConns := 2
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbname)
	conn, err := pg.NewPool(psqlInfo, dbOpenConns, dbIdleConns)
	if err != nil {
		return err
	}

	defer conn.Close()

	rtr, _ := initRouter(ctx, conn)

	log.Println("App initialization completed")

	// err = httpserv.NewServer(rtr.Handler()).Start(ctx)
	// Using TempHandler for now to test flow, will refactor once I know
	// where the routing is meant to go
	err = httpserv.NewServer(rtr.Handler()).Start(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func initRouter(
	ctx context.Context,
	dbConn pg.BeginnerExecutor) (router.Router, error) {
	return router.New(
		ctx,
		strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ","),
		os.Getenv("GQL_INTROSPECTION_ENABLED") == "true",
		systemController.New(repository.New(dbConn)),
		actorController.New(repository.New(dbConn)),
		movieController.New(repository.New(dbConn)),
	), nil
}
