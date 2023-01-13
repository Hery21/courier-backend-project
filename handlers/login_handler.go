package handlers

import (
	" hery-ciaputra/final-project-backend/dto"
	" hery-ciaputra/final-project-backend/httperror"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) LogIn(c *gin.Context) {
	value, _ := c.Get("payload")
	logInReq := value.(*dto.LogInReq)

	result, err := h.authService.LogIn(logInReq)
	if err != nil {
		_ = c.Error(httperror.InternalServerError("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, result)
}
