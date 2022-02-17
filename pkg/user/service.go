package user

import (
	"time"

	"go.uber.org/zap"
)

type UserService struct {
	logger *zap.Logger
	users  Users
}

func createUsers() *Users {

	JoysBirthday, _ := time.Parse("2006-01-02", "1991-08-30")
	katiesBirthday, _ := time.Parse("2006-01-02", "1988-12-01")
	dansBirthday, _ := time.Parse("2006-01-02", "1990-12-01")

	users := Users{
		User{
			Id:        "1",
			FirstName: "Joy",
			LastName:  "Lin",
			Age:       30,
			Birthday:  JoysBirthday},
		User{
			Id:          "2",
			FirstName:   "Katie",
			LastName:    "Hawcutt",
			Age:         33,
			Birthday:    katiesBirthday,
			PhoneNumber: 01234567,
		},
		User{
			Id:        "3",
			FirstName: "Dan",
			LastName:  "xx",
			Age:       23,
			Birthday:  dansBirthday,
		},
		User{
			Id:        "4",
			FirstName: "Jen",
			LastName:  "vv",
			Age:       43,
		},
	}
	return &users
}

func NewService(logger *zap.Logger) *UserService {
	return &UserService{
		logger: logger,
		users:  *createUsers(),
	}
}
