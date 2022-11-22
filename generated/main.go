package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/andrewbaxter/terraform-provider-stripe/shared"
	"github.com/go-playground/form/v4"
	"github.com/go-resty/resty/v2"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	plugin "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: func() *schema.Provider {
		return &schema.Provider{
			ConfigureFunc: func(d *schema.ResourceData) (any, error) {
				c := resty.New()
				c.SetBasicAuth(d.Get("token").(string), "")
				c.SetHeader("User-Agent", fmt.Sprintf("github.com/andrewbaxter/terraform-provider-stripe %s", gitHash()))
				c.SetBaseURL("https://api.stripe.com")
				return Facilitator{c: c}, nil
			},
			Schema:       map[string]*schema.Schema{"token": {Type: schema.TypeString}},
			ResourcesMap: resources(),
		}
	}})
}

type Facilitator struct {
	c   *resty.Client
	enc *form.Encoder
}

func (f *Facilitator) Post(ctx context.Context, path string, data map[string]any) (any, error) {
	res, err := f.c.R().SetFormDataFromValues(shared.Must(f.enc.Encode(data))).Post(path)
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
	res, err := f.c.R().Delete(path)
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
