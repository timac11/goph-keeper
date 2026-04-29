package http

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"

	clientModel "github.com/timac11/goph-keeper/internal/client/model"
	"github.com/timac11/goph-keeper/internal/common/model"
)

type HTTPClient struct {
	client resty.Client
}

// NewHTTPClient constructor
func NewHTTPClient(address string) *HTTPClient {
	client := resty.New()

	if !strings.HasPrefix(address, "http") {
		address = "http://" + address
	}

	client.SetBaseURL(address)

	return &HTTPClient{client: *client}
}

func (client *HTTPClient) Login(ctx context.Context, login, password string) (*model.UserLoginResultDto, error) {
	payload := model.UserLoginDto{Login: login, Password: password}
	var result model.UserLoginResultDto

	response, err := client.client.R().SetBody(payload).SetResult(&result).Post("/api/login")
	if err != nil {
		return nil, err
	}

	if response.IsError() {
		return nil, errors.New("login user error")
	}

	return &result, nil
}

func (client *HTTPClient) Register(ctx context.Context, login, password string) (*model.UserLoginResultDto, error) {
	payload := model.UserLoginDto{Login: login, Password: password}
	var result model.UserLoginResultDto

	response, err := client.client.R().SetBody(payload).SetResult(&result).Post("/api/register")
	if err != nil {
		return nil, err
	}

	if response.IsError() {
		return nil, errors.New("register user error")
	}

	return &result, nil
}

func (client *HTTPClient) GetList(ctx context.Context, token string) (*[]model.SecretInfoDto, error) {
	var result []model.SecretInfoDto

	response, err := client.client.R().SetHeader("Authorization", token).
		SetResult(&result).
		Get("/api/secrets")
	if err != nil {
		return nil, err
	}

	if response.IsError() {
		return nil, errors.New("secrets list error")
	}

	return &result, nil
}

func (client *HTTPClient) GetSecret(ctx context.Context, token, id string) (*clientModel.GetSecretDto, error) {
	response, err := client.client.R().SetHeader("Authorization", token).
		SetContext(ctx).
		SetDoNotParseResponse(true).
		Get(fmt.Sprintf("/api/secrets/%s", id))
	if err != nil {
		return nil, err
	}

	if response.IsError() {
		return nil, errors.New("get secret error")
	}

	if err != nil {
		return nil, err
	}

	return &clientModel.GetSecretDto{
		Name:      response.Header().Get("X-Secret-Name"),
		Type:      response.Header().Get("X-Secret-Type"),
		Metadata:  response.Header().Get("X-Secret-Metadata"),
		PublicKey: response.Header().Get("X-Secret-PublicKey"),
		File:      response.RawBody(),
	}, nil
}

func (client *HTTPClient) DeleteSecret(ctx context.Context, token, id string) error {
	response, err := client.client.R().SetHeader("Authorization", token).
		Delete(fmt.Sprintf("/api/secrets/%s", id))
	if err != nil {
		return err
	}

	if response.IsError() {
		return errors.New("delete secret error")
	}

	return nil
}

func (client *HTTPClient) UploadSecret(ctx context.Context, token string, data clientModel.UploadData) error {
	response, err := client.client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/octet-stream").
		SetHeader("Authorization", token).
		SetHeader("X-Secret-Name", data.Name).
		SetHeader("X-Secret-Type", data.Type).
		SetHeader("X-Secret-Metadata", data.Metadata).
		SetHeader("X-Secret-Public-Key", data.PublicKey).
		SetFile(data.Name, data.File.Name()).
		Post("/api/secrets")
	if err != nil {
		return err
	}

	if response.IsError() {
		return errors.New("upload secret error")
	}

	return nil
}
