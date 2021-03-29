package main

import "fmt"

func avg(param [5]float64) float64 {
	total:=0.0
	for _,number:= range param {
		total+=number
	}
	return total/float64(len(param))
}

func doesPetExist(myMap map[int]string,petIndex int) bool {
	_,ok:= myMap[petIndex]
	return ok
}

func addToGroceryList(slice []string,newGroceries ...string) []string {
	for _,g := range newGroceries {
		slice = append(slice, g)
	}
	return slice
}

func main() {
	scores :=[5]float64{1.2,2.4,4.6,6.3,7.0}
	// var myArray []int = make([]int, 5)
	fmt.Println(avg(scores))

	var petNames map[int]string
	petNames = make(map[int]string)
	petNames[1] = "Dog"
	petNames[2] = "Cat"
	fmt.Println(petNames)
	fmt.Println(doesPetExist(petNames,1))
	// var mySlice []int = make([]int, 5)
	mySlice:= []string{"apple","chocolate"}
	fmt.Println(mySlice)
	newSlice:= addToGroceryList(mySlice, "ice-cream", "banana" )
	fmt.Println(newSlice)

}