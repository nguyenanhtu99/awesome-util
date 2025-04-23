package main

func main() {
	transport := Transport{
		distant: 100,
		vehicle: &Car{
			feePerHour:  10,
			speed:       100,
			gasFeePerKm: 1,
		},
	}

	transport.ShowFee()
}
