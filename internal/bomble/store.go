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
	return &GameService{
		GameState: models.Game{
			Players: make(map[string]models.Player),
			Combatants: []models.Combatant{
				models.NewCombatant(),
				models.NewCombatant()},
			Bets: make(map[string]models.Bet),
		},
	}
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

func (service *GameService) AddBet(b models.Bet) (models.Bet, error) {
	if _, ok := service.GameState.Players[b.PlayerId]; !ok {
		return models.Bet{}, errors.New("not a valid player id")
	}

	// change this to support the combatant array
	// if _, ok := service.GameState.Combatants[b.CombatantId]; !ok {
	// 	return models.Bet{}, errors.New("not a valid combatant id")
	// }

	service.GameState.Bets[b.PlayerId] = b
	return b, nil
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
