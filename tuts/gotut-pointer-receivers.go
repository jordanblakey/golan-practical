package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmhMultiple float64 = 1.60934

type car2 struct {
	gasPedal      uint16
	brakePedal    uint16
	steeringWheel int16
	topSpeedKmh   float64
}

func (c car2) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax)
}

func (c car2) mph() float64 { // using a value receiver, creates a instance var copies
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax / kmhMultiple)
}

func (c *car2) newTopSpeed(newspeed float64) { // Using a pointer receiver. Operates on the actual struct, car2
	c.topSpeedKmh = newspeed
}

func newerTopSpeed(c car2, speed float64) car2 {
	c.topSpeedKmh = speed
	return c
}

func main() {
	aCar := car2{65000, 0, 12561, 225.0}

	fmt.Println(aCar.gasPedal)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
	aCar.newTopSpeed(500)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
	aCar = newerTopSpeed(aCar, 800) // replace the entire car2 struct stored in car2 with new top speed
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())

}
