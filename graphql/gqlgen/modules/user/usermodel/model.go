package usermodel

import (
	"awesome-util/utils/db/mongox"
)

type User struct {
	mongox.IDField `       bson:",inline"`
	Name           string `bson:"name"    json:"name"`
}
