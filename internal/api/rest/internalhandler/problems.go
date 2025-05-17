package internalhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) problems(c *gin.Context) {
	problems, err := h.app.Problems(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	res := make([]Problem, 0, len(problems))
	for _, problem := range problems {
		res = append(res, problemModelToResp(problem))
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) problem(c *gin.Context) {
	problem, err := h.app.Problem(c.Request.Context(), uuid.UUID{})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, problemModelToResp(problem))
}
