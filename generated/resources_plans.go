package main

import (
	"context"
	"fmt"
	shared "github.com/andrewbaxter/terraform-provider-stripe/shared"
	cty "github.com/hashicorp/go-cty/cty"
	diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resources_plans() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if v, exists := d.GetOk("id"); exists {
				params["id"] = v
			}
			if v, exists := d.GetOk("interval_count"); exists {
				params["interval_count"] = v
			}
			if v, exists := d.GetOk("product"); exists {
				params["product"] = v
			}
			if v, exists := d.GetOk("tiers"); exists {
				{
					path := cty.Path{}
					outerV := v
					for i, v := range outerV.([]any) {
						{
							path := path.IndexInt(i)
							outerV := v
							{
								path := path.IndexString("up_to")
								v := shared.Dig[any](outerV, "up_to")
								if v != nil {
								} else {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       "Field is missing but required.",
									})
								}
							}
						}
					}
				}
				params["tiers"] = v
			}
			if v, exists := d.GetOk("trial_period_days"); exists {
				params["trial_period_days"] = v
			}
			if v, exists := d.GetOk("amount"); exists {
				params["amount"] = v
			}
			if v, exists := d.GetOk("amount_decimal"); exists {
				params["amount_decimal"] = v
			}
			{
				v := d.Get("currency")
				params["currency"] = v
			}
			if v, exists := d.GetOk("nickname"); exists {
				params["nickname"] = v
			}
			if v, exists := d.GetOk("aggregate_usage"); exists {
				if inEnum(v.(string), []string{"last_during_period", "last_ever", "max", "sum"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: last_during_period, last_ever, max, sum"),
					})
				}
				params["aggregate_usage"] = v
			}
			if v, exists := d.GetOk("tiers_mode"); exists {
				if inEnum(v.(string), []string{"graduated", "volume"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: graduated, volume"),
					})
				}
				params["tiers_mode"] = v
			}
			if v, exists := d.GetOk("usage_type"); exists {
				if inEnum(v.(string), []string{"licensed", "metered"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: licensed, metered"),
					})
				}
				params["usage_type"] = v
			}
			if v, exists := d.GetOk("billing_scheme"); exists {
				if inEnum(v.(string), []string{"per_unit", "tiered"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: per_unit, tiered"),
					})
				}
				params["billing_scheme"] = v
			}
			{
				v := d.Get("interval")
				if inEnum(v.(string), []string{"day", "month", "week", "year"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: day, month, week, year"),
					})
				}
				params["interval"] = v
			}
			if v, exists := d.GetOk("transform_usage"); exists {
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("divide_by")
						v := shared.Dig[any](outerV, "divide_by")
						if v != nil {
						} else {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       "Field is missing but required.",
							})
						}
					}
					{
						path := path.IndexString("round")
						v := shared.Dig[any](outerV, "round")
						if v != nil {
							if inEnum(v.(string), []string{"down", "up"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: down, up"),
								})
							}
						} else {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       "Field is missing but required.",
							})
						}
					}
				}
				params["transform_usage"] = v
			}
			res, err := f.Post(ctx, "/v1/plans", params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to create new plan: %s", err),
				})
				return out
			}
			d.SetId(shared.Dig[string](res, "id"))
			return out
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			err := f.Delete(ctx, fmt.Sprintf("/v1/plans/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to delete plan %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			res, err := f.Get(ctx, fmt.Sprintf("/v1/plans/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to look up plan %s: %s", d.Id(), err),
				})
				return out
			}
			d.Set("id", shared.Dig[any](res, "id"))
			d.Set("interval_count", shared.Dig[any](res, "interval_count"))
			d.Set("product", shared.Dig[any](res, "product"))
			d.Set("tiers", shared.Dig[any](res, "tiers"))
			d.Set("trial_period_days", shared.Dig[any](res, "trial_period_days"))
			d.Set("amount", shared.Dig[any](res, "amount"))
			d.Set("amount_decimal", shared.Dig[any](res, "amount_decimal"))
			d.Set("currency", shared.Dig[any](res, "currency"))
			d.Set("nickname", shared.Dig[any](res, "nickname"))
			d.Set("aggregate_usage", shared.Dig[any](res, "aggregate_usage"))
			d.Set("tiers_mode", shared.Dig[any](res, "tiers_mode"))
			d.Set("usage_type", shared.Dig[any](res, "usage_type"))
			d.Set("billing_scheme", shared.Dig[any](res, "billing_scheme"))
			d.Set("interval", shared.Dig[any](res, "interval"))
			d.Set("transform_usage", shared.Dig[any](res, "transform_usage"))
			return out
		},
		Schema: map[string]*schema.Schema{
			"aggregate_usage": {
				Description: "Specifies a usage aggregation strategy for plans of `usage_type=metered`. Allowed values are `sum` for summing up all usage during a period, `last_during_period` for using the last usage record reported within a period, `last_ever` for using the last usage record ever (across period bounds) or `max` which uses the usage record with the maximum reported usage during a period. Defaults to `sum`.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"amount": {
				Description: "A positive integer in cents (or local equivalent) (or 0 for a free plan) representing how much to charge on a recurring basis.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"amount_decimal": {
				Description: "Same as `amount`, but accepts a decimal value with at most 12 decimal places. Only one of `amount` and `amount_decimal` can be set.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"billing_scheme": {
				Description: "Describes how to compute the price per period. Either `per_unit` or `tiered`. `per_unit` indicates that the fixed amount (specified in `amount`) will be charged per unit in `quantity` (for plans with `usage_type=licensed`), or per unit of total usage (for plans with `usage_type=metered`). `tiered` indicates that the unit pricing will be computed using a tiering strategy as defined using the `tiers` and `tiers_mode` attributes.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"currency": {
				Description: "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
			},
			"id": {
				Description: "An identifier randomly generated by Stripe. Used to identify this plan when subscribing a customer. You can optionally override this ID, but the ID must be unique across all plans in your Stripe account. You can, however, use the same plan ID in both live and test modes.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"interval": {
				Description: "Specifies billing frequency. Either `day`, `week`, `month` or `year`.",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
			},
			"interval_count": {
				Description: "The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of one year interval allowed (1 year, 12 months, or 52 weeks).",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"nickname": {
				Description: "A brief description of the plan, hidden from customers.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"product": {
				Description: "",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"tiers": {
				Description: "Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.",
				Elem: &schema.Schema{
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"flat_amount": {
							Description: "",
							Required:    false,
							Type:        schema.TypeInt,
						},
						"flat_amount_decimal": {
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"unit_amount": {
							Description: "",
							Required:    false,
							Type:        schema.TypeInt,
						},
						"unit_amount_decimal": {
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"up_to": {
							Description: "",
							Required:    true,
							Type:        schema.TypeInt,
						},
					}},
					Type: schema.TypeMap,
				},
				ForceNew: true,
				Required: false,
				Type:     schema.TypeList,
			},
			"tiers_mode": {
				Description: "Defines if the tiering price should be `graduated` or `volume` based. In `volume`-based tiering, the maximum quantity within a period determines the per unit price, in `graduated` tiering pricing can successively change as the quantity grows.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"transform_usage": {
				Description: "Apply a transformation to the reported usage or set quantity before computing the billed price. Cannot be combined with `tiers`.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"divide_by": {
						Description: "",
						Required:    true,
						Type:        schema.TypeInt,
					},
					"round": {
						Description: "",
						Required:    true,
						Type:        schema.TypeString,
					},
				}},
				ForceNew: true,
				Required: false,
				Type:     schema.TypeMap,
			},
			"trial_period_days": {
				Description: "Default number of trial days when subscribing a customer to this plan using [`trial_from_plan=true`](https://stripe.com/docs/api#create_subscription-trial_from_plan).",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"usage_type": {
				Description: "Configures how the quantity per period should be determined. Can be either `metered` or `licensed`. `licensed` automatically bills the `quantity` set when adding it to a subscription. `metered` aggregates the total usage based on usage records. Defaults to `licensed`.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if d.HasChange("product") {
				v := d.Get("product")
				params["product"] = v
			}
			if d.HasChange("trial_period_days") {
				v := d.Get("trial_period_days")
				params["trial_period_days"] = v
			}
			if d.HasChange("nickname") {
				v := d.Get("nickname")
				params["nickname"] = v
			}
			_, err := f.Post(ctx, fmt.Sprintf("/v1/plans/%v", d.Get("id")), params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to update plan %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
	}
}
