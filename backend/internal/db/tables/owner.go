package tables

import "github.com/google/uuid"

type Owner struct {
	ID   uuid.UUID
	Name string
}

func (_ Owner) StructName() string {
	return "Owner"
}
