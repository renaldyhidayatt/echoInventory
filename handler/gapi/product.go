package gapi

import (
	"context"
	"echoinventory/entity"
	"echoinventory/pb"
	"echoinventory/schemas"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductHandlerGrpc struct {
	pb.UnimplementedProductServiceServer
	product entity.EntityProduct
}

func NewProductHandlerGrpcHandler(product entity.EntityProduct) *ProductHandlerGrpc {
	productServer := ProductHandlerGrpc{
		product: product,
	}
	return &productServer
}

func (p *ProductHandlerGrpc) CreateProduct(ctx context.Context, req *pb.CreateProductInput) (*pb.ProductResponse, error) {
	newProduct := &schemas.SchemaProduct{
		ID:         uuid.New().String(),
		Name:       req.Name,
		Image:      req.Image,
		Qty:        req.Qty,
		CategoryID: req.CategoryId,
	}

	createdProduct, err := p.product.EntityCreate(newProduct)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ProductResponse{
		Product: &pb.Product{
			Id:         createdProduct.ID,
			Name:       createdProduct.Name,
			Image:      createdProduct.Image,
			Qty:        createdProduct.Qty,
			CategoryId: createdProduct.CategoryID,
		},
	}

	return res, nil
}

func (p *ProductHandlerGrpc) GetProducts(ctx context.Context, req *pb.ProductsRequest) (*pb.ProductsResponse, error) {
	products, err := p.product.EntityResults()

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var pbProducts []*pb.Product

	for _, product := range *products {
		pbProduct := &pb.Product{
			Id:    product.ID,
			Name:  product.Name,
			Image: product.Image,
			Qty:   product.Qty,
			Category: &pb.Category{
				Name: product.Category.Name,
			},
		}
		pbProducts = append(pbProducts, pbProduct)
	}

	return &pb.ProductsResponse{
		Products: pbProducts,
	}, nil
}

func (p *ProductHandlerGrpc) GetProduct(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	productID := req.GetId()

	data := &schemas.SchemaProduct{
		ID: productID,
	}

	product, err := p.product.EntityResult(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ProductResponse{
		Product: &pb.Product{
			Id:         product.ID,
			Name:       product.Name,
			Image:      product.Image,
			Qty:        product.Qty,
			CategoryId: product.CategoryID,
		},
	}

	return res, nil
}

func (p *ProductHandlerGrpc) UpdateProduct(ctx context.Context, req *pb.UpdateProductInput) (*pb.ProductResponse, error) {
	productID := req.GetId()

	productToUpdate := &schemas.SchemaProduct{
		ID:         productID,
		Name:       req.Name,
		Image:      req.Image,
		Qty:        req.Qty,
		CategoryID: req.CategoryId,
	}

	updatedProduct, err := p.product.EntityUpdate(productToUpdate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ProductResponse{
		Product: &pb.Product{
			Id:         updatedProduct.ID,
			Name:       updatedProduct.Name,
			Image:      updatedProduct.Image,
			Qty:        updatedProduct.Qty,
			CategoryId: updatedProduct.CategoryID,
		},
	}

	return res, nil
}

func (p *ProductHandlerGrpc) DeleteProduct(ctx context.Context, req *pb.ProductRequest) (*pb.DeleteProductResponse, error) {
	productID := req.GetId()

	data := &schemas.SchemaProduct{
		ID: productID,
	}

	_, err := p.product.EntityDelete(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.DeleteProductResponse{
		Success: true,
	}

	return res, nil
}
