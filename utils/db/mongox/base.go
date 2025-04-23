package mongox

import (
	"strconv"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DefaultModel struct contains IDField.
type DefaultModel struct {
	IDField `bson:",inline" json:",inline"`
}

// IntIDModel struct contains IDIntField.
type IntIDModel struct {
	IDIntField `bson:",inline" json:",inline"`
}

// DefaultModelWithDate struct contains IDField and DateFields.
type DefaultModelWithDate struct {
	IDField    `bson:",inline" json:",inline"`
	DateFields `bson:",inline" json:",inline"`
}

// Creating function calls the inner fields' defined hooks.
func (model *DefaultModelWithDate) Creating() error {
	return model.DateFields.Creating()
}

// Updating function calls the inner fields' defined hooks.
func (model *DefaultModelWithDate) Updating() error {
	return model.DateFields.Updating()
}

// IDField struct contains a model's ID field.
type IDField struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}

// e.g convert hex-string ID value to bson.ObjectId.
func (f *IDField) PrepareID(id interface{}) (interface{}, error) {
	if idStr, ok := id.(string); ok {
		return primitive.ObjectIDFromHex(idStr)
	}

	// Otherwise id must be ObjectId
	return id, nil
}

// GetID method returns a model's ID.
func (f *IDField) GetID() interface{} {
	return f.ID
}

// SetID sets the value of a model's ID field.
func (f *IDField) SetID(id interface{}) {
	f.ID = id.(primitive.ObjectID)
}

// IDIntField struct contains a model's ID field.
type IDIntField struct {
	ID int64 `json:"id" bson:"_id,omitempty"`
}

// e.g convert hex-string ID value to bson.ObjectId.
func (f *IDIntField) PrepareID(id interface{}) (interface{}, error) {
	if idStr, ok := id.(string); ok {
		idInt, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, err
		}
		return int64(idInt), nil
	}

	// Otherwise id must be ObjectId
	return id, nil
}

// GetID method return model's id.
func (f *IDIntField) GetID() interface{} {
	return f.ID
}

// SetID set id value of model's id field.
func (f *IDIntField) SetID(id interface{}) {
	switch v := id.(type) {
	case uint8:
		f.ID = int64(v)
	case uint16:
		f.ID = int64(v)
	case uint32:
		f.ID = int64(v)
	case uint64:
		f.ID = int64(v)
	case int8:
		f.ID = int64(v)
	case int16:
		f.ID = int64(v)
	case int32:
		f.ID = int64(v)
	case int:
		f.ID = int64(v)
	case int64:
		f.ID = v
	case string:
		intID, _ := strconv.Atoi(v)
		f.ID = int64(intID)
	}
}

// DateFields struct contains the `created_at` and `updated_at`
// fields that autofill when inserting or updating a model.
type DateFields struct {
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

// Creating hook is used here to set the `created_at` field
// value when inserting a new model into the database.
func (f *DateFields) Creating() error {
	now := time.Now().UTC()
	f.CreatedAt = now
	f.UpdatedAt = now

	return nil
}

// Updating hook is used here to set the `updated_at` field
// value when creating or updateing a model.
func (f *DateFields) Updating() error {
	f.UpdatedAt = time.Now().UTC()

	return nil
}

type BaseProps interface {
	mgm.CollectionNameGetter
}
