package bomble

import (
	"encoding/json"
	"net/http"
	"strconv"

	"bomble-fight/internal/bomble/models"
	"bomble-fight/pkg/health"
	"bomble-fight/pkg/status"

	"github.com/gorilla/mux"
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
		AppName: "bomble-fight",
		Version: appEnv.Version,
	}
	appEnv.Render.JSON(w, http.StatusOK, check)
}

func GetUserStateHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	vars := mux.Vars(req)
	state, err := appEnv.GameStore.GetUserState(vars["id"])
	if err != nil {
		response := status.Response{
			Status:  strconv.Itoa(http.StatusNotFound),
			Message: err.Error(),
		}
		appEnv.Render.JSON(w, http.StatusNotFound, response)
		return
	}
	appEnv.Render.JSON(w, http.StatusOK, state)
}

func AddPlayerHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	decoder := json.NewDecoder(req.Body)
	var p models.Player
	err := decoder.Decode(&p)
	if err != nil {
		response := status.Response{
			Status:  strconv.Itoa(http.StatusBadRequest),
			Message: "malformed player object",
		}
		appEnv.Render.JSON(w, http.StatusBadRequest, response)
		return
	}
	p, err = appEnv.GameStore.AddPlayer(p)
	if err != nil {
		response := status.Response{
			Status:  strconv.Itoa(http.StatusBadRequest),
			Message: err.Error(),
		}
		appEnv.Render.JSON(w, http.StatusBadRequest, response)
		return
	}
	appEnv.Render.JSON(w, http.StatusCreated, p)
}

func AddBetHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	decoder := json.NewDecoder(req.Body)
	var b models.Bet
	err := decoder.Decode(&b)
	if err != nil {
		response := status.Response{
			Status:  strconv.Itoa(http.StatusBadRequest),
			Message: "malformed bet object",
		}
		appEnv.Render.JSON(w, http.StatusBadRequest, response)
		return
	}

	b, err = appEnv.GameStore.AddBet(b)
	if err != nil {
		response := status.Response{
			Status:  strconv.Itoa(http.StatusBadRequest),
			Message: err.Error(),
		}
		appEnv.Render.JSON(w, http.StatusBadRequest, response)
		return
	}

	appEnv.Render.JSON(w, http.StatusCreated, b)
}

func ListPlayersHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	list, err := appEnv.GameStore.ListPlayers()
	if err != nil {
		response := status.Response{
			Status:  strconv.Itoa(http.StatusNotFound),
			Message: "can't find any players",
		}
		appEnv.Render.JSON(w, http.StatusNotFound, response)
		return
	}
	responseObject := make(map[string]interface{})
	responseObject["players"] = list
	responseObject["count"] = len(list)
	appEnv.Render.JSON(w, http.StatusOK, responseObject)
}

func ListCombatantsHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	list, err := appEnv.GameStore.ListCombatants()
	if err != nil {
		response := status.Response{
			Status:  strconv.Itoa(http.StatusNotFound),
			Message: "can't find any combatants",
		}
		appEnv.Render.JSON(w, http.StatusNotFound, response)
		return
	}
	responseObject := make(map[string]interface{})
	responseObject["combatants"] = list
	responseObject["count"] = len(list)
	appEnv.Render.JSON(w, http.StatusOK, responseObject)
}

func FightStepHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	step, _ := appEnv.GameStore.GetFightStep()
	appEnv.Render.JSON(w, http.StatusOK, step)
}

func ResetFightHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	appEnv.GameStore.ResetFight()
	appEnv.Render.JSON(w, http.StatusOK, nil)
}
