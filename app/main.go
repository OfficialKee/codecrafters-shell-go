package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//repl := true
	// Uncomment this block to pass the first stage
	//fmt.Fprint(os.Stdout, "$ ")

	//// Wait for user input
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		args := strings.Fields(command)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if args[0] == "exit" {
			if args[1] == "0" {
				break
			}
		}
		if args[0] == "echo" {
			output := strings.Join(args[1:len(args)], " ")
			fmt.Println(output)
			continue
		}
		fmt.Println(command[:len(command)-1] + ": command not found")

	}
}
