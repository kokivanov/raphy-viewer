package main

import (
	"context"
	"raphyviewer/env"
)

type LocalSettings struct {
	Locale *string `json:"locale"`
}

// App struct
type App struct {
	ctx context.Context
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
