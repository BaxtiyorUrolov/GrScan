package handler

import (
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

	// Check if login already exists
	if exists, err := check.IsLoginExist(createUser.Login, h.storage.User()); err != nil {
		handleResponse(c, "Error while checking login existence", http.StatusInternalServerError, err)
		return
	} else if exists {
		handleResponse(c, "Login already exists", http.StatusBadRequest, "This login already exists") 
		return
	}

	// Generate verification code
	code := check.GenerateVerificationCode()

	// Send verification code via SMS
	if err := check.Send(createUser.Phone, "+998333346767", code); err != nil {
		handleResponse(c, "Error while sending verification code", http.StatusInternalServerError, err)
		return
	}

	// Save verification code to session or database for later verification

	// Proceed with user registration
	pKey, err := h.storage.User().Create(createUser)
	if err != nil {
		handleResponse(c, "Error while creating user", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "User created successfully", http.StatusCreated, pKey)
}
