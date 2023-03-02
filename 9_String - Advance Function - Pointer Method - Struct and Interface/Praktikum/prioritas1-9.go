package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64
}

func (c *Car) EstimateDistance() float64 {
	fuelEfficiency := 1.5 // 1.5 L/mill
	distance := fuelEfficiency * c.FuelIn
	return distance
}

func main() {
	myCar := Car{
		Type:   "stragazer",
		FuelIn: 27.0,
	}

	estimatedDistance := myCar.EstimateDistance()

	fmt.Printf("Mobil ku %s menempuh jarak %.2f mill dengan bahan bakar yang tersedia.\n", myCar.Type, estimatedDistance)
}
