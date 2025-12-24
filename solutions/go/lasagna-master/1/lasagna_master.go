package lasagna

// TODO: define the 'PreparationTime()' function
	func PreparationTime(layers []string, average int) int {
        if average == 0 {
            average = 2
        }
        return average * len(layers)
    }

// TODO: define the 'Quantities()' function
	func Quantities(layers []string) (int, float64) {
        var grams int
        var sauce float64
        for i := 0; i < len(layers); i++ {
            if layers[i] == "noodles" {
                grams += 50
            }
            if layers[i] == "sauce" {
                sauce += 0.2
            }
        }
        return grams, sauce
    }
// TODO: define the 'AddSecretIngredient()' function
	func AddSecretIngredient(friendsList, myList []string) {
        myLastItem := len(myList) - 1
        friendsLastItem := len(friendsList) - 1
        (myList)[myLastItem] = friendsList[friendsLastItem]
    }

// TODO: define the 'ScaleRecipe()' function
	func ScaleRecipe(quantities []float64, portions int) []float64 {
        
        multiplier := float64(portions) / 2.0
        multipliedQuantities := make([]float64, len(quantities))
                                     
        for i:= 0; i < len(quantities); i++ {
            multipliedQuantities[i] = quantities[i] * multiplier
        }
        return multipliedQuantities
    }
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
