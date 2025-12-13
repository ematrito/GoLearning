package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	
    simplifiedSuccessRate := successRate / 100
    
    return float64(productionRate) * simplifiedSuccessRate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {

    var successfullyCarsPerHour = CalculateWorkingCarsPerHour(productionRate, successRate)

    return int(successfullyCarsPerHour) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	if (carsCount < 0) {
        return uint(0)
    }
    uintCarsCount := uint(carsCount)
   
    const costOfOne uint = 10000
    const costOfTenBatch uint = 95000

    numberOfSingles := uintCarsCount % 10
    numberOfBatches := uintCarsCount / 10

   	costOfBatches := numberOfBatches * costOfTenBatch
    costOfSingles := numberOfSingles * costOfOne
  
    return costOfBatches + costOfSingles
}
