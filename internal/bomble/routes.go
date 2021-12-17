package bomble

// Route is the model for the router setup
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc HandlerFunc
}

// Routes are the main setup for our Router
type Routes []Route

var routes = Routes{
	Route{"Healthcheck", "GET", "/healthcheck", HealthcheckHandler},

	// PLAYERS
	Route{"AddPlayer", "POST", "/player", AddPlayerHandler},
	Route{"GetPlayers", "GET", "/players", GetPlayersHandler},

	// BETS
	Route{"GetBets", "GET", "/bets", GetBetsHandler},
}
