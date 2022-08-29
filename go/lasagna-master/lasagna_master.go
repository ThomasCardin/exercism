package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averagePreparationTime int) int {
	if averagePreparationTime != 0 {
		return len(layers) * averagePreparationTime
	}

	return len(layers) * 2
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	sauce := float64(0)
	noodlesLayers := 0
	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "noodles":
			noodlesLayers += 50
			break
		case "sauce":
			sauce += 0.2
			break
		}
	}

	return noodlesLayers, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
	myList = myList[:len((myList))-1]
	myList = append(myList, friendList[len(friendList)-1])
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(portions []float64, amounts int) []float64 {
	newPortions := []float64{}
	for i := 0; i < len(portions); i++ {
		newPortions = append(newPortions, portions[i]*(float64(amounts)/2))
	}

	return newPortions
}
