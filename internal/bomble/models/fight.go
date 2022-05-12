package models

type FightStatus int

const (
	Pending  FightStatus = 0
	Starting FightStatus = 1
	Active   FightStatus = 2
	Finished FightStatus = 3
)

type FightAction int

const (
	Nothing FightAction = 0
	Jab     FightAction = 1
	Sweep   FightAction = 2
	Dodge   FightAction = 3
	Block   FightAction = 4
)

type CombatantStatus struct {
	Id     string
	Health int
	Loss   int
	Action FightAction
}

type FightStep struct {
	Left        CombatantStatus
	Right       CombatantStatus
	FightStatus FightStatus
}
