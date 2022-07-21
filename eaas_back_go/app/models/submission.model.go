package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Submission struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string        `json:"name,omitempty" bson:"name,omitempty"`
	Email     string        `json:"email,omitempty" bson:"email,omitempty"`
	Status    string        `json:"status,omitempty" bson:"status,omitempty"`
	CreatedBy string        `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	CreatedAt time.Time     `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

// SubmissionUpdate , defines submission update model
type SubmissionUpdate struct {
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	IsActive bool   `json:"isActive,omitempty" bson:"isActive,omitempty"`
}
