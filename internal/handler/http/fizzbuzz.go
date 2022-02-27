package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/TestardR/fizz-buzz/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var (
	errInvalidQuery = errors.New("failed to validate input query string")
	errStorage      = errors.New("failed to perform storage operation")
)

// GetFizzbuzz
// @Summary Return fizzbuzz result
// @Description Fizzbuzz operation
// @Tags Fizzbuzz Operation
// @Produce json
// @Param int1 query int true "int1 query parameter"
// @Param int2 query int true "int2 query parameter"
// @Param limit query int true "limit query parameter"
// @Param str1 query string true "str1 query parameter"
// @Param str2 query string true "str2 query parameter"
// @Success 200 {object} ResponseSuccess{data=[]string}
// @Failure 400 {object} ResponseError
// @Router /fizzbuzz [get].
func (h *handler) GetFizzbuzz(c *gin.Context) {
	var input domain.Fizzbuzz

	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, newResponseError(fmt.Errorf("%w: %s", errInvalidQuery, err)))

		return
	}

	h.log.Info(fmt.Sprintf("request received with payload: %v", input))

	if err := validator.New().Struct(&input); err != nil {
		c.JSON(http.StatusBadRequest, newResponseError(fmt.Errorf("%w: %s", errInvalidQuery, err)))

		return
	}

	result := input.ComputeSequence()

	err := h.store.IncrementCount(c.Request.Context(), input.ToString())
	if err != nil {
		h.log.Error(fmt.Errorf("%w: %s", errStorage, err))
	}

	c.JSON(http.StatusOK, ResponseSuccess{Data: result})
}
