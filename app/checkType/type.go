package shell

import (
	"fmt"
	"os/exec"
)

func CheckType(shellType string) string {
	switch shellType {
	case "exit":
		return fmt.Sprintf("exit is a shell builtin")
	case "echo":
		return fmt.Sprintf("echo is a shell builtin")
	case "type":
		return fmt.Sprintf("type is a shell builtin")
	case "pwd":
		return fmt.Sprintf("pwd is a shell builtin")
	default:
		result, err := exec.LookPath(shellType)
		if err != nil {
			return fmt.Sprintf(shellType + ": not found")
		} else {
			return fmt.Sprintf(shellType + " is " + result)
		}
	}
}
