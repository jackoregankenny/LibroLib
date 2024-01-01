package main

import (
	"context"
	"fmt"
)

// App struct holds application state and methods
type App struct {
	ctx context.Context
}

// NewApp creates a new instance of App
func NewApp() *App {
	return &App{}
}

// startup is called at application start
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	// Initialization code here
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Cleanup code here
}

// Greet is an example method that can be called from the frontend
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
