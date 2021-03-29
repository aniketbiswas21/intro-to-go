package main

import "fmt"

func main() {
	var sentence string = "This is a sentence"
	for index, value := range sentence {
		if index%2 != 0 {

			fmt.Println("Index", index, "Value:", string(value))
		}
	}
}
