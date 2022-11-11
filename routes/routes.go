package routes

import (
	"echoinventory/handler"
	"echoinventory/repository"
	"echoinventory/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRoute(db *gorm.DB, router *echo.Echo) {
	repositoryAuth := repository.NewRepositoryAuth(db)
	serviceAuth := services.NewServiceAuth(repositoryAuth)
	handlerAuth := handler.NewHandlerAuth(serviceAuth)

	repositoryCategory := repository.NewRepositoryCategory(db)
	serviceCategory := services.NewServiceCategory(repositoryCategory)
	handlerCategory := handler.NewHandlerCategory(serviceCategory)

	repositoryUser := repository.NewRepositoryUser(db)
	serviceUser := services.NewServiceUser(repositoryUser)
	handlerUser := handler.NewHandlerUser(serviceUser)

	repositoryProduct := repository.NewRepositoryUser(db)
	serviceProduct := services.NewServiceUser(repositoryProduct)
	handlerProduct := handler.NewHandlerUser(serviceProduct)

	repositoryCustomer := repository.NewRepositoryCustomer(db)
	serviceCustomer := services.NewServiceCustomer(repositoryCustomer)
	handlerCustomer := handler.NewHandlerCustomer(serviceCustomer)

	repositoryProductKeluar := repository.NewRepositoryProductKeluar(db)
	serviceProductKeluar := services.NewServiceProductKeluar(repositoryProductKeluar)
	handlerProductKeluar := handler.NewHandlerProductKeluar(serviceProductKeluar)

	// Product Masuk
	repositoryProductMasuk := repository.NewRepositoryProductMasuk(db)
	serviceProductMasuk := services.NewServiceProductMasuk(repositoryProductMasuk)
	handlerProductMasuk := handler.NewHandlerProductMasuk(serviceProductMasuk)

	// Sale
	repositorySale := repository.NewRepositorySale(db)
	serviceSale := services.NewServiceSale(repositorySale)
	handlerSale := handler.NewHandlerSale(serviceSale)

	// Supplier
	repositorySupplier := repository.NewRepositorySupplier(db)
	serviceSupplier := services.NewServiceSupplier(repositorySupplier)
	handlerSupplier := handler.NewHandlerSupplier(serviceSupplier)

	route := router.Group("/api/hello")
	routerAuth := router.Group("/api/auth")
	routerCategory := router.Group("/api/category")
	routerUser := router.Group("/api/users")
	routerProduct := router.Group("/api/product")
	routerCustomer := router.Group("/api/customer")
	routerProductKeluar := router.Group("/api/productkeluar")

	routerProductMasuk := router.Group("/api/productmasuk")

	routerSale := router.Group("/api/sale")

	routeSupplier := router.Group("/api/supplier")

	route.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	routerAuth.POST("/login", handlerAuth.HandlerLogin)
	routerAuth.POST("/register", handlerAuth.HandlerRegister)

	routerCategory.GET("/hello", handlerCategory.HandlerHello)
	routerCategory.POST("/create", handlerCategory.HandlerCreate)
	routerCategory.PUT("/update/:id", handlerCategory.HandlerUpdate)
	routerCategory.DELETE("/delete/:id", handlerCategory.HandlerDelete)
	routerCategory.GET("/", handlerCategory.HandlerResults)
	routerCategory.GET("/:id", handlerCategory.HandlerResult)

	routerUser.GET("/hello", handlerUser.HandlerHello)
	routerUser.POST("/create", handlerUser.HandlerCreate)
	routerUser.PUT("/update/:id", handlerUser.HandlerUpdate)
	routerUser.DELETE("/delete/:id", handlerUser.HandlerDelete)
	routerUser.GET("/", handlerUser.HandlerResults)
	routerUser.GET("/:id", handlerUser.HandlerResult)

	routerProduct.GET("/hello", handlerProduct.HandlerHello)
	routerProduct.POST("/create", handlerProduct.HandlerCreate)
	routerProduct.PUT("/update/:id", handlerProduct.HandlerUpdate)
	routerProduct.DELETE("/delete/:id", handlerProduct.HandlerDelete)
	routerProduct.GET("/", handlerProduct.HandlerResults)
	routerProduct.GET("/:id", handlerProduct.HandlerResult)

	routerCustomer.GET("/hello", handlerCustomer.HandlerHello)
	routerCustomer.POST("/create", handlerCustomer.HandlerCreate)
	routerCustomer.PUT("/update/:id", handlerCustomer.HandlerUpdate)
	routerCustomer.DELETE("/delete/:id", handlerCustomer.HandlerDelete)
	routerCustomer.GET("/", handlerCustomer.HandlerResults)
	routerCustomer.GET("/:id", handlerCustomer.HandlerResult)

	routerProductKeluar.GET("/", handlerProductKeluar.HandlerResults)
	routerProductKeluar.GET("/:id", handlerProductKeluar.HandlerResult)
	routerProductKeluar.POST("/create", handlerProductKeluar.HandlerCreate)
	routerProductKeluar.POST("/update/:id", handlerProductKeluar.HandlerUpdate)
	routerProductKeluar.POST("/delete/:id", handlerProductKeluar.HandlerDelete)

	routerProductMasuk.GET("/hello", handlerProductMasuk.HandlerHello)
	routerProductMasuk.GET("/", handlerProductMasuk.HandlerResults)
	routerProductMasuk.GET("/:id", handlerProductMasuk.HandlerResult)
	routerProductMasuk.POST("/create", handlerProductMasuk.HandlerCreate)
	routerProductMasuk.POST("/update/:id", handlerProductMasuk.HandlerUpdate)
	routerProductMasuk.POST("/delete/:id", handlerProductMasuk.HandlerDelete)

	routerSale.GET("/hello", handlerSale.HandlerHello)
	routerSale.GET("/", handlerSale.HandlerResults)
	routerSale.GET("/:id", handlerSale.HandlerResult)
	routerSale.POST("/create", handlerSale.HandlerCreate)
	routerSale.POST("/update/:id", handlerSale.HandlerUpdate)
	routerSale.POST("/delete/:id", handlerSale.HandlerDelete)

	routeSupplier.GET("/hello", handlerSupplier.HandlerHello)
	routeSupplier.GET("/", handlerSupplier.HandlerResults)
	routeSupplier.GET("/:id", handlerSupplier.HandlerResults)
	routeSupplier.POST("/create", handlerSupplier.HandlerCreate)
	routeSupplier.POST("/update/:id", handlerSupplier.HandlerUpdate)
	routeSupplier.POST("/delete/:id", handlerSupplier.HandlerDelete)

}
