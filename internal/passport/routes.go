package passport

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

	// BETS
	Route{"GetBets", "GET", "/bets", GetBetsHandler},

	//=== USERS ===
	Route{"ListUsers", "GET", "/users", ListUsersHandler},
	Route{"GetUser", "GET", "/users/{uid:[0-9]+}", GetUserHandler},
	Route{"CreateUser", "POST", "/users", CreateUserHandler},
	Route{"UpdateUser", "PUT", "/users/{uid:[0-9]+}", UpdateUserHandler},
	Route{"DeleteUser", "DELETE", "/users/{uid:[0-9]+}", DeleteUserHandler},
}
