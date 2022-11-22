package main

import (
	"context"
	"fmt"
	shared "github.com/andrewbaxter/terraform-provider-stripe/shared"
	cty "github.com/hashicorp/go-cty/cty"
	diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resources_invoiceitems() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
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
			if v, exists := d.GetOk("amount"); exists {
				params["amount"] = v
			}
			{
				v := d.Get("customer")
				params["customer"] = v
			}
			if v, exists := d.GetOk("description"); exists {
				params["description"] = v
			}
			if v, exists := d.GetOk("price"); exists {
				params["price"] = v
			}
			if v, exists := d.GetOk("discountable"); exists {
				params["discountable"] = v
			}
			if v, exists := d.GetOk("tax_rates"); exists {
				params["tax_rates"] = v
			}
			if v, exists := d.GetOk("unit_amount_decimal"); exists {
				params["unit_amount_decimal"] = v
			}
			if v, exists := d.GetOk("unit_amount"); exists {
				params["unit_amount"] = v
			}
			if v, exists := d.GetOk("period"); exists {
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("end")
						v := shared.Dig[any](outerV, "end")
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
						path := path.IndexString("start")
						v := shared.Dig[any](outerV, "start")
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
				params["period"] = v
			}
			if v, exists := d.GetOk("price_data"); exists {
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("currency")
						v := shared.Dig[any](outerV, "currency")
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
						path := path.IndexString("product")
						v := shared.Dig[any](outerV, "product")
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
						path := path.IndexString("tax_behavior")
						v := shared.Dig[any](outerV, "tax_behavior")
						if v != nil {
							if inEnum(v.(string), []string{"exclusive", "inclusive", "unspecified"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: exclusive, inclusive, unspecified"),
								})
							}
						}
					}
				}
				params["price_data"] = v
			}
			if v, exists := d.GetOk("quantity"); exists {
				params["quantity"] = v
			}
			if v, exists := d.GetOk("tax_code"); exists {
				params["tax_code"] = v
			}
			if v, exists := d.GetOk("currency"); exists {
				params["currency"] = v
			}
			if v, exists := d.GetOk("discounts"); exists {
				params["discounts"] = v
			}
			if v, exists := d.GetOk("invoice"); exists {
				params["invoice"] = v
			}
			if v, exists := d.GetOk("subscription"); exists {
				params["subscription"] = v
			}
			res, err := f.Post(ctx, "/v1/invoiceitems", params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to create new invoiceitem: %s", err),
				})
				return out
			}
			d.SetId(shared.Dig[string](res, "id"))
			return out
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			err := f.Delete(ctx, fmt.Sprintf("/v1/invoiceitems/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to delete invoiceitem %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			res, err := f.Get(ctx, fmt.Sprintf("/v1/invoiceitems/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to look up invoiceitem %s: %s", d.Id(), err),
				})
				return out
			}
			d.Set("tax_behavior", shared.Dig[any](res, "tax_behavior"))
			d.Set("amount", shared.Dig[any](res, "amount"))
			d.Set("customer", shared.Dig[any](res, "customer"))
			d.Set("description", shared.Dig[any](res, "description"))
			d.Set("price", shared.Dig[any](res, "price"))
			d.Set("discountable", shared.Dig[any](res, "discountable"))
			d.Set("tax_rates", shared.Dig[any](res, "tax_rates"))
			d.Set("unit_amount_decimal", shared.Dig[any](res, "unit_amount_decimal"))
			d.Set("unit_amount", shared.Dig[any](res, "unit_amount"))
			d.Set("period", shared.Dig[any](res, "period"))
			d.Set("price_data", shared.Dig[any](res, "price_data"))
			d.Set("quantity", shared.Dig[any](res, "quantity"))
			d.Set("tax_code", shared.Dig[any](res, "tax_code"))
			d.Set("currency", shared.Dig[any](res, "currency"))
			d.Set("discounts", shared.Dig[any](res, "discounts"))
			d.Set("invoice", shared.Dig[any](res, "invoice"))
			d.Set("subscription", shared.Dig[any](res, "subscription"))
			return out
		},
		Schema: map[string]*schema.Schema{
			"amount": {
				Description: "The integer amount in cents (or local equivalent) of the charge to be applied to the upcoming invoice. Passing in a negative `amount` will reduce the `amount_due` on the invoice.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"currency": {
				Description: "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"customer": {
				Description: "The ID of the customer who will be billed when this invoice item is billed.",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
			},
			"description": {
				Description: "An arbitrary string which you can attach to the invoice item. The description is displayed in the invoice for easy tracking.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"discountable": {
				Description: "Controls whether discounts apply to this invoice item. Defaults to false for prorations or negative invoice items, and true for all other invoice items.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeBool,
			},
			"discounts": {
				Description: "The coupons to redeem into discounts for the invoice item or invoice line item.",
				Elem: &schema.Schema{
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"coupon": {
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"discount": {
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
					}},
					Type: schema.TypeMap,
				},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeList,
			},
			"invoice": {
				Description: "The ID of an existing invoice to add this invoice item to. When left blank, the invoice item will be added to the next upcoming scheduled invoice. This is useful when adding invoice items in response to an invoice.created webhook. You can only add invoice items to draft invoices and there is a maximum of 250 items per invoice.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"period": {
				Description: "The period associated with this invoice item. When set to different values, the period will be rendered on the invoice.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"end": {
						Description: "",
						Required:    true,
						Type:        schema.TypeInt,
					},
					"start": {
						Description: "",
						Required:    true,
						Type:        schema.TypeInt,
					},
				}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"price": {
				Description: "The ID of the price object.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"price_data": {
				Description: "Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"currency": {
						Description: "",
						Required:    true,
						Type:        schema.TypeString,
					},
					"product": {
						Description: "",
						Required:    true,
						Type:        schema.TypeString,
					},
					"tax_behavior": {
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
				}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"quantity": {
				Description: "Non-negative integer. The quantity of units for the invoice item.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"subscription": {
				Description: "The ID of a subscription to add this invoice item to. When left blank, the invoice item will be be added to the next upcoming scheduled invoice. When set, scheduled invoices for subscriptions other than the specified subscription will ignore the invoice item. Use this when you want to express that an invoice item has been accrued within the context of a particular subscription.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"tax_behavior": {
				Description: "Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"tax_code": {
				Description: "A [tax code](https://stripe.com/docs/tax/tax-categories) ID.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"tax_rates": {
				Description: "The tax rates which apply to the invoice item. When set, the `default_tax_rates` on the invoice do not apply to this invoice item.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeList,
			},
			"unit_amount": {
				Description: "The integer unit amount in cents (or local equivalent) of the charge to be applied to the upcoming invoice. This `unit_amount` will be multiplied by the quantity to get the full amount. Passing in a negative `unit_amount` will reduce the `amount_due` on the invoice.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"unit_amount_decimal": {
				Description: "Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
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
			if d.HasChange("amount") {
				v := d.Get("amount")
				params["amount"] = v
			}
			if d.HasChange("description") {
				v := d.Get("description")
				params["description"] = v
			}
			if d.HasChange("price") {
				v := d.Get("price")
				params["price"] = v
			}
			if d.HasChange("discountable") {
				v := d.Get("discountable")
				params["discountable"] = v
			}
			if d.HasChange("tax_rates") {
				v := d.Get("tax_rates")
				params["tax_rates"] = v
			}
			if d.HasChange("unit_amount_decimal") {
				v := d.Get("unit_amount_decimal")
				params["unit_amount_decimal"] = v
			}
			if d.HasChange("unit_amount") {
				v := d.Get("unit_amount")
				params["unit_amount"] = v
			}
			if d.HasChange("period") {
				v := d.Get("period")
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("end")
						v := shared.Dig[any](outerV, "end")
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
						path := path.IndexString("start")
						v := shared.Dig[any](outerV, "start")
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
				params["period"] = v
			}
			if d.HasChange("price_data") {
				v := d.Get("price_data")
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("currency")
						v := shared.Dig[any](outerV, "currency")
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
						path := path.IndexString("product")
						v := shared.Dig[any](outerV, "product")
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
						path := path.IndexString("tax_behavior")
						v := shared.Dig[any](outerV, "tax_behavior")
						if v != nil {
							if inEnum(v.(string), []string{"exclusive", "inclusive", "unspecified"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: exclusive, inclusive, unspecified"),
								})
							}
						}
					}
				}
				params["price_data"] = v
			}
			if d.HasChange("quantity") {
				v := d.Get("quantity")
				params["quantity"] = v
			}
			if d.HasChange("tax_code") {
				v := d.Get("tax_code")
				params["tax_code"] = v
			}
			if d.HasChange("discounts") {
				v := d.Get("discounts")
				params["discounts"] = v
			}
			_, err := f.Post(ctx, fmt.Sprintf("/v1/invoiceitems/%v", d.Get("id")), params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to update invoiceitem %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
	}
}
