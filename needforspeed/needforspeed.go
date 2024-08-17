package main

import "fmt"

/*
1. Create a remote controlled car
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

Implement the Drive function that updates the number of meters driven based on the car's speed,
and reduces the battery according to the battery drainage. If there is not enough battery to
drive one more time the car will not move:

// => Car{speed: 5, batteryDrain: 2, battery: 98, distance: 5}
*/

func Drive(c Car, pedal int) Car {
	run := c.speed + pedal
	if c.battery < 0 {
		return Car{}
	} else {
		return Car{
			battery:      c.battery - c.batteryDrain,
			batteryDrain: c.batteryDrain,
			speed:        c.speed + pedal,
			distance:     c.distance + run,
		}
	}
}

/*
4. Check if a remote controlled car can finish a race

To finish a race, a car has to be able to drive the race's distance. This means not draining its battery
 before having crossed the finish line. Implement the CanFinish function that takes a Car and a Track
 instance as its parameter and returns true if the car can finish the race; otherwise, return false.

Assume that you are currently at the starting line of the race and start the engine of the car for the race.
Take into account that the car's battery might not necessarily be fully charged when starting the race:
*/

func CanFinish(c Car, track, progress int) bool {
	if c.battery == 100 && track < progress {
		return true
	} else {
		return false
	}
}

func main() {
	//fmt.Println("Hello there, welcome to nid for spid")

	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)

	fmt.Println("Características de tu bolido: ", car)

	distance := 100
	track := NewTrack(distance)
	fmt.Println("Tu pista mide (en metros)", track.distance)

	speed = 150
	batteryDrain = 5
	//car2 := NewCar(speed, batteryDrain)

	car = Drive(car, speed)
	fmt.Println("Características de tu bolido - actual: ", car)

	fmt.Println(CanFinish(car, distance, car.distance))
}
