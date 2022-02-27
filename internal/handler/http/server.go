package http

import (
	_ "github.com/TestardR/fizz-buzz/docs" // swagger documentation
	"github.com/TestardR/fizz-buzz/internal/storage/redis"
	"github.com/TestardR/fizz-buzz/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	healthRoute   = "/health"
	fizzbuzzRoute = "/fizzbuzz"
	statsRoute    = "/stats"
)

// Handler is the base struct for dependency injection.
type handler struct {
	log   logger.Logger
	store redis.Storager
}

// @title Fizzbuzz Rest Server
// @version 1.0
// @description This is a fizzbuz server with a statistic endpoint

// @contact.name Romain Testard
// @contact.email romain.rtestard@gmail.com

// @host localhost:3000
func NewServer(env string, log logger.Logger, store redis.Storager) *gin.Engine {
	h := handler{
		log:   log,
		store: store,
	}

	gin.SetMode(env)

	router := gin.New()
	router.Use(gin.Recovery())

	// swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// useful for monitoring our service and CI/CD tools
	router.GET(healthRoute, h.Health)

	router.GET(fizzbuzzRoute, h.GetFizzbuzz)

	router.GET(statsRoute, h.GetStats)
	router.DELETE(statsRoute, h.DeleteStats)

	return router
}
