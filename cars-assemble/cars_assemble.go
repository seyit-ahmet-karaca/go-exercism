package cars

const producedEachHour = 221

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	if speed == 0 {
		return 0.0
	} else if speed >= 1 && speed <= 4 {
		return 1.0
	} else if speed >= 5 && speed <= 8 {
		return 0.9
	} else if speed >= 9 && speed <= 10 {
		return 0.77
	}
	panic(speed)
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return float64(producedEachHour) * float64(speed) * SuccessRate(speed)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	if speed != 0 {
		return int(CalculateProductionRatePerHour(speed)) / 60
	}
	return 0
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	produceCarPerHour := CalculateProductionRatePerHour(speed) * float64(speed)
	if limit < produceCarPerHour {
		return limit
	} else if speed == 0 {
		return 0
	}
	return producedEachHour
}
