package main

import "fmt"

func calculateOfficeTripsDistance(distance int) int {
	totalDistance := distance * 2 // Round trip from home to office and back
	weeklyDistance := totalDistance * 5 // Distance traveled in a week with 5 working days
	return weeklyDistance
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var distance int
		fmt.Scan(&distance)
		weeklyDistance := calculateOfficeTripsDistance(distance)
		fmt.Println(weeklyDistance)
	}
}
