package cars

const productionRate = 221.0
// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
    if speed == 0 {
    	return 0.0
    } else if speed <= 4 {
    	return 1.0
    } else if speed <= 8 {
        return 0.9
    } else if speed <= 10 {
    	return 0.77
	}
	return 0.01
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return productionRate * float64(speed) * SuccessRate(speed)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
    return int(CalculateProductionRatePerHour(speed) / 60.00)
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
    rate := CalculateProductionRatePerHour(speed)
    if rate < limit {
        return float64(rate)
    }
    return float64(limit) 
}
