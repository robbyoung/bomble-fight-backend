package models

type Bet struct {
	Id          int
	PlayerId    string
	CombatantId string
	Amount      int
}

type BetStorage interface {
	GetBets() ([]Bet, error)
	// GetUser(i int) (User, error)
	// AddUser(u User) (User, error)
	// UpdateUser(u User) (User, error)
	// DeleteUser(i int) error
}
