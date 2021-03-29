package main

import "fmt"

// User is a user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

type Group struct {
	role             string
	Users            []User
	NewestUser       User
	IsSpaceAvailable bool
}

func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

func describeGroup(g Group) string {
	if len(g.Users) > 2 {
		g.IsSpaceAvailable = false
	}
	desc := fmt.Sprintf("This user group has %d. The newest user is %s %s. Accepting New Users: %t", len(g.Users), g.NewestUser.FirstName, g.NewestUser.LastName, g.IsSpaceAvailable)
	return desc
}

// // User is a user type
// type User struct {
// ID					       int
// FirstName, LastName, Email string
// }

func main() {
	u1 := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	u2 := User{ID: 2, FirstName: "Aniket", LastName: "Biswas", Email: "abiswas_be19@thapar.edu"}
	g1 := Group{role: "admin", Users: []User{u1, u2}, NewestUser: u2, IsSpaceAvailable: true}

	fmt.Println(describeGroup(g1))
	fmt.Println(describeUser(u1))

}
