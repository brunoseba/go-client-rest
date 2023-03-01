package customerservice

import (
	"client-rest/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CustomerServiceImpl struct {
	customerCollection *mongo.Collection
	ctx                context.Context
}

func NewCustomerService(customerCollection *mongo.Collection, ctx context.Context) CustomerService {
	return &CustomerServiceImpl{customerCollection, ctx}
}

func (c *CustomerServiceImpl) CreateCustomer(customer *models.CustomerDB) (*models.CustomerResponse, error) {
	customer.CreateAt = time.Now()
	customer.UpdatedAt = customer.CreateAt

	res, err := c.customerCollection.InsertOne(c.ctx, customer)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("the customer is already exist")
		}
		return nil, err
	}
	// TODO return debe retonar un error o nada (en el controlador respondera con un  status 200 ok)

	var newCustomer *models.CustomerResponse
	query := bson.M{"_id": res.InsertedID}
	//query := bson.D{{Key: "_id", Value: res.InsertedID}}
	if err = c.customerCollection.FindOne(c.ctx, query).Decode(&newCustomer); err != nil {
		return nil, err
	}
	return newCustomer, nil
}

func (c *CustomerServiceImpl) FindCustomerById(id string) (*models.CustomerResponse, error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	result := &models.CustomerResponse{}
	query := bson.M{"_id": oid}

	err := c.customerCollection.FindOne(c.ctx, query).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.CustomerResponse{}, err
		}
		return nil, err
	}

	return result, nil
}

func (c *CustomerServiceImpl) UpdateCustomer(id string, customer *models.CustomerUpdate) (*models.CustomerResponse, error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: oid}}
	updater := bson.D{{Key: "$set", Value: customer}}

	result := c.customerCollection.FindOneAndUpdate(c.ctx, query, updater, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedCustomer *models.CustomerResponse

	if err := result.Decode(&updatedCustomer); err != nil {
		return nil, errors.New("cloud no updates customer")
	}

	return updatedCustomer, nil

}

func (c *CustomerServiceImpl) GetAllCustomers() ([]*models.CustomerResponse, error) {
	var allCustomer []*models.CustomerResponse

	query := bson.M{}
	result, err := c.customerCollection.Find(c.ctx, query, &options.FindOptions{})

	if err != nil {
		return nil, err
	}
	defer result.Close(c.ctx)

	for result.Next(c.ctx) {
		customer := &models.CustomerResponse{}
		err := result.Decode(customer)

		if err != nil {
			return nil, err
		}
		allCustomer = append(allCustomer, customer)
	}

	if err := result.Err(); err != nil {
		return nil, err
	}

	if len(allCustomer) == 0 {
		return []*models.CustomerResponse{}, nil
	}

	return allCustomer, nil
}

func (c *CustomerServiceImpl) DeleteCustomer(id string) error {
	oid, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": oid}

	result, err := c.customerCollection.DeleteOne(c.ctx, query)

	// Error DB does not return
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("Customer not exists")
	}
	return nil
}
