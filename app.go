package main

import (
	"context"
	"fmt"
	"raphyviewer/env"
	app_state "raphyviewer/internals/core/state"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type LocalSettings struct {
	Locale *string `json:"locale"`
}

// App struct
type App struct {
	ctx   context.Context
	state app_state.AppStateSerivce

	connUnsub func()
}

func ptr(a string) *string {
	return &a
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.state = *app_state.NewAppStateService()

	var ch <-chan app_state.AppState
	go a.state.Heartbeat(ctx)

	ch, a.connUnsub = a.state.ConnectivityState.Subscribe()

	go func() {
		for state := range ch {
			runtime.EventsEmit(a.ctx, "app.connection", fmt.Sprintf("%v", state))
		}
	}()
}

func (a *App) GetPlatform() string {
	return env.PLATFORM
}

func (a *App) GetUserLocalSettings() (LS *LocalSettings) {
	LS = &LocalSettings{
		Locale: ptr("en"),
	}

	return
}

func (a *App) GetConnectionState() app_state.AppState {
	return a.state.Check(a.ctx, true)
}
