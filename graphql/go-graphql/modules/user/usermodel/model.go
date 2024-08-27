package usermodel

import (
	"awesome-util/utils/db/mongox"
)

type User struct {
	mongox.IDField `bson:",inline"`
	Name           string `json:"name" bson:"name"`
}
