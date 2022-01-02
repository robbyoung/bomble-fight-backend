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
	Streak int
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
		Streak: 0,
	}
}
