package models

type Player struct {
	Id    string
	Name  string
	Money int
}

type Bet struct {
	PlayerId    string
	CombatantId string
	Amount      int
}

type UserState struct {
	Player Player
	Bet    Bet
}

type Game struct {
	Players    map[string]Player
	Combatants map[string]Combatant
	Bets       map[string]Bet
}

type GameStorage interface {
	GetUserState(id string) (UserState, error)
	AddPlayer(p Player) (Player, error)
	AddBet(b Bet) (Bet, error)
	ListPlayers() ([]Player, error)
	ListCombatants() ([]Combatant, error)
}
