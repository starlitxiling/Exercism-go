package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	// panic("CalculateWorkingCarsPerHour not implemented")
	result := float64(productionRate) * successRate / 100
	return result

}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	// panic("CalculateWorkingCarsPerMinute not implemented")
	number := CalculateWorkingCarsPerHour(productionRate, successRate)
	success_produce := number / float64(60)
	return int(success_produce)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	// panic("CalculateCost not implemented")
	count := carsCount / 10
	a := carsCount % 10
	crash := count*95000 + a*10000
	return uint(crash)
}
