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

type PublisherHandler struct {
	App   *echo.Echo
	DB    *gorm.DB
	Viper *viper.Viper
	Log   *logrus.Logger
}

func NewPublisherHandler(app *echo.Echo, db *gorm.DB, viper *viper.Viper, log *logrus.Logger) *PublisherHandler {
	return &PublisherHandler{
		App:   app,
		DB:    db,
		Viper: viper,
		Log:   log,
	}
}

func (h *PublisherHandler) GetPublisher(c echo.Context) (err error) {
	search := c.QueryParam("search")
	page := c.QueryParam("page")
	rowPerPage := c.QueryParam("row_per_page")

	publisher := []models.Publisher{}
	query := h.DB.Preload("Books", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Author", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"id", "name"})
		}).Select([]string{"id", "title", "description", "author_id", "publisher_id"})
	}).Model(models.Publisher{})

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
		query = query.Where("city LIKE ?", "%"+search+"%")
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

	query = query.Debug().Find(&publisher)
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
		Rows:        publisher,
		TotalRows:   int(dataCount),
		RowPerPage:  rowPerPageInt,
		TotalPage:   int(totalPage),
		HasPrevPage: prevPage,
		HasNextPage: nextPage,
	}

	return c.JSON(http.StatusOK, utils.GenerateRes("Success", result))
}

func (h *PublisherHandler) GetPublisherById(c echo.Context) (err error) {
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.GenerateResErr("Internal Server Error", err))
	}

	authorDetail := models.Publisher{}

	err = h.DB.Preload("Books", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Author", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"id", "name"})
		}).Select([]string{"id", "title", "description", "author_id", "publisher_id"})
	}).Where(&models.Publisher{ID: uint(id)}).First(&authorDetail).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, utils.GenerateResErr("Data not found!", err))
		} else {
			return c.JSON(http.StatusInternalServerError, utils.GenerateResErr("Internal Server Error", err))
		}
	}

	return c.JSON(http.StatusOK, utils.GenerateRes("Success", authorDetail))
}

func (h *PublisherHandler) CreatePublisher(c echo.Context) (err error) {
	createPublisherPayload := dto.CreatePublisher{}
	err = c.Bind(&createPublisherPayload)
	if err != nil {
		h.Log.Error(err)
		return c.JSON(http.StatusBadRequest, utils.GenerateResErr("Bad Request", err))
	}

	validate := validator.New()
	err = validate.Struct(&createPublisherPayload)
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

	author := models.Publisher{
		Name: createPublisherPayload.Name,
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

	return c.JSON(http.StatusOK, utils.GenerateRes("Publisher Created", data))
}

func (h *PublisherHandler) UpdatePublisher(c echo.Context) (err error) {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, utils.GenerateResErr("No one selected book", err))
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.Log.Error(err)
		return c.JSON(http.StatusInternalServerError, utils.GenerateResErr("Internal Server Error", err))
	}

	authorPayload := dto.UpdatePublisher{}
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

	newPublisher := models.Publisher{
		Name: authorPayload.Name,
	}

	if err := tx.Where(models.Publisher{
		ID: uint(idInt),
	}).Updates(&newPublisher).Error; err != nil {
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

	return c.JSON(http.StatusOK, utils.GenerateRes("Publisher updated", nil))
}

func (h *PublisherHandler) DeletePublisher(c echo.Context) (err error) {
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

	if err := h.DB.Where(models.Publisher{ID: uint(idInt)}).Delete(&models.Publisher{}).Error; err != nil {
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
