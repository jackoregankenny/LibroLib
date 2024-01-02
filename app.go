package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// App struct holds application state and methods
type App struct {
	ctx    context.Context
	server *http.Server
}

// NewApp creates a new instance of App
func NewApp() *App {
	return &App{}
}

// startup is called at application start
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	// Initialization code here

	// Start HTTP server
	a.startServer()
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Cleanup code here

	// Stop HTTP server
	a.stopServer()
}

// Greet is an example method that can be called from the frontend
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

// startServer starts the HTTP server
func (a *App) startServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", a.handleRoot)

	a.server = &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
}

// stopServer stops the HTTP server gracefully
func (a *App) stopServer() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		log.Printf("Failed to stop server gracefully: %v", err)
	}
}

// handleRoot handles the root request
func (a *App) handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the app!")
}
