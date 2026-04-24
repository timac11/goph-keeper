package http

import (
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

func (client *HTTPClient) Login(login, password string) (*model.UserLoginResultDto, error) {
	payload := model.UserLoginDto{Login: login, Password: password}
	var result model.UserLoginResultDto

	response, err := client.client.R().SetBody(payload).SetResult(result).Post("/api/login")
	if err != nil {
		return nil, err
	}

	if response.IsError() {
		return nil, errors.New("login user error")
	}

	return &result, nil
}

func (client *HTTPClient) Register(login, password string) (*model.UserLoginResultDto, error) {
	payload := model.UserLoginDto{Login: login, Password: password}
	var result model.UserLoginResultDto

	response, err := client.client.R().SetBody(payload).SetResult(result).Post("/api/register")
	if err != nil {
		return nil, err
	}

	if response.IsError() {
		return nil, errors.New("register user error")
	}

	return &result, nil
}

func (client *HTTPClient) GetList(token string) (*[]model.SecretInfoDto, error) {
	var result []model.SecretInfoDto

	response, err := client.client.R().SetHeader("Authorization", token).
		SetResult(result).
		Get("/api/secrets")
	if err != nil {
		return nil, err
	}

	if response.IsError() {
		return nil, errors.New("secrets list error")
	}

	return &result, nil
}

func (client *HTTPClient) GetSecret(token, id string) (*model.Secret, error) {
	var result model.Secret

	response, err := client.client.R().SetHeader("Authorization", token).
		SetResult(result).
		Get(fmt.Sprintf("/api/secret/%s", id))
	if err != nil {
		return nil, err
	}

	if response.IsError() {
		return nil, errors.New("get secret error")
	}

	return &result, nil
}

func (client *HTTPClient) DeleteSecret(token, id string) error {
	response, err := client.client.R().SetHeader("Authorization", token).
		Delete(fmt.Sprintf("/api/secret/%s", id))
	if err != nil {
		return err
	}

	if response.IsError() {
		return errors.New("delete secret error")
	}

	return nil
}

func (client *HTTPClient) UploadSecret(token string, data clientModel.UploadData) error {
	response, err := client.client.R().SetHeader("Authorization", token).
		SetFile("data", data.FilePath).
		SetFile("publicKey", data.PublicKeyPath).
		SetFormData(map[string]string{"name": data.Name, "type": data.Type, "metadata": data.Metadata}).
		Post("/api/secret")
	if err != nil {
		return err
	}

	if response.IsError() {
		return errors.New("delete secret error")
	}

	return nil
}
