package main

import (
	"fmt"
	"project/lasagna"
)

func main() {
	// Call the functions from the `lasagna` package
	hornoTime := 30
	fmt.Println("Te quedan unos poquitos minutos:", lasagna.RemainingOvenTime(hornoTime))

	layers := 2
	fmt.Println("It's preparation time:", lasagna.PreparationTime(layers))

	fmt.Println("Tu lasagna lleva haciéndose: ", lasagna.ElapsedTime(layers, hornoTime))
}
