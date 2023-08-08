package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gitlabnew.nextcontact.ru/nextcontactcenter/services/highload-testers/db-highload-tester/config"
	v1 "gitlabnew.nextcontact.ru/nextcontactcenter/services/highload-testers/db-highload-tester/internal/controller/http/v1"
)

func Run(cfg *config.Config) {

	// Инициализация HTTP-сервера
	handler := gin.New()

	v1.NewRouter(handler, cfg)

	handler.Run(fmt.Sprintf("localhost:%d", cfg.HttpPort.HttpPort))
}
