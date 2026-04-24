package service

import (
	clientModel "github.com/timac11/goph-keeper/internal/client/model"
	"github.com/timac11/goph-keeper/internal/common/model"
)

func (s *Service) GetSecretsList(args clientModel.ListArgs) (*[]model.SecretInfoDto, error) {
	settings, err := s.cache.Restore()

	if err != nil {
		return nil, err
	}

	res, err := s.client.GetList(settings.Token)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) DeleteSecret(args clientModel.DeleteArgs) error {
	settings, err := s.cache.Restore()

	if err != nil {
		return err
	}

	return s.client.DeleteSecret(settings.Token, args.ID)
}

func (s *Service) UploadAuth(args clientModel.UploadAuthArgs) {
	// TODO: implement
}

func (s *Service) UploadCard(args clientModel.UploadCardArgs) {
	// TODO: implement
}

func (s *Service) UploadFile(args clientModel.UploadFileArgs) {
	// TODO: implement
}
