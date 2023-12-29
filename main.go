package main

import (
	"context"
	"embed"

	"LibroLib/library" // Update this import path as necessary

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

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
// so we can call the runtime methods.
// Paths are not set here anymore, waiting for user input through the UI.
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) SetLibraryPath(path string, dbPath string) {
	if a.LibManager == nil {
		a.LibManager = library.NewLibraryManager(path, dbPath)
	} else {
		// If the Library Manager is already initialized, handle updates to paths
		a.LibManager.LibraryPath = path
		// Update the database path if needed
		// a.LibManager.UpdateDatabasePath(dbPath) // This would be a new method in LibraryManager
	}
}

func main() {
	// Create an instance of the app struct
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "LibroLib",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
