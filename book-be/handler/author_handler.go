package handler

import (
	"book-be/dto"
	"book-be/models"
	"book-be/utils"
	"errors"
	"math"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AuthorHandler struct {
	App   *echo.Echo
	DB    *gorm.DB
	Viper *viper.Viper
	Log   *logrus.Logger
}

func NewAuthorHandler(app *echo.Echo, db *gorm.DB, viper *viper.Viper, log *logrus.Logger) *AuthorHandler {
	return &AuthorHandler{
		App:   app,
		DB:    db,
		Viper: viper,
		Log:   log,
	}
}

func (h *AuthorHandler) GetAuthor(c echo.Context) (err error) {
	search := c.QueryParam("search")
	page := c.QueryParam("page")
	rowPerPage := c.QueryParam("row_per_page")

	authors := []models.Author{}
	query := h.DB.Preload("Books", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Publisher", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"id", "name", "city"})
		}).Select([]string{"id", "title", "description", "author_id", "publisher_id"})
	}).Model(models.Author{})

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	dataCount := int64(0)

	rowPerPageInt, _ := strconv.Atoi(rowPerPage)
	if page != "" {
		query = query.Limit(rowPerPageInt)
	} else {
		query = query.Limit(10)
	}
	query = query.Count(&dataCount)
	totalPage := math.Ceil(float64(dataCount) / float64(rowPerPageInt))
	pageInt, _ := strconv.Atoi(page)
	if page != "" {
		offset := 0
		if pageInt > 1 {
			offset = (pageInt * rowPerPageInt) - 1
		}
		query = query.Offset(offset)
	}

	query = query.Find(&authors)
	err = query.Error
	if err != nil {
		h.Log.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, utils.GenerateResErr("Not Found", err))
		} else {
			return c.JSON(http.StatusNotFound, utils.GenerateResErr("Not Found", err))
		}
	}

	nextPage := true
	prevPage := true

	if pageInt >= int(totalPage) {
		nextPage = false
	}

	if pageInt == 1 {
		prevPage = false
	}

	result := dto.PaginationRes{
		Rows:        authors,
		TotalRows:   int(dataCount),
		RowPerPage:  rowPerPageInt,
		TotalPage:   int(totalPage),
		HasPrevPage: prevPage,
		HasNextPage: nextPage,
	}

	return c.JSON(http.StatusOK, utils.GenerateRes("Success", result))
}

func (h *AuthorHandler) GetAuthorById(c echo.Context) (err error) {
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.GenerateResErr("Internal Server Error", err))
	}

	authorDetail := models.Author{}

	err = h.DB.Preload("Books", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Publisher", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"id", "name", "city"})
		}).Select([]string{"id", "title", "description", "author_id", "publisher_id"})
	}).Where(&models.Author{ID: uint(id)}).First(&authorDetail).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, utils.GenerateResErr("Data not found!", err))
		} else {
			return c.JSON(http.StatusInternalServerError, utils.GenerateResErr("Internal Server Error", err))
		}
	}

	return c.JSON(http.StatusOK, utils.GenerateRes("Success", authorDetail))
}

func (h *AuthorHandler) CreateAuthor(c echo.Context) (err error) {
	createAuthorPayload := dto.CreateAuthor{}
	err = c.Bind(&createAuthorPayload)
	if err != nil {
		h.Log.Error(err)
		return c.JSON(http.StatusBadRequest, utils.GenerateResErr("Bad Request", err))
	}

	validate := validator.New()
	err = validate.Struct(&createAuthorPayload)
	if err != nil {
		h.Log.Error(err)
		validationErrors := utils.GetValidationErrorMsg(err.(validator.ValidationErrors))
		return c.JSON(http.StatusBadRequest, utils.GenerateResErr("Validation Error", validationErrors))
	}

	tx := h.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	author := models.Author{
		Name: createAuthorPayload.Name,
	}

	if err := tx.Create(&author).Error; err != nil {
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

	data := author

	return c.JSON(http.StatusOK, utils.GenerateRes("Author Created", data))
}

func (h *AuthorHandler) UpdateAuthor(c echo.Context) (err error) {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, utils.GenerateResErr("No one selected book", err))
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.Log.Error(err)
		return c.JSON(http.StatusInternalServerError, utils.GenerateResErr("Internal Server Error", err))
	}

	authorPayload := dto.UpdateAuthor{}
	err = c.Bind(&authorPayload)
	if err != nil {
		h.Log.Error(err)
		return c.JSON(http.StatusBadRequest, utils.GenerateResErr("Bad Request", err))
	}

	validate := validator.New()
	err = validate.Struct(&authorPayload)
	if err != nil {
		h.Log.Error(err)
		validationErrors := utils.GetValidationErrorMsg(err.(validator.ValidationErrors))
		return c.JSON(http.StatusBadRequest, utils.GenerateResErr("Validation Error", validationErrors))
	}

	tx := h.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	newAuthor := models.Author{
		Name: authorPayload.Name,
	}

	if err := tx.Where(models.Author{
		ID: uint(idInt),
	}).Updates(&newAuthor).Error; err != nil {
		tx.Rollback()
		h.Log.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, utils.GenerateResErr("Data not found", err))
		} else {
			return c.JSON(http.StatusInternalServerError, utils.GenerateResErr("Internal Server Error", err))
		}
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		h.Log.Error(err)
		return c.JSON(http.StatusInternalServerError, utils.GenerateResErr("Internal Server Error", err))
	}

	return c.JSON(http.StatusOK, utils.GenerateRes("Author updated", nil))
}

func (h *AuthorHandler) DeleteAuthor(c echo.Context) (err error) {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, utils.GenerateResErr("No one selected book", err))
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.Log.Error(err)
		return c.JSON(http.StatusInternalServerError, utils.GenerateResErr("Internal Server Error", err))
	}

	tx := h.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := h.DB.Where(models.Author{ID: uint(idInt)}).Delete(&models.Author{}).Error; err != nil {
		tx.Rollback()
		h.Log.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, utils.GenerateResErr("Data not found", err))
		} else {
			return c.JSON(http.StatusInternalServerError, utils.GenerateResErr("Internal Server Error", err))
		}
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		h.Log.Error(err)
		return c.JSON(http.StatusInternalServerError, utils.GenerateResErr("Internal Server Error", err))
	}

	return c.JSON(http.StatusOK, utils.GenerateRes("Book deleted", nil))
}
