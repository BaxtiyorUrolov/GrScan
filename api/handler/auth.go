package handler

import (
	"context"
	"grscan/api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CustomerLogin godoc
// @Router       /auth/customer/login [POST]
// @Summary      Customer login
// @Description  Customer login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login body models.CustomerLoginRequest false "login"
// @Success      201  {object}  models.User
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CustomerLogin(c *gin.Context) {
	loginRequest := models.CustomerLoginRequest{}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		handleResponse(c, h.log, "Error while binding body", http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := h.services.Auth().CustomerLogin(ctx, loginRequest); err != nil {
		handleResponse(c, h.log, "Incorrect password", http.StatusBadRequest, "Incorrect password")
		return
	}

	handleResponse(c, h.log, "Success", http.StatusOK, "")
}

