package lasagna

func PreparationTime(layers []string, averagePerLayer int) int {
	if averagePerLayer == 0 {
		averagePerLayer = 2
	}

	return len(layers) * averagePerLayer
}

func Quantities(layers []string) (int, float64) {
	flour := 0
	sauce := 0.0

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			flour += 50
		} else if layers[i] == "sauce" {
			sauce += 0.2
		}
	}

	return flour, sauce
}

func AddSecretIngredient(friendsList []string, myList []string) []string {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]

	return myList
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaled := append([]float64{}, quantities...)

	for i := 0; i < len(scaled); i++ {
		scaled[i] *= float64(portions) / 2.0
	}

	return scaled
}
