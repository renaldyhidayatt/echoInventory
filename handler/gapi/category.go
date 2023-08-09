package gapi

import (
	"context"
	"echoinventory/entity"
	"echoinventory/pb"
	"echoinventory/schemas"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CategoryHandlerGrpc struct {
	pb.UnimplementedCategoryServiceServer
	category entity.EntityCategory
}

func NewCategoryHandlerGrpcHandler(category entity.EntityCategory) *CategoryHandlerGrpc {

	categoryServer := CategoryHandlerGrpc{
		category: category,
	}

	return &categoryServer
}

func (c *CategoryHandlerGrpc) GetCategories(ctx context.Context, req *pb.CategoriesRequest) (*pb.CategoriesResponse, error) {
	categories, err := c.category.EntityResults()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	var pbCategories []*pb.Category

	for _, category := range *categories {
		createdAtProto := timestamppb.New(category.CreatedAt)
		updatedAtProto := timestamppb.New(category.UpdatedAt)

		pbCategory := &pb.Category{
			Id:        category.ID,
			Name:      category.Name,
			CreatedAt: createdAtProto,
			UpdatedAt: updatedAtProto,
		}
		pbCategories = append(pbCategories, pbCategory)
	}

	return &pb.CategoriesResponse{
		Categories: pbCategories,
	}, nil
}

func (c *CategoryHandlerGrpc) GetCategory(ctx context.Context, req *pb.CategoryRequest) (*pb.CategoryResponse, error) {
	catId := req.GetId()

	data := &schemas.SchemaCategory{
		ID: catId,
	}

	category, err := c.category.EntityResult(data)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	res := &pb.CategoryResponse{
		Category: &pb.Category{
			Id:   category.ID,
			Name: category.Name,
		},
	}

	return res, nil
}

func (c *CategoryHandlerGrpc) CreateCategory(ctx context.Context, req *pb.CreateCategoryInput) (*pb.CategoryResponse, error) {

	newCategory := &schemas.SchemaCategory{
		ID:   uuid.New().String(),
		Name: req.Name,
	}

	createdCategory, err := c.category.EntityCreate(newCategory)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CategoryResponse{
		Category: &pb.Category{
			Id:   createdCategory.ID,
			Name: createdCategory.Name,
		},
	}

	return res, nil
}

func (c *CategoryHandlerGrpc) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryInput) (*pb.CategoryResponse, error) {
	catId := req.GetId()

	cat := &schemas.SchemaCategory{
		ID:   catId,
		Name: req.Name,
	}

	update, err := c.category.EntityUpdate(cat)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CategoryResponse{
		Category: &pb.Category{
			Id:   update.ID,
			Name: update.Name,
		},
	}

	return res, nil
}

func (c *CategoryHandlerGrpc) DeleteCategory(ctx context.Context, req *pb.CategoryRequest) (*pb.DeleteCategoryResponse, error) {
	catId := req.GetId()

	data := &schemas.SchemaCategory{
		ID: catId,
	}

	_, err := c.category.EntityDelete(data)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.DeleteCategoryResponse{
		Success: true,
	}

	return res, nil
}
