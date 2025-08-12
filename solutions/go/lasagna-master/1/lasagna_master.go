package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minutes int) int {
    if minutes == 0 {
        return len(layers) * 2
    }
    return len(layers) * minutes
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    noodles := int(0)
    sauce := float64(0)
    for i := range layers {
        if layers[i] == "noodles" {
            noodles += 50
        }

        if layers[i] == "sauce" {
            sauce += 0.2
        }
    }
    return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, amount int) []float64 {
    times := float64(amount / 2)
	if amount%2 != 0 {
		times += 0.5
	}

	newQuantities := []float64{}
	for i := range quantities {
		newQuantities = append(newQuantities, quantities[i]*float64(times))
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
