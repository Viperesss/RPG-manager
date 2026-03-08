package cli

import (
	"fmt"
	"strconv"

	"github.com/Viperesss/rpg-manager/person"
	"github.com/Viperesss/rpg-manager/utils"
)

func CreatePersonage() {
	fmt.Print("Введите имя героя: ")
	input, err := utils.StringReader()
	if err != nil {
		fmt.Println(err)
	}

	hero := person.NewPerson(input)
	person.PersonageList = append(person.PersonageList, &hero)
	fmt.Printf("Персонаж успешно %q создан\n", input)
}

func CharacteristicUp() {
	fmt.Print("Введите имя героя: ")
	heroName, err := utils.StringReader()
	if err != nil {
		fmt.Println(err)
	}
	// реализовать fmt.Println("Персонаж с таким иминем не найден")

	fmt.Print("\nВыбирете одну характеристику для улучшения:\n1 - Сопротивление к урону\n2 - Сила\n3 - Жизнеспособность\n")
	input, err := utils.StringReader()
	if err != nil {
		fmt.Println(err)
	}
	chose, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	}

	for _, heroes := range person.PersonageList {
		if heroes.Nickname == heroName {
			switch chose {
			case 1:
				heroes.ResistancehUp()
			case 2:
				heroes.StrengthUp()
			case 3:
				heroes.ViabilityhUp()
			}
		}
	}

}

func ShowPeron() {
	fmt.Print("Введите имя героя: ")
	heroName, err := utils.StringReader()
	if err != nil {
		fmt.Println(err)
	}
	// реализовать fmt.Println("Персонаж с таким иминем не найден")

	for _, heroes := range person.PersonageList {
		if heroes.Nickname == heroName {
			heroes.Show()
		}
	}
}
