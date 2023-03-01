package controllers

import (
	"client-rest/models"
	"client-rest/services/customerservice"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	customerService customerservice.CustomerService
}

func NewCustomerController(customerService customerservice.CustomerService) CustomerController {
	return CustomerController{customerService}
}

func (cc *CustomerController) CreateCustomer(ctx *gin.Context) {
	var customer *models.CustomerDB

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	err = json.Unmarshal(body, &customer)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	newCustomer, err := cc.customerService.CreateCustomer(customer)
	if err != nil {
		if strings.Contains(err.Error(), "customer already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict, "message": err.Error()})
		}
		ctx.AbortWithStatusJSON(http.StatusBadGateway, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": newCustomer})
}

func (cc *CustomerController) FindCustomerById(ctx *gin.Context) {
	customerId := ctx.Param("id")
	if len(customerId) != 24 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User Id must be 24 characters"})
		return
	}

	customer, err := cc.customerService.FindCustomerById(customerId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": customer})
}

func (cc *CustomerController) UpdateCustomer(ctx *gin.Context) {
	var customer *models.CustomerUpdate
	oid := ctx.Param("id")
	if len(oid) != 24 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User Id must be 24 characters"})
		return
	}

	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}

	_, err := cc.customerService.UpdateCustomer(oid, customer)
	if err != nil {
		/*if strings.Contains(err.Error(), "customer not found") {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
			return
		}*/
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "customer updated"})

}

func (cc *CustomerController) GetAllCustomers(ctx *gin.Context) {
	customers, err := cc.customerService.GetAllCustomers()

	if err != nil {
		if strings.Contains(err.Error(), "customers not found") {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "customers not found"})
		}
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": customers})
}

func (cc *CustomerController) DeleteCustomer(ctx *gin.Context) {
	oid := ctx.Param("id")
	if len(oid) != 24 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "customer must be 24 characters"})
		return
	}

	err := cc.customerService.DeleteCustomer(oid)

	if err != nil {
		/*if strings.Contains(err.Error(), "customer not found") {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "customer not found"})
			return
		}*/
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "customer successfully deleted"})
}
