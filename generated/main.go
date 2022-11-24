// This file isn't actually generated
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andrewbaxter/terraform-provider-stripe/shared"
	"github.com/go-playground/form/v4"
	"github.com/go-resty/resty/v2"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	plugin "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/stripe/stripe-go/v74"
	"go.uber.org/ratelimit"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: func() *schema.Provider {
		return &schema.Provider{
			ConfigureFunc: func(d *schema.ResourceData) (any, error) {
				c := resty.New()
				c.SetBasicAuth(d.Get("token").(string), "")
				c.SetHeader("User-Agent", fmt.Sprintf("github.com/andrewbaxter/terraform-provider-stripe %s", gitHash()))
				c.SetBaseURL("https://api.stripe.com")
				c.SetRetryCount(3)
				c.AddRetryCondition(func(resp *resty.Response, err error) bool {
					// Based on https://github.com/stripe/stripe-go/blob/e6cd1eb064432902f0b195f613b00379bfb0b87a/stripe.go#L779
					req := resp.Request
					if req.Context() != nil && req.Context().Err() != nil {
						return false
					}
					if resp.Header().Get("Stripe-Should-Retry") == "false" {
						return false
					}
					if resp.Header().Get("Stripe-Should-Retry") == "true" {
						return true
					}
					switch resp.StatusCode() {
					case http.StatusConflict:
						return true
					case http.StatusTooManyRequests:
						return true
					case http.StatusServiceUnavailable:
						return true
					}
					if resp.StatusCode() >= http.StatusInternalServerError && req.Method != http.MethodPost {
						return true
					}
					return false
				})
				rateCount := d.Get("ratelimit").(int)
				return &Facilitator{
					c:     c,
					enc:   form.NewEncoder(),
					limit: ratelimit.New(rateCount, ratelimit.WithSlack(rateCount)),
				}, nil
			},
			Schema: map[string]*schema.Schema{
				"token": {
					Description: "Stripe API token",
					Type:        schema.TypeString,
					Required:    true,
				},
				"ratelimit": {
					Description: "Rate limit, max requests per second",
					Type:        schema.TypeInt,
					Optional:    true,
					Default:     10,
				},
			},
			ResourcesMap: resources(),
		}
	}})
}

const HeaderIdempotencyKey = "Idempotency-Key"

type Facilitator struct {
	c     *resty.Client
	enc   *form.Encoder
	limit ratelimit.Limiter
}

func (f *Facilitator) Post(ctx context.Context, path string, data map[string]any) (any, error) {
	f.limit.Take()
	res, err := f.c.R().
		SetHeader(HeaderIdempotencyKey, stripe.NewIdempotencyKey()).
		SetFormDataFromValues(shared.Must(f.enc.Encode(data))).
		Post(path)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, fmt.Errorf("request returned error response %d:\n%s", res.StatusCode(), res.Body())
	}
	var out any
	err = json.Unmarshal(res.Body(), &out)
	if err != nil {
		return nil, fmt.Errorf("api response was not valid json: %w\n%s", err, res.Body())
	}
	return out, nil
}

func (f *Facilitator) Get(ctx context.Context, path string) (any, error) {
	f.limit.Take()
	res, err := f.c.R().Get(path)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, fmt.Errorf("request returned error response %d:\n%s", res.StatusCode(), res.Body())
	}
	var out any
	err = json.Unmarshal(res.Body(), &out)
	if err != nil {
		return nil, fmt.Errorf("api response was not valid json: %w\n%s", err, res.Body())
	}
	return out, nil
}

func (f *Facilitator) Delete(ctx context.Context, path string) error {
	f.limit.Take()
	res, err := f.c.R().
		SetHeader(HeaderIdempotencyKey, stripe.NewIdempotencyKey()).
		Delete(path)
	if err != nil {
		return err
	}
	if res.IsError() {
		return fmt.Errorf("request returned error response %d:\n%s", res.StatusCode(), res.Body())
	}
	return nil
}

func inEnum(s string, values []string) bool {
	for _, v := range values {
		if v == s {
			return true
		}
	}
	return false
}

func inMap(key string, m map[string]any) bool {
	_, res := m[key]
	return res
}
