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
    	} else {
		if os.Args[1] != "run" && os.Args[1]!="child" {
			fmt.Println("Command not implemented yet")
		} else if os.Args[1] == "child" {
			fmt.Println("Initiate child process")
		} else if len(os.Args) < 3 {
			fmt.Println("Image not specified")
		} else if len(os.Args) < 4 {
			fmt.Println("Command not specified")
		}
	}
}

func doesContain(sl []string, name string) bool {
	    // iterate over the array and compare given string to each element
	for _, value := range sl {
	        if value == name {
	            return true
            }
       }
       return false
}
