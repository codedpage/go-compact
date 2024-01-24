package main

import "fmt"

func dayOfWeek(S string, K int) string {
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	// Find the index of the initial day S.
	initialDayIndex := -1
	for i, day := range daysOfWeek {
		if S == day {
			initialDayIndex = i
			break
		}
	}

	if initialDayIndex == -1 {
		return "Invalid Day"
	}

	// Calculate the new day's index after adding K days.
	newDayIndex := (initialDayIndex + K) % 7

	// Return the day of the week corresponding to the new index.
	return daysOfWeek[newDayIndex]
}

func main() {
	S1 := "Wed"
	K1 := 2
	result1 := dayOfWeek(S1, K1)
	fmt.Println(result1) // Output: "Fri"

	S2 := "Sat"
	K2 := 23
	result2 := dayOfWeek(S2, K2)
	fmt.Println(result2) // Output: "Mon"
}
