package main

type Mediator interface {
	registerArrive(Train)
	notifyAboutDeparture()
}
