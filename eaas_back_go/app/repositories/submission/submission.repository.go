package submission

import (
	"context"
	model "eaas_back/app/models"
	"eaas_back/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SubmissionRepository interface {

	// Create, will perform db opration to save user
	// Returns modified user and error if occurs
	Create(context.Context, *model.Submission) error

	// FildAll, returns all users in the system
	// It will return error also if occurs
	FindAll(context.Context) ([]model.Submission, error)

	// FindOneById, find the user by the provided id
	// return matched user and error if any
	FindOneById(context.Context, string) (*model.Submission, error)

	// Update, will update user data by id
	// return error if any
	Update(context.Context, interface{}, interface{}) error

	// Delete, will remove user entry from DB
	// Return error if any
	Delete(context.Context, *model.Submission) error

	// FindOne, will find one entry of user matched by the query
	// Query object is an interface type that can accept any object
	// return matched user and error if any
	FindOne(context.Context, interface{}) (*model.Submission, error)
}

type RepositoryImp struct {
	db     *mgo.Session
	config *config.Configuration
}

func New(db *mgo.Session, c *config.Configuration) SubmissionRepository {
	return &RepositoryImp{db: db, config: c}
}

func (service *RepositoryImp) Create(ctx context.Context, submission *model.Submission) error {
	return service.collection().Insert(submission)
}

func (service *RepositoryImp) FindAll(ctx context.Context) ([]model.Submission, error) {
	var submission []model.Submission
	query := bson.M{}
	e := service.collection().Find(query).All(&submission)
	return submission, e
}

func (service *RepositoryImp) Update(ctx context.Context, query, change interface{}) error {

	return service.collection().Update(query, change)
}

func (service *RepositoryImp) FindOneById(ctx context.Context, id string) (*model.Submission, error) {
	var submission model.Submission
	query := bson.M{}
	e := service.collection().Find(query).Select(bson.M{"password": 0, "salt": 0}).One(&submission)
	return &submission, e
}

func (service *RepositoryImp) Delete(ctx context.Context, user *model.Submission) error {
	return nil
}

func (service *RepositoryImp) FindOne(ctx context.Context, query interface{}) (*model.Submission, error) {
	var submission model.Submission
	e := service.collection().Find(query).One(&submission)
	return &submission, e
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
	return service.db.DB(service.config.DataBaseName).C("submissions")
}
