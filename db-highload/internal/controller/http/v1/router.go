package v1

import (
	"github.com/gin-gonic/gin"
	"gitlabnew.nextcontact.ru/nextcontactcenter/services/highload-testers/db-highload-tester/config"
)

func NewRouter(router *gin.Engine, cfg *config.Config) {

	h := router.Group("/v1")
	{
		NewHighloadRoutes(h, cfg)
	}
}
