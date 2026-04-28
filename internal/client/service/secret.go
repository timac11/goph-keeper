package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/timac11/goph-keeper/internal/client/encryption"
	clientModel "github.com/timac11/goph-keeper/internal/client/model"
	"github.com/timac11/goph-keeper/internal/common/model"
)

func (s *Service) GetSecretsList(ctx context.Context, args clientModel.ListArgs) (*[]model.SecretInfoDto, error) {
	settings, err := s.cache.Restore(ctx)

	if err != nil {
		return nil, err
	}

	res, err := s.client.GetList(ctx, settings.Token)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) DeleteSecret(ctx context.Context, args clientModel.DeleteArgs) error {
	settings, err := s.cache.Restore(ctx)

	if err != nil {
		return err
	}

	return s.client.DeleteSecret(ctx, settings.Token, args.ID)
}

func (s *Service) GetSecret(ctx context.Context, args clientModel.GetArgs) (string, error) {
	settings, err := s.cache.Restore(ctx)
	if err != nil {
		return "", err
	}

	res, err := s.client.GetSecret(ctx, settings.Token, args.ID)
	if err != nil {
		return "", err
	}

	// store encrypted file
	defer res.File.Close()
	tempFile, err := os.CreateTemp(s.config.StoreDir, "*")
	if err != nil {
		return "", err
	}

	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	_, err = io.Copy(tempFile, res.File)
	if err != nil {
		return "", err
	}

	// decode and decrypt public key
	decodedPublicKey, err := base64.StdEncoding.DecodeString(res.PublicKey)
	if err != nil {
		return "", err
	}

	decryptedPublicKey, err := encryption.AsymmetricDecrypt(s.config.PrivateKeyPath, decodedPublicKey)
	if err != nil {
		return "", err
	}

	// decrypt and store file
	dst, err := os.Create(filepath.Join(s.config.StoreDir, res.Name))
	if err != nil {
		return "", err
	}

	err = encryption.SymmetricDecryptFile(tempFile, dst, *decryptedPublicKey)
	if err != nil {
		return "", err
	}

	return dst.Name(), nil
}

func (s *Service) UploadFile(ctx context.Context, args clientModel.UploadFileArgs) error {
	return s.uploadEntry(ctx, args.Path, args.Metadata, "FILE")
}

func (s *Service) UploadAuth(ctx context.Context, args clientModel.UploadAuthArgs) error {
	fileName := fmt.Sprintf("%d.json", args.UploadName)

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(f)
	err = encoder.Encode(args)
	if err != nil {
		return err
	}

	defer f.Close()
	defer os.Remove(f.Name())

	return s.uploadEntry(ctx, f.Name(), args.Metadata, "AUTH")
}

func (s *Service) UploadCard(ctx context.Context, args clientModel.UploadCardArgs) error {
	fileName := fmt.Sprintf("%d.json", args.UploadName)

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(f)
	err = encoder.Encode(args)
	if err != nil {
		return err
	}

	defer f.Close()
	defer os.Remove(f.Name())

	return s.uploadEntry(ctx, f.Name(), args.Metadata, "CARD")
}

func (s *Service) uploadEntry(ctx context.Context, path, metadata, entryType string) error {
	settings, err := s.cache.Restore(ctx)
	if err != nil {
		return err
	}

	tempFile, err := os.CreateTemp(s.config.StoreDir, "*")
	if err != nil {
		return err
	}

	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	in, err := os.Open(path)
	if err != nil {
		return err
	}

	publicKeyBytes, err := encryption.SymmetricEncryptFile(in, tempFile)
	if err != nil {
		return err
	}

	encryptedPublicKey, err := encryption.AsymmetricEncrypt(s.config.RecPath, *publicKeyBytes)
	if err != nil {
		return err
	}

	publicKey := base64.StdEncoding.EncodeToString(*encryptedPublicKey)

	uploadData := clientModel.UploadData{
		Name:      filepath.Base(path),
		File:      tempFile,
		Type:      entryType,
		PublicKey: publicKey,
		Metadata:  metadata,
	}

	return s.client.UploadSecret(ctx, settings.Token, uploadData)
}
