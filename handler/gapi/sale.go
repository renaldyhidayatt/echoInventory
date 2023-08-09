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

type SaleHandlerGrpc struct {
	pb.UnimplementedSaleServiceServer
	sale entity.EntitySale
}

func NewSaleHandlerGrpc(sale entity.EntitySale) *SaleHandlerGrpc {
	saleServer := SaleHandlerGrpc{
		sale: sale,
	}

	return &saleServer
}

func (s *SaleHandlerGrpc) CreateSale(ctx context.Context, req *pb.CreateSaleInput) (*pb.SaleResponse, error) {
	newSale := &schemas.SchemaSale{
		ID:      uuid.New().String(),
		Name:    req.Name,
		Alamat:  req.Alamat,
		Email:   req.Email,
		Telepon: req.Telepon,
	}

	createdSale, err := s.sale.EntityCreate(newSale)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.SaleResponse{
		Sale: &pb.Sale{
			Id:      createdSale.ID,
			Name:    createdSale.Name,
			Alamat:  createdSale.Alamat,
			Email:   createdSale.Email,
			Telepon: createdSale.Telepon,
		},
	}

	return res, nil
}

func (s *SaleHandlerGrpc) GetSales(ctx context.Context, req *pb.SalesRequest) (*pb.SalesResponse, error) {
	sales, err := s.sale.EntityResults()

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var pbSales []*pb.Sale

	for _, sale := range *sales {
		pbSale := &pb.Sale{
			Id:        sale.ID,
			Name:      sale.Name,
			Alamat:    sale.Alamat,
			Email:     sale.Email,
			Telepon:   sale.Telepon,
			CreatedAt: timestamppb.New(sale.CreatedAt),
			UpdatedAt: timestamppb.New(sale.UpdatedAt),
		}
		pbSales = append(pbSales, pbSale)
	}

	return &pb.SalesResponse{
		Sales: pbSales,
	}, nil
}

func (s *SaleHandlerGrpc) GetSale(ctx context.Context, req *pb.SaleRequest) (*pb.SaleResponse, error) {
	saleId := req.GetId()

	data := &schemas.SchemaSale{
		ID: saleId,
	}

	sale, err := s.sale.EntityResult(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.SaleResponse{
		Sale: &pb.Sale{
			Id:      sale.ID,
			Name:    sale.Name,
			Alamat:  sale.Alamat,
			Email:   sale.Email,
			Telepon: sale.Telepon,
		},
	}

	return res, nil
}

func (s *SaleHandlerGrpc) UpdateSale(ctx context.Context, req *pb.UpdateSaleInput) (*pb.SaleResponse, error) {
	saleId := req.GetId()

	saleToUpdate := &schemas.SchemaSale{
		ID:      saleId,
		Name:    req.Name,
		Alamat:  req.Alamat,
		Email:   req.Email,
		Telepon: req.Telepon,
	}

	updatedSale, err := s.sale.EntityUpdate(saleToUpdate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.SaleResponse{
		Sale: &pb.Sale{
			Id:      updatedSale.ID,
			Name:    updatedSale.Name,
			Alamat:  updatedSale.Alamat,
			Email:   updatedSale.Email,
			Telepon: updatedSale.Telepon,
		},
	}

	return res, nil
}

func (s *SaleHandlerGrpc) DeleteSale(ctx context.Context, req *pb.SaleRequest) (*pb.DeleteSaleResponse, error) {
	saleId := req.GetId()

	data := &schemas.SchemaSale{
		ID: saleId,
	}

	_, err := s.sale.EntityDelete(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.DeleteSaleResponse{
		Success: true,
	}

	return res, nil
}
