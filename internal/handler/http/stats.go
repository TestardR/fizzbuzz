package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const templateStatsResponse = `The most popular fizzbuzz request with %d hits is [%s]`

// @Summary Retrieve statistic regarding fizzbuzz request
// @Description Returns the query that was performed the most and also the number of times
// @Tags Statistic
// @Produce json
// @Success 200 {object} ResponseSuccess
// @Success 204
// @Failure 500,400 {object} ResponseError
// @Router /stats [get].
func (h *handler) GetStats(c *gin.Context) {
	key, hits, err := h.store.GetMaxEntries(c.Request.Context())
	if err != nil {
		err = fmt.Errorf("%w: %s", errStorage, err)
		h.log.Error(err)
		c.JSON(http.StatusInternalServerError, newResponseError(err))

		return
	}

	if hits == 0 {
		c.Status(http.StatusNoContent)

		return
	}

	c.JSON(http.StatusOK, ResponseSuccess{Data: fmt.Sprintf(templateStatsResponse, hits, key)})
}

// @Summary Reset statistics
// @Description Reset all statistics regarding the queries made in the past
// @Tags Statistic
// @Success 200
// @Failure 500 {object} ResponseError
// @Router /stats [delete].
func (h *handler) DeleteStats(c *gin.Context) {
	if err := h.store.Reset(c.Request.Context()); err != nil {
		err = fmt.Errorf("%w: %s", errStorage, err)
		h.log.Error(err)
		c.JSON(http.StatusInternalServerError, newResponseError(err))

		return
	}

	c.Status(http.StatusOK)
}
