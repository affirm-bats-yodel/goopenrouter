package openrouter_test

import (
	"context"
	"testing"
	"time"

	"github.com/affirm-bats-yodel/openrouter"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	c, err := openrouter.NewClient()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("GetLimits", GetLimits(c))
}

func GetLimits(c *openrouter.Client) func(t *testing.T) {
	return func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		res, err := c.GetLimits(ctx)
		if assert.NoError(t, err) {
			t.Logf("%+v", res)
			assert.NotEmpty(t, res.Label)
			if assert.NotEmpty(t, res.RateLimit) {
				assert.NotEmpty(t, res.RateLimit.Requests)
				assert.NotEmpty(t, res.RateLimit.Interval)
				assert.NotEmpty(t, res.RateLimit.GetInterval())
			}
		}
	}
}
