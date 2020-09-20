// Introduction to variables in Go
package main

import (
	"fmt"
)

var i int = 69 // Can declare variables outside of main scope

// Declaring a group of variables ...
var actorName string = "Nicolas Cage"
var movie string = "National Treasure"
var year int = 2004
// ... can become ...
var (
	actorName2 string = "Matt Damon"
	movie2 string = "Martian"
	year2 int = 2015
)
// This technique can be used to group variables in logical blocks

func main() {
	var i int // Longhand: Declare variable and type but don't set i
	i = 27 // Set variable i
	// Note that this was declared above but the variable with the innermost scope takes precendence
	// However defining i again inside here is not allowed

	// var j int = 40
	// var j float32 = 40 // Shorterhand: Declare and set j

	// k := 42 // Shortesthand: Declare and set shorthand k

	fmt.Printf("%v, %T", i, i)
	// fmt.Println(i)

}