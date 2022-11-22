package main

import (
	"context"
	"fmt"
	shared "github.com/andrewbaxter/terraform-provider-stripe/shared"
	cty "github.com/hashicorp/go-cty/cty"
	diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resources_invoices() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if v, exists := d.GetOk("days_until_due"); exists {
				params["days_until_due"] = v
			}
			if v, exists := d.GetOk("description"); exists {
				params["description"] = v
			}
			if v, exists := d.GetOk("discounts"); exists {
				params["discounts"] = v
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
								{
									path := path.IndexString("card")
									v := shared.Dig[any](outerV, "card")
									if v != nil {
										{
											path := path
											outerV := v
											{
												path := path.IndexString("installments")
												v := shared.Dig[any](outerV, "installments")
												if v != nil {
													{
														path := path
														outerV := v
														{
															path := path.IndexString("plan")
															v := shared.Dig[any](outerV, "plan")
															if v != nil {
																{
																	path := path
																	outerV := v
																	{
																		path := path.IndexString("count")
																		v := shared.Dig[any](outerV, "count")
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
																		path := path.IndexString("interval")
																		v := shared.Dig[any](outerV, "interval")
																		if v != nil {
																			if inEnum(v.(string), []string{"month"}) {
																				out = append(out, diag.Diagnostic{
																					AttributePath: path,
																					Severity:      diag.Error,
																					Summary:       fmt.Sprintf("Field must be one of: month"),
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
																		path := path.IndexString("type")
																		v := shared.Dig[any](outerV, "type")
																		if v != nil {
																			if inEnum(v.(string), []string{"fixed_count"}) {
																				out = append(out, diag.Diagnostic{
																					AttributePath: path,
																					Severity:      diag.Error,
																					Summary:       fmt.Sprintf("Field must be one of: fixed_count"),
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
															}
														}
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
				}
				params["payment_settings"] = v
			}
			if v, exists := d.GetOk("pending_invoice_items_behavior"); exists {
				if inEnum(v.(string), []string{"exclude", "include", "include_and_require"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: exclude, include, include_and_require"),
					})
				}
				params["pending_invoice_items_behavior"] = v
			}
			if v, exists := d.GetOk("statement_descriptor"); exists {
				params["statement_descriptor"] = v
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
			if v, exists := d.GetOk("default_payment_method"); exists {
				params["default_payment_method"] = v
			}
			if v, exists := d.GetOk("default_tax_rates"); exists {
				params["default_tax_rates"] = v
			}
			if v, exists := d.GetOk("due_date"); exists {
				params["due_date"] = v
			}
			if v, exists := d.GetOk("on_behalf_of"); exists {
				params["on_behalf_of"] = v
			}
			if v, exists := d.GetOk("application_fee_amount"); exists {
				params["application_fee_amount"] = v
			}
			if v, exists := d.GetOk("auto_advance"); exists {
				params["auto_advance"] = v
			}
			if v, exists := d.GetOk("currency"); exists {
				params["currency"] = v
			}
			if v, exists := d.GetOk("custom_fields"); exists {
				{
					path := cty.Path{}
					outerV := v
					for i, v := range outerV.([]any) {
						{
							path := path.IndexInt(i)
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
							{
								path := path.IndexString("value")
								v := shared.Dig[any](outerV, "value")
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
				params["custom_fields"] = v
			}
			if v, exists := d.GetOk("customer"); exists {
				params["customer"] = v
			}
			if v, exists := d.GetOk("from_invoice"); exists {
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("action")
						v := shared.Dig[any](outerV, "action")
						if v != nil {
							if inEnum(v.(string), []string{"revision"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: revision"),
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
						path := path.IndexString("invoice")
						v := shared.Dig[any](outerV, "invoice")
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
				params["from_invoice"] = v
			}
			if v, exists := d.GetOk("rendering_options"); exists {
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("amount_tax_display")
						v := shared.Dig[any](outerV, "amount_tax_display")
						if v != nil {
							if inEnum(v.(string), []string{"", "exclude_tax", "include_inclusive_tax"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: , exclude_tax, include_inclusive_tax"),
								})
							}
						}
					}
				}
				params["rendering_options"] = v
			}
			if v, exists := d.GetOk("subscription"); exists {
				params["subscription"] = v
			}
			if v, exists := d.GetOk("account_tax_ids"); exists {
				params["account_tax_ids"] = v
			}
			if v, exists := d.GetOk("default_source"); exists {
				params["default_source"] = v
			}
			if v, exists := d.GetOk("footer"); exists {
				params["footer"] = v
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
			res, err := f.Post(ctx, "/v1/invoices", params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to create new invoice: %s", err),
				})
				return out
			}
			d.SetId(shared.Dig[string](res, "id"))
			return out
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			err := f.Delete(ctx, fmt.Sprintf("/v1/invoices/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to delete invoice %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			res, err := f.Get(ctx, fmt.Sprintf("/v1/invoices/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to look up invoice %s: %s", d.Id(), err),
				})
				return out
			}
			d.Set("days_until_due", shared.Dig[any](res, "days_until_due"))
			d.Set("description", shared.Dig[any](res, "description"))
			d.Set("discounts", shared.Dig[any](res, "discounts"))
			d.Set("payment_settings", shared.Dig[any](res, "payment_settings"))
			d.Set("pending_invoice_items_behavior", shared.Dig[any](res, "pending_invoice_items_behavior"))
			d.Set("statement_descriptor", shared.Dig[any](res, "statement_descriptor"))
			d.Set("transfer_data", shared.Dig[any](res, "transfer_data"))
			d.Set("automatic_tax", shared.Dig[any](res, "automatic_tax"))
			d.Set("default_payment_method", shared.Dig[any](res, "default_payment_method"))
			d.Set("default_tax_rates", shared.Dig[any](res, "default_tax_rates"))
			d.Set("due_date", shared.Dig[any](res, "due_date"))
			d.Set("on_behalf_of", shared.Dig[any](res, "on_behalf_of"))
			d.Set("application_fee_amount", shared.Dig[any](res, "application_fee_amount"))
			d.Set("auto_advance", shared.Dig[any](res, "auto_advance"))
			d.Set("currency", shared.Dig[any](res, "currency"))
			d.Set("custom_fields", shared.Dig[any](res, "custom_fields"))
			d.Set("customer", shared.Dig[any](res, "customer"))
			d.Set("from_invoice", shared.Dig[any](res, "from_invoice"))
			d.Set("rendering_options", shared.Dig[any](res, "rendering_options"))
			d.Set("subscription", shared.Dig[any](res, "subscription"))
			d.Set("account_tax_ids", shared.Dig[any](res, "account_tax_ids"))
			d.Set("default_source", shared.Dig[any](res, "default_source"))
			d.Set("footer", shared.Dig[any](res, "footer"))
			d.Set("collection_method", shared.Dig[any](res, "collection_method"))
			return out
		},
		Schema: map[string]*schema.Schema{
			"account_tax_ids": {
				Description: "The account tax IDs associated with the invoice. Only editable when the invoice is a draft.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeList,
			},
			"application_fee_amount": {
				Description: "A fee in cents (or local equivalent) that will be applied to the invoice and transferred to the application owner's Stripe account. The request must be made with an OAuth key or the Stripe-Account header in order to take an application fee. For more information, see the application fees [documentation](https://stripe.com/docs/billing/invoices/connect#collecting-fees).",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"auto_advance": {
				Description: "Controls whether Stripe will perform [automatic collection](https://stripe.com/docs/billing/invoices/workflow/#auto_advance) of the invoice. When `false`, the invoice's state will not automatically advance without an explicit action.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeBool,
			},
			"automatic_tax": {
				Description: "Settings for automatic tax lookup for this invoice.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"enabled": {
					Description: "",
					Required:    true,
					Type:        schema.TypeBool,
				}}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"collection_method": {
				Description: "Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay this invoice using the default source attached to the customer. When sending an invoice, Stripe will email this invoice to the customer with payment instructions. Defaults to `charge_automatically`.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"currency": {
				Description: "The currency to create this invoice in. Defaults to that of `customer` if not specified.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"custom_fields": {
				Description: "A list of up to 4 custom fields to be displayed on the invoice.",
				Elem: &schema.Schema{
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"name": {
							Description: "",
							Required:    true,
							Type:        schema.TypeString,
						},
						"value": {
							Description: "",
							Required:    true,
							Type:        schema.TypeString,
						},
					}},
					Type: schema.TypeMap,
				},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeList,
			},
			"customer": {
				Description: "The ID of the customer who will be billed.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"days_until_due": {
				Description: "The number of days from when the invoice is created until it is due. Valid only for invoices where `collection_method=send_invoice`.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"default_payment_method": {
				Description: "ID of the default payment method for the invoice. It must belong to the customer associated with the invoice. If not set, defaults to the subscription's default payment method, if any, or to the default payment method in the customer's invoice settings.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"default_source": {
				Description: "ID of the default payment source for the invoice. It must belong to the customer associated with the invoice and be in a chargeable state. If not set, defaults to the subscription's default source, if any, or to the customer's default source.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"default_tax_rates": {
				Description: "The tax rates that will apply to any line item that does not have `tax_rates` set.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeList,
			},
			"description": {
				Description: "An arbitrary string attached to the object. Often useful for displaying to users. Referenced as 'memo' in the Dashboard.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"discounts": {
				Description: "The coupons to redeem into discounts for the invoice. If not specified, inherits the discount from the invoice's customer. Pass an empty string to avoid inheriting any discounts.",
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
			"due_date": {
				Description: "The date on which payment for this invoice is due. Valid only for invoices where `collection_method=send_invoice`.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"footer": {
				Description: "Footer to be displayed on the invoice.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"from_invoice": {
				Description: "Revise an existing invoice. The new invoice will be created in `status=draft`. See the [revision documentation](https://stripe.com/docs/invoicing/invoice-revisions) for more details.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"action": {
						Description: "",
						Required:    true,
						Type:        schema.TypeString,
					},
					"invoice": {
						Description: "",
						Required:    true,
						Type:        schema.TypeString,
					},
				}},
				ForceNew: true,
				Required: false,
				Type:     schema.TypeMap,
			},
			"on_behalf_of": {
				Description: "The account (if any) for which the funds of the invoice payment are intended. If set, the invoice will be presented with the branding and support information of the specified account. See the [Invoices with Connect](https://stripe.com/docs/billing/invoices/connect) documentation for details.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"payment_settings": {
				Description: "Configuration settings for the PaymentIntent that is generated when the invoice is finalized.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"default_mandate": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
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
									"installments": {
										Description: "",
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"enabled": {
												Description: "",
												Required:    false,
												Type:        schema.TypeBool,
											},
											"plan": {
												Description: "",
												Elem: &schema.Resource{Schema: map[string]*schema.Schema{
													"count": {
														Description: "",
														Required:    true,
														Type:        schema.TypeInt,
													},
													"interval": {
														Description: "",
														Required:    true,
														Type:        schema.TypeString,
													},
													"type": {
														Description: "",
														Required:    true,
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
				}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"pending_invoice_items_behavior": {
				Description: "How to handle pending invoice items on invoice creation. One of `include` or `exclude`. `include` will include any pending invoice items, and will create an empty draft invoice if no pending invoice items exist. `exclude` will always create an empty invoice draft regardless if there are pending invoice items or not. Defaults to `exclude` if the parameter is omitted.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"rendering_options": {
				Description: "Options for invoice PDF rendering.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"amount_tax_display": {
					Description: "",
					Required:    false,
					Type:        schema.TypeString,
				}}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"statement_descriptor": {
				Description: "Extra information about a charge for the customer's credit card statement. It must contain at least one letter. If not specified and this invoice is part of a subscription, the default `statement_descriptor` will be set to the first subscription item's product's `statement_descriptor`.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"subscription": {
				Description: "The ID of the subscription to invoice, if any. If set, the created invoice will only include pending invoice items for that subscription and pending invoice items not associated with any subscription if `pending_invoice_items_behavior` is `include`. The subscription's billing cycle and regular subscription events won't be affected.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"transfer_data": {
				Description: "If specified, the funds from the invoice will be transferred to the destination and the ID of the resulting transfer will be found on the invoice's charge.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"amount": {
						Description: "",
						Required:    false,
						Type:        schema.TypeInt,
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
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if d.HasChange("days_until_due") {
				v := d.Get("days_until_due")
				params["days_until_due"] = v
			}
			if d.HasChange("description") {
				v := d.Get("description")
				params["description"] = v
			}
			if d.HasChange("discounts") {
				v := d.Get("discounts")
				params["discounts"] = v
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
								{
									path := path.IndexString("card")
									v := shared.Dig[any](outerV, "card")
									if v != nil {
										{
											path := path
											outerV := v
											{
												path := path.IndexString("installments")
												v := shared.Dig[any](outerV, "installments")
												if v != nil {
													{
														path := path
														outerV := v
														{
															path := path.IndexString("plan")
															v := shared.Dig[any](outerV, "plan")
															if v != nil {
																{
																	path := path
																	outerV := v
																	{
																		path := path.IndexString("count")
																		v := shared.Dig[any](outerV, "count")
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
																		path := path.IndexString("interval")
																		v := shared.Dig[any](outerV, "interval")
																		if v != nil {
																			if inEnum(v.(string), []string{"month"}) {
																				out = append(out, diag.Diagnostic{
																					AttributePath: path,
																					Severity:      diag.Error,
																					Summary:       fmt.Sprintf("Field must be one of: month"),
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
																		path := path.IndexString("type")
																		v := shared.Dig[any](outerV, "type")
																		if v != nil {
																			if inEnum(v.(string), []string{"fixed_count"}) {
																				out = append(out, diag.Diagnostic{
																					AttributePath: path,
																					Severity:      diag.Error,
																					Summary:       fmt.Sprintf("Field must be one of: fixed_count"),
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
															}
														}
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
				}
				params["payment_settings"] = v
			}
			if d.HasChange("statement_descriptor") {
				v := d.Get("statement_descriptor")
				params["statement_descriptor"] = v
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
			if d.HasChange("default_payment_method") {
				v := d.Get("default_payment_method")
				params["default_payment_method"] = v
			}
			if d.HasChange("default_tax_rates") {
				v := d.Get("default_tax_rates")
				params["default_tax_rates"] = v
			}
			if d.HasChange("due_date") {
				v := d.Get("due_date")
				params["due_date"] = v
			}
			if d.HasChange("on_behalf_of") {
				v := d.Get("on_behalf_of")
				params["on_behalf_of"] = v
			}
			if d.HasChange("application_fee_amount") {
				v := d.Get("application_fee_amount")
				params["application_fee_amount"] = v
			}
			if d.HasChange("auto_advance") {
				v := d.Get("auto_advance")
				params["auto_advance"] = v
			}
			if d.HasChange("custom_fields") {
				v := d.Get("custom_fields")
				{
					path := cty.Path{}
					outerV := v
					for i, v := range outerV.([]any) {
						{
							path := path.IndexInt(i)
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
							{
								path := path.IndexString("value")
								v := shared.Dig[any](outerV, "value")
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
				params["custom_fields"] = v
			}
			if d.HasChange("rendering_options") {
				v := d.Get("rendering_options")
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("amount_tax_display")
						v := shared.Dig[any](outerV, "amount_tax_display")
						if v != nil {
							if inEnum(v.(string), []string{"", "exclude_tax", "include_inclusive_tax"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: , exclude_tax, include_inclusive_tax"),
								})
							}
						}
					}
				}
				params["rendering_options"] = v
			}
			if d.HasChange("account_tax_ids") {
				v := d.Get("account_tax_ids")
				params["account_tax_ids"] = v
			}
			if d.HasChange("default_source") {
				v := d.Get("default_source")
				params["default_source"] = v
			}
			if d.HasChange("footer") {
				v := d.Get("footer")
				params["footer"] = v
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
			_, err := f.Post(ctx, fmt.Sprintf("/v1/invoices/%v", d.Get("id")), params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to update invoice %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
	}
}
