package main 

import (
    "fmt"
    "os"
)


func main() {
    	acceptedCommands := []string{"run","build","compose","child"}
    	if len(os.Args) < 2 {
		fmt.Println("Usage: minidocker run <image> <command>")
		return
    	}

	if !doesContain(acceptedCommands, os.Args[1]) {
		fmt.Println("Unknown command: ", os.Args[1])
		return
	}


	switch (os.Args[1]) {
		case "run":
			run()

		case "child":
			child()
	}

}

