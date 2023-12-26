package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

var Heading = color.New(color.FgCyan, color.Bold).PrintlnFunc()

func ReadString() string {
	reader := bufio.NewReader(os.Stdin)
	var str string
	str, _ = reader.ReadString('\n')
	reader.Reset(os.Stdin)
	return strings.Split(str, "\n")[0]
}

func Add(num1 int, num2 int) int {
	return num1 + num2
}

func Subtract(num1 int, num2 int) int {
	return num1 - num2
}

func Multiply(num1 int, num2 int) int {
	return num1 * num2
}

func Divide(num1 int, num2 int) (float32, error) {
	if num2 == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return float32(num1) / float32(num2), nil
}

func Log(str interface{}) {
	switch v := str.(type) {
	case map[string]interface{}:
		for key, value := range v {
			fmt.Printf("%v: %v\n", key, value)
		}
	default:
		fmt.Println(v)
	}
}
