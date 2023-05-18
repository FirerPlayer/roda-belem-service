package entity

import "github.com/google/uuid"

type Image struct {
	Id          uuid.UUID
	Data        []byte
	Description string
}

func NewImage(data []byte, description string) *Image {
	return &Image{
		Id:          uuid.New(),
		Data:        data,
		Description: description,
	}
}