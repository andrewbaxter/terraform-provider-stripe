// This file isn't actually generated
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/andrewbaxter/terraform-provider-stripe/shared"
	"github.com/go-playground/form/v4"
	"github.com/go-resty/resty/v2"
	cty "github.com/hashicorp/go-cty/cty"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	plugin "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/stripe/stripe-go/v74"
	"go.uber.org/ratelimit"
)

func Json(x any) string {
	out, _ := json.MarshalIndent(x, "", "  ")
	return string(out)
}

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

func (f *Facilitator) Post(ctx context.Context, path string, data any) (any, error) {
	f.limit.Take()
	log.Printf("STRIPE POST %s\n%s\n-> %s", path, Json(data), shared.Must(f.enc.Encode(data.(map[string]any))).Encode())
	res, err := f.c.R().
		SetHeader(HeaderIdempotencyKey, stripe.NewIdempotencyKey()).
		SetFormDataFromValues(shared.Must(f.enc.Encode(data.(map[string]any)))).
		Post(path)
	if err != nil {
		return nil, err
	}
	log.Printf("STRIPE POST RESP %s %d\n[[%s]]", path, res.StatusCode(), res.Body())
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
	log.Printf("STRIPE GET %s", path)
	res, err := f.c.R().Get(path)
	if err != nil {
		return nil, err
	}
	log.Printf("STRIPE GET RESP %s %d\n[[%s]]", path, res.StatusCode(), res.Body())
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
	log.Printf("STRIPE DELETE %s", path)
	res, err := f.c.R().
		SetHeader(HeaderIdempotencyKey, stripe.NewIdempotencyKey()).
		Delete(path)
	if err != nil {
		return err
	}
	log.Printf("STRIPE DELETE RESP %s %d\n[[%s]]", path, res.StatusCode(), res.Body())
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

func inResourceData(key string, d *schema.ResourceData) bool {
	_, ok := d.GetOk(key)
	return ok
}

func fmtPath(path cty.Path) string {
	out := []string{}
	for _, k := range path {
		var e1 string
		switch e := k.(type) {
		case cty.GetAttrStep:
			e1 = e.Name
		case cty.IndexStep:
			switch e.Key.Type() {
			case cty.Number:
				e1 = e.Key.AsBigFloat().String()
			case cty.String:
				e1 = e.Key.AsString()
			default:
				e1 = "(unknown type of path index step value)"
			}
		default:
			e1 = "(unknown type of path index step)"
		}
		out = append(out, "/"+e1)
	}
	return strings.Join(out, "")
}

func NewMutList[T any]() *MutList[T] {
	return shared.Pointer(MutList[T]([]T{}))
}

type MutList[T any] []T

func (m *MutList[T]) Add(v T, ok bool) {
	if ok {
		(*m) = append(*m, v)
	}
}

func IsNotDefault[T comparable](v T) bool {
	var d T
	return v == d
}
