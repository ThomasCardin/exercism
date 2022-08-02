package cars

/*
	Note to self: Toujours convertir la variable
	vers la valeur la plus grosse. Exemple: int < float < double etc..
	Donc, si j'ai un parametre qui est un int et une autre qui est un float
	dans la meme fonction. Je dois convertir le int en float
*/

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate/60) * successRate / 100)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	return uint((carsCount % 10 * 10000) + (carsCount / 10 * 95000))
}
