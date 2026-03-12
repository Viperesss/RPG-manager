package main

import (
	"fmt"
	"strconv"

	"github.com/Viperesss/rpg-manager/cli"
	"github.com/Viperesss/rpg-manager/utils"
)

func main() {
	fmt.Println("CLI - RPG game")
	for {
		fmt.Println("Выберите действие:\n1 - Создать нового персонажа\n2 - Улучшить характеристику\n3 - Показать информацию о персонаже\n0 - Выйти")
		input, err := utils.StringReader()
		if err != nil {
			fmt.Println(err)
			continue
		}
		chose, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch chose {

		case 1:
			cli.CreatePersonage()
		case 2:
			cli.CharacteristicUp()
		case 3:
			cli.ShowPerson()
		case 0:
			return
		}
	}

}
