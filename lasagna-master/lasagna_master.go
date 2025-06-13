package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prptime int) int {
	if prptime == 0 {
		prptime = 2
	}
	return len(layers) * prptime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodlePerLayer := 50
	saucePerLayer := 0.2
	noodleLayer := 0
	sauceLayer := 0

	for i := 0; i < len(layers); i++ {
		if layers[i] == "sauce" {
			sauceLayer = sauceLayer + 1
		} else if layers[i] == "noodles" {
			noodleLayer = noodleLayer + 1
		}
	}

	totalNoodle := noodleLayer * noodlePerLayer
	totalSauce := float64(sauceLayer) * saucePerLayer

	return totalNoodle, totalSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, num int) []float64 {
	var newQuant []float64
	for i := 0; i < len(quantities); i++ {
		newQuant = append(newQuant, quantities[i]*(float64(num)/2))
	}
	return newQuant
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
