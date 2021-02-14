package main

import "fmt"

// d := 99 *Implicit Declarations outside of function is invalid.


func main(){
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	
	fmt.Println(i, j, k, c, python, java)
}
