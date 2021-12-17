package bomble

import (
	"bomble-fight/internal/bomble/models"
)

// Compile-time proof of interface implementation
var _ models.BetStorage = (*BetService)(nil)

// BetService will hold the connection and key db info
type BetService struct {
	BetList  map[int]models.Bet
	MaxBetID int
}

// NewBetService creates a new Carer Service with the system's database connection
func NewBetService(list map[int]models.Bet, count int) models.BetStorage {
	return &BetService{
		BetList:  list,
		MaxBetID: count,
	}
}

// ListBets returns a list of JSON documents
func (service *BetService) GetBets() ([]models.Bet, error) {
	var list []models.Bet
	for _, v := range service.BetList {
		list = append(list, v)
	}
	return list, nil
}

// // GetBet returns a single JSON document
// func (service *BetService) GetBet(i int) (models.Bet, error) {
// 	Bet, ok := service.BetList[i]
// 	if !ok {
// 		return models.Bet{}, stacktrace.NewError("Failure trying to retrieve bet")
// 	}
// 	return Bet, nil
// }

// AddBet adds a Bet JSON document, returns the JSON document with the generated id
func (service *BetService) AddBet(b models.Bet) (models.Bet, error) {
	service.MaxBetID = service.MaxBetID + 1
	b.Id = service.MaxBetID
	service.BetList[service.MaxBetID] = b
	return b, nil
}

// // UpdateBet updates an existing Bet
// func (service *BetService) UpdateBet(u models.Bet) (models.Bet, error) {
// 	id := u.ID
// 	_, ok := service.BetList[id]
// 	if !ok {
// 		return u, stacktrace.NewError("Failure trying to update Bet")
// 	}
// 	service.BetList[id] = u
// 	return service.BetList[id], nil
// }

// // DeleteBet deletes a Bet
// func (service *BetService) DeleteBet(i int) error {
// 	_, ok := service.BetList[i]
// 	if !ok {
// 		return stacktrace.NewError("Failure trying to delete Bet")
// 	}
// 	delete(service.BetList, i)
// 	return nil
// }

func CreateMockBets() (map[int]models.Bet, int) {
	list := make(map[int]models.Bet)
	list[0] = models.Bet{
		Id:          0,
		CombatantId: "Volgarr",
		PlayerId:    "Bert",
		Amount:      100,
	}
	list[1] = models.Bet{
		Id:          1,
		CombatantId: "Hurburt",
		PlayerId:    "Len",
		Amount:      120,
	}
	return list, len(list) - 1
}
