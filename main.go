package main

import (
	"behavioural/command"
	"behavioural/mediator"
	"behavioural/chainofresp"
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

	//command pattern
	tv := command.NewDevice("tv")
	oncommand := command.GetOnCommand(tv)
	button := command.NewButton(oncommand)
	command.Executor(button)
	offcommand := command.GetOffCommand(tv)
	button = command.NewButton(offcommand)
	command.Executor(button)

	//chain of responsibilty
	cashier := chainofresp.NewCashier()
	medicine := chainofresp.NewReception()
	medicine.SetNext(cashier)

	doctor := chainofresp.NewDoctor()
	doctor.SetNext(medicine)

	reception := chainofresp.NewReception()
	reception.SetNext(doctor)
	patient := chainofresp.NewPatient("harsha")
	chainofresp.Executor(reception,patient)

}
