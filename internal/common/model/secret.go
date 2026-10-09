package model

import (
	"time"
)

const (
	File = "FILE"
	Auth = "AUTH"
	Card = "CARD"
)

type Secret struct {
	ID        string
	Name      string
	DataPath  string
	Metadata  string
	PublicKey string
	Type      string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type SecretCreateDto struct {
	Name      string
	DataPath  string
	Metadata  string
	PublicKey string
	Type      string
}

type SecretInfoDto struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
