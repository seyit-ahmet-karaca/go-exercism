package prime

func Factors(num int64) []int64 {
	primeFactors := []int64{}

	factor := int64(2)

	for num != 1 {
		if num%factor == 0 {
			primeFactors = append(primeFactors, factor)
			num /= factor
		} else {
			factor++
		}
	}
	return primeFactors
}
