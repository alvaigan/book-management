package middleware

import (
	"book-be/models"
	"book-be/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthMiddleware struct {
	DB *gorm.DB
}

func NewAuthMiddleware(db *gorm.DB) *AuthMiddleware {
	return &AuthMiddleware{
		DB: db,
	}
}

func (a *AuthMiddleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		header := c.Request().Header
		if header["Authorization"] != nil {
			var token string
			for _, val := range header["Authorization"] {
				splitVal := strings.Split(val, " ")
				if len(splitVal) > 1 {
					token = splitVal[1]
					break
				} else {
					return c.JSON(http.StatusUnauthorized, utils.GenerateResErr("Unauthorized", nil))
				}
			}

			isValid, userClaims, err := utils.ValidateToken(token)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, utils.GenerateResErr("Internal Server Error", err))
			}

			if !isValid {
				return c.JSON(http.StatusUnauthorized, utils.GenerateResErr("Unauthorized", nil))
			}

			if userClaims == nil {
				return c.JSON(http.StatusUnauthorized, utils.GenerateResErr("Unauthorized", nil))
			}

			user := models.User{}
			if err := a.DB.Where(models.User{
				ID:       userClaims.ID,
				Username: userClaims.Username,
			}).First(&user).Error; err != nil {
				return c.JSON(http.StatusUnauthorized, utils.GenerateResErr("Unauthorized", nil))
			}

			c.Set("user", user)
			c.Set("token", token)
			return next(c)
		} else {
			return c.JSON(http.StatusUnauthorized, utils.GenerateResErr("Unauthorized", nil))
		}
	}
}
