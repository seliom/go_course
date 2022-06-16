//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
package main

import "fmt"

//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
type Product struct{
	name string
	price float64
}

func printListInfo(list [4]Product) {
	totalCost := 0.0
	lastItem := list[0].name
	numItems := 0
	for i := 0; i < len(list); i++{
		item := list[i]
		itemName := item.name
		totalCost += item.price
		if itemName != ""{
			lastItem = itemName
			numItems ++
		} 
	}
	fmt.Println("The last item on the list:", lastItem)
	fmt.Println("The total number of items:", numItems)
	fmt.Println("The total cost of the items, in dollars:", (totalCost))	
		
}

func main() {
	//* Insert 3 products into the array


	shoppingList := [4]Product{
		{"Milk", 3.24},
		{"Chicken", 13.99},
		{"Broccoli", 5.39},
	}
	//* Print to the terminal:
	//  - The last item on the list
	//  - The total number of items
	//  - The total cost of the items
	//* Add a fourth product to the list and print out the
	//  information again
	printListInfo(shoppingList)
	
	// totalCost := 0
	// for i := 0; i < len(shoppingList); i++{
	// 	item := shoppingList[i]
	// 	itemName := item.name
	// 	totalCost += item.price
	// 	if itemName != ""{
	// 		lastItem := itemName
	// 	} else if itemName == "" || i == len(shoppingList) - 1{
	// 		fmt.Println("The last item on the list", itemName)
	// 		fmt.Println("The total number of items", len(shoppingList))
	// 		fmt.Println("The total cost of the items", totalCost)
	// 	}
	// }
	shoppingList[3] = Product{"Bread", 1.99}
	printListInfo(shoppingList)
	
}
