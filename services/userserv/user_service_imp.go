package userserv

import (
	"client-rest/models"
	"context"
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{userCollection, ctx}
}

func (u *UserServiceImpl) FindUserById(id string) (*models.UserResponse, error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	result := &models.UserResponse{}
	query := bson.M{"_id": oid}

	err := u.userCollection.FindOne(u.ctx, query).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.UserResponse{}, errors.New("user not exist") //err
		}
		return nil, err
	}
	return result, err
}

func (u *UserServiceImpl) FindUserByEmail(email string) (*models.UserResponse, error) {
	result := &models.UserResponse{}
	query := bson.M{"email": strings.ToLower(email)}

	err := u.userCollection.FindOne(u.ctx, query).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.UserResponse{}, err
		}
		return nil, err
	}
	return result, nil
}

func (u *UserServiceImpl) DeleteUser(id string) error {
	oid, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": oid}

	res, err := u.userCollection.DeleteOne(u.ctx, query)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("document is not exist")
	}
	return nil

}

func (u *UserServiceImpl) CreateUser(user *models.UserDB) (*models.UserResponse, error) {
	user.CreateAt = time.Now()
	user.UpdatedAt = user.CreateAt
	res, err := u.userCollection.InsertOne(u.ctx, user)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("the user is already exist")
		}
		return nil, err
	}

	//opt := options.Index()
	//opt.SetUnique(true)

	//index := mongo.IndexModel{keys: bson.M{"value": 1}, Options: opt}

	/*if _, err := u.collection.Indexes().CreateOne(u.ctx, index); err != nil {
		return nil, errors.New("could not create index")
	}*/

	var newUser *models.UserResponse
	query := bson.M{"_id": res.InsertedID}
	//query := bson.D{{Key: "_id", Value: res.InsertedID}}
	if err = u.userCollection.FindOne(u.ctx, query).Decode(&newUser); err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u *UserServiceImpl) UpdateUser(id string, user *models.UserUpdate) (*models.UserResponse, error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: user}}

	resp := u.userCollection.FindOneAndUpdate(u.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedUser *models.UserResponse
	if err := resp.Decode(&updatedUser); err != nil {
		return nil, errors.New("could not update user")
	}
	return updatedUser, nil
}

func (u *UserServiceImpl) GetAllUsers() ([]*models.UserResponse, error) {
	query := bson.M{}

	res, err := u.userCollection.Find(u.ctx, query, &options.FindOptions{})
	if err != nil {
		return nil, err
	}

	defer res.Close(u.ctx)

	var users []*models.UserResponse

	for res.Next(u.ctx) {
		user := &models.UserResponse{}
		err := res.Decode(user)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return []*models.UserResponse{}, nil
	}

	return users, nil
}
