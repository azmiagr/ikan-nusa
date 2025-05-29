package rest

import (
	"fmt"
	"ikan-nusa/internal/service"
	"ikan-nusa/pkg/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.Interface
}

func NewRest(service *service.Service, middleware middleware.Interface) *Rest {
	return &Rest{
		router:     gin.Default(),
		service:    service,
		middleware: middleware,
	}
}

func (r *Rest) MountEndpoint() {
	baseURL := r.router.Group("/api/v1")
	baseURL.GET("/product-types", r.GetAllTypes)
	baseURL.GET("/product/:product_type_id", r.GetProductsByType)

	auth := baseURL.Group("/auth")
	auth.POST("/register", r.Register)
	auth.POST("/register/add-address", r.AddAddressAfterRegister)
	auth.PATCH("/register", r.VerifyUser)
	auth.POST("/login", r.Login)

	user := baseURL.Group("/users")
	user.Use(r.middleware.AuthenticateUser)
	user.GET("/address", r.GetUserAddresses)
	user.POST("/register-store", r.RegisterStore)
	user.POST("/add-to-cart", r.AddToCart)
	user.POST("/review", r.AddReview)

	store := baseURL.Group("/stores")
	store.Use(r.middleware.AuthenticateUser)
	store.GET("/products/:category", r.GetProductsByCategory)
	store.GET("/products/detail/:product_id", r.GetProductsDetail)
	store.GET("/products", r.GetProductsByName)
	store.GET("/all-products", r.GetAllProducts)
	store.POST("/add-product", r.AddProduct)
	store.POST("/upload-photo/:product_id", r.UploadPhoto)

}

func (r *Rest) Run() {
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	r.router.Run(fmt.Sprintf("%s:%s", addr, port))
}
