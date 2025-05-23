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

	auth := baseURL.Group("/auth")
	auth.POST("/register", r.Register)
	auth.POST("/register/add-address", r.AddAddressAfterRegister)
	auth.PATCH("/register", r.VerifyUser)
	auth.POST("/login", r.Login)

	user := baseURL.Group("/users")
	user.Use(r.middleware.AuthenticateUser)
	user.POST("/register-store", r.RegisterStore)

	store := baseURL.Group("/stores")
	store.Use(r.middleware.AuthenticateUser)
	store.GET("/products/:category", r.GetProductsByCategory)
	store.POST("/add-product", r.AddProduct)

}

func (r *Rest) Run() {
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	r.router.Run(fmt.Sprintf("%s:%s", addr, port))
}
