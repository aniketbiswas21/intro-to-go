package main

import "fmt"

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

// YOUR CODE HERE

func updateEmail(u *User, email string) {
	u.Email = email
}

func main() {
	fmt.Println("Pointers!")
	u := User{ID: 1, FirstName: "Aniket", LastName: "Biswas", Email: "abiswas_be19@thapar.edu"}
	customEmail := "aniket.biswas75@gmail.com"
	updateEmail(&u, customEmail)
	fmt.Println(u)

}
