package main

import "fmt"

/*
1. Define the expected oven time in minutes

Define the OvenTime constant with how many minutes the lasagna should be in the oven.
According to the cooking book, the expected oven time in minutes is 40
*/
const OvenTime int = 40

func RemainingOvenTime(actual int) int {
	/*
		2. Calculate the remaining oven time in minutes

		Define the RemainingOvenTime() function that takes the actual minutes the lasagna
		has been in the oven as a parameter and returns how many minutes the lasagna still
		has to remain in the oven, based on the expected oven time in minutes from the previous task.
	*/

	return OvenTime - actual

}

/*
3. Calculate the preparation time in minutes

Define the PreparationTime function that takes the number of layers you added to the
lasagna as a parameter and returns how many minutes you spent preparing the lasagna,
assuming each layer takes you 2 minutes to prepare.
*/

func PreparationTime(numberOfLayers int) int {

	return numberOfLayers * 2

}

/*
4. Calculate the elapsed working time in minutes

Define the ElapsedTime function that takes two parameters: the first parameter is the
number of layers you added to the lasagna, and the second parameter is the number of
minutes the lasagna has been in the oven. The function should return how many minutes
in total you've worked on cooking the lasagna, which is the sum of the preparation
time in minutes, and the time in minutes the lasagna has spent in the oven at the moment.
*/

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	timeLayers := numberOfLayers * 2 // Preparation
	actualMinutesInOven = OvenTime   //Oven

	return timeLayers + actualMinutesInOven

}

func main() {

	/*2*/
	hornoTime := 30
	fmt.Println("Te quedan unos porquitos minutos:", RemainingOvenTime(hornoTime))

	/*3*/
	layers := 2
	fmt.Println("It's preparation time:", (PreparationTime(layers)))

	/*4*/
	fmt.Println("Tu lasagna lleva haciendose: ", ElapsedTime(layers, hornoTime))
}
