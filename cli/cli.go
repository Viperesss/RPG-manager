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
		return
	}

	fmt.Print("Выберите класс героя\n1 - Воин\n2 - Маг\n3 - Стрелок\n")
	heroClass, err := utils.StringReader()
	if err != nil {
		fmt.Println(err)
		return
	}
	heroChose, err := strconv.Atoi(heroClass)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch heroChose {
	case 1:
		newWarrior := person.Warrior{}
		newHero := newWarrior.Create(input)
		person.PersonageList[input] = newHero
	case 2:
		newMage := person.Mage{}
		newHero := newMage.Create(input)
		person.PersonageList[input] = newHero
	case 3:
		newRanger := person.Ranger{}
		newHero := newRanger.Create(input)
		person.PersonageList[input] = newHero
	default:
		fmt.Println("Ошибка, выберите из предложенного списка")
		return
	}
	fmt.Println("Персонаж успешно создан")
}

func CharacteristicUp() {
	// fmt.Print("Введите имя героя: ")
	// heroName, err := utils.StringReader()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // реализовать fmt.Println("Персонаж с таким иминем не найден")

	// fmt.Print("\nВыбирете одну характеристику для улучшения:\n1 - Сопротивление к урону\n2 - Сила\n3 - Жизнеспособность\n")
	// input, err := utils.StringReader()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// chose, err := strconv.Atoi(input)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// for _, heroes := range person.PersonageList {
	// 	if heroes.Nickname == heroName {
	// 		switch chose {
	// 		case 1:
	// 			heroes.ResistancehUp()
	// 		case 2:
	// 			heroes.StrengthUp()
	// 		case 3:
	// 			heroes.ViabilityhUp()
	// 		}
	// 	}
	// }

}

func ShowPerson() {
	if len(person.PersonageList) == 0 {
		fmt.Println("Список пуст")
	} else {
		fmt.Print("Введите имя героя: ")
		heroName, err := utils.StringReader()
		if err != nil {
			fmt.Println(err)
			return
		}
		value, ok := person.PersonageList[heroName]
		if !ok {
			fmt.Println("Персонаж с таким иминем не найден")
			return
		}
		value.Show()
	}
}
