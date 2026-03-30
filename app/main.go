package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"

	//"io"
	"os"
	//"os/exec"
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
		switch args[0] {

		case "exit":
			os.Exit(0)
		case "echo":
			output := strings.Join(args[1:len(args)], " ")
			fmt.Println(output)
			continue
		case "type":
			switch args[1] {
			case "exit":
				fmt.Println("exit is a shell builtin")
			case "echo":
				fmt.Println("echo is a shell builtin")
			case "type":
				fmt.Println("type is a shell builtin")
			case "pwd":
				fmt.Println("pwd is a shell builtin")
			default:
				result, err := exec.LookPath(args[1])
				if err != nil {
					fmt.Println(args[1] + ": not found")
				} else {
					fmt.Println(args[1] + " is " + result)
				}
			}
		case "pwd":
			fmt.Println(os.Getwd())

		default:
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = io.Discard
			err = cmd.Run()
			if err != nil {
				fmt.Println(command[:len(command)-1] + ": command not found")
			}
		}

	}
}
