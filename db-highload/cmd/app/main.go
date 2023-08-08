package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlabnew.nextcontact.ru/nextcontactcenter/services/highload-testers/db-highload-tester/config"
	"gitlabnew.nextcontact.ru/nextcontactcenter/services/highload-testers/db-highload-tester/internal/app"
)

const envPrefix = "dbhlt"

func main() {
	// Настройка логгера
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "02.01.2006 15:04:05",
	}
	log.Logger = log.Output(output)
	log.With().Timestamp().Logger()
	// END Настройка логгера

	// Загрузка файла конфигурации
	cfg := config.Load(envPrefix)
	// END Загрузка файла конфигурации

	// Запуск приложения
	log.Info().Msgf("Конфиги: %v", cfg)
	app.Run(cfg)
	// END Запуск приложения
}
