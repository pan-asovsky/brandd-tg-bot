package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pan-asovsky/brandd-tg-bot/internal/app"
)

func main() {
	ctx := context.Background()
	a := app.NewApp(ctx)

	if err := a.Init(); err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	if err := a.Run(); err != nil {
		log.Fatalf("failed to run app: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down...")

	a.Close()
	log.Println("App stopped gracefully")
}
