package main

import "fmt"

/*
1. Estimate the preparation time
Implement a function PreparationTime that accepts a slice
of layers as a []string and the average preparation time
per layer in minutes as an int.

The function should return the estimate for the total
preparation time based on the number of layers as an int.

Go has no default values for functions. If the average
preparation time is passed as 0 (the default initial value
for an int),then the default value of 2 should be used.
*/

func PreparationTime(a []string, b int) int {
	if b == 0 {
		b = 2
	}
	count := len(a)
	total := count * b
	return total

}

/*
2. Compute the amounts of noodles and sauce needed

For each noodle layer in your lasagna, you will need 50 grams of noodles.

For each sauce layer in your lasagna, you will need 0.2 liters of sauce.

Define the function Quantities that takes a slice of layers as parameter
as a []string. The function will then determine the quantity of noodles
and sauce needed to make your meal.

The result should be returned as two values of noodles as an int and
sauce as a float64.
*/

func Quantities(s []string) (noodles int, sauce float64) {

	i := 0
	j := 0
	for _, value := range s {

		if value == "noodles" {
			i++
			// default is 2
			//[]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
		} else if value == "sauce" {
			j++
			// default is 2
		}
	}

	noodles = i * 50
	// 100?
	calcj := float64(j)
	sauce = calcj * 0.2

	return
}

/*3. Add the secret ingredient*/

/*
Write a function AddSecretIngredient that accepts two slices of ingredients
of type []string as parameters. The first parameter is the list your friend sent you,
the second is the ingredient list of your own recipe.

The last element in your ingredient list is always "?". The function should replace it with the last item
from your friends list. Note: AddSecretIngredient does not return anything -
you should modify the list of your ingredients directly. The list with your
friend's ingredients should not be modified.

Also, since slice is passed into a function as pointers, changes to the two []string
arguments passed into AddSecretIngredient will be modified directly.
*/
func AddSecretIngredient() {

}

func main() {
	fmt.Println("You are on this council, but we don't grant you the rank of lasagna master")

	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	fmt.Printf("%v\n", PreparationTime(layers, 3))
	// => 18
	fmt.Printf("%v\n", PreparationTime(layers, 0))
	// => 12
	noodles, sauce := Quantities(layers)
	//Quantities([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"})
	//fmt.Printf("%v, %f\n", Quantities(layers))
	fmt.Printf("%v, %v\n", noodles, sauce)
	// => 100, 0.4

	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}

	AddSecretIngredient(friendsList, myList)
	// myList => []string{"noodles", "meat", "sauce", "mozzarella", "kampot pepper"}
}
