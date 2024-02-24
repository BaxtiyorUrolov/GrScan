package api

import (
	"grscan/api/handler"
	"grscan/pkg/logger"
	"grscan/service"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_"grscan/api/docs"
)

func New(services service.IServiceManager, log logger.ILogger) *gin.Engine {
	h := handler.New(services, log)

	r := gin.New()

	r.Use(traceRequest(log))

	r.POST("/user", h.CreateUser)
	// r.GET("/user/:id", h.Getuser)
	// r.GET("/users", h.GetuserList)
	// r.PUT("/user/:id", h.Updateuser)
	// r.DELETE("/user/:id", h.Deleteuser)
	// r.PATCH("/user/:id", h.UpdatePageNumber)

	// Swagger documentation route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func traceRequest(log logger.ILogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		beforeRequest(c, log)

		c.Next()

		afterRequest(c)
	}
}

func beforeRequest(c *gin.Context, log logger.ILogger) {
	startTime := time.Now()

	c.Set("start_time", startTime)

	log.Info("Request started", logger.String("start_time", startTime.Format(time.RFC3339)), logger.String("path", c.Request.URL.Path))
}

func afterRequest(c *gin.Context) {
	startTime, exists := c.Get("start_time")
	if !exists {
		startTime = time.Now()
	}

	duration := time.Since(startTime.(time.Time)).Nanoseconds()

	statusCode := c.Writer.Status() 

	log.Println("end time:", time.Now().Format("2006-01-02 15:04:05.0000"), "method:", c.Request.Method, "Status code:", statusCode, "duration nanoseconds:", duration)
	fmt.Println()
}