package main

import (
	"fmt"
)

type car struct {
	gasPedal      uint16 // min 0 max 65535
	brakePedal    uint16
	steeringWheel int16 // -32k to +32k
	topSpeedKmh   float64
}

func main() {
	// aCar := car{ // Instansiate struct, long form
	// 	gasPedal:      22341,
	// 	brakePedal:    0,
	// 	steeringWheel: 1256,
	// 	topSpeedKmh:   22.3,
	// }

	aCar := car{22341, 0, 1256, 22.3} // instantiate struct, short form

	fmt.Println(aCar.gasPedal)
}
