package captcha

import (
	"context"
	"sync"
)

type HCaptcha struct {
	client         *client
	resPool        *sync.Pool
	provider       Provider
	secret         string
	verifyEndpoint string
}

func (h *HCaptcha) Provider() string {
	return h.provider.String()
}

func (h *HCaptcha) Validate(ctx context.Context, challengeKey, remoteIP string) (bool, error) {
	challenge, err := validate(ctx, h.client, h.resPool, h.verifyEndpoint, h.secret, challengeKey, remoteIP)
	if err != nil {
		return false, err
	}

	return challenge.Success, nil
}

func (h *HCaptcha) ValidateWithResponse(ctx context.Context, challengeKey, remoteIP string) (Challenge, error) {
	challenge, err := validate(ctx, h.client, h.resPool, h.verifyEndpoint, h.secret, challengeKey, remoteIP)
	if err != nil {
		return Challenge{}, err
	}

	return *challenge, nil
}
