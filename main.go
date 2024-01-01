package main

import (
	"LibroLib/library" // Ensure correct import path
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
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
// so we can call the runtime methods. Now with Wails bindings.
func (a *App) startup(ctx context.Context, runtime *wails.Runtime) {
	a.ctx = ctx
	a.LibManager = library.NewLibraryManager("default library path", "default db path")
	runtime.Bind(a.LibManager.AddBookToLibrary) // Binding AddBookToLibrary method
	// Bind other methods as needed
}

// SetLibraryPath updates the library path based on user input
func (a *App) SetLibraryPath(path string, dbPath string) {
	if a.LibManager == nil {
		a.LibManager = library.NewLibraryManager(path, dbPath)
	} else {
		// If the Library Manager is already initialized, update the paths
		a.LibManager.SetLibraryPath(path, dbPath)
	}
}

func main() {
	// Wails app configuration
	app := NewApp()
	wails.Run(&options.App{
		Title:             "LibroLib E-Book Library",
		Width:             1024,
		Height:            768,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		MinWidth:          400,
		MinHeight:         400,
		MaxWidth:          1920,
		MaxHeight:         1080,
		StartHidden:       false,
		HideWindowOnClose: false,
		Assets:            assets,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		OnStartup:         app.startup,
	})
}
