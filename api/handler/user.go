package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"grscan/api/models"
	"grscan/pkg/check"
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
		handleResponse(c, "Error while reading body from client", http.StatusBadRequest, err)
		return
	}

	if !check.PhoneNumber(createUser.Phone) {
		handleResponse(c, "Incorrect phone number", http.StatusBadRequest, nil)
		return
	}

	if err := check.ValidatePassword(createUser.Password); err != nil {
		handleResponse(c, "Invalid password", http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.User().Create(createUser)
	if err != nil {
		handleResponse(c, "Error while creating user", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "User created successfully", http.StatusCreated, pKey)
}
