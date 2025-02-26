package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	productionRatePerMinute := float64(productionRate / 60)
	carsPerMinutes := productionRatePerMinute * (successRate / 100)
	return int(carsPerMinutes)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupedCars := carsCount / 10
	individualCars := carsCount % 10
	cost := (groupedCars * 95000) + (individualCars * 10000)
	return uint(cost)
}
