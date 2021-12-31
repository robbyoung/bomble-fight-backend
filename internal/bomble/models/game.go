package models

type Game struct {
	Players map[string]Player
	Bets    map[string]Bet
}

type UserState struct {
	Player Player
	Bet    Bet
}

type GameStorage interface {
	GetUserState(id string) (UserState, error)
}
