package user

import "time"

type User struct {
	Id          string
	FirstName   string
	LastName    string
	Age         int32
	Birthday    time.Time
	PhoneNumber int32
	Address     Address
}

type CreateUser struct {
	FirstName   string
	LastName    string
	Age         int32
	Birthday    time.Time
	PhoneNumber int32
	Address     Address
}

type Address struct {
	AddressLine1 string
	AddressLine2 string
	City         string
	PostCode     string
}

type Users []User
