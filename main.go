package main

import (
	"behavioural/chainofresp"
	"behavioural/command"
	"behavioural/mediator"
	"behavioural/strategy"
	"behavioural/template"
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
	chainofresp.Executor(reception, patient)

	//startegy
	lru := strategy.NewLRU()
	cache := strategy.NewCache()
	cache.SetEviction(lru)
	cache.Add("key1", "val1")
	cache.Add("key2", "val2")
	cache.Add("key3", "val3")

	//template pattern]
	smsOTP := &template.Sms{}
	o := template.Otp{IOtp: smsOTP}
	o.GenAndSendOTP(4)

	emailOTP := &template.Email{}
	o = template.Otp{
		IOtp: emailOTP,
	}
	o.GenAndSendOTP(4)

}
