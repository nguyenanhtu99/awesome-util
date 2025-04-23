package main

type Car struct {
	feePerHour  int
	speed       int
	gasFeePerKm int
}

func (c *Car) calculateFee(distant int) int {
	return c.calculateTime(distant)*c.feePerHour + c.calculateGasFee(distant)
}

func (c *Car) calculateTime(distant int) int {
	return distant / c.speed
}

func (c *Car) calculateGasFee(distant int) int {
	return distant * c.gasFeePerKm
}
