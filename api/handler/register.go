package handler

import (
	"context"
	"grscan/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// VerifyRegister godoc
// @Router       /verify-register [POST]
// @Summary      Creates a new verify register
// @Description  create a new verify register
// @Tags         verify register
// @Accept       json
// @Produce      json
// @Param        verify body models.CreateRegister false "verify"
// @Success      201  {object}  models.Register
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) VerifyRegister(c *gin.Context) {
	createUser := models.CreateRegister{}
	if err := c.ShouldBindJSON(&createUser); err != nil {
		handleResponse(c, h.log, "Error while reading body from client", http.StatusBadRequest, err)
		return
	}

	resp, err := h.services.Register().Verify(context.Background(), createUser)
	if err != nil {
		handleResponse(c, h.log, "Error while creating user", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.log, "User created successfully", http.StatusCreated, resp)
}
