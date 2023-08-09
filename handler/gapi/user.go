package gapi

import (
	"context"
	"echoinventory/entity"
	"echoinventory/pb"
	"echoinventory/schemas"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserHandlerGrpc struct {
	pb.UnimplementedUserServiceServer
	user entity.EntityUsers
}

func NewUserHandlerGrpcHandler(user entity.EntityUsers) *UserHandlerGrpc {
	userServer := UserHandlerGrpc{
		user: user,
	}

	return &userServer
}

func (u *UserHandlerGrpc) CreateUser(ctx context.Context, req *pb.CreateUserInput) (*pb.UserResponse, error) {
	// Create a SchemaUser instance from the request data
	newUser := &schemas.SchemaUser{
		ID:        uuid.New().String(),
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Email:     req.Email,
		Password:  req.Password,
		Role:      req.Role,
	}

	// Insert the new user into the database
	createdUser, err := u.user.EntityCreate(newUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	// Convert the created user to a protobuf response
	res := &pb.UserResponse{
		User: &pb.User{
			Id:        createdUser.ID,
			Firstname: createdUser.FirstName,
			Lastname:  createdUser.LastName,
			Email:     createdUser.Email,
			Role:      createdUser.Role,
		},
	}

	return res, nil
}

func (u *UserHandlerGrpc) GetUsers(ctx context.Context, req *pb.UsersRequest) (*pb.UsersResponse, error) {
	users, err := u.user.EntityResults()

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var pbUsers []*pb.User

	for _, user := range *users {
		pbUser := &pb.User{
			Id:        user.ID,
			Firstname: user.FirstName,
			Lastname:  user.LastName,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		}
		pbUsers = append(pbUsers, pbUser)
	}

	return &pb.UsersResponse{
		Users: pbUsers,
	}, nil
}

func (u *UserHandlerGrpc) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	userId := req.GetId()

	data := &schemas.SchemaUser{
		ID: userId,
	}

	user, err := u.user.EntityResult(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.UserResponse{
		User: &pb.User{
			Id:        user.ID,
			Firstname: user.FirstName,
			Lastname:  user.LastName,
			Email:     user.Email,
			Role:      user.Role,
		},
	}

	return res, nil
}

func (u *UserHandlerGrpc) UpdateUser(ctx context.Context, req *pb.UpdateUserInput) (*pb.UserResponse, error) {
	userId := req.GetId()

	userToUpdate := &schemas.SchemaUser{
		ID:        userId,
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Email:     req.Email,
		Role:      req.Role,
	}

	updatedUser, err := u.user.EntityUpdate(userToUpdate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.UserResponse{
		User: &pb.User{
			Id:        updatedUser.ID,
			Firstname: updatedUser.FirstName,
			Lastname:  updatedUser.LastName,
			Email:     updatedUser.Email,
			Role:      updatedUser.Role,
		},
	}

	return res, nil
}

func (u *UserHandlerGrpc) DeleteUser(ctx context.Context, req *pb.UserRequest) (*pb.DeleteUserResponse, error) {
	userId := req.GetId()

	data := &schemas.SchemaUser{
		ID: userId,
	}

	_, err := u.user.EntityDelete(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.DeleteUserResponse{
		Success: true,
	}

	return res, nil
}
