package mediator

import "log"

type FreightTrain struct {
	Mediator Mediator
}

func (p *FreightTrain) Arrive() {
	if p.Mediator.CanArrive(p) {
		log.Println("Freight train Arrived")
	} else {
		log.Println("Freight train blocked , waiting")
	}
}

func (p *FreightTrain) Depart() {
	log.Println("Freight train Departed")
	p.Mediator.Notifyaboutdepature()
}

func (p *FreightTrain) Permitarrival() {
	log.Println("Permit recived we can start now")
	p.Arrive()
}
