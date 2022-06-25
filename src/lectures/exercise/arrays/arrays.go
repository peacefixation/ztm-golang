//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price uint64
}

func printStats(list [4]Product) {
	count := 0
	totalPrice := uint64(0)

	for i := 0; i < len(list); i++ {
		item := list[i]
		totalPrice += item.price

		if item.name != "" {
			count++
		}
	}

	fmt.Println("Last item", list[count-1])
	fmt.Println("Count", count)
	fmt.Println("Total", totalPrice)

}

func main() {
	shoppingList := [4]Product{
		{name: "Bottle Water", price: 2},
		{name: "Sandwich", price: 7},
		{name: "Muffin", price: 4},
	}

	printStats(shoppingList)

	shoppingList[3] = Product{name: "Pepsi", price: 3}

	printStats(shoppingList)
}
