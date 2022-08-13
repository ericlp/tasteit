package tables

import "github.com/google/uuid"

type UserOwner struct {
	OwnerID       uuid.UUID
	TasteitUserID uuid.UUID
}

func (_ UserOwner) StructName() string {
	return "UserOwner"
}
