package model

import (
	"io"
	"os"
)

type UploadData struct {
	Name      string
	File      *os.File
	PublicKey string
	Metadata  string
	Type      string
}

type GetSecretDto struct {
	Name      string
	Type      string
	Metadata  string
	PublicKey string
	File      io.ReadCloser
}
