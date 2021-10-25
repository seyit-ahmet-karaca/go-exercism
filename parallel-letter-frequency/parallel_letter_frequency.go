package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// ConcurrentFrequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func ConcurrentFrequency(strings []string) FreqMap {
	m := FreqMap{}
	freqChan := make(chan FreqMap)

	for _, s := range strings {
		go frequency(s, freqChan)
	}

	for i := 0; i < len(strings); i++ {
		res := <-freqChan
		m.merge(res)
	}

	return m
}

func (m FreqMap) merge(fm FreqMap) {
	for key, val := range fm {
		_, ok := m[key]
		if ok {
			m[key] = m[key] + val
		} else {
			m[key] = val
		}
	}
}

func frequency(s string, fc chan<- FreqMap) {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	fc <- m
}

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}
