package server

import (
	"context"
	"fmt"

	pb "github.com/lintzuyun/protorepo-user-go-practice/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	fmt.Println("We're at CreateUser =============================")
	s.logger.Debug("gRPC call: CreateUser")
	// convert pbUser to modelCreateUser
	userToCreate := s.userService.PbToCreateUser(req)

	// call a function from the user service to create a user (user.go)
	//s.userSrv.CreateUser()
	user, err := s.userService.CreateUser(userToCreate)
	fmt.Println(user)
	if err != nil {
		return nil, err
	}

	// convert modelUser to pbUser
	//🔺
	pbUser := s.userService.MakePBUser(user)
	return &pb.UserResponse{User: pbUser}, nil
}

func (s *Service) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.UserResponse, error) {
	fmt.Println("We're at GetUserById =============================")
	s.logger.Debug("gRPC call: GetUserByID")

	if req.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid ID")
	}

	user, err := s.userService.GetUser(req.GetId())
	if err != nil {
		return nil, err
	}

	pbUser := s.userService.MakePBUser(user)

	return &pb.UserResponse{
		User: pbUser,
	}, nil
}

func (s *Service) UpdateUserById(ctx context.Context, req *pb.UpdateUserByIdRequest) (*pb.UserResponse, error) {
	fmt.Println("We're at UpdateUserById =============================")
	s.logger.Debug("gRPC call: UpdateUserByID")

	if req.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid ID")
	}

	//🔺
	// userToUpdate := s.userService.PbToUpdateUser(req)

	// user, err := s.userService.UpdateUser(userToUpdate)
	// if err != nil {
	// 	return nil, err
	// }

	return &pb.UserResponse{}, nil
}

func (s *Service) DeleteUserById(ctx context.Context, req *pb.DeleteUserByIdRequest) (*pb.DeleteUserResponse, error) {
	fmt.Println("We're at DeleteUserById =============================")
	s.logger.Debug("gRPC call: DeleteUserByID")

	if req.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid ID")
	}

	id, err := s.userService.DeleteUser(req.GetId())
	if err != nil {
		return &pb.DeleteUserResponse{
			Id: id,
		}, err
	}

	return &pb.DeleteUserResponse{Id: id}, nil
}
