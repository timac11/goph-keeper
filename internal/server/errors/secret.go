package errors

import (
	"errors"
)

var SecretInvalidForm = errors.New("invalid form data")
var SecretNameIsRequired = errors.New("secret name is required")
var SecretInvalidType = errors.New("secret type is invalid")
var SecretFileIsRequired = errors.New("secret file is required")
var SecretFileFailedToRead = errors.New("secret file failed to read")
var PublicKeyIsRequired = errors.New("public key is required")
var PublicKeyFailedToRead = errors.New("public key failed to read")
