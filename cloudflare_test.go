//go:build integration

package captcha

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"sync"
	"testing"
)

func TestCloudflare(t *testing.T) {
	cf := &Cloudflare{
		client:         newClient(),
		verifyEndpoint: "https://challenges.cloudflare.com/turnstile/v0/siteverify",
		secret:         os.Getenv("CLOUDFLARE_SECRET"),
		resPool: &sync.Pool{
			New: func() interface{} {
				return new(Challenge)
			},
		},
	}

	challengeKey := os.Getenv("CLOUDFLARE_CHALLENGEKEY")

	ok, err := cf.Validate(context.Background(), challengeKey, "")
	assert.NoError(t, err)
	assert.True(t, ok)
}
