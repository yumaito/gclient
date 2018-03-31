package gclient

import (
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	googleApi "golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

type Config struct {
	PrivateKey   string `yaml:"private_key" json:"private_key"`
	PrivateKeyID string `yaml:"private_key_id" json:"private_key_id"`
	ClientEmail  string `yaml:"client_email" json:"client_email"`
}

func New(config Config, scopes ...string) (*http.Client, error) {
	if config.ClientEmail == "" {
		return nil, errors.New("client_email is required")
	}
	if config.PrivateKey == "" {
		return nil, errors.New("private_key is required")
	}
	if config.PrivateKeyID == "" {
		return nil, errors.New("private_key_id is required")
	}

	jwtConfig := &jwt.Config{
		Email:        config.ClientEmail,
		PrivateKey:   []byte(config.PrivateKey),
		PrivateKeyID: config.PrivateKeyID,
		Scopes:       scopes,
		TokenURL:     googleApi.JWTTokenURL,
	}

	client := jwtConfig.Client(oauth2.NoContext)
	return client, nil
}
