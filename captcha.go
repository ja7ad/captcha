package captcha

import (
	"context"
	"net/http"
	"net/url"
	"sync"
)

// Captcha defines the interface for validating CAPTCHA challenges from different providers.
type Captcha interface {
	// Provider returns the name of the CAPTCHA provider (e.g., "cloudflare", "google", "hcaptcha").
	Provider() string

	// Validate verifies the given challenge key with the CAPTCHA provider and
	// returns true if the challenge was successful, or false with an error if not.
	Validate(ctx context.Context, challengeKey, remoteIP string) (bool, error)

	// ValidateWithResponse verifies the challenge key and returns the full Challenge
	// response object from the provider, including success, error codes, and metadata.
	ValidateWithResponse(ctx context.Context, challengeKey, remoteIP string) (Challenge, error)
}

// New creates a new Captcha instance for the specified provider using the given secret key.
func New(provider Provider, secretKey string) (Captcha, error) {
	if secretKey == "" {
		return nil, ErrInvalidSecretKey
	}

	switch provider {
	case GoogleProvider:
		return &Google{
			provider:       provider,
			secret:         secretKey,
			client:         newClient(),
			verifyEndpoint: "https://www.google.com/recaptcha/api/siteverify",
			resPool: &sync.Pool{
				New: func() interface{} {
					return new(Challenge)
				},
			},
		}, nil
	case CloudflareProvider:
		return &Cloudflare{
			provider:       provider,
			secret:         secretKey,
			client:         newClient(),
			verifyEndpoint: "https://challenges.cloudflare.com/turnstile/v0/siteverify",
			resPool: &sync.Pool{
				New: func() interface{} {
					return new(Challenge)
				},
			},
		}, nil
	case HCaptchaProvider:
		return &HCaptcha{
			provider:       provider,
			secret:         secretKey,
			client:         newClient(),
			verifyEndpoint: "https://api.hcaptcha.com/siteverify",
			resPool: &sync.Pool{
				New: func() interface{} {
					return new(Challenge)
				},
			},
		}, nil
	default:
		return nil, ErrUnsupportedProvider
	}
}

func validate(
	ctx context.Context,
	client *client,
	pool *sync.Pool,
	endpoint string,
	secretKey string,
	challengeKey, remoteIP string,
) (*Challenge, error) {
	var res *Challenge

	obj := pool.Get()
	if obj != nil {
		res = obj.(*Challenge)
	} else {
		res = new(Challenge)
	}

	defer pool.Put(res)

	form := url.Values{}
	form.Set("secret", secretKey)
	form.Set("response", challengeKey)

	if remoteIP != "" {
		form.Set("remoteip", remoteIP)
	}

	if err := client.Request(ctx, endpoint, http.MethodPost, form, res, nil, nil); err != nil {
		return nil, err
	}

	return res, nil
}
