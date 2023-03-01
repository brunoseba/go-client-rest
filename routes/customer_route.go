package routes

import (
	"client-rest/controllers"

	"github.com/gin-gonic/gin"
)

type CustomerRoute struct {
	customerController controllers.CustomerController
}

func NewCustomerRouteController(customerController controllers.CustomerController) CustomerRoute {
	return CustomerRoute{customerController}
}

func (rc *CustomerRoute) CustomerRoute(rg *gin.RouterGroup) {
	router := rg.Group("/v1")

	router.GET("/customer/:id", rc.customerController.FindCustomerById)
	router.GET("/customers", rc.customerController.GetAllCustomers)
	router.POST("/customer", rc.customerController.CreateCustomer)
	router.PUT("/customer/:id", rc.customerController.UpdateCustomer)
	router.DELETE("/customer/:id", rc.customerController.DeleteCustomer)

}
