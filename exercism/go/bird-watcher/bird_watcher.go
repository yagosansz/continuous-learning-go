package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	totalCount := 0
	for i := 0; i < len(birdsPerDay); i++ {
		totalCount += birdsPerDay[i]
	}
	return totalCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	totalCount := 0
	weekEnd := week * 7
	weekStart := weekEnd - 7
	targetWeek := birdsPerDay[weekStart:weekEnd]

	for i := 0; i < len(targetWeek); i++ {
		totalCount += targetWeek[i]
	}

	return totalCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i] += 1
		}
	}

	return birdsPerDay
}
