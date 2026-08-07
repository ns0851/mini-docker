
package main

func doesContain(sl []string, name string) bool {
	    // iterate over the array and compare given string to each element
	for _, value := range sl {
	        if value == name {
	            return true
            }
       }
       return false
}
