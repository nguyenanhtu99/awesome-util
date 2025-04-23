package main

import "fmt"

type Transport struct {
	distant int
	vehicle Vehicle
}

func (t *Transport) ShowFee() {
	fmt.Printf("Fee: %d$\n", t.vehicle.calculateFee(t.distant))
}
