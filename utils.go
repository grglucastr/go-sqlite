package main

import (
	"bufio"
	"fmt"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n') // read until the user press enter
	return strings.TrimSpace(input), err
}
