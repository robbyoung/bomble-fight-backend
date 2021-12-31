package bomble

import (
	"bomble-fight/internal/bomble/models"
)

var _ models.PlayerStorage = (*PlayerService)(nil)

type PlayerService struct {
	PlayerList  map[int]models.Player
	MaxPlayerID int
}

func NewPlayerService(list map[int]models.Player, count int) models.PlayerStorage {
	return &PlayerService{
		PlayerList:  list,
		MaxPlayerID: count,
	}
}

func (service *PlayerService) GetPlayers() ([]models.Player, error) {
	var list []models.Player
	for _, v := range service.PlayerList {
		list = append(list, v)
	}
	return list, nil
}

func (service *PlayerService) AddPlayer(p models.Player) (models.Player, error) {
	service.MaxPlayerID = service.MaxPlayerID + 1
	p.Id = ""
	service.PlayerList[service.MaxPlayerID] = p
	return p, nil
}

func CreateMockPlayers() (map[int]models.Player, int) {
	list := make(map[int]models.Player)
	list[0] = models.Player{
		Id:    "",
		Name:  "Wallace",
		Money: 50,
	}
	return list, len(list) - 1
}
