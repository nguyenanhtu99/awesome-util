package ordermodel

import "awesome-util/utils/db/mongox"

type Order struct {
	mongox.IDField `json:",inline" bson:",inline"`

	UserID string `json:"userId" bson:"userId"`
}
