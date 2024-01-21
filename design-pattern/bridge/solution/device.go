package main

type Device interface {
	setVolume(volume int)
	getVolume() int
}
