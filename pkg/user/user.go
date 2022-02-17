package user

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func (s *UserService) CreateUser(userToCreate *CreateUser) (*User, error) {

	user := User{
		Id:          uuid.New().String(),
		FirstName:   userToCreate.FirstName,
		LastName:    userToCreate.LastName,
		Age:         userToCreate.Age,
		Birthday:    userToCreate.Birthday,
		PhoneNumber: userToCreate.PhoneNumber,
		Address:     userToCreate.Address,
	}

	s.users = append(s.users, user)
	fmt.Println(s.users)

	return &user, nil
}

func (s *UserService) GetUser(id string) (*User, error) {

	var user User

	for _, u := range s.users {
		if u.Id == id {
			user = u
		}
	}

	emptyUser := User{}

	if user == emptyUser {
		return nil, errors.New("User not found")
	}

	return &user, nil
}

// 🔺
func (s *UserService) UpdateUser(userToUpdate *User) (*User, error) {

	// for _, u := range s.users {
	// 	if u.Id == id {

	// 	}
	// }

	return &User{}, nil
}

func (s *UserService) DeleteUser(id string) (string, error) {

	var index int

	for i, u := range s.users {
		if u.Id == id {
			index = i
		}
	}

	s.users[index] = s.users[len(s.users)-1]
	s.users = s.users[:len(s.users)-1]

	return id, nil
}
