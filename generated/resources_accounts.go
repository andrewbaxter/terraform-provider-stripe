package main

import (
	"context"
	"fmt"
	shared "github.com/andrewbaxter/terraform-provider-stripe/shared"
	cty "github.com/hashicorp/go-cty/cty"
	diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resources_accounts() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if v, exists := d.GetOk("account_token"); exists {
				params["account_token"] = v
			}
			if v, exists := d.GetOk("capabilities"); exists {
				params["capabilities"] = v
			}
			if v, exists := d.GetOk("company"); exists {
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("structure")
						v := shared.Dig[any](outerV, "structure")
						if v != nil {
							if inEnum(v.(string), []string{"", "free_zone_establishment", "free_zone_llc", "government_instrumentality", "governmental_unit", "incorporated_non_profit", "limited_liability_partnership", "llc", "multi_member_llc", "private_company", "private_corporation", "private_partnership", "public_company", "public_corporation", "public_partnership", "single_member_llc", "sole_establishment", "sole_proprietorship", "tax_exempt_government_instrumentality", "unincorporated_association", "unincorporated_non_profit"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: , free_zone_establishment, free_zone_llc, government_instrumentality, governmental_unit, incorporated_non_profit, limited_liability_partnership, llc, multi_member_llc, private_company, private_corporation, private_partnership, public_company, public_corporation, public_partnership, single_member_llc, sole_establishment, sole_proprietorship, tax_exempt_government_instrumentality, unincorporated_association, unincorporated_non_profit"),
								})
							}
						}
					}
				}
				params["company"] = v
			}
			if v, exists := d.GetOk("settings"); exists {
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("payouts")
						v := shared.Dig[any](outerV, "payouts")
						if v != nil {
							{
								path := path
								outerV := v
								{
									path := path.IndexString("schedule")
									v := shared.Dig[any](outerV, "schedule")
									if v != nil {
										{
											path := path
											outerV := v
											{
												path := path.IndexString("interval")
												v := shared.Dig[any](outerV, "interval")
												if v != nil {
													if inEnum(v.(string), []string{"daily", "manual", "monthly", "weekly"}) {
														out = append(out, diag.Diagnostic{
															AttributePath: path,
															Severity:      diag.Error,
															Summary:       fmt.Sprintf("Field must be one of: daily, manual, monthly, weekly"),
														})
													}
												}
											}
											{
												path := path.IndexString("weekly_anchor")
												v := shared.Dig[any](outerV, "weekly_anchor")
												if v != nil {
													if inEnum(v.(string), []string{"friday", "monday", "saturday", "sunday", "thursday", "tuesday", "wednesday"}) {
														out = append(out, diag.Diagnostic{
															AttributePath: path,
															Severity:      diag.Error,
															Summary:       fmt.Sprintf("Field must be one of: friday, monday, saturday, sunday, thursday, tuesday, wednesday"),
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
				params["settings"] = v
			}
			if v, exists := d.GetOk("type"); exists {
				if inEnum(v.(string), []string{"custom", "express", "standard"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: custom, express, standard"),
					})
				}
				params["type"] = v
			}
			if v, exists := d.GetOk("business_type"); exists {
				if inEnum(v.(string), []string{"company", "government_entity", "individual", "non_profit"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: company, government_entity, individual, non_profit"),
					})
				}
				params["business_type"] = v
			}
			if v, exists := d.GetOk("documents"); exists {
				params["documents"] = v
			}
			if v, exists := d.GetOk("email"); exists {
				params["email"] = v
			}
			if v, exists := d.GetOk("external_account"); exists {
				params["external_account"] = v
			}
			if v, exists := d.GetOk("individual"); exists {
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("dob")
						v := shared.Dig[any](outerV, "dob")
						if v != nil {
							{
								path := path
								outerV := v
								{
									path := path.IndexString("day")
									v := shared.Dig[any](outerV, "day")
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
									path := path.IndexString("month")
									v := shared.Dig[any](outerV, "month")
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
									path := path.IndexString("year")
									v := shared.Dig[any](outerV, "year")
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
					{
						path := path.IndexString("political_exposure")
						v := shared.Dig[any](outerV, "political_exposure")
						if v != nil {
							if inEnum(v.(string), []string{"existing", "none"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: existing, none"),
								})
							}
						}
					}
				}
				params["individual"] = v
			}
			if v, exists := d.GetOk("bank_account"); exists {
				params["bank_account"] = v
			}
			if v, exists := d.GetOk("business_profile"); exists {
				params["business_profile"] = v
			}
			if v, exists := d.GetOk("country"); exists {
				params["country"] = v
			}
			if v, exists := d.GetOk("default_currency"); exists {
				params["default_currency"] = v
			}
			if v, exists := d.GetOk("tos_acceptance"); exists {
				params["tos_acceptance"] = v
			}
			res, err := f.Post(ctx, "/v1/accounts", params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to create new account: %s", err),
				})
				return out
			}
			d.SetId(shared.Dig[string](res, "id"))
			return out
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			err := f.Delete(ctx, fmt.Sprintf("/v1/accounts/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to delete account %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			res, err := f.Get(ctx, fmt.Sprintf("/v1/accounts/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to look up account %s: %s", d.Id(), err),
				})
				return out
			}
			d.Set("account_token", shared.Dig[any](res, "account_token"))
			d.Set("capabilities", shared.Dig[any](res, "capabilities"))
			d.Set("company", shared.Dig[any](res, "company"))
			d.Set("settings", shared.Dig[any](res, "settings"))
			d.Set("type", shared.Dig[any](res, "type"))
			d.Set("business_type", shared.Dig[any](res, "business_type"))
			d.Set("documents", shared.Dig[any](res, "documents"))
			d.Set("email", shared.Dig[any](res, "email"))
			d.Set("external_account", shared.Dig[any](res, "external_account"))
			d.Set("individual", shared.Dig[any](res, "individual"))
			d.Set("bank_account", shared.Dig[any](res, "bank_account"))
			d.Set("business_profile", shared.Dig[any](res, "business_profile"))
			d.Set("country", shared.Dig[any](res, "country"))
			d.Set("default_currency", shared.Dig[any](res, "default_currency"))
			d.Set("tos_acceptance", shared.Dig[any](res, "tos_acceptance"))
			return out
		},
		Schema: map[string]*schema.Schema{
			"account_token": {
				Description: "An [account token](https://stripe.com/docs/api#create_account_token), used to securely provide details to the account.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"bank_account": {
				Description: "Either a token, like the ones returned by [Stripe.js](https://stripe.com/docs/js), or a dictionary containing a user's bank account details.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"business_profile": {
				Description: "Business information about the account.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"mcc": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"name": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"product_description": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"support_address": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"city": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"country": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"line1": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"line2": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"postal_code": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"state": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"support_email": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"support_phone": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"support_url": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"url": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
				}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"business_type": {
				Description: "The business type.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"capabilities": {
				Description: "Each key of the dictionary represents a capability, and each capability maps to its settings (e.g. whether it has been requested or not). Each capability will be inactive until you have provided its specific requirements and Stripe has verified them. An account may have some of its requested capabilities be active and some be inactive.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"acss_debit_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"affirm_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"afterpay_clearpay_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"au_becs_debit_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"bacs_debit_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"bancontact_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"bank_transfer_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"blik_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"boleto_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"card_issuing": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"card_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"cartes_bancaires_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"eps_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"fpx_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"giropay_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"grabpay_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"ideal_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"jcb_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"klarna_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"konbini_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"legacy_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"link_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"oxxo_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"p24_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"paynow_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"promptpay_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"sepa_debit_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"sofort_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"tax_reporting_us_1099_k": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"tax_reporting_us_1099_misc": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"transfers": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"treasury": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"us_bank_account_ach_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": {
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
				}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"company": {
				Description: "Information about the company or business. This field is available for any `business_type`.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"address": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"city": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"country": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"line1": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"line2": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"postal_code": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"state": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"address_kana": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"city": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"country": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"line1": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"line2": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"postal_code": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"state": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"town": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"address_kanji": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"city": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"country": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"line1": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"line2": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"postal_code": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"state": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"town": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"directors_provided": {
						Description: "",
						Required:    false,
						Type:        schema.TypeBool,
					},
					"executives_provided": {
						Description: "",
						Required:    false,
						Type:        schema.TypeBool,
					},
					"name": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"name_kana": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"name_kanji": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"owners_provided": {
						Description: "",
						Required:    false,
						Type:        schema.TypeBool,
					},
					"ownership_declaration": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"date": {
								Description: "",
								Required:    false,
								Type:        schema.TypeInt,
							},
							"ip": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"user_agent": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"phone": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"registration_number": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"structure": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"tax_id": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"tax_id_registrar": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"vat_id": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"verification": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"document": {
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"back": {
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"front": {
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
				}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"country": {
				Description: "The country in which the account holder resides, or in which the business is legally established. This should be an ISO 3166-1 alpha-2 country code. For example, if you are in the United States and the business for which you're creating an account is legally represented in Canada, you would use `CA` as the country for the account being created. Available countries include [Stripe's global markets](https://stripe.com/global) as well as countries where [cross-border payouts](https://stripe.com/docs/connect/cross-border-payouts) are supported.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"default_currency": {
				Description: "Three-letter ISO currency code representing the default currency for the account. This must be a currency that [Stripe supports in the account's country](https://stripe.com/docs/payouts).",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"documents": {
				Description: "Documents that may be submitted to satisfy various informational requests.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"bank_account_ownership_verification": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"files": {
							Description: "",
							Elem:        &schema.Schema{Type: schema.TypeString},
							Required:    false,
							Type:        schema.TypeList,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"company_license": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"files": {
							Description: "",
							Elem:        &schema.Schema{Type: schema.TypeString},
							Required:    false,
							Type:        schema.TypeList,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"company_memorandum_of_association": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"files": {
							Description: "",
							Elem:        &schema.Schema{Type: schema.TypeString},
							Required:    false,
							Type:        schema.TypeList,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"company_ministerial_decree": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"files": {
							Description: "",
							Elem:        &schema.Schema{Type: schema.TypeString},
							Required:    false,
							Type:        schema.TypeList,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"company_registration_verification": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"files": {
							Description: "",
							Elem:        &schema.Schema{Type: schema.TypeString},
							Required:    false,
							Type:        schema.TypeList,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"company_tax_id_verification": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"files": {
							Description: "",
							Elem:        &schema.Schema{Type: schema.TypeString},
							Required:    false,
							Type:        schema.TypeList,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"proof_of_registration": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"files": {
							Description: "",
							Elem:        &schema.Schema{Type: schema.TypeString},
							Required:    false,
							Type:        schema.TypeList,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
				}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"email": {
				Description: "The email address of the account holder. This is only to make the account easier to identify to you. Stripe only emails Custom accounts with your consent.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"external_account": {
				Description: "A card or bank account to attach to the account for receiving [payouts](https://stripe.com/docs/connect/bank-debit-card-payouts) (you won’t be able to use it for top-ups). You can provide either a token, like the ones returned by [Stripe.js](https://stripe.com/docs/js), or a dictionary, as documented in the `external_account` parameter for [bank account](https://stripe.com/docs/api#account_create_bank_account) creation. <br><br>By default, providing an external account sets it as the new default external account for its currency, and deletes the old default if one exists. To add additional external accounts without replacing the existing default for the currency, use the [bank account](https://stripe.com/docs/api#account_create_bank_account) or [card creation](https://stripe.com/docs/api#account_create_card) APIs.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"individual": {
				Description: "Information about the person represented by the account. This field is null unless `business_type` is set to `individual`.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"address": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"city": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"country": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"line1": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"line2": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"postal_code": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"state": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"address_kana": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"city": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"country": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"line1": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"line2": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"postal_code": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"state": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"town": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"address_kanji": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"city": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"country": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"line1": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"line2": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"postal_code": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"state": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"town": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"dob": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"day": {
								Description: "",
								Required:    true,
								Type:        schema.TypeInt,
							},
							"month": {
								Description: "",
								Required:    true,
								Type:        schema.TypeInt,
							},
							"year": {
								Description: "",
								Required:    true,
								Type:        schema.TypeInt,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"email": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"first_name": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"first_name_kana": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"first_name_kanji": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"full_name_aliases": {
						Description: "",
						Elem:        &schema.Schema{Type: schema.TypeString},
						Required:    false,
						Type:        schema.TypeList,
					},
					"gender": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"id_number": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"id_number_secondary": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"last_name": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"last_name_kana": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"last_name_kanji": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"maiden_name": {
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
					"phone": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"political_exposure": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"registered_address": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"city": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"country": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"line1": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"line2": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"postal_code": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"state": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"ssn_last_4": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"verification": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"additional_document": {
								Description: "",
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"back": {
										Description: "",
										Required:    false,
										Type:        schema.TypeString,
									},
									"front": {
										Description: "",
										Required:    false,
										Type:        schema.TypeString,
									},
								}},
								Required: false,
								Type:     schema.TypeMap,
							},
							"document": {
								Description: "",
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"back": {
										Description: "",
										Required:    false,
										Type:        schema.TypeString,
									},
									"front": {
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
				}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"settings": {
				Description: "Options for customizing how the account functions within Stripe.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"branding": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"icon": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"logo": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"primary_color": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"secondary_color": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"card_issuing": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"tos_acceptance": {
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"date": {
									Description: "",
									Required:    false,
									Type:        schema.TypeInt,
								},
								"ip": {
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"user_agent": {
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"card_payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"decline_on": {
								Description: "",
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"avs_failure": {
										Description: "",
										Required:    false,
										Type:        schema.TypeBool,
									},
									"cvc_failure": {
										Description: "",
										Required:    false,
										Type:        schema.TypeBool,
									},
								}},
								Required: false,
								Type:     schema.TypeMap,
							},
							"statement_descriptor_prefix": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"statement_descriptor_prefix_kana": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"statement_descriptor_prefix_kanji": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"payments": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"statement_descriptor": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"statement_descriptor_kana": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"statement_descriptor_kanji": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"payouts": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"debit_negative_balances": {
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							},
							"schedule": {
								Description: "",
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"delay_days": {
										Description: "",
										Required:    false,
										Type:        schema.TypeInt,
									},
									"interval": {
										Description: "",
										Required:    false,
										Type:        schema.TypeString,
									},
									"monthly_anchor": {
										Description: "",
										Required:    false,
										Type:        schema.TypeInt,
									},
									"weekly_anchor": {
										Description: "",
										Required:    false,
										Type:        schema.TypeString,
									},
								}},
								Required: false,
								Type:     schema.TypeMap,
							},
							"statement_descriptor": {
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
						}},
						Required: false,
						Type:     schema.TypeMap,
					},
					"treasury": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"tos_acceptance": {
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"date": {
									Description: "",
									Required:    false,
									Type:        schema.TypeInt,
								},
								"ip": {
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"user_agent": {
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					},
				}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"tos_acceptance": {
				Description: "Details on the account's acceptance of the [Stripe Services Agreement](https://stripe.com/docs/connect/updating-accounts#tos-acceptance).",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"date": {
						Description: "",
						Required:    false,
						Type:        schema.TypeInt,
					},
					"ip": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"service_agreement": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"user_agent": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
				}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"type": {
				Description: "The type of Stripe account to create. May be one of `custom`, `express` or `standard`.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if d.HasChange("account_token") {
				v := d.Get("account_token")
				params["account_token"] = v
			}
			if d.HasChange("capabilities") {
				v := d.Get("capabilities")
				params["capabilities"] = v
			}
			if d.HasChange("company") {
				v := d.Get("company")
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("structure")
						v := shared.Dig[any](outerV, "structure")
						if v != nil {
							if inEnum(v.(string), []string{"", "free_zone_establishment", "free_zone_llc", "government_instrumentality", "governmental_unit", "incorporated_non_profit", "limited_liability_partnership", "llc", "multi_member_llc", "private_company", "private_corporation", "private_partnership", "public_company", "public_corporation", "public_partnership", "single_member_llc", "sole_establishment", "sole_proprietorship", "tax_exempt_government_instrumentality", "unincorporated_association", "unincorporated_non_profit"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: , free_zone_establishment, free_zone_llc, government_instrumentality, governmental_unit, incorporated_non_profit, limited_liability_partnership, llc, multi_member_llc, private_company, private_corporation, private_partnership, public_company, public_corporation, public_partnership, single_member_llc, sole_establishment, sole_proprietorship, tax_exempt_government_instrumentality, unincorporated_association, unincorporated_non_profit"),
								})
							}
						}
					}
				}
				params["company"] = v
			}
			if d.HasChange("settings") {
				v := d.Get("settings")
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("payouts")
						v := shared.Dig[any](outerV, "payouts")
						if v != nil {
							{
								path := path
								outerV := v
								{
									path := path.IndexString("schedule")
									v := shared.Dig[any](outerV, "schedule")
									if v != nil {
										{
											path := path
											outerV := v
											{
												path := path.IndexString("interval")
												v := shared.Dig[any](outerV, "interval")
												if v != nil {
													if inEnum(v.(string), []string{"daily", "manual", "monthly", "weekly"}) {
														out = append(out, diag.Diagnostic{
															AttributePath: path,
															Severity:      diag.Error,
															Summary:       fmt.Sprintf("Field must be one of: daily, manual, monthly, weekly"),
														})
													}
												}
											}
											{
												path := path.IndexString("weekly_anchor")
												v := shared.Dig[any](outerV, "weekly_anchor")
												if v != nil {
													if inEnum(v.(string), []string{"friday", "monday", "saturday", "sunday", "thursday", "tuesday", "wednesday"}) {
														out = append(out, diag.Diagnostic{
															AttributePath: path,
															Severity:      diag.Error,
															Summary:       fmt.Sprintf("Field must be one of: friday, monday, saturday, sunday, thursday, tuesday, wednesday"),
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
				params["settings"] = v
			}
			if d.HasChange("business_type") {
				v := d.Get("business_type")
				if inEnum(v.(string), []string{"company", "government_entity", "individual", "non_profit"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: company, government_entity, individual, non_profit"),
					})
				}
				params["business_type"] = v
			}
			if d.HasChange("documents") {
				v := d.Get("documents")
				params["documents"] = v
			}
			if d.HasChange("email") {
				v := d.Get("email")
				params["email"] = v
			}
			if d.HasChange("external_account") {
				v := d.Get("external_account")
				params["external_account"] = v
			}
			if d.HasChange("individual") {
				v := d.Get("individual")
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("dob")
						v := shared.Dig[any](outerV, "dob")
						if v != nil {
							{
								path := path
								outerV := v
								{
									path := path.IndexString("day")
									v := shared.Dig[any](outerV, "day")
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
									path := path.IndexString("month")
									v := shared.Dig[any](outerV, "month")
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
									path := path.IndexString("year")
									v := shared.Dig[any](outerV, "year")
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
					{
						path := path.IndexString("political_exposure")
						v := shared.Dig[any](outerV, "political_exposure")
						if v != nil {
							if inEnum(v.(string), []string{"existing", "none"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: existing, none"),
								})
							}
						}
					}
				}
				params["individual"] = v
			}
			if d.HasChange("business_profile") {
				v := d.Get("business_profile")
				params["business_profile"] = v
			}
			if d.HasChange("default_currency") {
				v := d.Get("default_currency")
				params["default_currency"] = v
			}
			if d.HasChange("tos_acceptance") {
				v := d.Get("tos_acceptance")
				params["tos_acceptance"] = v
			}
			_, err := f.Post(ctx, fmt.Sprintf("/v1/accounts/%v", d.Get("id")), params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to update account %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
	}
}
