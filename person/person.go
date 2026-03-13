package person

type person interface {
	LevelUp()
	ResistanceUp()
	StrengthUp()
	ViabilityUp()
	Show()
}

var PersonageList = make(map[string]person)

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

type Warrior struct {
	personage
}

type Mage struct {
	personage
}

type Ranger struct {
	personage
}

// Пользователь может:

// создавать персонажей
// прокачивать характеристики
// наносить урон
// лечить персонажа
// смотреть состояние персонажа
