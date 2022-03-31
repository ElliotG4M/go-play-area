package models

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	// A slice of pointers to User objects
	users []*User
	// At the package level, no need to use :=, Go will figure out from the assignment that it's an int
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(user User) (User, error) {
	user.ID = nextID
	nextID++
	users = append(users, &user)
	return user, nil
}

func GetUser(ID int) (User, error) {
	if ID < 1 {
		return User{}, fmt.Errorf("Invalid ID provided")
	}
	for _, user := range users {
		if user.ID == ID {
			return *user, nil
		}
	}
	return User{}, fmt.Errorf("No user found with ID %v", ID)
}

func UpdateUser(updatedUser User) (User, error) {
	for i, user := range users {
		if user.ID == updatedUser.ID {
			users[i] = &updatedUser
			return updatedUser, nil
		}
	}
	return User{}, fmt.Errorf("No user found with ID %v", updatedUser.ID)
}

func RemoveUser(ID int) error {
	if ID < 1 {
		return fmt.Errorf("Invalid ID provided")
	}
	for i, user := range users {
		if user.ID == ID {
			// Get a slice of all the users up to index i, and append the slice of all users after i+1, effectively cutting out the one at i we want deleting
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("No user found with ID %v", ID)
}
