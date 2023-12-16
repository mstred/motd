package message

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func Read(message string) string {
	fmt.Print(message)
	phrase, _ := reader.ReadString('\n')
	return strings.TrimSpace(phrase)
}

func Readtovar(_var *string, message string) {
	*_var = Read(message)
}
