package utils

import (
	"book-be/models"
	"book-be/types"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateRes(message string, data any) *types.ResponseSuccess {

	return &types.ResponseSuccess{
		Status:  "success",
		Message: message,
		Data:    &data,
	}
}

func GenerateResErr(message string, error any) *types.ResponseError {
	return &types.ResponseError{
		Status:  "error",
		Message: message,
		Error:   &error,
	}
}

func GetValidationErrorMsg(ve validator.ValidationErrors) map[string][]string {
	msg := make(map[string][]string)
	for _, err := range ve {
		fieldName := err.Field()
		msg[fieldName] = append(msg[fieldName], MakeValidationErrorMsg(fieldName, err.Tag()))
	}

	return msg
}

func MakeValidationErrorMsg(fieldName string, tag string) (msg string) {
	switch tag {
	case "required":
		msg = fmt.Sprintf("%s could not be empty", fieldName)
	case "email":
		msg = fmt.Sprintf("%s is not valid email format", fieldName)
	case "min":
		msg = fmt.Sprintf("%s should be minimal 6 characters", fieldName)
	case "numeric":
		msg = fmt.Sprintf("%s should contain numeric characters", fieldName)
	case "alpha":
		msg = fmt.Sprintf("%s should contain alphabetic characters", fieldName)
	}

	return msg
}

func GenerateToken(user models.User) (token string, err error) {
	secret := viper.GetString("jwt.secret")
	signKey := []byte(secret)
	claims := models.UserClaims{
		ID:       user.ID,
		Username: user.Username,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = t.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateToken(token string) (bool, *models.UserClaims, error) {
	secret := viper.GetString("jwt.secret")
	signKey := []byte(secret)

	t, err := jwt.ParseWithClaims(token, &models.UserClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(signKey), nil
	})

	if err != nil {
		logrus.Error(err)
		return false, nil, err
	} else if claims, ok := t.Claims.(*models.UserClaims); ok {
		return true, claims, nil
	} else {
		logrus.Error("unknown claims type, cannot proceed")
		return false, nil, err
	}
}
