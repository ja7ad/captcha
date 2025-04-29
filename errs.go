package captcha

import "errors"

var (
	ErrUnsupportedProvider = errors.New("unsupported provider")
	ErrInvalidSecretKey    = errors.New("invalid secret key")
)
