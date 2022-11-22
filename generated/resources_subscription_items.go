package main

import (
	"context"
	"fmt"
	shared "github.com/andrewbaxter/terraform-provider-stripe/shared"
	cty "github.com/hashicorp/go-cty/cty"
	diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resources_subscription_items() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if v, exists := d.GetOk("billing_thresholds"); exists {
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("usage_gte")
						v := shared.Dig[any](outerV, "usage_gte")
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
				params["billing_thresholds"] = v
			}
			if v, exists := d.GetOk("payment_behavior"); exists {
				if inEnum(v.(string), []string{"allow_incomplete", "default_incomplete", "error_if_incomplete", "pending_if_incomplete"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: allow_incomplete, default_incomplete, error_if_incomplete, pending_if_incomplete"),
					})
				}
				params["payment_behavior"] = v
			}
			if v, exists := d.GetOk("proration_date"); exists {
				params["proration_date"] = v
			}
			if v, exists := d.GetOk("quantity"); exists {
				params["quantity"] = v
			}
			if v, exists := d.GetOk("tax_rates"); exists {
				params["tax_rates"] = v
			}
			if v, exists := d.GetOk("price"); exists {
				params["price"] = v
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
						path := path.IndexString("recurring")
						v := shared.Dig[any](outerV, "recurring")
						if v != nil {
							{
								path := path
								outerV := v
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
			if v, exists := d.GetOk("proration_behavior"); exists {
				if inEnum(v.(string), []string{"always_invoice", "create_prorations", "none"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: always_invoice, create_prorations, none"),
					})
				}
				params["proration_behavior"] = v
			}
			{
				v := d.Get("subscription")
				params["subscription"] = v
			}
			res, err := f.Post(ctx, "/v1/subscription_items", params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to create new subscription_item: %s", err),
				})
				return out
			}
			d.SetId(shared.Dig[string](res, "id"))
			return out
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			err := f.Delete(ctx, fmt.Sprintf("/v1/subscription_items/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to delete subscription_item %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			res, err := f.Get(ctx, fmt.Sprintf("/v1/subscription_items/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to look up subscription_item %s: %s", d.Id(), err),
				})
				return out
			}
			d.Set("billing_thresholds", shared.Dig[any](res, "billing_thresholds"))
			d.Set("payment_behavior", shared.Dig[any](res, "payment_behavior"))
			d.Set("proration_date", shared.Dig[any](res, "proration_date"))
			d.Set("quantity", shared.Dig[any](res, "quantity"))
			d.Set("tax_rates", shared.Dig[any](res, "tax_rates"))
			d.Set("price", shared.Dig[any](res, "price"))
			d.Set("price_data", shared.Dig[any](res, "price_data"))
			d.Set("proration_behavior", shared.Dig[any](res, "proration_behavior"))
			d.Set("subscription", shared.Dig[any](res, "subscription"))
			return out
		},
		Schema: map[string]*schema.Schema{
			"billing_thresholds": {
				Description: "Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. When updating, pass an empty string to remove previously-defined thresholds.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"usage_gte": {
					Description: "",
					Required:    true,
					Type:        schema.TypeInt,
				}}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"payment_behavior": {
				Description: "Use `allow_incomplete` to transition the subscription to `status=past_due` if a payment is required but cannot be paid. This allows you to manage scenarios where additional user actions are needed to pay a subscription's invoice. For example, SCA regulation may require 3DS authentication to complete payment. See the [SCA Migration Guide](https://stripe.com/docs/billing/migration/strong-customer-authentication) for Billing to learn more. This is the default behavior.\n\nUse `default_incomplete` to transition the subscription to `status=past_due` when payment is required and await explicit confirmation of the invoice's payment intent. This allows simpler management of scenarios where additional user actions are needed to pay a subscription’s invoice. Such as failed payments, [SCA regulation](https://stripe.com/docs/billing/migration/strong-customer-authentication), or collecting a mandate for a bank debit payment method.\n\nUse `pending_if_incomplete` to update the subscription using [pending updates](https://stripe.com/docs/billing/subscriptions/pending-updates). When you use `pending_if_incomplete` you can only pass the parameters [supported by pending updates](https://stripe.com/docs/billing/pending-updates-reference#supported-attributes).\n\nUse `error_if_incomplete` if you want Stripe to return an HTTP 402 status code if a subscription's invoice cannot be paid. For example, if a payment method requires 3DS authentication due to SCA regulation and further user action is needed, this parameter does not update the subscription and returns an error instead. This was the default behavior for API versions prior to 2019-03-14. See the [changelog](https://stripe.com/docs/upgrades#2019-03-14) to learn more.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
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
					"recurring": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
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
						}},
						Required: true,
						Type:     schema.TypeMap,
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
			"proration_behavior": {
				Description: "Determines how to handle [prorations](https://stripe.com/docs/subscriptions/billing-cycle#prorations) when the billing cycle changes (e.g., when switching plans, resetting `billing_cycle_anchor=now`, or starting a trial), or if an item's `quantity` changes.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"proration_date": {
				Description: "If set, the proration will be calculated as though the subscription was updated at the given time. This can be used to apply the same proration that was previewed with the [upcoming invoice](https://stripe.com/docs/api#retrieve_customer_invoice) endpoint.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"quantity": {
				Description: "The quantity you'd like to apply to the subscription item you're creating.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"subscription": {
				Description: "The identifier of the subscription to modify.",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
			},
			"tax_rates": {
				Description: "A list of [Tax Rate](https://stripe.com/docs/api/tax_rates) ids. These Tax Rates will override the [`default_tax_rates`](https://stripe.com/docs/api/subscriptions/create#create_subscription-default_tax_rates) on the Subscription. When updating, pass an empty string to remove previously-defined tax rates.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeList,
			},
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if d.HasChange("billing_thresholds") {
				v := d.Get("billing_thresholds")
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("usage_gte")
						v := shared.Dig[any](outerV, "usage_gte")
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
				params["billing_thresholds"] = v
			}
			if d.HasChange("payment_behavior") {
				v := d.Get("payment_behavior")
				if inEnum(v.(string), []string{"allow_incomplete", "default_incomplete", "error_if_incomplete", "pending_if_incomplete"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: allow_incomplete, default_incomplete, error_if_incomplete, pending_if_incomplete"),
					})
				}
				params["payment_behavior"] = v
			}
			if d.HasChange("proration_date") {
				v := d.Get("proration_date")
				params["proration_date"] = v
			}
			if d.HasChange("quantity") {
				v := d.Get("quantity")
				params["quantity"] = v
			}
			if d.HasChange("tax_rates") {
				v := d.Get("tax_rates")
				params["tax_rates"] = v
			}
			if d.HasChange("price") {
				v := d.Get("price")
				params["price"] = v
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
						path := path.IndexString("recurring")
						v := shared.Dig[any](outerV, "recurring")
						if v != nil {
							{
								path := path
								outerV := v
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
			if d.HasChange("proration_behavior") {
				v := d.Get("proration_behavior")
				if inEnum(v.(string), []string{"always_invoice", "create_prorations", "none"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: always_invoice, create_prorations, none"),
					})
				}
				params["proration_behavior"] = v
			}
			_, err := f.Post(ctx, fmt.Sprintf("/v1/subscription_items/%v", d.Get("id")), params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to update subscription_item %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
	}
}
