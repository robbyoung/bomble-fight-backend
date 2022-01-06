package bomble

import (
	"bomble-fight/internal/bomble/models"
	"math"
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

func (service *GameService) SetupFight(attackerId string, defenderId string) {
	service.GameState.Fight = models.FightStep{
		AttackerId:     attackerId,
		DefenderId:     defenderId,
		AttackerHealth: service.GameState.Combatants[attackerId].Health,
		DefenderHealth: service.GameState.Combatants[attackerId].Health,
		FightStatus:    models.Pending,
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

	ad := int(math.Min(float64(service.getDamage()), float64(currentStep.DefenderHealth)))
	dd := int(math.Min(float64(service.getDamage()), float64(currentStep.AttackerHealth)))

	status := models.Active
	if currentStep.AttackerHealth-dd <= 0 || currentStep.DefenderHealth-ad <= 0 {
		service.resolveBets()
		status = models.Finished
	}

	service.GameState.Fight = models.FightStep{
		AttackerId:     currentStep.AttackerId,
		DefenderId:     currentStep.DefenderId,
		AttackerHealth: currentStep.AttackerHealth - dd,
		DefenderHealth: currentStep.DefenderHealth - ad,
		AttackerDamage: ad,
		DefenderDamage: dd,
		FightStatus:    status,
	}
}

func (service *GameService) resolveBets() {
	fight := service.GameState.Fight
	winnerId := fight.AttackerId
	if fight.AttackerHealth == 0 {
		winnerId = fight.DefenderId
	}

	for _, b := range service.GameState.Bets {
		if b.CombatantId == winnerId {
			updatedPlayer := service.GameState.Players[b.PlayerId]
			updatedPlayer.Money = updatedPlayer.Money + (b.Amount * 2)
			service.GameState.Players[b.PlayerId] = updatedPlayer
		}
	}

	service.GameState.Bets = make(map[string]models.Bet)
	service.GameState.BetCount = 0
}

func (service *GameService) getDamage() int {
	return r.Intn(21) + 5
}
