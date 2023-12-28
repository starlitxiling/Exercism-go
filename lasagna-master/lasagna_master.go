package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, preparation_time int) int {
	if preparation_time == 0 {
		return 2 * len(layers)
	} else {
		return preparation_time * len(layers)
	}
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	grams := 0
	var liters float64
	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "noodles":
			grams += 50
		case "sauce":
			liters += 0.2
		default:
			break
		}
	}
	return grams, liters
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scale := float64(portions) / 2.0
	// modify := quantities    // just point ,not copy, also modify quantities slice
	modify := make([]float64, len(quantities))
	for i, qty := range quantities {
		modify[i] = qty * scale
	}
	return modify
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
