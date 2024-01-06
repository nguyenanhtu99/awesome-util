package main

type Bicycle struct {
	feePerHour int
	speed      int
}

func (b *Bicycle) calculateFee(distant int) int {
	return b.calculateTime(distant) * b.feePerHour
}

func (b *Bicycle) calculateTime(distant int) int {
	return distant / b.speed
}
