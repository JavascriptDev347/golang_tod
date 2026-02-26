package controllers

import (
	"net/http"
	"to-do/dto"
	"to-do/services"
	"to-do/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

// @Summary     Ro'yxatdan o'tish
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       request body dto.RegisterRequest true "Register ma'lumotlari"
// @Success     201 {object} utils.Response
// @Failure     400 {object} utils.Response
// @Router      /auth/register [post]
func (ctrl *UserController) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := ctrl.service.Register(&req)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, http.StatusCreated, "Muvaffaqiyatli ro'yxatdan o'tdingiz", response)
}

// @Summary     Kirish
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       request body dto.LoginRequest true "Login ma'lumotlari"
// @Success     200 {object} utils.Response
// @Failure     401 {object} utils.Response
// @Router      /auth/login [post]
func (ctrl *UserController) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := ctrl.service.Login(&req)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Muvaffaqiyatli kirdingiz", response)

}
