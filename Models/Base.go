package Models

import "github.com/google/uuid"


type BinaryUUID uuid.UUID

type ModelBase struct{
	ID BinaryUUID `json:"id"`

}