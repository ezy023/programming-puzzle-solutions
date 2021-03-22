package main

func lateRide(n int) int {
	sumDigits := func(d int) int {
		return (d / 10) + (d % 10)
	}
	hours := n / 60
	minutes := n % 60

	return sumDigits(hours) + sumDigits(minutes)
}
