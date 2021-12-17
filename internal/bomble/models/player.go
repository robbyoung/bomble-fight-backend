package models

type Player struct {
	Id    int
	Name  string
	Money int
}

type PlayerStorage interface {
	AddPlayer(p Player) (Player, error)
	GetPlayers() ([]Player, error)
}
