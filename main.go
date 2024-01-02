package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS // Embedded assets

func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		Title:      "Librolib",
		Width:      800,
		Height:     600,
		OnStartup:  app.startup,
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,
		},
		Assets: assets, // Use the embedded assets
	})
	if err != nil {
		log.Fatal(err)
	}
}
