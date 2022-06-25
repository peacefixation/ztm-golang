//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

func printAssemblyLine(line []Part) {
	fmt.Println()
	fmt.Println("---", "Assembly Line", "---")

	for i := 0; i < len(line); i++ {
		part := line[i]
		fmt.Println(part)
	}
}

type Part string

func main() {
	line := []Part{"Chassis", "Body", "Wheels"}

	printAssemblyLine(line)

	line = append(line, "Mirrors")
	line = append(line, "Seats")

	printAssemblyLine(line)

	line = line[3:]

	printAssemblyLine(line)
}
