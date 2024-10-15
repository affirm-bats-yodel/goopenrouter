package openrouter_test

import (
	"context"
	"testing"
	"time"

	openrouter "github.com/affirm-bats-yodel/goopenrouter"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	c, err := openrouter.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("GetLimits", GetLimits(c))
	t.Run("GetModels", GetModels(c))
	t.Run("GetParameters", GetParameters(c))
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

func GetModels(c *openrouter.Client) func(t *testing.T) {
	return func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		m, err := c.GetModels(ctx)
		if assert.NoError(t, err) && assert.NotEmpty(t, m) {
			for _, elem := range m {
				if assert.NotEmpty(t, elem.ID) {
					t.Logf("model id: %s", elem.ID)
				}
				assert.NotEmpty(t, elem.Description) // just assert
				assert.NotEmpty(t, elem.Created)     // Seems like Unix timestamp
				if assert.NotEmpty(t, elem.ContextLength) {
					t.Logf("context length: %d", elem.ContextLength)
				}
				if assert.NotEmpty(t, elem.Pricing) {
					t.Logf("pricing: %+v", elem.Pricing)
					assert.NotEmpty(t, elem.Pricing.Prompt)     // "0" or floating point
					assert.NotEmpty(t, elem.Pricing.Completion) // "0" or floating point
					assert.NotEmpty(t, elem.Pricing.Request)    // "0" or floating point
					assert.NotEmpty(t, elem.Pricing.Image)      // "0" or floating point
				}
				if assert.NotEmpty(t, elem.Architecture) {
					t.Logf("architecture: %+v", elem.Architecture)
					assert.NotEmpty(t, elem.Architecture.Tokenizer)
					// assert.NotEmpty(t, elem.Architecture.InstructType) // can be empty
					assert.NotEmpty(t, elem.Architecture.Modality)
				}
				// can be empty on some models
				// uncomment if you want to check
				//
				// if assert.NotEmpty(t, elem.TopProvider) {
				// 	t.Logf("top provider: %+v", elem.TopProvider)
				// }
				// if assert.NotEmpty(t, elem.PerRequestLimits) {
				// 	t.Logf("per request limits: %+v", elem.PerRequestLimits)
				// }
			}
		}
	}
}

func GetParameters(c *openrouter.Client) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("ShouldError", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			_, err := c.GetParameters(ctx, "some/model", "SomeProvider")
			assert.Error(t, err)
		})
		t.Run("Success", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			res, err := c.GetParameters(ctx, "openai/gpt-4o-mini-2024-07-18", "OpenAI")
			if assert.NoError(t, err) {
				t.Logf("%+v", res)

				// Check Results are correct
				//
				// Reference:
				// https://openrouter.ai/openai/gpt-4o-mini-2024-07-18/parameters

				/**
				 * temperature
				 ** p10: 0
				 ** p50: 0.5
				 ** p90: 1
				 */
				assert.Equal(t, res.TemperatureP10, float64(0))
				assert.Equal(t, res.TemperatureP50, float64(0.5))
				assert.Equal(t, res.TemperatureP90, float64(1))

				/**
				 * top_p
				 ** p10: 1
				 ** p50: 1
				 ** p90: 1
				 */
				assert.Equal(t, res.TopPP10, float64(1))
				assert.Equal(t, res.TopPP50, float64(1))
				assert.Equal(t, res.TopPP90, float64(1))

				/**
				 * repetition_penalty
				 ** p10: 1
				 ** p50: 1
				 ** p90: 1
				 */
				assert.Equal(t, res.RepetitionPenaltyP10, float64(1))
				assert.Equal(t, res.RepetitionPenaltyP50, float64(1))
				assert.Equal(t, res.RepetitionPenaltyP90, float64(1))
			}
		})
	}
}
