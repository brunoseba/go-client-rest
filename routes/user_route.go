package routes

import (
	"client-rest/controllers"

	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewUserRouteController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (r *UserRouteController) UserRoute(rg *gin.RouterGroup) {
	router := rg.Group("/v1")

	router.GET("/users", r.userController.GetAllUsers)
	router.GET("/user/:id", r.userController.FindUserById)
	router.POST("/user", r.userController.CreateUser)
	router.DELETE("/user/:id", r.userController.DeleteUser)
	router.PUT("/user/:id", r.userController.UpdateUser)
}
