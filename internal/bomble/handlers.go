package bomble

import (
	"net/http"
	"strconv"

	"bomble-fight/pkg/health"
	"bomble-fight/pkg/status"

	log "github.com/sirupsen/logrus"
)

// HandlerFunc is a custom implementation of the http.HandlerFunc
type HandlerFunc func(http.ResponseWriter, *http.Request, AppEnv)

// MakeHandler allows us to pass an environment struct to our handlers, without resorting to global
// variables. It accepts an environment (Env) struct and our own handler function. It returns
// a function of the type http.HandlerFunc so can be passed on to the HandlerFunc in main.go.
func MakeHandler(appEnv AppEnv, fn func(http.ResponseWriter, *http.Request, AppEnv)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Terry Pratchett tribute
		w.Header().Set("X-Clacks-Overhead", "GNU Terry Pratchett")
		// return function with AppEnv
		fn(w, r, appEnv)
	}
}

// HealthcheckHandler returns useful info about the app
func HealthcheckHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	check := health.Check{
		AppName: "go-rest-api-template",
		Version: appEnv.Version,
	}
	appEnv.Render.JSON(w, http.StatusOK, check)
}

func GetBetsHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	list, err := appEnv.BetStore.GetBets()
	if err != nil {
		response := status.Response{
			Status:  strconv.Itoa(http.StatusNotFound),
			Message: "can't find any bets",
		}
		log.WithFields(log.Fields{
			"env":    appEnv.Env,
			"status": http.StatusNotFound,
		}).Error("Can't find any bets")
		appEnv.Render.JSON(w, http.StatusNotFound, response)
		return
	}
	responseObject := make(map[string]interface{})
	responseObject["bets"] = list
	responseObject["count"] = len(list)
	appEnv.Render.JSON(w, http.StatusOK, responseObject)
}
