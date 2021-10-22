package lasagna

const gramsOfNoodle = 50
const litersOfSauce = 0.2

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {

	for _, layer := range layers {
		if layer == "sauce" {
			sauce++
		} else if layer == "noodles" {
			noodles++
		}
	}
	return noodles * gramsOfNoodle, sauce * litersOfSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) []string {
	return append(myList, friendsList[len(friendsList)-1])
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int ) []float64 {
	x := float64(portions) / 2
	newList := []float64{}
	for _, quantity := range quantities {
		newList = append(newList, quantity * x)
	}
	return newList
}
