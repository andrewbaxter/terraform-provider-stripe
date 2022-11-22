package main

import (
	"context"
	"fmt"
	shared "github.com/andrewbaxter/terraform-provider-stripe/shared"
	cty "github.com/hashicorp/go-cty/cty"
	diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resources_promotion_codes() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if v, exists := d.GetOk("expires_at"); exists {
				params["expires_at"] = v
			}
			if v, exists := d.GetOk("customer"); exists {
				params["customer"] = v
			}
			if v, exists := d.GetOk("max_redemptions"); exists {
				params["max_redemptions"] = v
			}
			if v, exists := d.GetOk("restrictions"); exists {
				params["restrictions"] = v
			}
			if v, exists := d.GetOk("code"); exists {
				params["code"] = v
			}
			{
				v := d.Get("coupon")
				params["coupon"] = v
			}
			res, err := f.Post(ctx, "/v1/promotion_codes", params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to create new promotion_code: %s", err),
				})
				return out
			}
			d.SetId(shared.Dig[string](res, "id"))
			return out
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			_, err := f.Post(ctx, fmt.Sprintf("/v1/promotion_codes/%v", d.Get("id")), map[string]any{"active": false})
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to set active to false on promotion_code %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			res, err := f.Get(ctx, fmt.Sprintf("/v1/promotion_codes/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to look up promotion_code %s: %s", d.Id(), err),
				})
				return out
			}
			d.Set("expires_at", shared.Dig[any](res, "expires_at"))
			d.Set("customer", shared.Dig[any](res, "customer"))
			d.Set("max_redemptions", shared.Dig[any](res, "max_redemptions"))
			d.Set("restrictions", shared.Dig[any](res, "restrictions"))
			d.Set("code", shared.Dig[any](res, "code"))
			d.Set("coupon", shared.Dig[any](res, "coupon"))
			return out
		},
		Schema: map[string]*schema.Schema{
			"code": {
				Description: "The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for a specific customer. If left blank, we will generate one automatically.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"coupon": {
				Description: "The coupon for this promotion code.",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
			},
			"customer": {
				Description: "The customer that this promotion code can be used by. If not set, the promotion code can be used by all customers.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"expires_at": {
				Description: "The timestamp at which this promotion code will expire. If the coupon has specified a `redeems_by`, then this value cannot be after the coupon's `redeems_by`.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"max_redemptions": {
				Description: "A positive integer specifying the number of times the promotion code can be redeemed. If the coupon has specified a `max_redemptions`, then this value cannot be greater than the coupon's `max_redemptions`.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"restrictions": {
				Description: "Settings that restrict the redemption of the promotion code.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"currency_options": {
						Description: "",
						Elem:        &schema.Resource{Schema: map[string]*schema.Schema{}},
						Required:    false,
						Type:        schema.TypeMap,
					},
					"first_time_transaction": {
						Description: "",
						Required:    false,
						Type:        schema.TypeBool,
					},
					"minimum_amount": {
						Description: "",
						Required:    false,
						Type:        schema.TypeInt,
					},
					"minimum_amount_currency": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
				}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if d.HasChange("restrictions") {
				v := d.Get("restrictions")
				params["restrictions"] = v
			}
			_, err := f.Post(ctx, fmt.Sprintf("/v1/promotion_codes/%v", d.Get("id")), params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to update promotion_code %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
	}
}
