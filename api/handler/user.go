package handler

import (
	"context"
	"grscan/api/models"
	"grscan/pkg/check"
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

	if !check.PhoneNumber(createUser.Phone) {
		handleResponse(c, h.log, "Incorrect phone number", http.StatusBadRequest, nil)
		return
	}

	if !check.ValidatePassword(createUser.Password) {
		handleResponse(c, h.log, "Invalid password", http.StatusBadRequest, nil)
		return
	}

	exists, err := check.IsLoginExist(createUser.Login, h.storage.User())
    if err != nil {
        handleResponse(c, h.log, "Error while checking login existence", http.StatusInternalServerError, nil)
        return
    }
    if exists {
        handleResponse(c, h.log, "Login already exists", http.StatusBadRequest, "This login already exists")
        return
    }

	resp, err := h.services.User().Create(context.Background(), createUser)
	if err != nil {
		handleResponse(c, h.log, "Error while creating user", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.log, "User created successfully", http.StatusCreated, resp)
}
