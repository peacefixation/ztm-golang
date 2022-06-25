//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

type Item struct {
	name string
	tag  SecurityTag
}

func activate(tag *SecurityTag) {
	*tag = Active
}

func deactivate(tag *SecurityTag) {
	*tag = Inactive
}

func checkout(items []Item) {
	for i, _ := range items {
		deactivate(&items[i].tag)
	}
}

func print(items []Item) {
	fmt.Println("---", "Items", "---")
	for _, item := range items {
		fmt.Println(item.name, item.tag)
	}
}

func main() {
	items := []Item{{
		name: "book",
		tag:  Active,
	}, {
		name: "cup",
		tag:  Active,
	}, {
		name: "pen",
		tag:  Active,
	}, {
		name: "ruler",
		tag:  Active,
	},
	}
	print(items)

	deactivate(&items[1].tag)
	print(items)

	checkout(items)
	print(items)
}
