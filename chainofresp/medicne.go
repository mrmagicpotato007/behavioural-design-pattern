package chainofresp

import "log"

type medicine struct {
	next Department
}

func (r *medicine) execute(patient *patient) {
	if patient.medicinedone {
		log.Println("Patient  already recieved medicine")
		r.next.execute(patient)
		return
	}
	patient.medicinedone = true
	log.Printf("mr: %s took medicie", patient.name)
	r.next.execute(patient)
}

func NewMedicine() *medicine {
	return &medicine{}
}

func (c *medicine) SetNext(next Department) {
    c.next = next
}