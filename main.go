package main

import (
	"embed"
	"log"
	"net/http"
	"strings"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"raphyviewer/bindings/api"
	"raphyviewer/bindings/data"
	"raphyviewer/internals/api/defaultApi"
	"raphyviewer/internals/core/cache"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {

	lg := logger.NewDefaultLogger()

	// Create an instance of the app structure
	app := NewApp()
	cacheService, err := cache.NewCacheManager("./cache", lg)
	if err != nil {
		log.Fatal(err)
	}

	apiProvider := defaultApi.NewDefaultProvider()
	apiManager, err := api.NewApiService(apiProvider, lg)
	if err != nil {
		log.Fatal(err)
	}

	dataService, err := data.NewDataService("./data", lg)
	if err != nil {
		log.Fatal(err)
	}

	// Create application with options
	err = wails.Run(&options.App{
		Title:     "raphyviewer",
		Width:     1024,
		Height:    768,
		Frameless: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
			Middleware: func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/local") {
						cacheService.ServeHTTP(w, r) //TODO: Replace with data service to firstly check if media is downloaded
						return
					}
					next.ServeHTTP(w, r)
				})
			},
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
			apiManager,
			dataService,
		},
		Logger: lg,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
