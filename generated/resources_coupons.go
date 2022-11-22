package main

import (
	"context"
	"fmt"
	shared "github.com/andrewbaxter/terraform-provider-stripe/shared"
	cty "github.com/hashicorp/go-cty/cty"
	diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resources_coupons() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if v, exists := d.GetOk("applies_to"); exists {
				params["applies_to"] = v
			}
			if v, exists := d.GetOk("currency"); exists {
				params["currency"] = v
			}
			if v, exists := d.GetOk("currency_options"); exists {
				params["currency_options"] = v
			}
			if v, exists := d.GetOk("duration"); exists {
				if inEnum(v.(string), []string{"forever", "once", "repeating"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: forever, once, repeating"),
					})
				}
				params["duration"] = v
			}
			if v, exists := d.GetOk("duration_in_months"); exists {
				params["duration_in_months"] = v
			}
			if v, exists := d.GetOk("name"); exists {
				params["name"] = v
			}
			if v, exists := d.GetOk("percent_off"); exists {
				params["percent_off"] = v
			}
			if v, exists := d.GetOk("amount_off"); exists {
				params["amount_off"] = v
			}
			if v, exists := d.GetOk("id"); exists {
				params["id"] = v
			}
			if v, exists := d.GetOk("max_redemptions"); exists {
				params["max_redemptions"] = v
			}
			if v, exists := d.GetOk("redeem_by"); exists {
				params["redeem_by"] = v
			}
			res, err := f.Post(ctx, "/v1/coupons", params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to create new coupon: %s", err),
				})
				return out
			}
			d.SetId(shared.Dig[string](res, "id"))
			return out
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			err := f.Delete(ctx, fmt.Sprintf("/v1/coupons/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to delete coupon %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			res, err := f.Get(ctx, fmt.Sprintf("/v1/coupons/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to look up coupon %s: %s", d.Id(), err),
				})
				return out
			}
			d.Set("applies_to", shared.Dig[any](res, "applies_to"))
			d.Set("currency", shared.Dig[any](res, "currency"))
			d.Set("currency_options", shared.Dig[any](res, "currency_options"))
			d.Set("duration", shared.Dig[any](res, "duration"))
			d.Set("duration_in_months", shared.Dig[any](res, "duration_in_months"))
			d.Set("name", shared.Dig[any](res, "name"))
			d.Set("percent_off", shared.Dig[any](res, "percent_off"))
			d.Set("amount_off", shared.Dig[any](res, "amount_off"))
			d.Set("id", shared.Dig[any](res, "id"))
			d.Set("max_redemptions", shared.Dig[any](res, "max_redemptions"))
			d.Set("redeem_by", shared.Dig[any](res, "redeem_by"))
			return out
		},
		Schema: map[string]*schema.Schema{
			"amount_off": {
				Description: "A positive integer representing the amount to subtract from an invoice total (required if `percent_off` is not passed).",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"applies_to": {
				Description: "A hash containing directions for what this Coupon will apply discounts to.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"products": {
					Description: "",
					Elem:        &schema.Schema{Type: schema.TypeString},
					Required:    false,
					Type:        schema.TypeList,
				}}},
				ForceNew: true,
				Required: false,
				Type:     schema.TypeMap,
			},
			"currency": {
				Description: "Three-letter [ISO code for the currency](https://stripe.com/docs/currencies) of the `amount_off` parameter (required if `amount_off` is passed).",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"currency_options": {
				Description: "Coupons defined in each available currency option (only supported if `amount_off` is passed). Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).",
				Elem:        &schema.Resource{Schema: map[string]*schema.Schema{}},
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeMap,
			},
			"duration": {
				Description: "Specifies how long the discount will be in effect if used on a subscription. Defaults to `once`.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"duration_in_months": {
				Description: "Required only if `duration` is `repeating`, in which case it must be a positive integer that specifies the number of months the discount will be in effect.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"id": {
				Description: "Unique string of your choice that will be used to identify this coupon when applying it to a customer. If you don't want to specify a particular code, you can leave the ID blank and we'll generate a random code for you.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"max_redemptions": {
				Description: "A positive integer specifying the number of times the coupon can be redeemed before it's no longer valid. For example, you might have a 50% off coupon that the first 20 readers of your blog can use.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"name": {
				Description: "Name of the coupon displayed to customers on, for instance invoices, or receipts. By default the `id` is shown if `name` is not set.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"percent_off": {
				Description: "A positive float larger than 0, and smaller or equal to 100, that represents the discount the coupon will apply (required if `amount_off` is not passed).",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeFloat,
			},
			"redeem_by": {
				Description: "Unix timestamp specifying the last time at which the coupon can be redeemed. After the redeem_by date, the coupon can no longer be applied to new customers.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeInt,
			},
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if d.HasChange("currency_options") {
				v := d.Get("currency_options")
				params["currency_options"] = v
			}
			if d.HasChange("name") {
				v := d.Get("name")
				params["name"] = v
			}
			_, err := f.Post(ctx, fmt.Sprintf("/v1/coupons/%v", d.Get("id")), params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to update coupon %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
	}
}
