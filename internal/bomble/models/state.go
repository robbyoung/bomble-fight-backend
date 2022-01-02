package models

type Player struct {
	Id    string
	Name  string
	Money int
}

type Bet struct {
	Id          int
	PlayerId    string
	CombatantId string
	Amount      int
}

type UserState struct {
	Player Player
	Bet    Bet
}

type Game struct {
	Players map[string]Player
	Bets    map[string]Bet
}

type GameStorage interface {
	GetUserState(id string) (UserState, error)
	AddPlayer(p Player) (Player, error)
	ListPlayers() ([]Player, error)
}
