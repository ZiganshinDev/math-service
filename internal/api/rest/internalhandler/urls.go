package internalhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Urls() map[string]map[string]gin.HandlerFunc {
	return map[string]map[string]gin.HandlerFunc{
		"/problems": {
			http.MethodGet: h.problems,
		},
		"/problems/{id}": {
			http.MethodGet: h.problem,
		},
	}
}
