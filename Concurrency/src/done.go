package main

import (
	"fmt"
)

// START OMIT
func functionIWantToRunConcurrently() {
	//Do Work here
}

func main() {
	go functionIWantToRunConcurrently()
	fmt.Println("Working")
}

// END OMIT
