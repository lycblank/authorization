package interfaces

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type HttpService struct {
}

type HttpServiceArg struct {
	dig.In
	Engine *gin.Engine
}

func (hs *HttpService) Run(args HttpServiceArg) {
	v1 := args.Engine.Group("/v1")
	v1.POST("/account", hs.Register)
	v1.POST("/account/login", hs.Login)
	v1.POST("/account/logout", hs.Logout)
	v1.PUT("/account/token", hs.RefreshToken)

	v1.POST("/tenant", hs.TenantRegister)
}
