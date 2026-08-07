package main
import (
	"fmt"
	"os"
	"os/exec"
	"log"
	"syscall"
)

func run() {
	if len(os.Args) < 3 {
		fmt.Println("Image not specified")
	}

	// creating a slice that includes "child" along with all the arguments
	args := append([]string{"child"}, os.Args[2:]...)	
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr


	// Adding PID namespace so that the child process is mapped with PID of 1
	cmd.SysProcAttr = &syscall.SysProcAttr {
		Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS,	
	}

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
