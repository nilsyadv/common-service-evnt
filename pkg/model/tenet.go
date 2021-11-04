package model

import uuid "github.com/satori/go.uuid"

type Tenent struct {
	ID   uuid.UUID
	Name string
}
