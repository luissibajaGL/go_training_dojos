package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	args := os.Args
	name := ""

	if !(len(args) > 1) {
		fmt.Println("Whats your name?")
		input, err := readInput()

		if err != nil {
			fmt.Println("There was an error reading the input", err)
			panic(err)
		}

		name = input
	} else {
		name = args[1]
	}

	fmt.Println("Hello " + name)
}

func readInput() (string, error) {
	//return reader.ReadString('\n')
	return "", errors.New("Error happened")
}
