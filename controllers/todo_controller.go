package controllers

import (
	"net/http"
	"strconv"
	"to-do/dto"
	"to-do/services"
	"to-do/utils"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	service services.TodoService
}

func NewTodoController(service services.TodoService) *TodoController {
	return &TodoController{service: service}
}

// ==================== CREATE ====================
// @Summary     Todo yaratish
// @Tags        todos
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       request body dto.CreateTodoRequest true "Todo ma'lumotlari"
// @Success     201 {object} utils.Response
// @Failure     400 {object} utils.Response
// @Failure     401 {object} utils.Response
// @Router      /todos [post]
func (ctrl *TodoController) Create(c *gin.Context) {
	var req dto.CreateTodoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// Token dan userID olish
	userID, _ := c.Get("user_id")

	response, err := ctrl.service.Create(userID.(uint), &req)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, http.StatusCreated, "Todo muvaffaqiyatli yaratildi", response)
}

// ==================== FIND ALL ====================

// @Summary     Barcha todolar
// @Tags        todos
// @Produce     json
// @Security    BearerAuth
// @Param       search query string false "Qidirish"
// @Param       status query string false "Status (pending, in_progress, completed)"
// @Param       priority query string false "Priority (low, medium, high)"
// @Param       sort_by query string false "Saralash (title, status, priority, due_date, created_at)"
// @Param       sort_dir query string false "Tartib (asc, desc)"
// @Success     200 {object} utils.Response
// @Failure     401 {object} utils.Response
// @Router      /todos [get]
func (ctrl *TodoController) FindAll(c *gin.Context) {
	var filter dto.TodoFilter

	if err := c.ShouldBindQuery(&filter); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("user_id")

	response, err := ctrl.service.FindAll(userID.(uint), filter)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Todolar muvaffaqiyatli olindi", response)
}

// ==================== FIND BY ID ====================
// @Summary     Todo ID bo'yicha
// @Tags        todos
// @Produce     json
// @Security    BearerAuth
// @Param       id path int true "Todo ID"
// @Success     200 {object} utils.Response
// @Failure     401 {object} utils.Response
// @Failure     404 {object} utils.Response
// @Router      /todos/{id} [get]
func (ctrl *TodoController) FindByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "ID noto'g'ri formatda")
		return
	}

	userID, _ := c.Get("user_id")

	response, err := ctrl.service.FindByID(userID.(uint), uint(id))
	if err != nil {
		utils.Error(c, http.StatusNotFound, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Todo muvaffaqiyatli olindi", response)
}

// ==================== UPDATE ====================
// @Summary     Todo yangilash
// @Tags        todos
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       id path int true "Todo ID"
// @Param       request body dto.UpdateTodoRequest true "Yangilanadigan ma'lumotlar"
// @Success     200 {object} utils.Response
// @Failure     400 {object} utils.Response
// @Failure     401 {object} utils.Response
// @Failure     404 {object} utils.Response
// @Router      /todos/{id} [put]
func (ctrl *TodoController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "ID noto'g'ri formatda")
		return
	}

	var req dto.UpdateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("user_id")

	response, err := ctrl.service.Update(userID.(uint), uint(id), &req)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Todo muvaffaqiyatli yangilandi", response)
}

// ==================== DELETE ====================

// @Summary     Todo o'chirish
// @Tags        todos
// @Produce     json
// @Security    BearerAuth
// @Param       id path int true "Todo ID"
// @Success     200 {object} utils.Response
// @Failure     401 {object} utils.Response
// @Failure     404 {object} utils.Response
// @Router      /todos/{id} [delete]
func (ctrl *TodoController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "ID noto'g'ri formatda")
		return
	}

	userID, _ := c.Get("user_id")

	if err := ctrl.service.Delete(userID.(uint), uint(id)); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Todo muvaffaqiyatli o'chirildi", nil)
}
