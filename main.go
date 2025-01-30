package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("expense-tracker>")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.Split(input, " ")
		if len(args) == 0 {
			continue
		}

		command := strings.ToLower(args[0])

		switch {
		case command == "add" && isValidAddInput(args):
			fmt.Println("add")
		case command == "list":
			fmt.Println("list")
		case command == "summary":
			fmt.Println("summary")
		case command == "delete" && isValidDeleteInput(args):
			fmt.Println("delete")
		case command == "exit":
			fmt.Println("exiting program")
			return
		case command == "clear":
			fmt.Print("\033[H\033[2J")
		default:
			fmt.Println("unknown command", args)
		}
	}
}
