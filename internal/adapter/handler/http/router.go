package http

import (
	"strings"

	"github.com/davifrjose/My_Turn/internal/adapter/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	config *config.HTTP,
	serviceTypeHandler ServiceTypeHandler,
) (*Router, error) {
	if config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	ginConfig := cors.DefaultConfig()
	allowedOrigins := config.AllowedOrigins
	originsList := strings.Split(allowedOrigins, ",")
	ginConfig.AllowOrigins = originsList

	router := gin.New()

	v1 := router.Group("/v1")
	{
		serviceType := v1.Group("/service-type")
		{
			serviceType.POST("/", serviceTypeHandler.CreateServiceType)
		}
	}

	return &Router{
		router,
	}, nil
}

func (router *Router) Serve(listenAddress string) error {
	return router.Run(listenAddress)
}
