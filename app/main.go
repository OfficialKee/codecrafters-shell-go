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

	shell "github.com/codecrafters-io/shell-starter-go/app/checkType"
)

func main() {

	//repl := true
	// Uncomment this block to pass the first stage
	//fmt.Fprint(os.Stdout, "$ ")

	//// Wait for user input
	for {
		fmt.Fprint(os.Stdout, "$ ")
		commands, err := bufio.NewReader(os.Stdin).ReadString('\n')
		args := strings.Fields(commands)
		command := args[0]

		switch command {
		case "exit":
			os.Exit(0)
		case "echo":
			output := strings.Join(args[1:len(args)], " ")
			fmt.Println(output)
			continue
		case "type":
			result := shell.CheckType(args[1])
			fmt.Println(result)
		case "pwd":
			result, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(result)
		case "cd":
			if _, err := os.Stat(args[1]); os.IsNotExist(err) {
				fmt.Println("cd: " + args[1] + ": No such file or directory")
				continue
			}
			os.Chdir(args[1])
		default:
			cmd := exec.Command(command, args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = io.Discard
			err = cmd.Run()
			if err != nil {
				fmt.Println(command[:len(command)-1] + ": command not found")
			}
		}

	}
}
