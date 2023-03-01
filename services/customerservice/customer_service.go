package customerservice

import (
	"client-rest/models"
)

type CustomerService interface {
	FindCustomerById(id string) (*models.CustomerResponse, error)
	GetAllCustomers() ([]*models.CustomerResponse, error)
	CreateCustomer(*models.CustomerDB) (*models.CustomerResponse, error)
	UpdateCustomer(id string, customer *models.CustomerUpdate) (*models.CustomerResponse, error)
	DeleteCustomer(id string) error
	//FindCustomerByEmail(email string) (*models.CustomerResponse, error)
}
