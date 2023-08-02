package chainofresp

import "log"

type reception struct {
	next Department
}

func (r *reception) execute(patient *patient) {
	if patient.appointment {
		log.Println("Patient registration already done")
		r.next.execute(patient)
		return
	}
	patient.appointment = true
	log.Printf("Appointment booked for mr: %s", patient.name)
	r.next.execute(patient)
}

func NewReception() *reception {
	return &reception{}
}

func (c *reception) SetNext(next Department) {
	c.next = next
}
