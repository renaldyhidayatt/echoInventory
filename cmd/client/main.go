package main

import (
	"echoinventory/pb"
	"echoinventory/route"
	"flag"
	"log"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	e := echo.New()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	clientAuth := pb.NewAuthServiceClient(conn)
	clientCategory := pb.NewCategoryServiceClient(conn)
	clientCustomer := pb.NewCustomerServiceClient(conn)
	clientUser := pb.NewUserServiceClient(conn)
	clientProduct := pb.NewProductServiceClient(conn)
	clientProductKeluar := pb.NewProductKeluarServiceClient(conn)
	clientSale := pb.NewSaleServiceClient(conn)
	clientSupplier := pb.NewSupplierServiceClient(conn)
	clientProductMasuk := pb.NewProductMasukServiceClient(conn)

	route.NewAuthRoute(clientAuth, e)
	route.NewCategoryRoute(clientCategory, e)
	route.NewCustomerRoute(clientCustomer, e)
	route.NewUserRoute(clientUser, e)
	route.NewProductRoute(clientProduct, e)
	route.NewProductMasukRoute(clientProductMasuk, e)
	route.NewProductKeluarRoute(clientProductKeluar, e)
	route.NewSaleRoute(clientSale, e)
	route.NewSupplierRoute(clientSupplier, e)

	e.Logger.Fatal(e.Start(":5000"))
}
