package bomble

import (
	"bomble-fight/internal/bomble/models"
	"errors"
)

var _ models.GameStorage = (*GameService)(nil)

type GameService struct {
	GameState models.Game
}

func NewGameService() models.GameStorage {
	service := GameService{
		GameState: models.Game{
			Players:    make(map[string]models.Player),
			Combatants: make(map[string]models.Combatant),
			Bets:       make(map[string]models.Bet),
		},
	}

	c1 := models.NewCombatant()
	c2 := models.NewCombatant()
	service.GameState.Combatants[c1.Id] = c1
	service.GameState.Combatants[c2.Id] = c2

	return &service
}

func (service *GameService) PopulateCombatants() {

}

func (service *GameService) GetUserState(id string) (models.UserState, error) {
	if p, ok := service.GameState.Players[id]; ok {
		b := service.GameState.Bets[id]
		return models.UserState{
			Player: p,
			Bet:    b,
		}, nil
	}

	return models.UserState{}, errors.New("player data not found")
}

func (service *GameService) AddPlayer(p models.Player) (models.Player, error) {
	service.GameState.Players[p.Id] = p
	return p, nil
}

func (service *GameService) ListPlayers() ([]models.Player, error) {
	var list []models.Player
	for _, p := range service.GameState.Players {
		list = append(list, p)
	}
	return list, nil
}

func (service *GameService) ListCombatants() ([]models.Combatant, error) {
	var list []models.Combatant
	for _, c := range service.GameState.Combatants {
		list = append(list, c)
	}
	return list, nil
}
