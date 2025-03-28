package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, preparationTime int) int {
	if (preparationTime == 0) {
		return len(layers) * 2
	}

	return len(layers) * preparationTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0;
	sauce = 0;

	for i := 0; i < len(layers); i++ {
		if (layers[i] == "noodles") {
			noodles += 50
		} else if (layers[i] == "sauce") {
			sauce += 0.2
		}
	}

	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, scale int) []float64 {
	newQuantities := []float64{}
	newQuantities = append(newQuantities, quantities...)

	for i := 0; i < len(newQuantities); i++ {
		newQuantities[i] = quantities[i] * (float64(scale) / 2)
	}

	return newQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
