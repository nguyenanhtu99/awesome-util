package main

type Remote interface {
	Device
	volumeUp()
	volumeDown()
}