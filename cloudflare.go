package captcha

import (
	"context"
	"sync"
)

type Cloudflare struct {
	client         *client
	resPool        *sync.Pool
	provider       Provider
	secret         string
	verifyEndpoint string
}

func (c *Cloudflare) Provider() string {
	return c.provider.String()
}

func (c *Cloudflare) Validate(ctx context.Context, challengeKey, remoteIP string) (bool, error) {
	challenge, err := validate(ctx, c.client, c.resPool, c.verifyEndpoint, c.secret, challengeKey, remoteIP)
	if err != nil {
		return false, err
	}

	return challenge.Success, nil
}

func (c *Cloudflare) ValidateWithResponse(ctx context.Context, challengeKey, remoteIP string) (Challenge, error) {
	challenge, err := validate(ctx, c.client, c.resPool, c.verifyEndpoint, c.secret, challengeKey, remoteIP)
	if err != nil {
		return Challenge{}, err
	}

	return *challenge, nil
}
