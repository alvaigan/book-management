package handler

import (
	"book-be/dto"
	"book-be/models"
	"book-be/utils"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (h *Handler) Login(c echo.Context) (err error) {
	loginPayload := dto.Login{}
	err = c.Bind(&loginPayload)
	if err != nil {
		h.Log.Error(err)
		return c.JSON(http.StatusBadRequest, utils.GenerateResErr("Bad Request", err))
	}

	validate := validator.New()
	err = validate.Struct(&loginPayload)
	if err != nil {
		h.Log.Error(err)
		validationErrors := utils.GetValidationErrorMsg(err.(validator.ValidationErrors))
		return c.JSON(http.StatusBadRequest, utils.GenerateResErr("Validation Error", validationErrors))
	}

	user := models.User{}
	if err := h.DB.Where(&models.User{
		Username: loginPayload.Username,
	}).First(&user).Error; err != nil {
		h.Log.Error(err)
		return c.JSON(http.StatusBadRequest, utils.GenerateResErr("Username and password are not valid!", nil))
	}

	isValidPassword := utils.CheckPasswordHash(loginPayload.Password, user.Password)
	if !isValidPassword {
		h.Log.Error("Invalid Password")
		return c.JSON(http.StatusBadRequest, utils.GenerateResErr("Username and password are not valid!", nil))
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		h.Log.Error(err)
		return c.JSON(http.StatusBadRequest, utils.GenerateResErr("Username and password are not valid!", nil))
	}

	data := map[string]interface{}{
		"token": token,
	}

	return c.JSON(http.StatusOK, utils.GenerateRes("Login succeed", data))
}

func (h *Handler) Register(c echo.Context) (err error) {
	registerPayload := dto.RegisterReq{}
	err = c.Bind(&registerPayload)
	if err != nil {
		h.Log.Error(err)
		return c.JSON(http.StatusBadRequest, utils.GenerateResErr("Bad Request", err))
	}

	validate := validator.New()
	err = validate.Struct(&registerPayload)
	if err != nil {
		h.Log.Error(err)
		validationErrors := utils.GetValidationErrorMsg(err.(validator.ValidationErrors))
		return c.JSON(http.StatusBadRequest, utils.GenerateResErr("Validation Error", validationErrors))
	}

	user := models.User{}
	userCheck := h.DB.Where(&models.User{
		Username: registerPayload.Username,
	}).First(&user)
	err = userCheck.Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		h.Log.Error(err)
		return c.JSON(http.StatusInternalServerError, utils.GenerateResErr("Internal Server Error", err))
	}

	hashPassword, err := utils.HashPassword(registerPayload.Password)

	user = models.User{
		Username: registerPayload.Username,
		Password: hashPassword,
	}

	tx := h.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		h.Log.Error(err)
		return c.JSON(http.StatusInternalServerError, utils.GenerateResErr("Internal Server Error", err))
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		h.Log.Error(err)
		return c.JSON(http.StatusInternalServerError, utils.GenerateResErr("Internal Server Error", err))
	}

	return c.JSON(http.StatusOK, utils.GenerateRes("Registration succeed", nil))
}

func (h *Handler) Logout(c echo.Context) (err error) {
	fmt.Println(c.Get("user"))

	return nil
}
