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

	if len(os.Args) < 3 {
		fmt.Println("Image not specified")
	} else if len(os.Args) < 4 {
		fmt.Println("Command not specified")
	}

	switch (os.Args[1]) {
		case "run":
			fmt.Println("implement run here")
		case "child":
			fmt.Println("implement child here")
		case default:
			fmt.Println("not implemented yet")
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
