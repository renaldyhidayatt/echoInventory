package main

import (
	"echoinventory/config"
	"echoinventory/handler/gapi"
	"echoinventory/pb"
	"echoinventory/repository"
	"echoinventory/services"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "gRPC server port")
)

func main() {
	fmt.Println("gRPC server running ...")

	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	db := config.DatabaseConnect()

	authRepo := repository.NewRepositoryAuth(db)
	authService := services.NewServiceAuth(authRepo)

	categoryRepo := repository.NewRepositoryCategory(db)
	categoryService := services.NewServiceCategory(categoryRepo)

	customerRepo := repository.NewRepositoryCustomer(db)
	customerService := services.NewServiceCustomer(customerRepo)

	productRepo := repository.NewRepositoryProduct(db)
	productService := services.NewServiceProduct(productRepo)

	productKeluarRepo := repository.NewRepositoryProductKeluar(db)
	productKeluarService := services.NewServiceProductKeluar(productKeluarRepo)

	userRepo := repository.NewRepositoryUser(db)
	userService := services.NewServiceUser(userRepo)

	saleRepo := repository.NewRepositorySale(db)
	saleService := services.NewServiceSale(saleRepo)

	supplierRepo := repository.NewRepositorySupplier(db)
	supplierService := services.NewServiceSupplier(supplierRepo)

	authGapi := gapi.NewAuthHandlerGrpcServer(authService)
	categoryGapi := gapi.NewCategoryHandlerGrpcHandler(categoryService)

	customerGapi := gapi.NewCustomerHandlerGrpc(customerService)

	productGapi := gapi.NewProductHandlerGrpcHandler(productService)

	userGapi := gapi.NewUserHandlerGrpcHandler(userService)

	productKeluarGapi := gapi.NewProductKeluarHandlerGrpcHandler(productKeluarService)

	saleGapi := gapi.NewSaleHandlerGrpc(saleService)

	supplierGapi := gapi.NewSupplierHandlerGrpc(supplierService)

	if err != nil {
		return
	}

	s := grpc.NewServer()

	pb.RegisterAuthServiceServer(s, authGapi)
	pb.RegisterCategoryServiceServer(s, categoryGapi)
	pb.RegisterCustomerServiceServer(s, customerGapi)
	pb.RegisterProductServiceServer(s, productGapi)
	pb.RegisterUserServiceServer(s, userGapi)
	pb.RegisterProductKeluarServiceServer(s, productKeluarGapi)
	pb.RegisterSaleServiceServer(s, saleGapi)
	pb.RegisterSupplierServiceServer(s, supplierGapi)

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
