package command

type Device interface {
    on()
    off()
}

func NewDevice(device string) Device{
 if device =="tv"{
	return &tv{}
 }
 return nil
}
