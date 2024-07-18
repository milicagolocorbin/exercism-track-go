package lasagna

import "slices"

func PreparationTime(layers []string, avg int) int {
	if avg == 0 {
		return 2 * len(layers)
	}
	return avg * len(layers)
}

func Quantities(layers []string) (int, float64) {
	var (
		noodles int
		sauce   float64
	)
	for _, l := range layers {
		if l == "noodles" {
			noodles += 50
		}
		if l == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendIngredients, myIngredients []string) {
	_ = slices.Replace(myIngredients, len(myIngredients)-1, len(myIngredients), friendIngredients[len(friendIngredients)-1])
}

func ScaleRecipe(quantities []float64, numPortions int) []float64 {
	var numScaled float64 = float64(numPortions) / 2
	quantitiesScaled := make([]float64, 0)

	for _, q := range quantities {
		quantitiesScaled = append(quantitiesScaled, q*numScaled)
	}
	return quantitiesScaled
}
