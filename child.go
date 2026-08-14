package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"golang.org/x/sys/unix"
)

func child() {
	if len(os.Args) < 4 {
		log.Fatal("No Command Specified")
	}

	fmt.Println("Child Implementing")
	fmt.Println(os.Args)
	cmd := exec.Command(os.Args[3], os.Args[4:]...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	fmt.Println("Child PID:", os.Getpid())
	if err:=syscall.Sethostname([]byte("mini-docker")); err != nil {
		log.Fatal(err)
	}

	if err := unix.Mount("", "/", "", unix.MS_PRIVATE|unix.MS_REC, ""); err != nil {
		log.Fatal(err)
	}

	if err := syscall.Chroot("/root/mini-docker/images/alpine"); err != nil {
		log.Fatal(err)
	}
	os.Chdir("/")

	if err := unix.Mount("", "/proc", "proc",0,""); err != nil {
		log.Fatal(err)
	}


	if err:=cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
