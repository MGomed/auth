package main

import (
	"context"
	"log"

	app "github.com/MGomed/auth/internal/app"
)

func main() {
	ctx := context.Background()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("couldn't create app: %v", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("couldn't run app: %v", err)
	}
}
