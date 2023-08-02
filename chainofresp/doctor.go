package chainofresp

import "log"

type doctor struct {
	next Department
}

func (r *doctor) execute(patient *patient) {
	if patient.checkupdone {
		log.Println("checkupdone already ")
		r.next.execute(patient)
		return
	}
	patient.checkupdone = true
	log.Printf("mr: %s just finsihed checkup", patient.name)
	r.next.execute(patient)
}

func NewDoctor() *doctor {
	return &doctor{}
}

func (c *doctor) SetNext(next Department) {
    c.next = next
}