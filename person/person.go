package person

import (
	"fmt"
)

var PersonageList []*personage

type characteristics struct {
	resistance float64
	strength   float64
	viability  float64
}

type personage struct {
	Nickname string
	damage   float64
	hp       float64
	level    int
	characteristics
}

func NewPerson(name string) personage {
	person := personage{
		Nickname: name,
		damage:   3,
		hp:       15,
		level:    1,
		characteristics: characteristics{
			resistance: 5,
			strength:   10,
			viability:  2,
		},
	}
	return person
}

func (p *personage) LevelUp() {
	p.level++
	p.damage = p.characteristics.strength*0.2 + float64(p.level)
	p.hp += p.viability * 0.8
}

func (p *personage) ResistancehUp() {
	p.resistance++
}

func (p *personage) StrengthUp() {
	p.strength++
}

func (p *personage) ViabilityhUp() {
	p.viability++
}

func (p *personage) Show() {
	fmt.Printf(
		"\n---------------------------\nГерой %q - имеет:\n%.2f здоровья\n%.2f урона\n---------------------------\nХарактеристики:\n%.2f сопротивления к урону\n%.2f силы\n%.2f жизнеспособности\n---------------------------\nУровень персонажа: %v\n---------------------------\n",
		p.Nickname, p.hp, p.damage, p.resistance, p.strength, p.viability, p.level)
}

// Пользователь может:

// создавать персонажей
// прокачивать характеристики
// наносить урон
// лечить персонажа
// смотреть состояние персонажа
