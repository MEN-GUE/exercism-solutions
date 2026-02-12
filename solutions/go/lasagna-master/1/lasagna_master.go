package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgMinPerLayer int) int{
    if(avgMinPerLayer == 0) {
        return len(layers) * 2
    }
    return len(layers) * avgMinPerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64){
    for i:=0; i<len(layers);i++ {
        switch {
            case layers[i] == "sauce":
            	sauce += 0.2
            case layers[i] == "noodles":
            	noodles += 50
        }
    }
    return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string){
    myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64{
    var scaledQuantities []float64
    for i:=0 ; i<len(quantities) ;i++{
        scaledQuantities = append(scaledQuantities, (quantities[i]/2) * float64(portions))
    }
    return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
