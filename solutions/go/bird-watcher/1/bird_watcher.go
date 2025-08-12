package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	birds := 0
    for i := range birdsPerDay {
        birds = birds + birdsPerDay[i]
    }
    return birds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	firstDay := 7 * (week - 1)
    lastDay := firstDay + 6
    
    birds := 0
    for i := firstDay; i <= lastDay; i++ {
        birds = birds + birdsPerDay[i]
    }
    return birds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range birdsPerDay {
        if i % 2 == 1 {
            continue
        }
        birdsPerDay[i] = birdsPerDay[i] + 1
    }
    return birdsPerDay
}
