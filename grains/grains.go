package grains

import (
	"errors"
	"math"
	"sync"
)

const NUMBER_OF_SQUARE = 64
const ZERO = 0
const BASE_TWO = 2.0

func Square(num int) (uint64 ,error) {
	if num > NUMBER_OF_SQUARE || num <= ZERO {
		return ZERO, errors.New("error")
	}
	return uint64(math.Pow(BASE_TWO,float64(num-1))), nil
}

func Total() uint64 {
	mtx := sync.Mutex{}
	var totalGrain uint64
	var wg sync.WaitGroup

	for i := 0 ; i < NUMBER_OF_SQUARE ; i++ {
		wg.Add(1)
		go func(square float64) {
			defer wg.Done()
			gratin := uint64(math.Pow(BASE_TWO,square))
			mtx.Lock()
			totalGrain = totalGrain + gratin
			mtx.Unlock()
		}(float64(i))
	}

	wg.Wait()
	return totalGrain
}