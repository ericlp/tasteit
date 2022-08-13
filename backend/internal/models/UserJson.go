package models

import "github.com/google/uuid"

type User struct {
	Id   uuid.UUID `json:"id"`
	Nick string    `json:"nick"`
	CID  string    `json:"cid"`
}
