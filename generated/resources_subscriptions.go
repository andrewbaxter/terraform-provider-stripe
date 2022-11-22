package main

import (
	"context"
	"fmt"
	shared "github.com/andrewbaxter/terraform-provider-stripe/shared"
	cty "github.com/hashicorp/go-cty/cty"
	diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resources_subscriptions() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if v, exists := d.GetOk("trial_from_plan"); exists {
				params["trial_from_plan"] = v
			}
			if v, exists := d.GetOk("trial_period_days"); exists {
				params["trial_period_days"] = v
			}
			if v, exists := d.GetOk("collection_method"); exists {
				if inEnum(v.(string), []string{"charge_automatically", "send_invoice"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: charge_automatically, send_invoice"),
					})
				}
				params["collection_method"] = v
			}
			if v, exists := d.GetOk("description"); exists {
				params["description"] = v
			}
			if v, exists := d.GetOk("promotion_code"); exists {
				params["promotion_code"] = v
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
			if v, exists := d.GetOk("cancel_at_period_end"); exists {
				params["cancel_at_period_end"] = v
			}
			if v, exists := d.GetOk("default_source"); exists {
				params["default_source"] = v
			}
			if v, exists := d.GetOk("trial_end"); exists {
				params["trial_end"] = v
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
			if v, exists := d.GetOk("payment_settings"); exists {
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("payment_method_options")
						v := shared.Dig[any](outerV, "payment_method_options")
						if v != nil {
							{
								path := path
								outerV := v
								{
									path := path.IndexString("card")
									v := shared.Dig[any](outerV, "card")
									if v != nil {
										{
											path := path
											outerV := v
											{
												path := path.IndexString("mandate_options")
												v := shared.Dig[any](outerV, "mandate_options")
												if v != nil {
													{
														path := path
														outerV := v
														{
															path := path.IndexString("amount_type")
															v := shared.Dig[any](outerV, "amount_type")
															if v != nil {
																if inEnum(v.(string), []string{"fixed", "maximum"}) {
																	out = append(out, diag.Diagnostic{
																		AttributePath: path,
																		Severity:      diag.Error,
																		Summary:       fmt.Sprintf("Field must be one of: fixed, maximum"),
																	})
																}
															}
														}
													}
												}
											}
											{
												path := path.IndexString("network")
												v := shared.Dig[any](outerV, "network")
												if v != nil {
													if inEnum(v.(string), []string{"amex", "cartes_bancaires", "diners", "discover", "interac", "jcb", "mastercard", "unionpay", "unknown", "visa"}) {
														out = append(out, diag.Diagnostic{
															AttributePath: path,
															Severity:      diag.Error,
															Summary:       fmt.Sprintf("Field must be one of: amex, cartes_bancaires, diners, discover, interac, jcb, mastercard, unionpay, unknown, visa"),
														})
													}
												}
											}
											{
												path := path.IndexString("request_three_d_secure")
												v := shared.Dig[any](outerV, "request_three_d_secure")
												if v != nil {
													if inEnum(v.(string), []string{"any", "automatic"}) {
														out = append(out, diag.Diagnostic{
															AttributePath: path,
															Severity:      diag.Error,
															Summary:       fmt.Sprintf("Field must be one of: any, automatic"),
														})
													}
												}
											}
										}
									}
								}
								{
									path := path.IndexString("customer_balance")
									v := shared.Dig[any](outerV, "customer_balance")
									if v != nil {
										{
											path := path
											outerV := v
											{
												path := path.IndexString("bank_transfer")
												v := shared.Dig[any](outerV, "bank_transfer")
												if v != nil {
													{
														path := path
														outerV := v
														{
															path := path.IndexString("eu_bank_transfer")
															v := shared.Dig[any](outerV, "eu_bank_transfer")
															if v != nil {
																{
																	path := path
																	outerV := v
																	{
																		path := path.IndexString("country")
																		v := shared.Dig[any](outerV, "country")
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
													}
												}
											}
										}
									}
								}
								{
									path := path.IndexString("us_bank_account")
									v := shared.Dig[any](outerV, "us_bank_account")
									if v != nil {
										{
											path := path
											outerV := v
											{
												path := path.IndexString("financial_connections")
												v := shared.Dig[any](outerV, "financial_connections")
												if v != nil {
													{
														path := path
														outerV := v
														{
															path := path.IndexString("permissions")
															v := shared.Dig[any](outerV, "permissions")
															if v != nil {
																{
																	path := path
																	outerV := v
																	for i, v := range outerV.([]any) {
																		if inEnum(v.(string), []string{"balances", "ownership", "payment_method", "transactions"}) {
																			out = append(out, diag.Diagnostic{
																				AttributePath: path.IndexInt(i),
																				Severity:      diag.Error,
																				Summary:       fmt.Sprintf("Field must be one of: balances, ownership, payment_method, transactions"),
																			})
																		}
																	}
																}
															}
														}
													}
												}
											}
											{
												path := path.IndexString("verification_method")
												v := shared.Dig[any](outerV, "verification_method")
												if v != nil {
													if inEnum(v.(string), []string{"automatic", "instant", "microdeposits"}) {
														out = append(out, diag.Diagnostic{
															AttributePath: path,
															Severity:      diag.Error,
															Summary:       fmt.Sprintf("Field must be one of: automatic, instant, microdeposits"),
														})
													}
												}
											}
										}
									}
								}
								{
									path := path.IndexString("acss_debit")
									v := shared.Dig[any](outerV, "acss_debit")
									if v != nil {
										{
											path := path
											outerV := v
											{
												path := path.IndexString("mandate_options")
												v := shared.Dig[any](outerV, "mandate_options")
												if v != nil {
													{
														path := path
														outerV := v
														{
															path := path.IndexString("transaction_type")
															v := shared.Dig[any](outerV, "transaction_type")
															if v != nil {
																if inEnum(v.(string), []string{"business", "personal"}) {
																	out = append(out, diag.Diagnostic{
																		AttributePath: path,
																		Severity:      diag.Error,
																		Summary:       fmt.Sprintf("Field must be one of: business, personal"),
																	})
																}
															}
														}
													}
												}
											}
											{
												path := path.IndexString("verification_method")
												v := shared.Dig[any](outerV, "verification_method")
												if v != nil {
													if inEnum(v.(string), []string{"automatic", "instant", "microdeposits"}) {
														out = append(out, diag.Diagnostic{
															AttributePath: path,
															Severity:      diag.Error,
															Summary:       fmt.Sprintf("Field must be one of: automatic, instant, microdeposits"),
														})
													}
												}
											}
										}
									}
								}
								{
									path := path.IndexString("bancontact")
									v := shared.Dig[any](outerV, "bancontact")
									if v != nil {
										{
											path := path
											outerV := v
											{
												path := path.IndexString("preferred_language")
												v := shared.Dig[any](outerV, "preferred_language")
												if v != nil {
													if inEnum(v.(string), []string{"de", "en", "fr", "nl"}) {
														out = append(out, diag.Diagnostic{
															AttributePath: path,
															Severity:      diag.Error,
															Summary:       fmt.Sprintf("Field must be one of: de, en, fr, nl"),
														})
													}
												}
											}
										}
									}
								}
							}
						}
					}
					{
						path := path.IndexString("payment_method_types")
						v := shared.Dig[any](outerV, "payment_method_types")
						if v != nil {
							{
								path := path
								outerV := v
								for i, v := range outerV.([]any) {
									if inEnum(v.(string), []string{"ach_credit_transfer", "ach_debit", "acss_debit", "au_becs_debit", "bacs_debit", "bancontact", "boleto", "card", "customer_balance", "fpx", "giropay", "grabpay", "ideal", "konbini", "link", "paynow", "promptpay", "sepa_debit", "sofort", "us_bank_account", "wechat_pay"}) {
										out = append(out, diag.Diagnostic{
											AttributePath: path.IndexInt(i),
											Severity:      diag.Error,
											Summary:       fmt.Sprintf("Field must be one of: ach_credit_transfer, ach_debit, acss_debit, au_becs_debit, bacs_debit, bancontact, boleto, card, customer_balance, fpx, giropay, grabpay, ideal, konbini, link, paynow, promptpay, sepa_debit, sofort, us_bank_account, wechat_pay"),
										})
									}
								}
							}
						}
					}
					{
						path := path.IndexString("save_default_payment_method")
						v := shared.Dig[any](outerV, "save_default_payment_method")
						if v != nil {
							if inEnum(v.(string), []string{"off", "on_subscription"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: off, on_subscription"),
								})
							}
						}
					}
				}
				params["payment_settings"] = v
			}
			if v, exists := d.GetOk("cancel_at"); exists {
				params["cancel_at"] = v
			}
			if v, exists := d.GetOk("currency"); exists {
				params["currency"] = v
			}
			if v, exists := d.GetOk("backdate_start_date"); exists {
				params["backdate_start_date"] = v
			}
			if v, exists := d.GetOk("billing_cycle_anchor"); exists {
				params["billing_cycle_anchor"] = v
			}
			if v, exists := d.GetOk("days_until_due"); exists {
				params["days_until_due"] = v
			}
			if v, exists := d.GetOk("add_invoice_items"); exists {
				{
					path := cty.Path{}
					outerV := v
					for i, v := range outerV.([]any) {
						{
							path := path.IndexInt(i)
							outerV := v
							{
								path := path.IndexString("price_data")
								v := shared.Dig[any](outerV, "price_data")
								if v != nil {
									{
										path := path
										outerV := v
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
									}
								}
							}
						}
					}
				}
				params["add_invoice_items"] = v
			}
			if v, exists := d.GetOk("application_fee_percent"); exists {
				params["application_fee_percent"] = v
			}
			{
				v := d.Get("customer")
				params["customer"] = v
			}
			if v, exists := d.GetOk("on_behalf_of"); exists {
				params["on_behalf_of"] = v
			}
			if v, exists := d.GetOk("billing_thresholds"); exists {
				params["billing_thresholds"] = v
			}
			if v, exists := d.GetOk("default_payment_method"); exists {
				params["default_payment_method"] = v
			}
			if v, exists := d.GetOk("default_tax_rates"); exists {
				params["default_tax_rates"] = v
			}
			if v, exists := d.GetOk("transfer_data"); exists {
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("destination")
						v := shared.Dig[any](outerV, "destination")
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
				params["transfer_data"] = v
			}
			if v, exists := d.GetOk("pending_invoice_item_interval"); exists {
				{
					path := cty.Path{}
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
				params["pending_invoice_item_interval"] = v
			}
			if v, exists := d.GetOk("automatic_tax"); exists {
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
				params["automatic_tax"] = v
			}
			if v, exists := d.GetOk("coupon"); exists {
				params["coupon"] = v
			}
			if v, exists := d.GetOk("items"); exists {
				{
					path := cty.Path{}
					outerV := v
					for i, v := range outerV.([]any) {
						{
							path := path.IndexInt(i)
							outerV := v
							{
								path := path.IndexString("price_data")
								v := shared.Dig[any](outerV, "price_data")
								if v != nil {
									{
										path := path
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
								}
							}
							{
								path := path.IndexString("billing_thresholds")
								v := shared.Dig[any](outerV, "billing_thresholds")
								if v != nil {
									{
										path := path
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
								}
							}
						}
					}
				}
				params["items"] = v
			}
			if v, exists := d.GetOk("off_session"); exists {
				params["off_session"] = v
			}
			res, err := f.Post(ctx, "/v1/subscriptions", params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to create new subscription: %s", err),
				})
				return out
			}
			d.SetId(shared.Dig[string](res, "id"))
			return out
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			err := f.Delete(ctx, fmt.Sprintf("/v1/subscriptions/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to delete subscription %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			res, err := f.Get(ctx, fmt.Sprintf("/v1/subscriptions/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to look up subscription %s: %s", d.Id(), err),
				})
				return out
			}
			d.Set("trial_from_plan", shared.Dig[any](res, "trial_from_plan"))
			d.Set("trial_period_days", shared.Dig[any](res, "trial_period_days"))
			d.Set("collection_method", shared.Dig[any](res, "collection_method"))
			d.Set("description", shared.Dig[any](res, "description"))
			d.Set("promotion_code", shared.Dig[any](res, "promotion_code"))
			d.Set("proration_behavior", shared.Dig[any](res, "proration_behavior"))
			d.Set("cancel_at_period_end", shared.Dig[any](res, "cancel_at_period_end"))
			d.Set("default_source", shared.Dig[any](res, "default_source"))
			d.Set("trial_end", shared.Dig[any](res, "trial_end"))
			d.Set("payment_behavior", shared.Dig[any](res, "payment_behavior"))
			d.Set("payment_settings", shared.Dig[any](res, "payment_settings"))
			d.Set("cancel_at", shared.Dig[any](res, "cancel_at"))
			d.Set("currency", shared.Dig[any](res, "currency"))
			d.Set("backdate_start_date", shared.Dig[any](res, "backdate_start_date"))
			d.Set("billing_cycle_anchor", shared.Dig[any](res, "billing_cycle_anchor"))
			d.Set("days_until_due", shared.Dig[any](res, "days_until_due"))
			d.Set("add_invoice_items", shared.Dig[any](res, "add_invoice_items"))
			d.Set("application_fee_percent", shared.Dig[any](res, "application_fee_percent"))
			d.Set("customer", shared.Dig[any](res, "customer"))
			d.Set("on_behalf_of", shared.Dig[any](res, "on_behalf_of"))
			d.Set("billing_thresholds", shared.Dig[any](res, "billing_thresholds"))
			d.Set("default_payment_method", shared.Dig[any](res, "default_payment_method"))
			d.Set("default_tax_rates", shared.Dig[any](res, "default_tax_rates"))
			d.Set("transfer_data", shared.Dig[any](res, "transfer_data"))
			d.Set("pending_invoice_item_interval", shared.Dig[any](res, "pending_invoice_item_interval"))
			d.Set("automatic_tax", shared.Dig[any](res, "automatic_tax"))
			d.Set("coupon", shared.Dig[any](res, "coupon"))
			d.Set("items", shared.Dig[any](res, "items"))
			d.Set("off_session", shared.Dig[any](res, "off_session"))
			return out
		},
		Schema: map[string]*schema.Schema{
			"add_invoice_items": {
				Description: "A list of prices and quantities that will generate invoice items appended to the next invoice for this subscription. You may pass up to 20 items.",
				Elem: &schema.Schema{
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"price": {
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"price_data": {
							Description: "",
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
							Required: false,
							Type:     schema.TypeMap,
						},
						"quantity": {
							Description: "",
							Required:    false,
							Type:        schema.TypeInt,
						},
						"tax_rates": {
							Description: "",
							Elem:        &schema.Schema{Type: schema.TypeString},
							Required:    false,
							Type:        schema.TypeList,
						},
					}},
					Type: schema.TypeMap,
				},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeList,
			},
			"application_fee_percent": {
				Description: "A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice subtotal that will be transferred to the application owner's Stripe account. The request must be made by a platform account on a connected account in order to set an application fee percentage. For more information, see the application fees [documentation](https://stripe.com/docs/connect/subscriptions#collecting-fees-on-subscriptions).",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeFloat,
			},
			"automatic_tax": {
				Description: "Automatic tax settings for this subscription. We recommend you only include this parameter when the existing value is being changed.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"enabled": {
					Description: "",
					Required:    true,
					Type:        schema.TypeBool,
				}}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"backdate_start_date": {
				Description: "For new subscriptions, a past timestamp to backdate the subscription's start date to. If set, the first invoice will contain a proration for the timespan between the start date and the current time. Can be combined with trials and the billing cycle anchor.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"billing_cycle_anchor": {
				Description: "A future timestamp to anchor the subscription's [billing cycle](https://stripe.com/docs/subscriptions/billing-cycle). This is used to determine the date of the first full invoice, and, for plans with `month` or `year` intervals, the day of the month for subsequent invoices. The timestamp is in UTC format.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"billing_thresholds": {
				Description: "Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"amount_gte": {
						Description: "",
						Required:    false,
						Type:        schema.TypeInt,
					},
					"reset_billing_cycle_anchor": {
						Description: "",
						Required:    false,
						Type:        schema.TypeBool,
					},
				}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"cancel_at": {
				Description: "A timestamp at which the subscription should cancel. If set to a date before the current period ends, this will cause a proration if prorations have been enabled using `proration_behavior`. If set during a future period, this will always cause a proration for that period.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"cancel_at_period_end": {
				Description: "Boolean indicating whether this subscription should cancel at the end of the current period.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeBool,
			},
			"collection_method": {
				Description: "Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay this subscription at the end of the cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`. Defaults to `charge_automatically`.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"coupon": {
				Description: "The ID of the coupon to apply to this subscription. A coupon applied to a subscription will only affect invoices created for that particular subscription.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"currency": {
				Description: "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"customer": {
				Description: "The identifier of the customer to subscribe.",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
			},
			"days_until_due": {
				Description: "Number of days a customer has to pay invoices generated by this subscription. Valid only for subscriptions where `collection_method` is set to `send_invoice`.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"default_payment_method": {
				Description: "ID of the default payment method for the subscription. It must belong to the customer associated with the subscription. This takes precedence over `default_source`. If neither are set, invoices will use the customer's [invoice_settings.default_payment_method](https://stripe.com/docs/api/customers/object#customer_object-invoice_settings-default_payment_method) or [default_source](https://stripe.com/docs/api/customers/object#customer_object-default_source).",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"default_source": {
				Description: "ID of the default payment source for the subscription. It must belong to the customer associated with the subscription and be in a chargeable state. If `default_payment_method` is also set, `default_payment_method` will take precedence. If neither are set, invoices will use the customer's [invoice_settings.default_payment_method](https://stripe.com/docs/api/customers/object#customer_object-invoice_settings-default_payment_method) or [default_source](https://stripe.com/docs/api/customers/object#customer_object-default_source).",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"default_tax_rates": {
				Description: "The tax rates that will apply to any subscription item that does not have `tax_rates` set. Invoices created will have their `default_tax_rates` populated from the subscription.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeList,
			},
			"description": {
				Description: "The subscription's description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"items": {
				Description: "A list of up to 20 subscription items, each with an attached price.",
				Elem: &schema.Schema{
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"billing_thresholds": {
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"usage_gte": {
								Description: "",
								Required:    true,
								Type:        schema.TypeInt,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"metadata": {
							Description: "",
							Elem:        &schema.Resource{Schema: map[string]*schema.Schema{}},
							Required:    false,
							Type:        schema.TypeMap,
						},
						"price": {
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"price_data": {
							Description: "",
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
							Required: false,
							Type:     schema.TypeMap,
						},
						"quantity": {
							Description: "",
							Required:    false,
							Type:        schema.TypeInt,
						},
						"tax_rates": {
							Description: "",
							Elem:        &schema.Schema{Type: schema.TypeString},
							Required:    false,
							Type:        schema.TypeList,
						},
					}},
					Type: schema.TypeMap,
				},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeList,
			},
			"off_session": {
				Description: "Indicates if a customer is on or off-session while an invoice payment is attempted.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeBool,
			},
			"on_behalf_of": {
				Description: "The account on behalf of which to charge, for each of the subscription's invoices.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"payment_behavior": {
				Description: "Only applies to subscriptions with `collection_method=charge_automatically`.\n\nUse `allow_incomplete` to create subscriptions with `status=incomplete` if the first invoice cannot be paid. Creating subscriptions with this status allows you to manage scenarios where additional user actions are needed to pay a subscription's invoice. For example, SCA regulation may require 3DS authentication to complete payment. See the [SCA Migration Guide](https://stripe.com/docs/billing/migration/strong-customer-authentication) for Billing to learn more. This is the default behavior.\n\nUse `default_incomplete` to create Subscriptions with `status=incomplete` when the first invoice requires payment, otherwise start as active. Subscriptions transition to `status=active` when successfully confirming the payment intent on the first invoice. This allows simpler management of scenarios where additional user actions are needed to pay a subscription’s invoice. Such as failed payments, [SCA regulation](https://stripe.com/docs/billing/migration/strong-customer-authentication), or collecting a mandate for a bank debit payment method. If the payment intent is not confirmed within 23 hours subscriptions transition to `status=incomplete_expired`, which is a terminal state.\n\nUse `error_if_incomplete` if you want Stripe to return an HTTP 402 status code if a subscription's first invoice cannot be paid. For example, if a payment method requires 3DS authentication due to SCA regulation and further user action is needed, this parameter does not create a subscription and returns an error instead. This was the default behavior for API versions prior to 2019-03-14. See the [changelog](https://stripe.com/docs/upgrades#2019-03-14) to learn more.\n\n`pending_if_incomplete` is only used with updates and cannot be passed when creating a subscription.\n\nSubscriptions with `collection_method=send_invoice` are automatically activated regardless of the first invoice status.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"payment_settings": {
				Description: "Payment settings to pass to invoices created by the subscription.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"payment_method_options": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"acss_debit": {
								Description: "",
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"mandate_options": {
										Description: "",
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{"transaction_type": {
											Description: "",
											Required:    false,
											Type:        schema.TypeString,
										}}},
										Required: false,
										Type:     schema.TypeMap,
									},
									"verification_method": {
										Description: "",
										Required:    false,
										Type:        schema.TypeString,
									},
								}},
								Required: false,
								Type:     schema.TypeMap,
							},
							"bancontact": {
								Description: "",
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{"preferred_language": {
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								}}},
								Required: false,
								Type:     schema.TypeMap,
							},
							"card": {
								Description: "",
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"mandate_options": {
										Description: "",
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"amount": {
												Description: "",
												Required:    false,
												Type:        schema.TypeInt,
											},
											"amount_type": {
												Description: "",
												Required:    false,
												Type:        schema.TypeString,
											},
											"description": {
												Description: "",
												Required:    false,
												Type:        schema.TypeString,
											},
										}},
										Required: false,
										Type:     schema.TypeMap,
									},
									"network": {
										Description: "",
										Required:    false,
										Type:        schema.TypeString,
									},
									"request_three_d_secure": {
										Description: "",
										Required:    false,
										Type:        schema.TypeString,
									},
								}},
								Required: false,
								Type:     schema.TypeMap,
							},
							"customer_balance": {
								Description: "",
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"bank_transfer": {
										Description: "",
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"eu_bank_transfer": {
												Description: "",
												Elem: &schema.Resource{Schema: map[string]*schema.Schema{"country": {
													Description: "",
													Required:    true,
													Type:        schema.TypeString,
												}}},
												Required: false,
												Type:     schema.TypeMap,
											},
											"type": {
												Description: "",
												Required:    false,
												Type:        schema.TypeString,
											},
										}},
										Required: false,
										Type:     schema.TypeMap,
									},
									"funding_type": {
										Description: "",
										Required:    false,
										Type:        schema.TypeString,
									},
								}},
								Required: false,
								Type:     schema.TypeMap,
							},
							"konbini": {
								Description: "",
								Elem:        &schema.Resource{Schema: map[string]*schema.Schema{}},
								Required:    false,
								Type:        schema.TypeMap,
							},
							"us_bank_account": {
								Description: "",
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"financial_connections": {
										Description: "",
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{"permissions": {
											Description: "",
											Elem:        &schema.Schema{Type: schema.TypeString},
											Required:    false,
											Type:        schema.TypeList,
										}}},
										Required: false,
										Type:     schema.TypeMap,
									},
									"verification_method": {
										Description: "",
										Required:    false,
										Type:        schema.TypeString,
									},
								}},
								Required: false,
								Type:     schema.TypeMap,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"payment_method_types": {
						Description: "",
						Elem:        &schema.Schema{Type: schema.TypeString},
						Required:    false,
						Type:        schema.TypeList,
					},
					"save_default_payment_method": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
				}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"pending_invoice_item_interval": {
				Description: "Specifies an interval for how often to bill for any pending invoice items. It is analogous to calling [Create an invoice](https://stripe.com/docs/api#create_invoice) for the given subscription at the specified interval.",
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
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"promotion_code": {
				Description: "The API ID of a promotion code to apply to this subscription. A promotion code applied to a subscription will only affect invoices created for that particular subscription.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"proration_behavior": {
				Description: "Determines how to handle [prorations](https://stripe.com/docs/subscriptions/billing-cycle#prorations) resulting from the `billing_cycle_anchor`. If no value is passed, the default is `create_prorations`.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"transfer_data": {
				Description: "If specified, the funds from the subscription's invoices will be transferred to the destination and the ID of the resulting transfers will be found on the resulting charges.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"amount_percent": {
						Description: "",
						Required:    false,
						Type:        schema.TypeFloat,
					},
					"destination": {
						Description: "",
						Required:    true,
						Type:        schema.TypeString,
					},
				}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"trial_end": {
				Description: "Unix timestamp representing the end of the trial period the customer will get before being charged for the first time. This will always overwrite any trials that might apply via a subscribed plan. If set, trial_end will override the default trial period of the plan the customer is being subscribed to. The special value `now` can be provided to end the customer's trial immediately. Can be at most two years from `billing_cycle_anchor`. See [Using trial periods on subscriptions](https://stripe.com/docs/billing/subscriptions/trials) to learn more.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"trial_from_plan": {
				Description: "Indicates if a plan's `trial_period_days` should be applied to the subscription. Setting `trial_end` per subscription is preferred, and this defaults to `false`. Setting this flag to `true` together with `trial_end` is not allowed. See [Using trial periods on subscriptions](https://stripe.com/docs/billing/subscriptions/trials) to learn more.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeBool,
			},
			"trial_period_days": {
				Description: "Integer representing the number of trial period days before the customer is charged for the first time. This will always overwrite any trials that might apply via a subscribed plan. See [Using trial periods on subscriptions](https://stripe.com/docs/billing/subscriptions/trials) to learn more.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeInt,
			},
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if d.HasChange("trial_from_plan") {
				v := d.Get("trial_from_plan")
				params["trial_from_plan"] = v
			}
			if d.HasChange("collection_method") {
				v := d.Get("collection_method")
				if inEnum(v.(string), []string{"charge_automatically", "send_invoice"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: charge_automatically, send_invoice"),
					})
				}
				params["collection_method"] = v
			}
			if d.HasChange("description") {
				v := d.Get("description")
				params["description"] = v
			}
			if d.HasChange("promotion_code") {
				v := d.Get("promotion_code")
				params["promotion_code"] = v
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
			if d.HasChange("cancel_at_period_end") {
				v := d.Get("cancel_at_period_end")
				params["cancel_at_period_end"] = v
			}
			if d.HasChange("default_source") {
				v := d.Get("default_source")
				params["default_source"] = v
			}
			if d.HasChange("trial_end") {
				v := d.Get("trial_end")
				params["trial_end"] = v
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
			if d.HasChange("payment_settings") {
				v := d.Get("payment_settings")
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("payment_method_options")
						v := shared.Dig[any](outerV, "payment_method_options")
						if v != nil {
							{
								path := path
								outerV := v
								{
									path := path.IndexString("card")
									v := shared.Dig[any](outerV, "card")
									if v != nil {
										{
											path := path
											outerV := v
											{
												path := path.IndexString("mandate_options")
												v := shared.Dig[any](outerV, "mandate_options")
												if v != nil {
													{
														path := path
														outerV := v
														{
															path := path.IndexString("amount_type")
															v := shared.Dig[any](outerV, "amount_type")
															if v != nil {
																if inEnum(v.(string), []string{"fixed", "maximum"}) {
																	out = append(out, diag.Diagnostic{
																		AttributePath: path,
																		Severity:      diag.Error,
																		Summary:       fmt.Sprintf("Field must be one of: fixed, maximum"),
																	})
																}
															}
														}
													}
												}
											}
											{
												path := path.IndexString("network")
												v := shared.Dig[any](outerV, "network")
												if v != nil {
													if inEnum(v.(string), []string{"amex", "cartes_bancaires", "diners", "discover", "interac", "jcb", "mastercard", "unionpay", "unknown", "visa"}) {
														out = append(out, diag.Diagnostic{
															AttributePath: path,
															Severity:      diag.Error,
															Summary:       fmt.Sprintf("Field must be one of: amex, cartes_bancaires, diners, discover, interac, jcb, mastercard, unionpay, unknown, visa"),
														})
													}
												}
											}
											{
												path := path.IndexString("request_three_d_secure")
												v := shared.Dig[any](outerV, "request_three_d_secure")
												if v != nil {
													if inEnum(v.(string), []string{"any", "automatic"}) {
														out = append(out, diag.Diagnostic{
															AttributePath: path,
															Severity:      diag.Error,
															Summary:       fmt.Sprintf("Field must be one of: any, automatic"),
														})
													}
												}
											}
										}
									}
								}
								{
									path := path.IndexString("customer_balance")
									v := shared.Dig[any](outerV, "customer_balance")
									if v != nil {
										{
											path := path
											outerV := v
											{
												path := path.IndexString("bank_transfer")
												v := shared.Dig[any](outerV, "bank_transfer")
												if v != nil {
													{
														path := path
														outerV := v
														{
															path := path.IndexString("eu_bank_transfer")
															v := shared.Dig[any](outerV, "eu_bank_transfer")
															if v != nil {
																{
																	path := path
																	outerV := v
																	{
																		path := path.IndexString("country")
																		v := shared.Dig[any](outerV, "country")
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
													}
												}
											}
										}
									}
								}
								{
									path := path.IndexString("us_bank_account")
									v := shared.Dig[any](outerV, "us_bank_account")
									if v != nil {
										{
											path := path
											outerV := v
											{
												path := path.IndexString("financial_connections")
												v := shared.Dig[any](outerV, "financial_connections")
												if v != nil {
													{
														path := path
														outerV := v
														{
															path := path.IndexString("permissions")
															v := shared.Dig[any](outerV, "permissions")
															if v != nil {
																{
																	path := path
																	outerV := v
																	for i, v := range outerV.([]any) {
																		if inEnum(v.(string), []string{"balances", "ownership", "payment_method", "transactions"}) {
																			out = append(out, diag.Diagnostic{
																				AttributePath: path.IndexInt(i),
																				Severity:      diag.Error,
																				Summary:       fmt.Sprintf("Field must be one of: balances, ownership, payment_method, transactions"),
																			})
																		}
																	}
																}
															}
														}
													}
												}
											}
											{
												path := path.IndexString("verification_method")
												v := shared.Dig[any](outerV, "verification_method")
												if v != nil {
													if inEnum(v.(string), []string{"automatic", "instant", "microdeposits"}) {
														out = append(out, diag.Diagnostic{
															AttributePath: path,
															Severity:      diag.Error,
															Summary:       fmt.Sprintf("Field must be one of: automatic, instant, microdeposits"),
														})
													}
												}
											}
										}
									}
								}
								{
									path := path.IndexString("acss_debit")
									v := shared.Dig[any](outerV, "acss_debit")
									if v != nil {
										{
											path := path
											outerV := v
											{
												path := path.IndexString("mandate_options")
												v := shared.Dig[any](outerV, "mandate_options")
												if v != nil {
													{
														path := path
														outerV := v
														{
															path := path.IndexString("transaction_type")
															v := shared.Dig[any](outerV, "transaction_type")
															if v != nil {
																if inEnum(v.(string), []string{"business", "personal"}) {
																	out = append(out, diag.Diagnostic{
																		AttributePath: path,
																		Severity:      diag.Error,
																		Summary:       fmt.Sprintf("Field must be one of: business, personal"),
																	})
																}
															}
														}
													}
												}
											}
											{
												path := path.IndexString("verification_method")
												v := shared.Dig[any](outerV, "verification_method")
												if v != nil {
													if inEnum(v.(string), []string{"automatic", "instant", "microdeposits"}) {
														out = append(out, diag.Diagnostic{
															AttributePath: path,
															Severity:      diag.Error,
															Summary:       fmt.Sprintf("Field must be one of: automatic, instant, microdeposits"),
														})
													}
												}
											}
										}
									}
								}
								{
									path := path.IndexString("bancontact")
									v := shared.Dig[any](outerV, "bancontact")
									if v != nil {
										{
											path := path
											outerV := v
											{
												path := path.IndexString("preferred_language")
												v := shared.Dig[any](outerV, "preferred_language")
												if v != nil {
													if inEnum(v.(string), []string{"de", "en", "fr", "nl"}) {
														out = append(out, diag.Diagnostic{
															AttributePath: path,
															Severity:      diag.Error,
															Summary:       fmt.Sprintf("Field must be one of: de, en, fr, nl"),
														})
													}
												}
											}
										}
									}
								}
							}
						}
					}
					{
						path := path.IndexString("payment_method_types")
						v := shared.Dig[any](outerV, "payment_method_types")
						if v != nil {
							{
								path := path
								outerV := v
								for i, v := range outerV.([]any) {
									if inEnum(v.(string), []string{"ach_credit_transfer", "ach_debit", "acss_debit", "au_becs_debit", "bacs_debit", "bancontact", "boleto", "card", "customer_balance", "fpx", "giropay", "grabpay", "ideal", "konbini", "link", "paynow", "promptpay", "sepa_debit", "sofort", "us_bank_account", "wechat_pay"}) {
										out = append(out, diag.Diagnostic{
											AttributePath: path.IndexInt(i),
											Severity:      diag.Error,
											Summary:       fmt.Sprintf("Field must be one of: ach_credit_transfer, ach_debit, acss_debit, au_becs_debit, bacs_debit, bancontact, boleto, card, customer_balance, fpx, giropay, grabpay, ideal, konbini, link, paynow, promptpay, sepa_debit, sofort, us_bank_account, wechat_pay"),
										})
									}
								}
							}
						}
					}
					{
						path := path.IndexString("save_default_payment_method")
						v := shared.Dig[any](outerV, "save_default_payment_method")
						if v != nil {
							if inEnum(v.(string), []string{"off", "on_subscription"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: off, on_subscription"),
								})
							}
						}
					}
				}
				params["payment_settings"] = v
			}
			if d.HasChange("cancel_at") {
				v := d.Get("cancel_at")
				params["cancel_at"] = v
			}
			if d.HasChange("billing_cycle_anchor") {
				v := d.Get("billing_cycle_anchor")
				params["billing_cycle_anchor"] = v
			}
			if d.HasChange("days_until_due") {
				v := d.Get("days_until_due")
				params["days_until_due"] = v
			}
			if d.HasChange("add_invoice_items") {
				v := d.Get("add_invoice_items")
				{
					path := cty.Path{}
					outerV := v
					for i, v := range outerV.([]any) {
						{
							path := path.IndexInt(i)
							outerV := v
							{
								path := path.IndexString("price_data")
								v := shared.Dig[any](outerV, "price_data")
								if v != nil {
									{
										path := path
										outerV := v
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
									}
								}
							}
						}
					}
				}
				params["add_invoice_items"] = v
			}
			if d.HasChange("application_fee_percent") {
				v := d.Get("application_fee_percent")
				params["application_fee_percent"] = v
			}
			if d.HasChange("on_behalf_of") {
				v := d.Get("on_behalf_of")
				params["on_behalf_of"] = v
			}
			if d.HasChange("billing_thresholds") {
				v := d.Get("billing_thresholds")
				params["billing_thresholds"] = v
			}
			if d.HasChange("default_payment_method") {
				v := d.Get("default_payment_method")
				params["default_payment_method"] = v
			}
			if d.HasChange("default_tax_rates") {
				v := d.Get("default_tax_rates")
				params["default_tax_rates"] = v
			}
			if d.HasChange("transfer_data") {
				v := d.Get("transfer_data")
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("destination")
						v := shared.Dig[any](outerV, "destination")
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
				params["transfer_data"] = v
			}
			if d.HasChange("pending_invoice_item_interval") {
				v := d.Get("pending_invoice_item_interval")
				{
					path := cty.Path{}
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
				params["pending_invoice_item_interval"] = v
			}
			if d.HasChange("automatic_tax") {
				v := d.Get("automatic_tax")
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
				params["automatic_tax"] = v
			}
			if d.HasChange("coupon") {
				v := d.Get("coupon")
				params["coupon"] = v
			}
			if d.HasChange("items") {
				v := d.Get("items")
				{
					path := cty.Path{}
					outerV := v
					for i, v := range outerV.([]any) {
						{
							path := path.IndexInt(i)
							outerV := v
							{
								path := path.IndexString("price_data")
								v := shared.Dig[any](outerV, "price_data")
								if v != nil {
									{
										path := path
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
								}
							}
							{
								path := path.IndexString("billing_thresholds")
								v := shared.Dig[any](outerV, "billing_thresholds")
								if v != nil {
									{
										path := path
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
								}
							}
						}
					}
				}
				params["items"] = v
			}
			if d.HasChange("off_session") {
				v := d.Get("off_session")
				params["off_session"] = v
			}
			_, err := f.Post(ctx, fmt.Sprintf("/v1/subscriptions/%v", d.Get("id")), params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to update subscription %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
	}
}
