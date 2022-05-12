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

	service.SetupFight(c1.Id, c2.Id)

	return &service
}

func (service *GameService) PopulateCombatants() {

}

func (service *GameService) GetUserState(id string) (models.UserState, error) {
	if p, ok := service.GameState.Players[id]; ok {
		b := service.GameState.Bets[id]
		return models.UserState{
			Player:      p,
			Bet:         b,
			FightStatus: service.GameState.Fight.FightStatus,
		}, nil
	}

	return models.UserState{}, errors.New("player data not found")
}

func (service *GameService) AddPlayer(p models.Player) (models.Player, error) {
	if _, ok := service.GameState.Players[p.Id]; ok {
		return models.Player{}, errors.New("this player has already been created")
	}

	service.GameState.Players[p.Id] = p
	service.GameState.PlayerCount++
	return p, nil
}

func (service *GameService) AddBet(b models.Bet) (models.Bet, error) {
	if _, ok := service.GameState.Bets[b.PlayerId]; ok {
		return models.Bet{}, errors.New("this player already has a bet placed")
	}

	if _, ok := service.GameState.Players[b.PlayerId]; !ok {
		return models.Bet{}, errors.New("not a valid player id")
	}

	if _, ok := service.GameState.Combatants[b.CombatantId]; !ok {
		return models.Bet{}, errors.New("not a valid combatant id")
	}

	if service.GameState.Players[b.PlayerId].Money < b.Amount {
		return models.Bet{}, errors.New("this player doesn't have enough money")
	}

	if b.Amount <= 0 {
		return models.Bet{}, errors.New("bet amount needs to be greater than zero")
	}

	service.GameState.Bets[b.PlayerId] = b

	updatedPlayer := service.GameState.Players[b.PlayerId]
	updatedPlayer.Money = updatedPlayer.Money - b.Amount
	service.GameState.Players[b.PlayerId] = updatedPlayer

	service.GameState.BetCount++
	if service.GameState.BetCount == service.GameState.PlayerCount {
		service.StartFight()
	}

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
