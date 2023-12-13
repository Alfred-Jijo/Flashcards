package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ClearTerminal() {
	switch runtime.GOOS {
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}

func ErrHandle(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func SuccessHandle(success bool) {
	if !success {
		fmt.Println("Didnt succeed")
	}
	fmt.Println("Successful")
}
