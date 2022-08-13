package tables

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID `json:"id"`
	Nick string    `json:"nick"`
	CID  string    `JSON:"cid"`
}

func (_ User) StructName() string {
	return "User"
}
