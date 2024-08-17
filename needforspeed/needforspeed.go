package main

import "fmt"

/*
1. Create a remote controlled car

Define a Car struct with the following int type fields:

	battery
	batteryDrain
	speed
	distance
*/
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

/*
Allow creating a remote controlled car by defining a function NewCar that
takes the speed of the car in meters, and the battery drain percentage as
its two parameters (both of type int) and returns a Car instance:


// => Car{speed: 5, batteryDrain: 2, battery:100, distance: 0}
*/

func NewCar(veloc, batteryd int) Car {
	bolido := Car{
		battery:      100,
		batteryDrain: batteryd,
		speed:        veloc,
		distance:     0,
	}

	return bolido
}

/*
2. Create a race track

Define another struct type called Track with the field distance of type integer.
Allow creating a race track by defining a function NewTrack that takes the track's
distance in meters as its sole parameter (which is of type int):
*/
type Track struct {
	distance int
}

func NewTrack(long int) Track {
	pista := Track{
		distance: long,
	}

	return pista
}

/*
3. Drive the car

Implement the Drive function that updates the number of meters driven based on the car's speed, and reduces the battery according to the battery drainage. If there is not enough battery to drive one more time the car will not move:

speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)
car = Drive(car)
// => Car{speed: 5, batteryDrain: 2, battery: 98, distance: 5}
*/

/*
4. Check if a remote controlled car can finish a race

To finish a race, a car has to be able to drive the race's distance. This means not draining its battery before having crossed the finish line. Implement the CanFinish function that takes a Car and a Track instance as its parameter and returns true if the car can finish the race; otherwise, return false.

Assume that you are currently at the starting line of the race and start the engine of the car for the race. Take into account that the car's battery might not necessarily be fully charged when starting the race:

speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)

distance := 100
track := NewTrack(distance)

CanFinish(car, track)
// => true
*/
func main() {
	//fmt.Println("Hello there, welcome to nid for spid")

	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)

	fmt.Println("Características de tu bolido: ", car)

	distance := 800
	track := NewTrack(distance)
	fmt.Println("Tu pista mide (en metros)", track.distance)
}
