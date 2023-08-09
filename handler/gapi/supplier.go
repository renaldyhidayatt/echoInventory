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

type SupplierHandlerGrpc struct {
	pb.UnimplementedSupplierServiceServer
	supplier entity.EntitySupplier
}

func NewSupplierHandlerGrpc(supplier entity.EntitySupplier) *SupplierHandlerGrpc {
	supplierServer := SupplierHandlerGrpc{
		supplier: supplier,
	}

	return &supplierServer
}

func (s *SupplierHandlerGrpc) CreateSupplier(ctx context.Context, req *pb.CreateSupplierInput) (*pb.SupplierResponse, error) {
	newSupplier := &schemas.SchemaSupplier{
		ID:      uuid.New().String(),
		Name:    req.Name,
		Alamat:  req.Alamat,
		Email:   req.Email,
		Telepon: req.Telepon,
	}

	createdSupplier, err := s.supplier.EntityCreate(newSupplier)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.SupplierResponse{
		Supplier: &pb.Supplier{
			Id:      createdSupplier.ID,
			Name:    createdSupplier.Name,
			Alamat:  createdSupplier.Alamat,
			Email:   createdSupplier.Email,
			Telepon: createdSupplier.Telepon,
		},
	}

	return res, nil
}

func (s *SupplierHandlerGrpc) GetSuppliers(ctx context.Context, req *pb.SuppliersRequest) (*pb.SuppliersResponse, error) {
	suppliers, err := s.supplier.EntityResults()

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var pbSuppliers []*pb.Supplier

	for _, supplier := range *suppliers {
		pbSupplier := &pb.Supplier{
			Id:        supplier.ID,
			Name:      supplier.Name,
			Alamat:    supplier.Alamat,
			Email:     supplier.Email,
			Telepon:   supplier.Telepon,
			CreatedAt: timestamppb.New(supplier.CreatedAt),
			UpdatedAt: timestamppb.New(supplier.UpdatedAt),
		}
		pbSuppliers = append(pbSuppliers, pbSupplier)
	}

	return &pb.SuppliersResponse{
		Suppliers: pbSuppliers,
	}, nil
}

func (s *SupplierHandlerGrpc) GetSupplier(ctx context.Context, req *pb.SupplierRequest) (*pb.SupplierResponse, error) {
	supplierId := req.GetId()

	data := &schemas.SchemaSupplier{
		ID: supplierId,
	}

	supplier, err := s.supplier.EntityResult(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.SupplierResponse{
		Supplier: &pb.Supplier{
			Id:      supplier.ID,
			Name:    supplier.Name,
			Alamat:  supplier.Alamat,
			Email:   supplier.Email,
			Telepon: supplier.Telepon,
		},
	}

	return res, nil
}

func (s *SupplierHandlerGrpc) UpdateSupplier(ctx context.Context, req *pb.UpdateSupplierInput) (*pb.SupplierResponse, error) {
	supplierId := req.GetId()

	supplierToUpdate := &schemas.SchemaSupplier{
		ID:      supplierId,
		Name:    req.Name,
		Alamat:  req.Alamat,
		Email:   req.Email,
		Telepon: req.Telepon,
	}

	updatedSupplier, err := s.supplier.EntityUpdate(supplierToUpdate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.SupplierResponse{
		Supplier: &pb.Supplier{
			Id:      updatedSupplier.ID,
			Name:    updatedSupplier.Name,
			Alamat:  updatedSupplier.Alamat,
			Email:   updatedSupplier.Email,
			Telepon: updatedSupplier.Telepon,
		},
	}

	return res, nil
}

func (s *SupplierHandlerGrpc) DeleteSupplier(ctx context.Context, req *pb.SupplierRequest) (*pb.DeleteSupplierResponse, error) {
	supplierId := req.GetId()

	data := &schemas.SchemaSupplier{
		ID: supplierId,
	}

	_, err := s.supplier.EntityDelete(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.DeleteSupplierResponse{
		Success: true,
	}

	return res, nil
}
