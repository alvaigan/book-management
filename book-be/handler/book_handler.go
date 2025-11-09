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
	"gorm.io/gorm"
)

func (h *Handler) GetBook(c echo.Context) (err error) {
	search := c.QueryParam("search")
	page := c.QueryParam("page")
	rowPerPage := c.QueryParam("row_per_page")
	// publisher_id := c.QueryParam("publisher_id")

	books := []models.Book{}
	query := h.DB.Model(models.Book{})

	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	rowPerPageInt, _ := strconv.Atoi(rowPerPage)
	if page != "" {
		query = query.Limit(rowPerPageInt)
	} else {
		query = query.Limit(10)
	}

	pageInt, _ := strconv.Atoi(page)
	if page != "" {
		offset := 0
		if pageInt > 1 {
			offset = (pageInt * rowPerPageInt) + 1
		}
		query = query.Offset(offset)
	}

	dataCount := int64(0)
	query = query.Joins("Publisher", h.DB.Select([]string{"id", "name", "city"})).Joins("Author", h.DB.Select([]string{"id", "name"})).Debug().Find(&books)
	query = query.Count(&dataCount)
	err = query.Error
	if err != nil {
		h.Log.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, utils.GenerateResErr("Not Found", err))
		} else {
			return c.JSON(http.StatusNotFound, utils.GenerateResErr("Not Found", err))
		}
	}

	totalPage := math.Ceil(float64(dataCount) / float64(rowPerPageInt))
	nextPage := true
	prevPage := true

	if pageInt > int(totalPage) {
		nextPage = false
	}

	if pageInt <= int(totalPage) {
		prevPage = false
	}

	result := dto.PaginationRes{
		Rows:        books,
		TotalRows:   int(dataCount),
		RowPerPage:  rowPerPageInt,
		TotalPage:   int(totalPage),
		HasPrevPage: prevPage,
		HasNextPage: nextPage,
	}

	return c.JSON(http.StatusOK, utils.GenerateRes("Success", result))
}

func (h *Handler) GetBookById(c echo.Context) (err error) {

	return c.JSON(http.StatusOK, utils.GenerateRes("Success", nil))
}

func (h *Handler) CreateBook(c echo.Context) (err error) {
	createPayload := dto.CreateBook{}
	err = c.Bind(&createPayload)
	if err != nil {
		h.Log.Error(err)
		return c.JSON(http.StatusBadRequest, utils.GenerateResErr("Bad Request", err))
	}

	validate := validator.New()
	err = validate.Struct(&createPayload)
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

	book := models.Book{
		Title:       createPayload.Title,
		Description: createPayload.Description,
		PublisherId: uint(createPayload.PublisherId),
	}

	if err := tx.Create(&book).Error; err != nil {
		tx.Rollback()
		h.Log.Error(err)
		return c.JSON(http.StatusInternalServerError, utils.GenerateResErr("Internal Server Error", err))
	}

	if err := tx.Model(models.Author{}).Where(models.Author{ID: uint(createPayload.AuthorId)}).Update("book_id", book.ID).Error; err != nil {
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

	data := book

	return c.JSON(http.StatusOK, utils.GenerateRes("Book Created", data))
}

func (h *Handler) UpdateBook(c echo.Context) (err error) {
	createPayload := dto.CreateBook{}
	err = c.Bind(&createPayload)
	if err != nil {
		h.Log.Error(err)
		return c.JSON(http.StatusBadRequest, utils.GenerateResErr("Bad Request", err))
	}

	validate := validator.New()
	err = validate.Struct(&createPayload)
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

	if err := tx.Create(&createPayload).Error; err != nil {
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

	data := createPayload

	return c.JSON(http.StatusOK, utils.GenerateRes("Book updated", data))
}

func (h *Handler) DeleteBook(c echo.Context) (err error) {

	return c.JSON(http.StatusOK, utils.GenerateRes("Book deleted", nil))
}
