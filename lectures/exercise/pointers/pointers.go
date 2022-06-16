//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
package main

import "fmt"

//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)

type SecurityTag struct{
	item string
	state bool
}

//* Create functions to activate and deactivate security tags using pointers
func activate(tag *SecurityTag){
	tag.state = true
}
func deactivate(tag *SecurityTag){
	tag.state = false
}

//* Create a checkout() function which can deactivate all tags in a slice
func checkout(slice []SecurityTag){
	for i, _ := range slice{
		deactivate(&slice[i])
	}
}

func main() {
	//* Perform the following:
	//  - Create at least 4 items, all with active security tags
	//  - Store them in a slice or array
	productTags := []SecurityTag{{"belt", true}, {"jeans", true}, {"t-shirt", true}, {"bag", true}}

	fmt.Println(productTags)
	//  - Deactivate any one security tag in the array/slice
	deactivate(&productTags[1])
	fmt.Println(productTags)

	//  - Call the checkout() function to deactivate all tags
	//  - Print out the array/slice after each change
	checkout(productTags)
	fmt.Println(productTags)

	

}
