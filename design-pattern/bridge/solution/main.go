package main

func main() {
	radio := radio{volume: 10}
	radioRemote := radioRemote{
		device: &radio,
	}
	radioRemote.volumeUp()
	radioRemote.volumeDown()

	tv := tv{volume: 50}
	tvRemote := tvRemote{
		device: &tv,
	}
	tvRemote.volumeUp()
	tvRemote.volumeDown()
}