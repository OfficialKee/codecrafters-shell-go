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
			result := shell.CheckType(args[1])
			fmt.Println(result)
		case "pwd":
			result, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(result)

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
