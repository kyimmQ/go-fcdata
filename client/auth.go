package client

import (
	"fmt"

	"github.com/kyimmQ/go-fcdata/models"
)

const AccessTokenEndpoint = "api/v2/Market/AccessToken"

func (c *FCDataClient) Login(consumerID, consumerSecret string) (string, error) {
	reqBody := models.AuthRequest{
		ConsumerID:     consumerID,
		ConsumerSecret: consumerSecret,
	}

	var respBody models.AuthResponse
	err := c.doRequest("POST", AccessTokenEndpoint, nil, reqBody, &respBody)
	if err != nil {
		return "", err
	}

	if respBody.Status != 200 {
		return "", fmt.Errorf("login failed: %s", respBody.Message)
	}

	c.SetToken(respBody.Data.AccessToken)
	return respBody.Data.AccessToken, nil
}
