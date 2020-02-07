package main

import (
	"fmt"
	"log"
)

const (
	ValueRock     = "rock"
	ValuePaper    = "paper"
	ValueScissors = "scissors"
	ValueLizard   = "lizard"
	ValueSpock    = "spock"
)

type Rock string
type Paper string
type Scissor string
type Lizard string
type Spock string

type Weapon interface {
	Beats(val string) bool
}

func (r *Rock) Beats(val string) bool {
	return val == ValueScissors || val == ValueLizard
}

func (r *Paper) Beats(val string) bool {
	return val == ValueRock || val == ValueSpock
}

func (r *Scissor) Beats(val string) bool {
	return val == ValuePaper || val == ValueLizard
}

func (r *Lizard) Beats(val string) bool {
	return val == ValuePaper || val == ValueSpock
}

func (r *Spock) Beats(val string) bool {
	return val == ValueScissors || val == ValueRock
}

func parseWeapon(str string) (Weapon, error) {
	switch str {
	case ValueRock:
		r := Rock(str)
		return &r, nil
	case ValuePaper:
		p := Paper(str)
		return &p, nil
	case ValueScissors:
		s := Scissor(str)
		return &s, nil
	case ValueLizard:
		l := Lizard(str)
		return &l, nil
	case ValueSpock:
		s := Spock(str)
		return &s, nil
	}

	return nil, fmt.Errorf("invalid value '%v' to parse as weapon", str)
}

func main() {
	p1 := ValuePaper
	p2 := ValueScissors

	lh, err := parseWeapon(p1)
	if err != nil {
		log.Fatal(err)
	}
	if lh.Beats(p2) {
		log.Println("P1 wins")
		return
	}
	log.Println("P2 wins")
}
