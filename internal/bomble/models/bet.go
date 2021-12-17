package models

type Bet struct {
	Id          int
	PlayerId    string
	CombatantId string
	Amount      int
}

type BetStorage interface {
	GetBets() ([]Bet, error)
}
