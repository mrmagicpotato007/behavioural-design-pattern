package chainofresp

type patient struct {
	name         string
	appointment  bool
	checkupdone  bool
	medicinedone bool
	paymentdone  bool
}


func NewPatient(name string) *patient{
   return &patient{name: name}
}

