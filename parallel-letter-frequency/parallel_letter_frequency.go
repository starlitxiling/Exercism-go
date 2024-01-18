package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	// panic("Implement the ConcurrentFrequency function")
	var waitGroup sync.WaitGroup
	freqmap := FreqMap{}
	var mutex sync.Mutex

	for i := 0; i < len(texts); i++ {
		waitGroup.Add(1)
		go func(text string) {
			defer waitGroup.Done()

			freq := Frequency(text)

			mutex.Lock()

			for r, f := range freq {
				freqmap[r] += f
			}

			mutex.Unlock()

		}(texts[i])
	}
	waitGroup.Wait()

	return freqmap
}
