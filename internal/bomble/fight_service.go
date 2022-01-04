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

func (service *GameService) progressFight() {
	var currentStep = service.GameState.Fight
	if currentStep.FightStatus == models.Pending || currentStep.FightStatus == models.Finished {
		return
	}

	ad := int(math.Min(float64(service.getDamage()), float64(currentStep.DefenderHealth)))
	dd := int(math.Min(float64(service.getDamage()), float64(currentStep.AttackerHealth)))

	status := models.Active
	if currentStep.AttackerHealth-dd <= 0 || currentStep.DefenderHealth-ad <= 0 {
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

func (service *GameService) getDamage() int {
	return r.Intn(21) + 5
}
