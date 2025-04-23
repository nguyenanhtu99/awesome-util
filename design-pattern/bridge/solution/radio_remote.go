package main

type radioRemote struct {
	device Device
}

func (r *radioRemote) volumeUp() {
	r.device.setVolume(r.device.getVolume() + 1)
}

func (r *radioRemote) volumeDown() {
	r.device.setVolume(r.device.getVolume() - 1)
}
