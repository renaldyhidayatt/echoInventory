package gapi

import (
	"context"
	"echoinventory/entity"
	"echoinventory/pb"
	"echoinventory/schemas"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProductKeluarGrpc struct {
	pb.UnimplementedProductKeluarServiceServer
	productkeluar entity.EntityProductKeluar
}

func NewProductKeluarHandlerGrpcHandler(product entity.EntityProductKeluar) *ProductKeluarGrpc {
	return &ProductKeluarGrpc{
		productkeluar: product,
	}
}

func (p *ProductKeluarGrpc) CreateProductKeluar(ctx context.Context, req *pb.CreateProductKeluarInput) (*pb.ProductKeluarResponse, error) {
	newProductKeluar := &schemas.SchemaProductKeluar{
		Qty:        req.Qty,
		ProductID:  req.ProductId,
		CategoryID: req.CategoryId,
	}

	createdProductKeluar, err := p.productkeluar.EntityCreate(newProductKeluar)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pbProductKeluar := &pb.ProductKeluar{
		Id:         createdProductKeluar.ID,
		Qty:        createdProductKeluar.Qty,
		ProductId:  createdProductKeluar.ProductID,
		CategoryId: createdProductKeluar.CategoryID,
		CreatedAt:  timestamppb.New(createdProductKeluar.CreatedAt),
		UpdatedAt:  timestamppb.New(createdProductKeluar.UpdatedAt),
	}

	return &pb.ProductKeluarResponse{
		ProductKeluar: pbProductKeluar,
	}, nil
}

func (p *ProductKeluarGrpc) GetProductKeluars(ctx context.Context, req *pb.ProductKeluarsRequest) (*pb.ProductKeluarsResponse, error) {
	productKeluar, err := p.productkeluar.EntityResults()

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var pbProductKeluars []*pb.ProductKeluar

	for _, product := range *productKeluar {
		pbProductKeluar := &pb.ProductKeluar{
			Id:         product.ID,
			Qty:        product.Qty,
			ProductId:  product.ProductID,
			CategoryId: product.CategoryID,
			CreatedAt:  timestamppb.New(product.CreatedAt),
			UpdatedAt:  timestamppb.New(product.UpdatedAt),
		}
		pbProductKeluars = append(pbProductKeluars, pbProductKeluar)
	}

	return &pb.ProductKeluarsResponse{
		ProductKeluars: pbProductKeluars,
	}, nil
}

func (p *ProductKeluarGrpc) GetProductKeluar(ctx context.Context, req *pb.ProductKeluarRequest) (*pb.ProductKeluarResponse, error) {
	productKeluarID := req.GetId()

	data := &schemas.SchemaProductKeluar{
		ID: productKeluarID,
	}

	productKeluar, err := p.productkeluar.EntityResult(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pbProductKeluar := &pb.ProductKeluar{
		Id:         productKeluar.ID,
		Qty:        productKeluar.Qty,
		ProductId:  productKeluar.ProductID,
		CategoryId: productKeluar.CategoryID,
		CreatedAt:  timestamppb.New(productKeluar.CreatedAt),
		UpdatedAt:  timestamppb.New(productKeluar.UpdatedAt),
	}

	return &pb.ProductKeluarResponse{
		ProductKeluar: pbProductKeluar,
	}, nil
}

func (p *ProductKeluarGrpc) UpdateProductKeluar(ctx context.Context, req *pb.UpdateProductKeluarInput) (*pb.ProductKeluarResponse, error) {
	productKeluarID := req.GetId()

	productKeluarToUpdate := &schemas.SchemaProductKeluar{
		ID:         productKeluarID,
		Qty:        req.Qty,
		ProductID:  req.ProductId,
		CategoryID: req.CategoryId,
	}

	updatedProductKeluar, err := p.productkeluar.EntityUpdate(productKeluarToUpdate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pbProductKeluar := &pb.ProductKeluar{
		Id:         updatedProductKeluar.ID,
		Qty:        updatedProductKeluar.Qty,
		ProductId:  updatedProductKeluar.ProductID,
		CategoryId: updatedProductKeluar.CategoryID,
		CreatedAt:  timestamppb.New(updatedProductKeluar.CreatedAt),
		UpdatedAt:  timestamppb.New(updatedProductKeluar.UpdatedAt),
	}

	return &pb.ProductKeluarResponse{
		ProductKeluar: pbProductKeluar,
	}, nil
}

func (p *ProductKeluarGrpc) DeleteProductKeluar(ctx context.Context, req *pb.ProductKeluarRequest) (*pb.DeleteProductKeluarResponse, error) {
	productKeluarID := req.GetId()

	data := &schemas.SchemaProductKeluar{
		ID: productKeluarID,
	}

	_, err := p.productkeluar.EntityDelete(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.DeleteProductKeluarResponse{
		Success: true,
	}, nil
}
