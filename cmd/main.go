package main

import (
	"dictionary/internal/service"
	"dictionary/pkg/logs"

	"github.com/labstack/echo/v4"
)

func main() {
	// создаем логгер
	logger := logs.NewLogger(false)

	// подключаемся к базе
	db, err := PostgresConnection()
	if err != nil {
		logger.Fatal(err)
	}

	svc := service.NewService(db, logger)

	router := echo.New()
	// создаем группу api
	api := router.Group("api")

	// прописываем пути
	api.GET("/word/:id", svc.GetWordById)
	api.POST("/words", svc.CreateWords)
	api.PUT("/update/:id", svc.UpdateWord)
	api.DELETE("/delete/:id", svc.DeleteWord)

	api.GET("/report/get/:id", svc.GetReport)
	api.POST("/report/create", svc.CreateReport)
	api.PUT("/report/update/:id", svc.UpdateReport)
	api.DELETE("/report/delete/:id", svc.DeleteReport)

	// запускаем сервер, чтобы слушал 8000 порт
	router.Logger.Fatal(router.Start(":8000"))
}
