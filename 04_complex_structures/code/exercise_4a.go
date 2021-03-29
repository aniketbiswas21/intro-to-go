package main

import "fmt"

func average(num1, num2, num3 float64) float64 {
	avg:= (num1+num2+num3)/3
	return avg
}

func awesomeAverage(nums ...float64) float64 {
	var total float64
	var totalElements int
	totalElements = len(nums)
	for _,value := range nums {
		total+=value
	}
	return total/float64((totalElements))
}

func main() {
	var no, no1, no2 float64
	fmt.Println("Enter three numbers")
	fmt.Scan(&no)
	fmt.Scan(&no1)
	fmt.Scan(&no2)
	fmt.Printf("The average is: %f", average(no,no1,no2))
	fmt.Printf("The average is: %f", awesomeAverage(no,no1,no2))
	
}