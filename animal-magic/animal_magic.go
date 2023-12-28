package chance

import (
	"math/rand"
)

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	// panic("Please implement the RollADie function")
	return rand.Intn(20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	// panic("Please implement the GenerateWandEnergy function")
	f := rand.Float64() * 12
	return f
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	// panic("Please implement the ShuffleAnimals function")
	animous := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(len(animous), func(i, j int) {
		animous[i], animous[j] = animous[j], animous[i]
	})
	return animous
}
