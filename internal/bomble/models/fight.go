package models

type FightStatus int

const (
	Pending  FightStatus = 0
	Active   FightStatus = 1
	Finished FightStatus = 2
)

type FightStep struct {
	AttackerId  string
	DefenderId  string
	Damage      int
	FightStatus FightStatus
}
