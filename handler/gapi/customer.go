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

type CustomerHandlerGrpc struct {
	pb.UnimplementedCustomerServiceServer
	customer entity.EntityCustomer
}

func NewCustomerHandlerGrpc(customer entity.EntityCustomer) *CustomerHandlerGrpc {
	customerServer := CustomerHandlerGrpc{
		customer: customer,
	}

	return &customerServer
}

func (c *CustomerHandlerGrpc) CreateCustomer(ctx context.Context, req *pb.CreateCustomerInput) (*pb.CustomerResponse, error) {
	// Create a SchemaCustomer instance from the request data
	newCustomer := &schemas.SchemaCustomer{
		ID:      uuid.New().String(),
		Name:    req.Name,
		Alamat:  req.Alamat,
		Email:   req.Email,
		Telepon: req.Telepon,
	}

	// Insert the new customer into the database
	createdCustomer, err := c.customer.EntityCreate(newCustomer)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	// Convert the created customer to a protobuf response
	res := &pb.CustomerResponse{
		Customer: &pb.Customer{
			Id:      createdCustomer.ID,
			Name:    createdCustomer.Name,
			Alamat:  createdCustomer.Alamat,
			Email:   createdCustomer.Email,
			Telepon: createdCustomer.Telepon,
		},
	}

	return res, nil
}

func (c *CustomerHandlerGrpc) GetCustomers(ctx context.Context, req *pb.CustomersRequest) (*pb.CustomersResponse, error) {
	customers, err := c.customer.EntityResults()

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var pbCustomers []*pb.Customer

	for _, customer := range *customers {
		pbCustomer := &pb.Customer{
			Id:        customer.ID,
			Name:      customer.Name,
			Alamat:    customer.Alamat,
			Email:     customer.Email,
			Telepon:   customer.Telepon,
			CreatedAt: timestamppb.New(customer.CreatedAt),
			UpdatedAt: timestamppb.New(customer.UpdatedAt),
		}
		pbCustomers = append(pbCustomers, pbCustomer)
	}

	return &pb.CustomersResponse{
		Customers: pbCustomers,
	}, nil
}

func (c *CustomerHandlerGrpc) GetCustomer(ctx context.Context, req *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	customerId := req.GetId()

	data := &schemas.SchemaCustomer{
		ID: customerId,
	}

	customer, err := c.customer.EntityResult(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CustomerResponse{
		Customer: &pb.Customer{
			Id:      customer.ID,
			Name:    customer.Name,
			Alamat:  customer.Alamat,
			Email:   customer.Email,
			Telepon: customer.Telepon,
		},
	}

	return res, nil
}

func (c *CustomerHandlerGrpc) UpdateCustomer(ctx context.Context, req *pb.UpdateCustomerInput) (*pb.CustomerResponse, error) {
	customerId := req.GetId()

	customerToUpdate := &schemas.SchemaCustomer{
		ID:      customerId,
		Name:    req.Name,
		Alamat:  req.Alamat,
		Email:   req.Email,
		Telepon: req.Telepon,
	}

	updatedCustomer, err := c.customer.EntityUpdate(customerToUpdate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CustomerResponse{
		Customer: &pb.Customer{
			Id:      updatedCustomer.ID,
			Name:    updatedCustomer.Name,
			Alamat:  updatedCustomer.Alamat,
			Email:   updatedCustomer.Email,
			Telepon: updatedCustomer.Telepon,
		},
	}

	return res, nil
}

func (c *CustomerHandlerGrpc) DeleteCustomer(ctx context.Context, req *pb.CustomerRequest) (*pb.DeleteCustomerResponse, error) {
	customerId := req.GetId()

	data := &schemas.SchemaCustomer{
		ID: customerId,
	}

	_, err := c.customer.EntityDelete(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.DeleteCustomerResponse{
		Success: true,
	}

	return res, nil
}
