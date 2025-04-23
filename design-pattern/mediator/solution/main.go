package main

import "sync"

func main() {
	var wG sync.WaitGroup
	stationManager := newStationManger(&wG)

	for i := range 5 {
		wG.Add(1)
		passengerTrain := &PassengerTrain{
			id:       i,
			mediator: stationManager,
		}

		go func() {
			passengerTrain.registerArrive()
		}()
	}

	for i := range 5 {
		wG.Add(1)
		freightTrain := &FreightTrain{
			id:       i,
			mediator: stationManager,
		}

		go func() {
			freightTrain.registerArrive()
		}()
	}

	wG.Wait()
}
