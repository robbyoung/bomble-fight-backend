package bomble

import (
	"bomble-fight/internal/bomble/models"

	"github.com/unrolled/render"
)

// AppEnv holds application configuration data
type AppEnv struct {
	Render   *render.Render
	Version  string
	Env      string
	Port     string
	BetStore models.BetStorage
}

// CreateContextForTestSetup initialises an application context struct
// for testing purposes
func CreateContextForTestSetup() AppEnv {
	testVersion := "0.0.0"
	appEnv := AppEnv{
		Render:   render.New(),
		Version:  testVersion,
		Env:      "LOCAL",
		Port:     "3001",
		BetStore: NewBetService(CreateMockBets()),
	}
	return appEnv
}
