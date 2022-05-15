package models

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Combatant struct {
	Id     string
	Name   string
	Health int
	Streak int

	Ferocity  int
	Agility   int
	Endurance int
	Skill     int
	Speed     int
}

var names = []string{
	"Winslow",
	"Eberhardt",
	"Otto",
	"Alured",
	"Crawford",
	"Willowby",
	"Rollo",
	"Augustus",
	"Forster",
	"Guy",
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewCombatant() Combatant {
	n := names[r.Intn(len(names))]
	id := fmt.Sprintf("%s_%d", n, r.Intn(900)+100)

	return Combatant{
		Id:     strings.ToLower(id),
		Name:   n,
		Health: 50,
		Streak: 0,

		Ferocity:  r.Intn(10) + 1,
		Endurance: r.Intn(10) + 1,
		Skill:     r.Intn(10) + 1,
		Agility:   r.Intn(10) + 1,
		Speed:     r.Intn(10) + 1,
	}
}
