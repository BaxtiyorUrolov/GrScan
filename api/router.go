package api

import (
	"grscan/api/handler"
	"grscan/pkg/logger"
	"grscan/service"
	"grscan/storage"

	_ "grscan/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New(services service.IServiceManager, storage storage.IStorage, log logger.ILogger) *gin.Engine {
	h := handler.New(services, storage, log)

	r := gin.New()

	r.Use(gin.Logger())

	// auth endpoints
	r.POST("/auth/customer/login", h.CustomerLogin)

	r.POST("/user", h.CreateUser)
	r.POST("/verify-register", h.VerifyRegister)

	// Swagger documentation route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
