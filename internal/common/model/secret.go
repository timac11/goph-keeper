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
	Data      []byte
	Metadata  string
	PublicKey []byte
	Type      string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type SecretCreateDto struct {
	Name      string
	Data      []byte
	Metadata  string
	PublicKey []byte
	Type      string
}

type SecretInfoDto struct {
	ID        string
	Name      string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
