package models

type FightStatus int

const (
	Pending  FightStatus = 0
	Starting FightStatus = 1
	Active   FightStatus = 2
	Finished FightStatus = 3
)

type FightStep struct {
	AttackerId     string
	DefenderId     string
	AttackerHealth int
	DefenderHealth int
	AttackerDamage int
	DefenderDamage int
	FightStatus    FightStatus
}
