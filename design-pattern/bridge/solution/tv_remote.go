package main

type tvRemote struct {
	device Device
}

func (r *tvRemote) volumeUp() {
	r.device.setVolume(r.device.getVolume() + 10)
}

func (r *tvRemote) volumeDown() {
	r.device.setVolume(r.device.getVolume() - 10)
}