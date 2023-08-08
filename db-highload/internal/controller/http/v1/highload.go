package v1

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gitlabnew.nextcontact.ru/nextcontactcenter/services/highload-testers/db-highload-tester/config"
	"gitlabnew.nextcontact.ru/nextcontactcenter/services/highload-testers/db-highload-tester/internal/helpers/databases/oracle"
	"gitlabnew.nextcontact.ru/nextcontactcenter/services/highload-testers/db-highload-tester/internal/usecase"
	"gitlabnew.nextcontact.ru/nextcontactcenter/services/highload-testers/db-highload-tester/internal/usecase/repository"
)

type HighloadRouters struct {
	handler *gin.RouterGroup
	cfg     *config.Config
}

func NewHighloadRoutes(handler *gin.RouterGroup, cfg *config.Config) {

	r := HighloadRouters{
		cfg: cfg,
	}

	// Запустить нагрузочное тестирование
	handler.GET("/test", r.Test)
}

func (h *HighloadRouters) Test(c *gin.Context) {

	type TestRequest struct {
		BeginDate    string `form:"date-start" binding:"required"`
		EndDate      string `form:"date-end" binding:"required"`
		SessionCount int    `form:"session-count" binding:"required"`
	}

	var req TestRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Print("req", req)

	// Основная логика нагрузочного тестирования
	if req.SessionCount <= 0 {
		log.Fatal().Msg("количество потоков не может быть меньше или равным нулю")
	}

	var wg sync.WaitGroup

	// Время запуска программы
	startTime := time.Now()
	defer func() {
		log.Info().Msgf("Общее время выполнения программы: %s", time.Now().Sub(startTime).String())
	}()

	// Запуск потоков
	for i := 0; i < int(req.SessionCount); i++ {

		wg.Add(1)

		go func(threadIdx int) {
			defer wg.Done()

			thID := threadIdx + 1

			// Создание нового подключения к Oracle
			conn, err := oracle.GenerateNewConnection(h.cfg.Databases.Oracle)
			if err != nil {
				log.Fatal().Err(err).Msg("не удалось подключиться к базе данных Oracle")
			}
			defer conn.Close()
			// END Создание нового подключения к Oracle

			// Инициализация репозиториев
			kedrReportRepo := repository.NewKedrReportOracle(conn)
			// END Инициализация Oracle репозиториев

			// Инициализация UseCase
			highload := usecase.NewHighload(kedrReportRepo)
			// END Инициализация UseCase

			var tWg sync.WaitGroup

			tWg.Add(3)

			// Запуск нагрузки
			go func() {
				defer tWg.Done()

				s := time.Now()

				log.Debug().Msgf("highload.GetEmployeeTimesheet в потоке №%d запущен", thID)

				if err := highload.GetEmployeeTimesheet(context.Background(), req.BeginDate, req.EndDate); err != nil {
					log.Err(err).Msgf("Во время выполнения highload.GetEmployeeTimesheet в потоке №%d произошла ошибка", thID)
					c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Во время выполнения highload.GetEmployeeTimesheet в потоке №%d произошла ошибка", thID)})
					return
				} else {
					log.Debug().Msgf("highload.GetEmployeeTimesheet в потоке №%d успешно выполнен за: %s", thID, time.Now().Sub(s).String())
				}
			}()
			go func() {
				defer tWg.Done()

				s := time.Now()

				log.Debug().Msgf("highload.GetAPIStatuses в потоке №%d запущен", thID)

				if err := highload.GetAPIStatuses(context.Background(), req.BeginDate, req.EndDate); err != nil {
					log.Err(err).Msgf("Во время выполнения highload.GetAPIStatuses в потоке №%d произошла ошибка", thID)
					c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Во время выполнения highload.GetAPIStatuses в потоке №%d произошла ошибка", thID)})
					return
				} else {
					log.Debug().Msgf("highload.GetAPIStatuses в потоке №%d успешно выполнен за: %s", thID, time.Now().Sub(s).String())
				}
			}()
			go func() {
				defer tWg.Done()

				s := time.Now()

				log.Debug().Msgf("highload.GetDetailCalls в потоке №%d запущен", thID)

				if err := highload.GetDetailCalls(context.Background(), req.BeginDate, req.EndDate); err != nil {
					log.Err(err).Msgf("Во время выполнения highload.GetDetailCalls в потоке №%d произошла ошибка", thID)
					c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Во время выполнения highload.GetDetailCalls в потоке №%d произошла ошибка", thID)})
					return
				} else {
					log.Debug().Msgf("highload.GetDetailCalls в потоке №%d успешно выполнен за: %s", thID, time.Now().Sub(s).String())
				}
			}()
			// END Запуск нагрузки

			tWg.Wait()

			log.Debug().Msgf("Все запросы в потоке №%d выполнены, поток завершен.", thID)
		}(i)
	}

	log.Info().Msg("Все потоки запущены. Ожидаем завершения.")

	wg.Wait()

	log.Info().Msg("Все потоки завершили работу.")
	// END Основная логика нагрузочного тестирования

	c.JSON(http.StatusOK, fmt.Sprintf("Общее время выполнения программы: %s", time.Now().Sub(startTime).String()))
}
