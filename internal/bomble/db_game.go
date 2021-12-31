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
	players := make(map[string]models.Player)

	players["test"] = models.Player{
		Id:    "test",
		Name:  "Test Player",
		Money: 200,
	}

	return &GameService{
		GameState: models.Game{
			Players: players,
			Bets:    make(map[string]models.Bet),
		},
	}
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
