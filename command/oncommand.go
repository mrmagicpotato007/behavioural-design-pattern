package command

type oncommand struct {
	device Device
}

func (o *oncommand) execute() {
	o.device.on()
}

func GetOnCommand(device Device) *oncommand {
	return &oncommand{device: device}
}
