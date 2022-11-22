package main

import (
	"context"
	"fmt"
	shared "github.com/andrewbaxter/terraform-provider-stripe/shared"
	cty "github.com/hashicorp/go-cty/cty"
	diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resources_prices() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if v, exists := d.GetOk("lookup_key"); exists {
				params["lookup_key"] = v
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
			if v, exists := d.GetOk("currency_options"); exists {
				params["currency_options"] = v
			}
			if v, exists := d.GetOk("product"); exists {
				params["product"] = v
			}
			if v, exists := d.GetOk("product_data"); exists {
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("name")
						v := shared.Dig[any](outerV, "name")
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
				params["product_data"] = v
			}
			if v, exists := d.GetOk("recurring"); exists {
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("aggregate_usage")
						v := shared.Dig[any](outerV, "aggregate_usage")
						if v != nil {
							if inEnum(v.(string), []string{"last_during_period", "last_ever", "max", "sum"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: last_during_period, last_ever, max, sum"),
								})
							}
						}
					}
					{
						path := path.IndexString("interval")
						v := shared.Dig[any](outerV, "interval")
						if v != nil {
							if inEnum(v.(string), []string{"day", "month", "week", "year"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: day, month, week, year"),
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
					{
						path := path.IndexString("usage_type")
						v := shared.Dig[any](outerV, "usage_type")
						if v != nil {
							if inEnum(v.(string), []string{"licensed", "metered"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: licensed, metered"),
								})
							}
						}
					}
				}
				params["recurring"] = v
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
			if v, exists := d.GetOk("transform_quantity"); exists {
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
				params["transform_quantity"] = v
			}
			if v, exists := d.GetOk("custom_unit_amount"); exists {
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("enabled")
						v := shared.Dig[any](outerV, "enabled")
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
				params["custom_unit_amount"] = v
			}
			if v, exists := d.GetOk("nickname"); exists {
				params["nickname"] = v
			}
			if v, exists := d.GetOk("transfer_lookup_key"); exists {
				params["transfer_lookup_key"] = v
			}
			if v, exists := d.GetOk("unit_amount"); exists {
				params["unit_amount"] = v
			}
			if v, exists := d.GetOk("unit_amount_decimal"); exists {
				params["unit_amount_decimal"] = v
			}
			{
				v := d.Get("currency")
				params["currency"] = v
			}
			if v, exists := d.GetOk("tax_behavior"); exists {
				if inEnum(v.(string), []string{"exclusive", "inclusive", "unspecified"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: exclusive, inclusive, unspecified"),
					})
				}
				params["tax_behavior"] = v
			}
			res, err := f.Post(ctx, "/v1/prices", params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to create new price: %s", err),
				})
				return out
			}
			d.SetId(shared.Dig[string](res, "id"))
			return out
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			_, err := f.Post(ctx, fmt.Sprintf("/v1/prices/%v", d.Get("id")), map[string]any{"active": false})
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to set active to false on price %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			res, err := f.Get(ctx, fmt.Sprintf("/v1/prices/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to look up price %s: %s", d.Id(), err),
				})
				return out
			}
			d.Set("lookup_key", shared.Dig[any](res, "lookup_key"))
			d.Set("billing_scheme", shared.Dig[any](res, "billing_scheme"))
			d.Set("currency_options", shared.Dig[any](res, "currency_options"))
			d.Set("product", shared.Dig[any](res, "product"))
			d.Set("product_data", shared.Dig[any](res, "product_data"))
			d.Set("recurring", shared.Dig[any](res, "recurring"))
			d.Set("tiers", shared.Dig[any](res, "tiers"))
			d.Set("tiers_mode", shared.Dig[any](res, "tiers_mode"))
			d.Set("transform_quantity", shared.Dig[any](res, "transform_quantity"))
			d.Set("custom_unit_amount", shared.Dig[any](res, "custom_unit_amount"))
			d.Set("nickname", shared.Dig[any](res, "nickname"))
			d.Set("transfer_lookup_key", shared.Dig[any](res, "transfer_lookup_key"))
			d.Set("unit_amount", shared.Dig[any](res, "unit_amount"))
			d.Set("unit_amount_decimal", shared.Dig[any](res, "unit_amount_decimal"))
			d.Set("currency", shared.Dig[any](res, "currency"))
			d.Set("tax_behavior", shared.Dig[any](res, "tax_behavior"))
			return out
		},
		Schema: map[string]*schema.Schema{
			"billing_scheme": {
				Description: "Describes how to compute the price per period. Either `per_unit` or `tiered`. `per_unit` indicates that the fixed amount (specified in `unit_amount` or `unit_amount_decimal`) will be charged per unit in `quantity` (for prices with `usage_type=licensed`), or per unit of total usage (for prices with `usage_type=metered`). `tiered` indicates that the unit pricing will be computed using a tiering strategy as defined using the `tiers` and `tiers_mode` attributes.",
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
			"currency_options": {
				Description: "Prices defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).",
				Elem:        &schema.Resource{Schema: map[string]*schema.Schema{}},
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeMap,
			},
			"custom_unit_amount": {
				Description: "When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"enabled": {
						Description: "",
						Required:    true,
						Type:        schema.TypeBool,
					},
					"maximum": {
						Description: "",
						Required:    false,
						Type:        schema.TypeInt,
					},
					"minimum": {
						Description: "",
						Required:    false,
						Type:        schema.TypeInt,
					},
					"preset": {
						Description: "",
						Required:    false,
						Type:        schema.TypeInt,
					},
				}},
				ForceNew: true,
				Required: false,
				Type:     schema.TypeMap,
			},
			"lookup_key": {
				Description: "A lookup key used to retrieve prices dynamically from a static string. This may be up to 200 characters.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"nickname": {
				Description: "A brief description of the price, hidden from customers.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"product": {
				Description: "The ID of the product that this price will belong to.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"product_data": {
				Description: "These fields can be used to create a new product that this price will belong to.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"active": {
						Description: "",
						Required:    false,
						Type:        schema.TypeBool,
					},
					"id": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"metadata": {
						Description: "",
						Elem:        &schema.Resource{Schema: map[string]*schema.Schema{}},
						Required:    false,
						Type:        schema.TypeMap,
					},
					"name": {
						Description: "",
						Required:    true,
						Type:        schema.TypeString,
					},
					"statement_descriptor": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"tax_code": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"unit_label": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
				}},
				ForceNew: true,
				Required: false,
				Type:     schema.TypeMap,
			},
			"recurring": {
				Description: "The recurring components of a price such as `interval` and `usage_type`.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"aggregate_usage": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"interval": {
						Description: "",
						Required:    true,
						Type:        schema.TypeString,
					},
					"interval_count": {
						Description: "",
						Required:    false,
						Type:        schema.TypeInt,
					},
					"usage_type": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
				}},
				ForceNew: true,
				Required: false,
				Type:     schema.TypeMap,
			},
			"tax_behavior": {
				Description: "Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
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
			"transfer_lookup_key": {
				Description: "If set to true, will atomically remove the lookup key from the existing price, and assign it to this price.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeBool,
			},
			"transform_quantity": {
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
			"unit_amount": {
				Description: "A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge. One of `unit_amount` or `custom_unit_amount` is required, unless `billing_scheme=tiered`.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"unit_amount_decimal": {
				Description: "Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if d.HasChange("lookup_key") {
				v := d.Get("lookup_key")
				params["lookup_key"] = v
			}
			if d.HasChange("currency_options") {
				v := d.Get("currency_options")
				params["currency_options"] = v
			}
			if d.HasChange("nickname") {
				v := d.Get("nickname")
				params["nickname"] = v
			}
			if d.HasChange("transfer_lookup_key") {
				v := d.Get("transfer_lookup_key")
				params["transfer_lookup_key"] = v
			}
			if d.HasChange("tax_behavior") {
				v := d.Get("tax_behavior")
				if inEnum(v.(string), []string{"exclusive", "inclusive", "unspecified"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: exclusive, inclusive, unspecified"),
					})
				}
				params["tax_behavior"] = v
			}
			_, err := f.Post(ctx, fmt.Sprintf("/v1/prices/%v", d.Get("id")), params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to update price %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
	}
}
