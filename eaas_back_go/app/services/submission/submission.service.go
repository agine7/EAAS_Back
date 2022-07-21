package submission

import (
	"context"
	model "eaas_back/app/models"
	repository "eaas_back/app/repositories/submission"
	"eaas_back/config"
	"eaas_back/utility"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SubmissionServiceInterface interface {
	Update(context.Context, string, *model.SubmissionUpdate) error
	Get(context.Context) ([]model.Submission, error)
	Create(context.Context, *model.Submission) error
}

// SubmissionService , implements SubmissionService
// and perform user related business logics
type SubmissionService struct {
	db         *mgo.Session
	repository repository.SubmissionRepository
	config     *config.Configuration
}

// New function will initialize SubmissionService
func New(submissionRepo repository.SubmissionRepository) SubmissionServiceInterface {
	return &SubmissionService{repository: submissionRepo}
}

// Update function will update the user info
// return error if any
func (service *SubmissionService) Update(ctx context.Context, id string, submission *model.SubmissionUpdate) error {
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	CustomBson := &utility.CustomBson{}
	change, err := CustomBson.Set(submission)
	if err != nil {
		return err
	}
	return service.repository.Update(ctx, query, change)
}

// Get function will find user by id
// return user and error if any
func (service *SubmissionService) Get(ctx context.Context) ([]model.Submission, error) {
	return service.repository.FindAll(ctx)
}

func (service *SubmissionService) Create(ctx context.Context, submission *model.Submission) error {
	return service.repository.Create(ctx, submission)
}
