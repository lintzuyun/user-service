package user

import (
	pb "github.com/lintzuyun/protorepo-user-go-practice/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *UserService) MakePBUser(user *User) *pb.User {

	pbAddress := &pb.Address{
		AddressLine_1: user.Address.AddressLine1,
		AddressLine_2: user.Address.AddressLine2,
		City:          user.Address.City,
		PostCode:      user.Address.PostCode,
	}

	pbUser := &pb.User{
		Id:          user.Id,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Age:         user.Age,
		Birthday:    timestamppb.New(user.Birthday),
		PhoneNumber: user.PhoneNumber,
		Address:     pbAddress,
	}

	return pbUser

}

func (s *UserService) PbToCreateUser(pbUserReq *pb.CreateUserRequest) *CreateUser {

	a := pbUserReq.GetAddress()

	address := Address{
		AddressLine1: a.AddressLine_1,
		AddressLine2: a.AddressLine_2,
		City:         a.City,
		PostCode:     a.PostCode,
	}

	user := &CreateUser{
		FirstName:   pbUserReq.GetFirstName(),
		LastName:    pbUserReq.GetLastName(),
		Age:         pbUserReq.GetAge(),
		Birthday:    pbUserReq.Birthday.AsTime(),
		PhoneNumber: pbUserReq.GetPhoneNumber(),
		Address:     address,
	}
	return user
}

func (s *UserService) PbToUpdateUser(pbUserReq *pb.UpdateUserByIdRequest) *User {
	a := pbUserReq.GetAddress()

	address := Address{
		AddressLine1: a.AddressLine_1,
		AddressLine2: a.AddressLine_2,
		City:         a.City,
		PostCode:     a.PostCode,
	}

	user := &User{
		Id:          pbUserReq.GetId(),
		FirstName:   pbUserReq.GetFirstName(),
		LastName:    pbUserReq.GetLastName(),
		Age:         pbUserReq.GetAge(),
		Birthday:    pbUserReq.Birthday.AsTime(),
		PhoneNumber: pbUserReq.GetPhoneNumber(),
		Address:     address,
	}
	return user
}
