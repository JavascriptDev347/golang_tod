package controllers

import (
	"net/http"
	"strconv"
	"to-do/dto"
	"to-do/services"
	"to-do/utils"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service services.CategoryService
}

func NewCategoryController(service services.CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

// ==================== CREATE ====================
// c.ShouldBind(&req)      // multipart/form-data — image bilan ishlaydi
// c.ShouldBindJSON(&req)  // application/json — faqat JSON

// @Summary     Category yaratish
// @Tags        categories
// @Accept      multipart/form-data
// @Produce     json
// @Security    BearerAuth
// @Param       name formData string true "Category nomi"
// @Param       description formData string false "Tavsif"
// @Param       color formData string false "Rang"
// @Param       image formData file false "Rasm"
// @Success     201 {object} utils.Response
// @Failure     400 {object} utils.Response
// @Failure     401 {object} utils.Response
// @Failure     403 {object} utils.Response
// @Router      /categories [post]
func (ctrl *CategoryController) Create(c *gin.Context) {

	var req dto.CreateCategoryRequest

	if err := c.ShouldBind(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
	}

	// Image olish (majburiy emas)
	fileHeader, err := c.FormFile("image")
	if err != nil {
		// Image yuborilmagan — nil beramiz
		response, err := ctrl.service.Create(&req, nil, nil)
		if err != nil {
			utils.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		utils.Success(c, http.StatusCreated, "Category muvaffaqiyatli yaratildi", response)
		return
	}
	// Image yuborilgan
	openedFile, err := fileHeader.Open()
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Imageni ochishda xatolik")
		return
	}
	defer openedFile.Close()

	response, err := ctrl.service.Create(&req, openedFile, fileHeader)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, http.StatusCreated, "Category muvaffaqiyatli yaratildi", response)
}

// @Summary     Barcha categorylar
// @Tags        categories
// @Produce     json
// @Param       search query string false "Qidirish"
// @Param       sort_by query string false "Saralash (name, created_at)"
// @Param       sort_dir query string false "Tartib (asc, desc)"
// @Success     200 {object} utils.Response
// @Router      /categories [get]
func (ctrl *CategoryController) FindAll(c *gin.Context) {
	var filter dto.CategoryFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := ctrl.service.FindAll(filter)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Categorylar muvaffaqiyatli olindi", response)
}

// @Summary     Category ID bo'yicha
// @Tags        categories
// @Produce     json
// @Param       id path int true "Category ID"
// @Success     200 {object} utils.Response
// @Failure     404 {object} utils.Response
// @Router      /categories/{id} [get]
func (ctrl *CategoryController) FindById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "ID noto'g'ri formatda")
		return
	}

	response, err := ctrl.service.FindByID(uint(id))
	if err != nil {
		utils.Error(c, http.StatusNotFound, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Category muvaffaqiyatli olindi", response)

}

// @Summary     Category yangilash
// @Tags        categories
// @Accept      multipart/form-data
// @Produce     json
// @Security    BearerAuth
// @Param       id path int true "Category ID"
// @Param       name formData string false "Category nomi"
// @Param       description formData string false "Tavsif"
// @Param       color formData string false "Rang"
// @Param       image formData file false "Rasm"
// @Success     200 {object} utils.Response
// @Failure     400 {object} utils.Response
// @Failure     404 {object} utils.Response
// @Router      /categories/{id} [put]
func (ctrl *CategoryController) Update(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "ID noto'g'ri formatda")
		return
	}

	var req dto.UpdateCategoryRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// Image olish (majburiy emas)
	fileHeader, err := c.FormFile("image")
	if err != nil {
		// Image yuborilmagan — service ga nil beramiz
		response, err := ctrl.service.Update(uint(id), &req, nil, nil)
		if err != nil {
			utils.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		utils.Success(c, http.StatusOK, "Category muvaffaqiyatli yangilandi", response)
		return
	}

	openedFile, err := fileHeader.Open()
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Imageni ochishda xatolik")
		return
	}
	defer openedFile.Close()

	response, err := ctrl.service.Update(uint(id), &req, openedFile, fileHeader)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Category muvaffaqiyatli yangilandi", response)

}

// @Summary     Category o'chirish
// @Tags        categories
// @Produce     json
// @Security    BearerAuth
// @Param       id path int true "Category ID"
// @Success     200 {object} utils.Response
// @Failure     404 {object} utils.Response
// @Router      /categories/{id} [delete]
func (ctrl *CategoryController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "ID noto'g'ri formatda")
		return
	}

	if err := ctrl.service.Delete(uint(id)); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Category muvaffaqiyatli o'chirildi", nil)
}
