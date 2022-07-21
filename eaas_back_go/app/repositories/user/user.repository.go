package user

import (
	"context"
	"eaas_back/config"

	model "eaas_back/app/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserRepository , used to perform DB oprations
// Inteface contains basic oprations on user document
// So that, db opration can be perform easily
type UserRepository interface {

	// Create, will perform db opration to save user
	// Returns modified user and error if occurs
	Create(context.Context, *model.User) error

	// FildAll, returns all users in the system
	// It will return error also if occurs
	FindAll(context.Context) ([]*model.User, error)

	// FindOneById, find the user by the provided id
	// return matched user and error if any
	FindOneById(context.Context, string) (*model.User, error)

	// Update, will update user data by id
	// return error if any
	Update(context.Context, interface{}, interface{}) error

	// Delete, will remove user entry from DB
	// Return error if any
	Delete(context.Context, *model.User) error

	// FindOne, will find one entry of user matched by the query
	// Query object is an interface type that can accept any object
	// return matched user and error if any
	FindOne(context.Context, interface{}) (*model.User, error)

	IsUserAlreadyExists(context.Context, string) bool
}

type RepositoryImp struct {
	db     *mgo.Session
	config *config.Configuration
}

func New(db *mgo.Session, c *config.Configuration) UserRepository {
	return &RepositoryImp{db: db, config: c}
}

func (service *RepositoryImp) Create(ctx context.Context, user *model.User) error {
	return service.collection().Insert(user)
}

func (service *RepositoryImp) FindAll(ctx context.Context) ([]*model.User, error) {
	return nil, nil
}

func (service *RepositoryImp) Update(ctx context.Context, query, change interface{}) error {

	return service.collection().Update(query, change)
}

func (service *RepositoryImp) FindOneById(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	e := service.collection().Find(query).Select(bson.M{"password": 0, "salt": 0}).One(&user)
	return &user, e
}

func (service *RepositoryImp) Delete(ctx context.Context, user *model.User) error {
	return nil
}

func (service *RepositoryImp) FindOne(ctx context.Context, query interface{}) (*model.User, error) {
	var user model.User
	e := service.collection().Find(query).One(&user)
	return &user, e
}

// IsUserAlreadyExists , checks if user already exists in DB
func (service *RepositoryImp) IsUserAlreadyExists(ctx context.Context, email string) bool {
	query := bson.M{"email": email}
	_, e := service.FindOne(ctx, query)
	if e != nil {
		return false
	}
	return true
}

func (service *RepositoryImp) collection() *mgo.Collection {
	return service.db.DB(service.config.DataBaseName).C("users")
}
