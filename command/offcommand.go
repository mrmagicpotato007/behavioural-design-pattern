package command

type offcommand struct {
	device Device
}

func (o *offcommand) execute() {
	o.device.off()
}

func GetOffCommand(device Device) *offcommand {
	return &offcommand{device: device}
}
