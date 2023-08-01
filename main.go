package main

import (
	"behavioural/mediator"
)

func main() {

	
	//mediator pattern
	stationManager := mediator.NewStationManager()
	passengerTrain := &mediator.PassengerTrain{Mediator: stationManager}
	freightTrain := &mediator.FreightTrain{
		Mediator: stationManager,
	}

	passengerTrain.Arrive()
	freightTrain.Arrive()
	passengerTrain.Depart()

}
