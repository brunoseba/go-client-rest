package main

import (
	"client-rest/config"
	"client-rest/controllers"
	"client-rest/routes"
	"client-rest/services/customerservice"
	"client-rest/services/userserv"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server *gin.Engine
	ctx    context.Context

	//User Services, controllers, routes
	userCollection     *mongo.Collection
	customerCollection *mongo.Collection

	userService         userserv.UserService
	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	customerService    customerservice.CustomerService
	CustomerController controllers.CustomerController
	CustomerRoute      routes.CustomerRoute
)

func init() {

	ctx = context.TODO()

	// connect with database --------------------------------------------------------
	mongoclient := config.DBconn(ctx)
	// -------------------------------------------------------------------------------

	//Collections
	userCollection = mongoclient.Database("clientApi").Collection("users")
	customerCollection = mongoclient.Database("clientApi").Collection("customers")

	//Route, Controller and Service --------------------------------------------------
	userService = userserv.NewUserService(userCollection, ctx)
	UserController = controllers.NewUserController(userService)
	UserRouteController = routes.NewUserRouteController(UserController)

	customerService = customerservice.NewCustomerService(customerCollection, ctx)
	CustomerController = controllers.NewCustomerController(customerService)
	CustomerRoute = routes.NewCustomerRouteController(CustomerController)

	// -------------------------------------------------------------------------------

	server = gin.Default()
}

func main() {

	// Routes
	router := server.Group("/api")
	UserRouteController.UserRoute(router)
	CustomerRoute.CustomerRoute(router)

	log.Fatal(server.Run(":8080"))

}
