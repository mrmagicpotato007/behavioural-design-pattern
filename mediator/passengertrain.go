package mediator

import "log"

type PassengerTrain struct {
	Mediator Mediator
}

func (p *PassengerTrain) Arrive() {
	if p.Mediator.CanArrive(p) {
		log.Println("Passenger train Arrived")
	} else {
		log.Println("Passenger train blocked , waiting")
	}
}

func (p *PassengerTrain) Depart() {
	log.Println("Passenger train Departed")
	p.Mediator.Notifyaboutdepature()
}

func (p *PassengerTrain) Permitarrival() {
	log.Println("Permit recived we can start now")
	p.Arrive()
}
