package handlers

import (
	" hery-ciaputra/final-project-backend/dto"
	" hery-ciaputra/final-project-backend/httperror"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) Register(c *gin.Context) {
	value, _ := c.Get("payload")
	registerReq := value.(*dto.RegisterReq)

	result, err := h.registerService.Register(registerReq)
	if err != nil {
		_ = c.Error(httperror.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result)
}
