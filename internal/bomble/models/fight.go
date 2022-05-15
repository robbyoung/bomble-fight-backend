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
	ActionNothing  FightAction = 0
	ActionAttack   FightAction = 1
	ActionCritical FightAction = 2
	ActionDodge    FightAction = 3
	ActionBlock    FightAction = 4
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
