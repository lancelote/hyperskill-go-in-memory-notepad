package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const Prompt = "Enter a command and data: "

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(Prompt)

		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		command, data := input[0], strings.Join(input[1:], " ")

		switch command {
		case "exit":
			fmt.Println("[Info] Bye!")
			return
		default:
			fmt.Println(command + " " + data)
		}
	}
}
