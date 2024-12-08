package mongoxv2

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Modeler interface {
	CollName() string
}

type IDer interface {
	SetID(id any)
}

type IDModeler interface {
	Modeler
	IDer
}

type ID = IDT[any]

type ObjectID = IDT[primitive.ObjectID]

type IntID = IDT[int]

type StringID = IDT[string]

type IDT[T any] struct {
	ID T `json:"id" bson:"_id,omitempty"`
}

// SetID sets the ID field of the IDT struct.
// It attempts to type assert the given value to the generic type T.
// If the type assertion fails, it logs an error message.
func (id *IDT[T]) SetID(t any) {
	var ok bool

	id.ID, ok = t.(T)
	if !ok {
		fmt.Println("SetID: type assertion failed")
	}
}
