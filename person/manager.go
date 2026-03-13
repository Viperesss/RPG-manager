package person

import (
	"fmt"
)

func (w *Warrior) LevelUp() {
	w.level++
	w.damage = w.characteristics.strength*0.2 + float64(w.level)
	w.hp += w.viability * 0.8
}

func CreateWarrior(name string) person {
	return &Warrior{
		personage: personage{
			Nickname: name,
			damage:   5,
			hp:       15,
			level:    1,
			characteristics: characteristics{
				resistance: 5,
				strength:   10,
				viability:  2,
			},
		},
	}
}

func (m *Mage) LevelUp() {
	m.level++
	m.damage = m.characteristics.strength*0.4 + float64(m.level)
	m.hp += m.viability * 0.7
}

func CreateMage(name string) person {
	return &Mage{personage: personage{
		Nickname: name,
		damage:   4,
		hp:       13,
		level:    1,
		characteristics: characteristics{
			resistance: 3,
			strength:   8,
			viability:  1.5,
		},
	},
	}
}

func (r *Ranger) LevelUp() {
	r.level++
	r.damage = r.characteristics.strength*0.3 + float64(r.level)
	r.hp += r.viability * 0.7
}

func CreateRanger(name string) person {
	return &Ranger{personage: personage{
		Nickname: name,
		damage:   4,
		hp:       14,
		level:    1,
		characteristics: characteristics{
			resistance: 4,
			strength:   9,
			viability:  1.5,
		},
	},
	}
}

func (p *personage) ResistanceUp() {
	p.resistance++
}

func (p *personage) StrengthUp() {
	p.strength++
}

func (p *personage) ViabilityUp() {
	p.viability++
}

func (p *personage) Show() {
	fmt.Printf(
		"\n---------------------------\nГерой %q - имеет:\n%.2f здоровья\n%.2f урона\n---------------------------\nХарактеристики:\n%.2f сопротивления к урону\n%.2f силы\n%.2f жизнеспособности\n---------------------------\nУровень персонажа: %v\n---------------------------\n",
		p.Nickname, p.hp, p.damage, p.resistance, p.strength, p.viability, p.level)
}
