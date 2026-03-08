package utils

import (
	"bufio"
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

	return str[0], nil
}
