package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string,time int)int{
	if time == 0{
		return len(layers) * 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string)(noodels int,sauce float64){
	for i := 0;i< len(layers);i++{
		if layers[i] == "noodles"{
			noodels += 50
		}
		if layers[i] == "sauce"{
			sauce += 0.2
		}
	}
	return 
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList,myList []string)[]string{
	myList[len(myList)-1] = friendList[len(friendList)-1] 
	return myList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portion int)[]float64{
	scaledQuantities := make([]float64,len(quantities))
	for i:=0;i < len(scaledQuantities);i++{
		scaledQuantities[i] = quantities[i] * float64(portion)/2
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
