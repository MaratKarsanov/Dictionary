package service

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetWordById ищем слово по id
// localhost:8000/api/word/:id
func (s *Service) GetWordById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.wordsRepo
	word, err := repo.RGetWordById(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, Response{Object: word})
}

// CreateWords добавляем в базу новые слова в базу
// localhost:8000/api/words
func (s *Service) CreateWords(c echo.Context) error {
	var wordSlice []Word
	err := c.Bind(&wordSlice)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.wordsRepo
	for _, word := range wordSlice {
		err = repo.CreateNewWords(word.Title, word.Translation)
	}
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "OK")
}

// UpdateWord обновляем слово в базе
// localhost:8000/updateword/:id
func (s *Service) UpdateWord(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	var word Word
	err = c.Bind(&word)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.wordsRepo
	err = repo.UpdateWord(id, word.Title, word.Translation)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "OK")
}

// DeleteWord удаляем слово из базы
// localhost:8000/deleteword/:id
func (s *Service) DeleteWord(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.wordsRepo
	err = repo.DeleteWordById(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "OK")
}
