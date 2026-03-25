package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StringReader() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	str := strings.Fields(input)
	if len(str) == 0 {
		return "", fmt.Errorf("Пустой ввод")
	}

	return str[0], nil
}
