package models

type Player struct {
	Id    string
	Name  string
	Money int
}

type PlayerStorage interface {
	AddPlayer(p Player) (Player, error)
	GetPlayers() ([]Player, error)
}
