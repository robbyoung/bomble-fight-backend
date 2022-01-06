package main

import (
	"os"
	"strings"

	bomble "bomble-fight/internal/bomble"
	vparse "bomble-fight/pkg/version"

	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

func init() {
	if strings.ToUpper(os.Getenv("ENV")) == "LOCAL" {
		log.SetFormatter(&log.TextFormatter{})
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	// ===========================================================================
	// Load environment variables
	// ===========================================================================
	var (
		env     = strings.ToUpper(os.Getenv("ENV")) // LOCAL, DEV, STG, PRD
		port    = os.Getenv("PORT")                 // server traffic on this port
		version = os.Getenv("VERSION")              // path to VERSION file
	)
	// ===========================================================================
	// Read version information
	// ===========================================================================
	version, err := vparse.ParseVersionFile(version)
	if err != nil {
		log.WithFields(log.Fields{
			"env":  env,
			"err":  err,
			"path": os.Getenv("VERSION"),
		}).Fatal("Can't find a VERSION file")
		return
	}
	log.WithFields(log.Fields{
		"env":     env,
		"path":    os.Getenv("VERSION"),
		"version": version,
	}).Info("Loaded VERSION file")
	// ===========================================================================
	// Initialise data storage
	// ===========================================================================
	store := bomble.NewGameService()
	// ===========================================================================
	// Initialise application context
	// ===========================================================================
	appEnv := bomble.AppEnv{
		Render:    render.New(),
		Version:   version,
		Env:       env,
		Port:      port,
		GameStore: store,
	}
	// ===========================================================================
	// Start application
	// ===========================================================================
	bomble.StartServer(appEnv)
}
