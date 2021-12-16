package passport

import (
	"net/http"
)

func GetBetsHandler(w http.ResponseWriter, req *http.Request, appEnv AppEnv) {
	list, _ := appEnv.BetStore.GetBets()
	// if err != nil {
	// 	response := status.Response{
	// 		Status:  strconv.Itoa(http.StatusNotFound),
	// 		Message: "can't find any bets",
	// 	}
	// 	log.WithFields(log.Fields{
	// 		"env":    appEnv.Env,
	// 		"status": http.StatusNotFound,
	// 	}).Error("Can't find any bets")
	// 	appEnv.Render.JSON(w, http.StatusNotFound, response)
	// 	return
	// }
	responseObject := make(map[string]interface{})
	responseObject["bets"] = list
	responseObject["count"] = len(list)
	appEnv.Render.JSON(w, http.StatusOK, responseObject)
}
