package service

import (
	"time"
)

const (
	KindMeowCreated = iota + 1
)

type MeowCreatedMessage struct {
	Kind      int64     `json:"kind"`
	Id        string    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func NewMeowCreatedMessage(id string, body string, createdAt time.Time) *MeowCreatedMessage {
	return &MeowCreatedMessage{
		Kind:      KindMeowCreated,
		Id:        id,
		Body:      body,
		CreatedAt: createdAt,
	}
}
