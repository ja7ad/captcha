package captcha

import (
	"context"
	"sync"
)

type Google struct {
	client         *client
	resPool        *sync.Pool
	provider       Provider
	secret         string
	verifyEndpoint string
}

func (g *Google) Provider() string {
	return g.provider.String()
}

func (g *Google) Validate(ctx context.Context, challengeKey, remoteIP string) (bool, error) {
	challenge, err := validate(ctx, g.client, g.resPool, g.verifyEndpoint, g.secret, challengeKey, remoteIP)
	if err != nil {
		return false, err
	}

	return challenge.Success, nil
}

func (g *Google) ValidateWithResponse(ctx context.Context, challengeKey, remoteIP string) (Challenge, error) {
	challenge, err := validate(ctx, g.client, g.resPool, g.verifyEndpoint, g.secret, challengeKey, remoteIP)
	if err != nil {
		return Challenge{}, err
	}

	return *challenge, nil
}
