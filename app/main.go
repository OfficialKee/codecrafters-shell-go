package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//repl := true
	// Uncomment this block to pass the first stage
	//fmt.Fprint(os.Stdout, "$ ")

	//// Wait for user input
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if command == "exit 0\n" {
			break
		}
		fmt.Println(command[:len(command)-1] + ": command not found")

	}
}
