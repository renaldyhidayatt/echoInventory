package gapi

import (
	"context"
	"echoinventory/entity"
	"echoinventory/pb"
	"echoinventory/pkg"
	"echoinventory/schemas"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthHandlerGrpc struct {
	pb.UnimplementedAuthServiceServer
	auth entity.EntityAuth
}

func NewAuthHandlerGrpcServer(auth entity.EntityAuth) *AuthHandlerGrpc {
	authServer := AuthHandlerGrpc{
		auth: auth,
	}

	return &authServer
}

func (a *AuthHandlerGrpc) RegisterUser(ctx context.Context, req *pb.SignUpUserInput) (*pb.SignUpUserResponse, error) {
	user := schemas.SchemaUser{
		FirstName: req.GetFirstname(),
		LastName:  req.GetLastname(),
		Email:     req.GetEmail(),
		Password:  req.GetPassword(),
		Role:      req.GetRole(),
	}

	newUser, err := a.auth.EntityRegister(&user)

	if err != nil {
		if strings.Contains(err.Error(), "email already exist") {
			return nil, status.Errorf(codes.AlreadyExists, "%s", err.Error())
		}
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	res := &pb.SignUpUserResponse{
		User: &pb.User{
			Firstname: newUser.FirstName,
			Lastname:  newUser.LastName,
			Email:     newUser.Email,
			Role:      newUser.Role,
		},
	}

	return res, nil
}

func (a *AuthHandlerGrpc) LoginUser(ctx context.Context, req *pb.SignInUserInput) (*pb.SignInUserResponse, error) {
	user := schemas.SchemaUser{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	res, err := a.auth.EntityLogin(&user)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	jwt, err := pkg.Sign(&schemas.SchemaJWT{
		ID:    res.ID,
		Email: res.Email,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	result := &pb.SignInUserResponse{
		Status: "Success",
		Token:  jwt,
	}

	return result, nil
}
