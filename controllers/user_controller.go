package controllers

import (
	"client-rest/models"
	"client-rest/services/userserv"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService userserv.UserService //services.UserService
}

func NewUserController(userService userserv.UserService) UserController {
	return UserController{userService}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user *models.UserDB

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	newUser, err := uc.UserService.CreateUser(user)

	if err != nil {
		if strings.Contains(err.Error(), "user already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "409", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "502", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newUser})
}

func (uc *UserController) FindUserById(ctx *gin.Context) {
	userId := ctx.Param("id")
	if len(userId) != 24 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User Id must be 24 characters"})
		return
	}

	user, err := uc.UserService.FindUserById(userId)

	if err != nil {
		/*if strings.Contains(err.Error(), "user not found") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}*/
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	if len(userId) != 24 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User Id must be 24 characters"})
		return
	}

	var user *models.UserUpdate
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "is not JSON -- " + err.Error()})
		return
	}

	userUpdate, err := uc.UserService.UpdateUser(userId, user)
	if err != nil {
		/*if strings.Contains(err.Error(), "user not found") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "404", "message": "usern not found " + err.Error()})
			return
		}*/
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": userUpdate})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	if len(userId) != 24 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User Id must be 24 characters"})
		return
	}

	err := uc.UserService.DeleteUser(userId)

	if err != nil {
		/*if strings.Contains(err.Error(), "user not found") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}*/
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "User successfully deleted"})
}

func (uc *UserController) GetAllUsers(ctx *gin.Context) {

	users, err := uc.UserService.GetAllUsers()

	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "404", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "502", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": users})
}
