package chainofresp

import "log"

type cashier struct {
	next Department
}

func (r *cashier) execute(patient *patient) {
	if patient.paymentdone {
		log.Println("paymentdone ")
		r.next.execute(patient)
		return
	}
	patient.paymentdone = true
	log.Printf("mr: %s did paymentdone", patient.name)
}

func NewCashier() *cashier {
	return &cashier{}
}

func (c *cashier) SetNext(next Department) {
    c.next = next
}
