package handler

import (
	"context"
	"grscan/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Router       /user [POST]
// @Summary      Creates a new user
// @Description  create a new user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user body models.CreateUser false "user"
// @Success      201  {object}  models.User
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateUser(c *gin.Context) {
	createUser := models.CreateUser{}
	if err := c.ShouldBindJSON(&createUser); err != nil {
		handleResponse(c, h.log, "Error while reading body from client", http.StatusBadRequest, err)
		return
	}

	resp, err := h.services.User().Create(context.Background(), createUser)
	if err != nil {
		handleResponse(c, h.log, "error is while creating category", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusCreated, resp)
}
