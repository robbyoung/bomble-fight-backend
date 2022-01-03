package bomble

import (
	"bomble-fight/internal/bomble/models"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func (service *GameService) GetFightStep() (models.FightStep, error) {
	return service.GameState.Fight, nil
}

func (service *GameService) GetFightStatus() (models.FightStatus, error) {
	return service.GameState.Fight.FightStatus, nil
}

func (service *GameService) startFight() {
	service.GameState.Fight = models.FightStep{
		AttackerId:  service.GameState.Combatants[0].Id,
		DefenderId:  service.GameState.Combatants[1].Id,
		Damage:      service.getDamage(),
		FightStatus: models.Active,
	}
}

func (service *GameService) endFight() {
	service.GameState.Fight.FightStatus = models.Finished
}

func (service *GameService) getDamage() int {
	return r.Intn(21) + 5
}
