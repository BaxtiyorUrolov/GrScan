package handler

import (
	"context"
	"fmt"
	"grscan/api/models"
	"grscan/pkg/check"
	"grscan/pkg/sms"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

// Global o'zgaruvchi uchun kesh obyekti
var cacheStore = cache.New(5*time.Minute, 10*time.Minute)

// Kod generatsiyasi funksiyasi
func generateCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// Kod yuborish funksiyasi (bu joyda faqat konsolga chiqariladi, aslida SMS xizmatidan foydalanishingiz kerak)
func sendCode(phoneNumber string, code string) {
	fmt.Printf("SMS kod: %s raqamiga yuborildi: %s\n", phoneNumber, code)
}

// CreateUser godoc
// @Router       /user [POST]
// @Summary      Creates a new user
// @Description  create a new user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user body models.CreateUserRequest false "user"
// @Success      201  {object}  models.User
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) CreateUser(c *gin.Context) {
	createUser := models.CreateUserRequest{}
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

	// SMS kod generatsiya qilish va yuborish
	code := generateCode()
	cacheStore.Set(createUser.Phone, code, cache.DefaultExpiration)
	cacheStore.Set(createUser.Phone+"_data", createUser, cache.DefaultExpiration)
	if err := sms.Send(createUser.Phone, code); err != nil {
		handleResponse(c, h.log, "Error while sending SMS code", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.log, "SMS kod yuborildi", http.StatusCreated, gin.H{
		"message": "SMS code sent, please verify",
		"phone":   createUser.Phone,
	})
}

// VerifyRegister godoc
// @Router       /verify-register [POST]
// @Summary      Verifies the SMS code
// @Description  verify the SMS code sent to user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        verification body models.VerifyCodeRequest false "verification"
// @Success      200  {object}  models.User
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) VerifyRegister(c *gin.Context) {
	verification := models.VerifyCodeRequest{}
	if err := c.ShouldBindJSON(&verification); err != nil {
		handleResponse(c, h.log, "Error while reading body from client", http.StatusBadRequest, err)
		return
	}

	code, found := cacheStore.Get(verification.Phone)
	if !found {
		handleResponse(c, h.log, "Verification code not found or expired", http.StatusBadRequest, gin.H{
			"error": "The verification code has expired or does not exist",
		})
		return
	}

	if code != verification.Code {
		handleResponse(c, h.log, "Invalid verification code", http.StatusBadRequest, gin.H{
			"error": "The provided verification code is incorrect",
		})
		return
	}

	// Avval saqlangan ma'lumotlarni olish
	userRequest, found := cacheStore.Get(verification.Phone + "_data")
	if !found {
		handleResponse(c, h.log, "User data not found", http.StatusInternalServerError, nil)
		return
	}

	createUser := userRequest.(models.CreateUserRequest)

	user, err := h.services.User().Create(context.Background(), models.CreateUser{
		Phone:    createUser.Phone,
		Login:    createUser.Login,
		Password: createUser.Password,
		UserType: createUser.UserType,
	})
	if err != nil {
		handleResponse(c, h.log, "Error while creating user", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.log, "User created successfully", http.StatusCreated, user)
}


