package utils

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func StringReader(r io.Reader) (string, error) {
	reader := bufio.NewReader(r)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	str := strings.TrimSpace(input)
	if len(str) == 0 {
		return "", fmt.Errorf("Пустой ввод")
	}

	return str, nil
}
