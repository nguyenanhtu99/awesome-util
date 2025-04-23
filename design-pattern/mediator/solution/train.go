package main

type Train interface {
	registerArrive()
	depart()
	permitArrival()
}
