package bomble

import (
	"bomble-fight/internal/bomble/models"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func (service *GameService) GetFightStep() (models.FightStep, error) {
	var currentStep = service.GameState.Fight
	service.progressFight()
	return currentStep, nil
}

func (service *GameService) GetFightStatus() (models.FightStatus, error) {
	return service.GameState.Fight.FightStatus, nil
}

func (service *GameService) SetupFight(idLeft string, idRight string) {
	service.GameState.Fight = models.FightStep{
		Left: models.CombatantStatus{
			Id:     idLeft,
			Health: service.GameState.Combatants[idLeft].Health,
			Loss:   0,
			Action: 0,
		},
		Right: models.CombatantStatus{
			Id:     idLeft,
			Health: service.GameState.Combatants[idRight].Health,
			Loss:   0,
			Action: 0,
		},
		FightStatus: models.Pending,
	}
}

func (service *GameService) StartFight() {
	service.GameState.Fight.FightStatus = models.Starting
}

func (service *GameService) ResetFight() {
	service.GameState.Combatants = make(map[string]models.Combatant)
	c1 := models.NewCombatant()
	c2 := models.NewCombatant()
	service.GameState.Combatants[c1.Id] = c1
	service.GameState.Combatants[c2.Id] = c2

	service.SetupFight(c1.Id, c2.Id)
}

func (service *GameService) progressFight() {
	var currentStep = service.GameState.Fight
	if currentStep.FightStatus == models.Pending || currentStep.FightStatus == models.Finished {
		return
	}

	left, right := service.GameState.Combatants[currentStep.Left.Id], service.GameState.Combatants[currentStep.Right.Id]
	lAct, lDam, rAct, rDam := service.getDamage(left, right)

	status := models.Active
	if currentStep.Left.Health-rDam <= 0 || currentStep.Right.Health-lDam <= 0 {
		service.resolveBets()
		status = models.Finished
	}

	service.GameState.Fight = models.FightStep{
		Left: models.CombatantStatus{
			Id:     currentStep.Left.Id,
			Health: currentStep.Left.Health - rDam,
			Loss:   rDam,
			Action: lAct,
		},
		Right: models.CombatantStatus{
			Id:     currentStep.Left.Id,
			Health: currentStep.Right.Health - lDam,
			Loss:   lDam,
			Action: rAct,
		},
		FightStatus: status,
	}
}

func (service *GameService) resolveBets() {
	fight := service.GameState.Fight

	winnerId := ""
	if fight.Left.Health > 0 {
		winnerId = fight.Left.Id
	} else if fight.Right.Health > 0 {
		winnerId = fight.Right.Id
	}

	for _, b := range service.GameState.Bets {
		if b.CombatantId == winnerId {
			updatedPlayer := service.GameState.Players[b.PlayerId]
			updatedPlayer.Money = updatedPlayer.Money + (b.Amount * 2)
			service.GameState.Players[b.PlayerId] = updatedPlayer
		} else if winnerId == "" {
			updatedPlayer := service.GameState.Players[b.PlayerId]
			updatedPlayer.Money = updatedPlayer.Money + b.Amount
			service.GameState.Players[b.PlayerId] = updatedPlayer
		}
	}

	service.GameState.Bets = make(map[string]models.Bet)
	service.GameState.BetCount = 0
}

func (service *GameService) getDamage(left models.Combatant, right models.Combatant) (models.FightAction, int, models.FightAction, int) {
	var att, def models.Combatant

	leftAtt := r.Intn(left.Speed+right.Speed) < left.Speed
	if leftAtt {
		att = left
		def = right
	} else {
		att = left
		def = right
	}

	attAction := models.ActionAttack
	damage := att.Ferocity*2 + r.Intn(10)
	if skillCheck(att.Skill) {
		damage = damage * 2
		attAction = models.ActionCritical
	}

	defAction := models.ActionNothing
	if skillCheck(def.Agility) {
		defAction = models.ActionDodge
		damage = 0
	} else if skillCheck(def.Endurance) {
		defAction = models.ActionBlock
		damage = 0
	}

	if leftAtt {
		return attAction, damage, defAction, 0
	} else {
		return defAction, 0, attAction, damage
	}
}

func skillCheck(stat int) bool {
	return r.Intn(stat+10) > 10
}
