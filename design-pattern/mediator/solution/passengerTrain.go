package main

import "fmt"

type PassengerTrain struct {
	id       int
	mediator Mediator
}

func (t *PassengerTrain) registerArrive() {
	t.mediator.registerArrive(t)
}

func (t *PassengerTrain) depart() {
	fmt.Printf("PassengerTrain %d: Leaving\n", t.id)
	t.mediator.notifyAboutDeparture()
}

func (t *PassengerTrain) permitArrival() {
	fmt.Printf("PassengerTrain %d: Arriving\n", t.id)
	go t.depart()
}
