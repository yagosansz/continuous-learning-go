package lasagna

func PreparationTime(layers []string, avgPrepTimePerLayer int) int {
	if avgPrepTimePerLayer == 0 {
		return len(layers) * 2
	}
	return len(layers) * avgPrepTimePerLayer
}

func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64
	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledRecipe := make([]float64, len(quantities))
	for i := 0; i < len(scaledRecipe); i++ {
		scaledRecipe[i] = quantities[i] * float64(portions) / 2
	}
	return scaledRecipe
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
