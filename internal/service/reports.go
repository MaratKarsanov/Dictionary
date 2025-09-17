package service

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// localhost:8000/api/report/get/:id
func (s *Service) GetReport(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	reportsRepository := s.reportsRepository
	report, err := reportsRepository.GetReport(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, Response{Object: report})
}

// localhost:8000/api/report/create
func (s *Service) CreateReport(c echo.Context) error {
	var report Report
	err := c.Bind(&report)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	reportsRepository := s.reportsRepository
	err = reportsRepository.CreateReport(report.Title, report.Description)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "OK")
}

// localhost:8000/report/update/:id
func (s *Service) UpdateReport(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	var report Report
	err = c.Bind(&report)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	reportsRepository := s.reportsRepository
	err = reportsRepository.UpdateReport(id, report.Title, report.Description)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "OK")
}

// localhost:8000/report/delete/:id
func (s *Service) DeleteReport(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	reportsRepository := s.reportsRepository
	err = reportsRepository.DeleteReport(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "OK")
}
