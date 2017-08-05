package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

// App contains the essential components of this application
// and provides their references to its methods.
type App struct {
	Config
	engine    *gin.Engine
	SteamPipe *SteamPipe
}

// Config contains the app's configuration options.
type Config struct {
	port             string
	SteamPipeAddress string
}

// NewApp creates a new App, loads environment variables into Config,
// registers routes, and sets up the SteamPipe grpc client.
func NewApp() *App {
	a := &App{}
	a.engine = gin.Default()

	// Set up default configuration.
	defaults := map[string]string{
		"PORT":               ":80",
		"STEAM_PIPE_ADDRESS": "localhost:9407",
	}

	// Load environment variables.
	for k := range defaults {
		env := os.Getenv(k)
		if env != "" {
			defaults[k] = env
		}
	}

	// Set Config.
	a.port = defaults["PORT"]
	a.SteamPipeAddress = defaults["STEAM_PIPE_ADDRESS"]

	a.SteamPipe = NewSteamPipe(a.SteamPipeAddress)

	a.initRoutes()

	return a
}
