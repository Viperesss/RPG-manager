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
		person.PersonageList[input] = person.CreateWarrior(input)
	case 2:
		person.PersonageList[input] = person.CreateMage(input)
	case 3:
		person.PersonageList[input] = person.CreateRanger(input)
	default:
		fmt.Println("Ошибка, выберите из предложенного списка")
		return
	}
	fmt.Println("Персонаж успешно создан")
}

func CharacteristicUp() {
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

	fmt.Print("\nВыбирете одну характеристику для улучшения:\n1 - Сопротивление к урону\n2 - Сила\n3 - Жизнеспособность\n")
	input, err := utils.StringReader()
	if err != nil {
		fmt.Println(err)
		return
	}
	chose, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch chose {
	case 1:
		value.ResistanceUp()
	case 2:
		value.StrengthUp()
	case 3:
		value.ViabilityUp()
	default:
		fmt.Println("Ошибка, выберите из предложенного списка")
		return
	}
	fmt.Println("Характеристика успешно улучшена")
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
