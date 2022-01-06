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
	Player      Player
	Bet         Bet
	FightStatus FightStatus
}

type Game struct {
	Players     map[string]Player
	Combatants  map[string]Combatant
	Bets        map[string]Bet
	Fight       FightStep
	PlayerCount int
	BetCount    int
}

type GameStorage interface {
	GetUserState(id string) (UserState, error)
	AddPlayer(p Player) (Player, error)
	AddBet(b Bet) (Bet, error)
	ListPlayers() ([]Player, error)
	ListCombatants() ([]Combatant, error)
	GetFightStatus() (FightStatus, error)
	GetFightStep() (FightStep, error)
	ResetFight()
}
