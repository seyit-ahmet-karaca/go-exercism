package sieve

func Sieve(limit int) []int {
	var numbers []int
	var marks []bool
	var primes []int

	if limit < 2 {
		return []int{}
	}
	for i:=2 ; i <= limit ; i++ {
		numbers = append(numbers, i)
		marks = append(marks,true)
	}

	for i,v := range numbers {
		if marks[i] {
			for j:= i+v; j < len(marks); j+=v {
				marks[j] = false
			}
		}
	}

	for i,_ := range marks{
		if marks[i] {
			primes = append(primes, numbers[i])
		}
	}

	return primes
}