package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

func AsymmetricEncrypt(recPath string, message []byte) (*[]byte, error) {
	certificateBytes, err := os.ReadFile(recPath)
	if err != nil {
		return nil, err
	}

	certificatePemBlock, _ := pem.Decode(certificateBytes)
	if certificatePemBlock == nil {
		return nil, errors.New("certificate not found")
	}

	certificate, err := x509.ParseCertificate(certificatePemBlock.Bytes)
	if err != nil {
		return nil, err
	}

	encryptedMessage, err := rsa.EncryptPKCS1v15(rand.Reader, certificate.PublicKey.(*rsa.PublicKey), message)
	if err != nil {
		return nil, err
	}

	return &encryptedMessage, nil
}

func AsymmetricDecrypt(keyPath string, message []byte) (*[]byte, error) {
	privateKeyBytes, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	privateKeyPemBlock, _ := pem.Decode(privateKeyBytes)
	if privateKeyPemBlock == nil {
		return nil, errors.New("private key not found")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyPemBlock.Bytes)
	if err != nil {
		return nil, err
	}

	decryptedMessage, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, message)
	if err != nil {
		return nil, err
	}

	return &decryptedMessage, nil
}
