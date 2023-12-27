package main

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/yourusername/yourproject/library"
)

// App struct
type App struct {
	ctx        context.Context
	LibManager *library.LibraryManager // Library Manager
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// Initialize Library Manager
	libraryPath := filepath.Join("path", "to", "your", "library", "folder")
	a.LibManager = library.NewLibraryManager(libraryPath)

	// Initialize your database here if needed
	// Example: a.LibManager.InitDatabase()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func main() {
	// Create an instance of the app struct
	app := NewApp()

	// Start your app
	wails.Run(&wails.AppConfig{
		Width:   1024,
		Height:  768,
		Title:   "E-Book Library",
		JS:      js,
		CSS:     css,
		Colour:  "#131313",
		Startup: app.startup,
		Bind: []interface{}{
			app,
			app.LibManager, // Bind your library manager here
		},
	})
}
