package main

import "fmt"

type FreightTrain struct {
	id       int
	mediator Mediator
}

func (t *FreightTrain) registerArrive() {
	t.mediator.registerArrive(t)
}

func (t *FreightTrain) depart() {
	fmt.Printf("FreightTrain %d: Leaving\n", t.id)
	t.mediator.notifyAboutDeparture()
}

func (t *FreightTrain) permitArrival() {
	fmt.Printf("PassengerTrain %d: Arriving\n", t.id)
	go t.depart()
}
