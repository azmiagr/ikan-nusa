package rest

import (
	"fmt"
	"ikan-nusa/internal/service"
	"os"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router  *gin.Engine
	service *service.Service
}

func NewRest(service *service.Service) *Rest {
	return &Rest{
		router:  gin.Default(),
		service: service,
	}
}

func (r *Rest) MountEndpoint() {
	baseURL := r.router.Group("/api/v1")

	auth := baseURL.Group("/auth")
	auth.POST("/register", r.Register)
	auth.POST("/register/add-address", r.AddAddressAfterRegister)
	auth.PATCH("/register", r.VerifyUser)
	auth.POST("/login", r.Login)

}

func (r *Rest) Run() {
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	r.router.Run(fmt.Sprintf("%s:%s", addr, port))
}
