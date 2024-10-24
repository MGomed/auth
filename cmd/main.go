package main

import (
	"context"
	"log"
	"time"

	app "github.com/MGomed/auth/internal/app"
)

const timeout = 15 * time.Second

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("couldn't create app: %v", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("couldn't run app: %v", err)
	}
}
