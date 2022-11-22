package main

import (
	"context"
	"fmt"
	shared "github.com/andrewbaxter/terraform-provider-stripe/shared"
	cty "github.com/hashicorp/go-cty/cty"
	diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"accounts": &schema.Resource{
			CreateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if v, exists := d.GetOk("documents"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
							path := path.IndexString("company_license")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("files")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
						}
						{
							path := path.IndexString("company_memorandum_of_association")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("files")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
						}
						{
							path := path.IndexString("company_ministerial_decree")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("files")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
						}
						{
							path := path.IndexString("company_registration_verification")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("files")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
						}
						{
							path := path.IndexString("company_tax_id_verification")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("files")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
						}
						{
							path := path.IndexString("proof_of_registration")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("files")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
						}
						{
							path := path.IndexString("bank_account_ownership_verification")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("files")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
						}
					}
					params["documents"] = v
				}
				if v, exists := d.GetOk("individual"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
							path := path.IndexString("address_kanji")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
						}
						{
						}
						{
						}
						{
						}
						{
						}
						{
							path := path.IndexString("metadata")
							outerV := shared.Dig[any](outerV)
						}
						if inEnum(shared.Dig[any](outerV), []string{"existing", "none"}) {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       fmt.Sprintf("Field must be one of: existing, none"),
							})
						}
						{
							path := path.IndexString("verification")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("document")
								outerV := shared.Dig[any](outerV)
								{
								}
								{
								}
							}
							{
								path := path.IndexString("additional_document")
								outerV := shared.Dig[any](outerV)
								{
								}
								{
								}
							}
						}
						{
							path := path.IndexString("address_kana")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
						}
						{
						}
						{
						}
						{
						}
						{
							path := path.IndexString("address")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
						}
						{
							path := path.IndexString("dob")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
						}
						{
						}
						{
							path := path.IndexString("full_name_aliases")
							outerV := shared.Dig[any](outerV)
							for i, v := range outerV.([]any) {
								{
								}
							}
						}
						{
						}
						{
						}
						{
							path := path.IndexString("registered_address")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
						}
						{
						}
						{
						}
						{
						}
					}
					params["individual"] = v
				}
				if v, exists := d.GetOk("business_type"); exists {
					if inEnum(v, []string{"company", "government_entity", "individual", "non_profit"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: company, government_entity, individual, non_profit"),
						})
					}
					params["business_type"] = v
				}
				if v, exists := d.GetOk("business_profile"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
						{
						}
						{
						}
						{
						}
						{
							path := path.IndexString("support_address")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
						}
						{
						}
						{
						}
					}
					params["business_profile"] = v
				}
				if v, exists := d.GetOk("capabilities"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
							path := path.IndexString("afterpay_clearpay_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("ideal_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("legacy_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("paynow_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("promptpay_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("affirm_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("tax_reporting_us_1099_k")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("transfers")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("oxxo_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("sofort_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("fpx_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("giropay_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("konbini_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("p24_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("us_bank_account_ach_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("eps_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("jcb_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("treasury")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("au_becs_debit_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("bacs_debit_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("bank_transfer_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("boleto_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("card_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("acss_debit_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("cartes_bancaires_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("sepa_debit_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("bancontact_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("blik_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("grabpay_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("tax_reporting_us_1099_misc")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("card_issuing")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("klarna_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("link_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
					}
					params["capabilities"] = v
				}
				if v, exists := d.GetOk("email"); exists {
					{
					}
					params["email"] = v
				}
				if v, exists := d.GetOk("settings"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
							path := path.IndexString("card_payments")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("decline_on")
								outerV := shared.Dig[any](outerV)
								{
								}
								{
								}
							}
							{
							}
							{
							}
							{
							}
						}
						{
							path := path.IndexString("payments")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
						}
						{
							path := path.IndexString("payouts")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("schedule")
								outerV := shared.Dig[any](outerV)
								{
								}
								if inEnum(shared.Dig[any](outerV), []string{"friday", "monday", "saturday", "sunday", "thursday", "tuesday", "wednesday"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: friday, monday, saturday, sunday, thursday, tuesday, wednesday"),
									})
								}
								{
								}
								if inEnum(shared.Dig[any](outerV), []string{"daily", "manual", "monthly", "weekly"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: daily, manual, monthly, weekly"),
									})
								}
							}
							{
							}
							{
							}
						}
						{
							path := path.IndexString("treasury")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("tos_acceptance")
								outerV := shared.Dig[any](outerV)
								{
								}
								{
								}
								{
								}
							}
						}
						{
							path := path.IndexString("branding")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
							{
							}
						}
						{
							path := path.IndexString("card_issuing")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("tos_acceptance")
								outerV := shared.Dig[any](outerV)
								{
								}
								{
								}
								{
								}
							}
						}
					}
					params["settings"] = v
				}
				if v, exists := d.GetOk("tos_acceptance"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
						{
						}
						{
						}
					}
					params["tos_acceptance"] = v
				}
				if v, exists := d.GetOk("type"); exists {
					if inEnum(v, []string{"custom", "express", "standard"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: custom, express, standard"),
						})
					}
					params["type"] = v
				}
				if v, exists := d.GetOk("account_token"); exists {
					{
					}
					params["account_token"] = v
				}
				if v, exists := d.GetOk("company"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
						{
						}
						{
							path := path.IndexString("address_kana")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
						}
						{
						}
						{
							path := path.IndexString("ownership_declaration")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
						}
						{
						}
						{
						}
						{
							path := path.IndexString("address_kanji")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
						}
						{
						}
						{
							path := path.IndexString("verification")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("document")
								outerV := shared.Dig[any](outerV)
								{
								}
								{
								}
							}
						}
						{
							path := path.IndexString("address")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
						}
						{
						}
						{
						}
						{
						}
						if inEnum(shared.Dig[any](outerV), []string{"", "free_zone_establishment", "free_zone_llc", "government_instrumentality", "governmental_unit", "incorporated_non_profit", "limited_liability_partnership", "llc", "multi_member_llc", "private_company", "private_corporation", "private_partnership", "public_company", "public_corporation", "public_partnership", "single_member_llc", "sole_establishment", "sole_proprietorship", "tax_exempt_government_instrumentality", "unincorporated_association", "unincorporated_non_profit"}) {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       fmt.Sprintf("Field must be one of: , free_zone_establishment, free_zone_llc, government_instrumentality, governmental_unit, incorporated_non_profit, limited_liability_partnership, llc, multi_member_llc, private_company, private_corporation, private_partnership, public_company, public_corporation, public_partnership, single_member_llc, sole_establishment, sole_proprietorship, tax_exempt_government_instrumentality, unincorporated_association, unincorporated_non_profit"),
							})
						}
						{
						}
					}
					params["company"] = v
				}
				if v, exists := d.GetOk("country"); exists {
					{
					}
					params["country"] = v
				}
				if v, exists := d.GetOk("external_account"); exists {
					{
					}
					params["external_account"] = v
				}
				if v, exists := d.GetOk("bank_account"); exists {
					{
					}
					params["bank_account"] = v
				}
				if v, exists := d.GetOk("default_currency"); exists {
					{
					}
					params["default_currency"] = v
				}
				res, err := f.Post("/v1/accounts", params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to create new account", err),
					})
					return out
				}
				d.SetId(Dig(res, "id"))
				return out
			},
			DeleteWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				err := f.Delete(fmt.Sprintf("/v1/accounts/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to delete account %s", d.Id(), err),
					})
					return out
				}
				return out
			},
			ReadWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				res, err := f.Get(fmt.Sprintf("/v1/accounts/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to look up account %s", d.Id(), err),
					})
					return out
				}
				d.Set("documents", Dig[any](res, "documents"))
				d.Set("individual", Dig[any](res, "individual"))
				d.Set("business_type", Dig[any](res, "business_type"))
				d.Set("business_profile", Dig[any](res, "business_profile"))
				d.Set("capabilities", Dig[any](res, "capabilities"))
				d.Set("email", Dig[any](res, "email"))
				d.Set("settings", Dig[any](res, "settings"))
				d.Set("tos_acceptance", Dig[any](res, "tos_acceptance"))
				d.Set("type", Dig[any](res, "type"))
				d.Set("account_token", Dig[any](res, "account_token"))
				d.Set("company", Dig[any](res, "company"))
				d.Set("country", Dig[any](res, "country"))
				d.Set("external_account", Dig[any](res, "external_account"))
				d.Set("bank_account", Dig[any](res, "bank_account"))
				d.Set("default_currency", Dig[any](res, "default_currency"))
				return out
			},
			Schema: map[string]*schema.Schema{
				"account_token": &schema.Schema{
					Description: "An [account token](https://stripe.com/docs/api#create_account_token), used to securely provide details to the account.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"bank_account": &schema.Schema{
					Description: "Either a token, like the ones returned by [Stripe.js](https://stripe.com/docs/js), or a dictionary containing a user's bank account details.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"business_profile": &schema.Schema{
					Description: "Business information about the account.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"mcc": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"name": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"product_description": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"support_address": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"city": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"country": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"line1": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"line2": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"postal_code": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"state": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"support_email": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"support_phone": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"support_url": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"url": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
					}},
					ForceNew: false,
					Required: false,
					Type:     schema.TypeMap,
				},
				"business_type": &schema.Schema{
					Description: "The business type.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"capabilities": &schema.Schema{
					Description: "Each key of the dictionary represents a capability, and each capability maps to its settings (e.g. whether it has been requested or not). Each capability will be inactive until you have provided its specific requirements and Stripe has verified them. An account may have some of its requested capabilities be active and some be inactive.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"acss_debit_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"affirm_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"afterpay_clearpay_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"au_becs_debit_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"bacs_debit_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"bancontact_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"bank_transfer_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"blik_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"boleto_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"card_issuing": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"card_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"cartes_bancaires_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"eps_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"fpx_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"giropay_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"grabpay_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"ideal_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"jcb_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"klarna_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"konbini_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"legacy_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"link_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"oxxo_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"p24_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"paynow_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"promptpay_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"sepa_debit_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"sofort_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"tax_reporting_us_1099_k": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"tax_reporting_us_1099_misc": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"transfers": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"treasury": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeBool,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"us_bank_account_ach_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"requested": &schema.Schema{
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
				"company": &schema.Schema{
					Description: "Information about the company or business. This field is available for any `business_type`.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"address": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"city": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"country": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"line1": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"line2": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"postal_code": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"state": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"address_kana": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"city": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"country": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"line1": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"line2": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"postal_code": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"state": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"town": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"address_kanji": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"city": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"country": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"line1": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"line2": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"postal_code": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"state": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"town": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"directors_provided": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						},
						"executives_provided": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						},
						"name": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"name_kana": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"name_kanji": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"owners_provided": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						},
						"ownership_declaration": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"date": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeInt,
								},
								"ip": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"user_agent": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"phone": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"registration_number": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"structure": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"tax_id": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"tax_id_registrar": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"vat_id": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"verification": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"document": &schema.Schema{
								Description: "",
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"back": &schema.Schema{
										Description: "",
										Required:    false,
										Type:        schema.TypeString,
									},
									"front": &schema.Schema{
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
				"country": &schema.Schema{
					Description: "The country in which the account holder resides, or in which the business is legally established. This should be an ISO 3166-1 alpha-2 country code. For example, if you are in the United States and the business for which you're creating an account is legally represented in Canada, you would use `CA` as the country for the account being created. Available countries include [Stripe's global markets](https://stripe.com/global) as well as countries where [cross-border payouts](https://stripe.com/docs/connect/cross-border-payouts) are supported.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"default_currency": &schema.Schema{
					Description: "Three-letter ISO currency code representing the default currency for the account. This must be a currency that [Stripe supports in the account's country](https://stripe.com/docs/payouts).",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"documents": &schema.Schema{
					Description: "Documents that may be submitted to satisfy various informational requests.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"bank_account_ownership_verification": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"files": &schema.Schema{
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeString},
								Required:    false,
								Type:        schema.TypeList,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"company_license": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"files": &schema.Schema{
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeString},
								Required:    false,
								Type:        schema.TypeList,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"company_memorandum_of_association": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"files": &schema.Schema{
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeString},
								Required:    false,
								Type:        schema.TypeList,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"company_ministerial_decree": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"files": &schema.Schema{
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeString},
								Required:    false,
								Type:        schema.TypeList,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"company_registration_verification": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"files": &schema.Schema{
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeString},
								Required:    false,
								Type:        schema.TypeList,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"company_tax_id_verification": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"files": &schema.Schema{
								Description: "",
								Elem:        &schema.Schema{Type: schema.TypeString},
								Required:    false,
								Type:        schema.TypeList,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"proof_of_registration": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"files": &schema.Schema{
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
				"email": &schema.Schema{
					Description: "The email address of the account holder. This is only to make the account easier to identify to you. Stripe only emails Custom accounts with your consent.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"external_account": &schema.Schema{
					Description: "A card or bank account to attach to the account for receiving [payouts](https://stripe.com/docs/connect/bank-debit-card-payouts) (you won’t be able to use it for top-ups). You can provide either a token, like the ones returned by [Stripe.js](https://stripe.com/docs/js), or a dictionary, as documented in the `external_account` parameter for [bank account](https://stripe.com/docs/api#account_create_bank_account) creation. <br><br>By default, providing an external account sets it as the new default external account for its currency, and deletes the old default if one exists. To add additional external accounts without replacing the existing default for the currency, use the [bank account](https://stripe.com/docs/api#account_create_bank_account) or [card creation](https://stripe.com/docs/api#account_create_card) APIs.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"individual": &schema.Schema{
					Description: "Information about the person represented by the account. This field is null unless `business_type` is set to `individual`.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"address": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"city": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"country": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"line1": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"line2": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"postal_code": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"state": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"address_kana": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"city": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"country": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"line1": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"line2": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"postal_code": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"state": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"town": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"address_kanji": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"city": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"country": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"line1": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"line2": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"postal_code": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"state": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"town": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"dob": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"day": &schema.Schema{
									Description: "",
									Required:    true,
									Type:        schema.TypeInt,
								},
								"month": &schema.Schema{
									Description: "",
									Required:    true,
									Type:        schema.TypeInt,
								},
								"year": &schema.Schema{
									Description: "",
									Required:    true,
									Type:        schema.TypeInt,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"email": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"first_name": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"first_name_kana": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"first_name_kanji": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"full_name_aliases": &schema.Schema{
							Description: "",
							Elem:        &schema.Schema{Type: schema.TypeString},
							Required:    false,
							Type:        schema.TypeList,
						},
						"gender": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"id_number": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"id_number_secondary": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"last_name": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"last_name_kana": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"last_name_kanji": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"maiden_name": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"metadata": &schema.Schema{
							Description: "",
							Elem:        &schema.Resource{Schema: map[string]*schema.Schema{}},
							Required:    false,
							Type:        schema.TypeMap,
						},
						"phone": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"political_exposure": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"registered_address": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"city": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"country": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"line1": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"line2": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"postal_code": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"state": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"ssn_last_4": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"verification": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"additional_document": &schema.Schema{
									Description: "",
									Elem: &schema.Resource{Schema: map[string]*schema.Schema{
										"back": &schema.Schema{
											Description: "",
											Required:    false,
											Type:        schema.TypeString,
										},
										"front": &schema.Schema{
											Description: "",
											Required:    false,
											Type:        schema.TypeString,
										},
									}},
									Required: false,
									Type:     schema.TypeMap,
								},
								"document": &schema.Schema{
									Description: "",
									Elem: &schema.Resource{Schema: map[string]*schema.Schema{
										"back": &schema.Schema{
											Description: "",
											Required:    false,
											Type:        schema.TypeString,
										},
										"front": &schema.Schema{
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
				"settings": &schema.Schema{
					Description: "Options for customizing how the account functions within Stripe.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"branding": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"icon": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"logo": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"primary_color": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"secondary_color": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"card_issuing": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"tos_acceptance": &schema.Schema{
								Description: "",
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"date": &schema.Schema{
										Description: "",
										Required:    false,
										Type:        schema.TypeInt,
									},
									"ip": &schema.Schema{
										Description: "",
										Required:    false,
										Type:        schema.TypeString,
									},
									"user_agent": &schema.Schema{
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
						"card_payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"decline_on": &schema.Schema{
									Description: "",
									Elem: &schema.Resource{Schema: map[string]*schema.Schema{
										"avs_failure": &schema.Schema{
											Description: "",
											Required:    false,
											Type:        schema.TypeBool,
										},
										"cvc_failure": &schema.Schema{
											Description: "",
											Required:    false,
											Type:        schema.TypeBool,
										},
									}},
									Required: false,
									Type:     schema.TypeMap,
								},
								"statement_descriptor_prefix": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"statement_descriptor_prefix_kana": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"statement_descriptor_prefix_kanji": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"payments": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"statement_descriptor": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"statement_descriptor_kana": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"statement_descriptor_kanji": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"payouts": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"debit_negative_balances": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeBool,
								},
								"schedule": &schema.Schema{
									Description: "",
									Elem: &schema.Resource{Schema: map[string]*schema.Schema{
										"delay_days": &schema.Schema{
											Description: "",
											Required:    false,
											Type:        schema.TypeInt,
										},
										"interval": &schema.Schema{
											Description: "",
											Required:    false,
											Type:        schema.TypeString,
										},
										"monthly_anchor": &schema.Schema{
											Description: "",
											Required:    false,
											Type:        schema.TypeInt,
										},
										"weekly_anchor": &schema.Schema{
											Description: "",
											Required:    false,
											Type:        schema.TypeString,
										},
									}},
									Required: false,
									Type:     schema.TypeMap,
								},
								"statement_descriptor": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"treasury": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"tos_acceptance": &schema.Schema{
								Description: "",
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"date": &schema.Schema{
										Description: "",
										Required:    false,
										Type:        schema.TypeInt,
									},
									"ip": &schema.Schema{
										Description: "",
										Required:    false,
										Type:        schema.TypeString,
									},
									"user_agent": &schema.Schema{
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
				"tos_acceptance": &schema.Schema{
					Description: "Details on the account's acceptance of the [Stripe Services Agreement](https://stripe.com/docs/connect/updating-accounts#tos-acceptance).",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"date": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeInt,
						},
						"ip": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"service_agreement": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"user_agent": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
					}},
					ForceNew: false,
					Required: false,
					Type:     schema.TypeMap,
				},
				"type": &schema.Schema{
					Description: "The type of Stripe account to create. May be one of `custom`, `express` or `standard`.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
			},
			UpdateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if d.HasChange("documents") {
					v := d.Get("documents")
					{
						path := cty.Path{}
						outerV := v
						{
							path := path.IndexString("company_license")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("files")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
						}
						{
							path := path.IndexString("company_memorandum_of_association")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("files")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
						}
						{
							path := path.IndexString("company_ministerial_decree")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("files")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
						}
						{
							path := path.IndexString("company_registration_verification")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("files")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
						}
						{
							path := path.IndexString("company_tax_id_verification")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("files")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
						}
						{
							path := path.IndexString("proof_of_registration")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("files")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
						}
						{
							path := path.IndexString("bank_account_ownership_verification")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("files")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
						}
					}
					params["documents"] = v
				}
				if d.HasChange("individual") {
					v := d.Get("individual")
					{
						path := cty.Path{}
						outerV := v
						{
							path := path.IndexString("address_kanji")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
						}
						{
						}
						{
						}
						{
						}
						{
						}
						{
							path := path.IndexString("metadata")
							outerV := shared.Dig[any](outerV)
						}
						if inEnum(shared.Dig[any](outerV), []string{"existing", "none"}) {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       fmt.Sprintf("Field must be one of: existing, none"),
							})
						}
						{
							path := path.IndexString("verification")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("document")
								outerV := shared.Dig[any](outerV)
								{
								}
								{
								}
							}
							{
								path := path.IndexString("additional_document")
								outerV := shared.Dig[any](outerV)
								{
								}
								{
								}
							}
						}
						{
							path := path.IndexString("address_kana")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
						}
						{
						}
						{
						}
						{
						}
						{
							path := path.IndexString("address")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
						}
						{
							path := path.IndexString("dob")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
						}
						{
						}
						{
							path := path.IndexString("full_name_aliases")
							outerV := shared.Dig[any](outerV)
							for i, v := range outerV.([]any) {
								{
								}
							}
						}
						{
						}
						{
						}
						{
							path := path.IndexString("registered_address")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
						}
						{
						}
						{
						}
						{
						}
					}
					params["individual"] = v
				}
				if d.HasChange("business_type") {
					v := d.Get("business_type")
					if inEnum(v, []string{"company", "government_entity", "individual", "non_profit"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: company, government_entity, individual, non_profit"),
						})
					}
					params["business_type"] = v
				}
				if d.HasChange("business_profile") {
					v := d.Get("business_profile")
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
						{
						}
						{
						}
						{
						}
						{
							path := path.IndexString("support_address")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
						}
						{
						}
						{
						}
					}
					params["business_profile"] = v
				}
				if d.HasChange("capabilities") {
					v := d.Get("capabilities")
					{
						path := cty.Path{}
						outerV := v
						{
							path := path.IndexString("afterpay_clearpay_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("ideal_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("legacy_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("paynow_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("promptpay_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("affirm_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("tax_reporting_us_1099_k")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("transfers")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("oxxo_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("sofort_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("fpx_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("giropay_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("konbini_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("p24_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("us_bank_account_ach_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("eps_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("jcb_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("treasury")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("au_becs_debit_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("bacs_debit_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("bank_transfer_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("boleto_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("card_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("acss_debit_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("cartes_bancaires_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("sepa_debit_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("bancontact_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("blik_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("grabpay_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("tax_reporting_us_1099_misc")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("card_issuing")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("klarna_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
						{
							path := path.IndexString("link_payments")
							outerV := shared.Dig[any](outerV)
							{
							}
						}
					}
					params["capabilities"] = v
				}
				if d.HasChange("email") {
					v := d.Get("email")
					{
					}
					params["email"] = v
				}
				if d.HasChange("settings") {
					v := d.Get("settings")
					{
						path := cty.Path{}
						outerV := v
						{
							path := path.IndexString("card_payments")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("decline_on")
								outerV := shared.Dig[any](outerV)
								{
								}
								{
								}
							}
							{
							}
							{
							}
							{
							}
						}
						{
							path := path.IndexString("payments")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
						}
						{
							path := path.IndexString("payouts")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("schedule")
								outerV := shared.Dig[any](outerV)
								{
								}
								if inEnum(shared.Dig[any](outerV), []string{"friday", "monday", "saturday", "sunday", "thursday", "tuesday", "wednesday"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: friday, monday, saturday, sunday, thursday, tuesday, wednesday"),
									})
								}
								{
								}
								if inEnum(shared.Dig[any](outerV), []string{"daily", "manual", "monthly", "weekly"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: daily, manual, monthly, weekly"),
									})
								}
							}
							{
							}
							{
							}
						}
						{
							path := path.IndexString("treasury")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("tos_acceptance")
								outerV := shared.Dig[any](outerV)
								{
								}
								{
								}
								{
								}
							}
						}
						{
							path := path.IndexString("branding")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
							{
							}
						}
						{
							path := path.IndexString("card_issuing")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("tos_acceptance")
								outerV := shared.Dig[any](outerV)
								{
								}
								{
								}
								{
								}
							}
						}
					}
					params["settings"] = v
				}
				if d.HasChange("tos_acceptance") {
					v := d.Get("tos_acceptance")
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
						{
						}
						{
						}
					}
					params["tos_acceptance"] = v
				}
				if d.HasChange("account_token") {
					v := d.Get("account_token")
					{
					}
					params["account_token"] = v
				}
				if d.HasChange("company") {
					v := d.Get("company")
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
						{
						}
						{
							path := path.IndexString("address_kana")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
						}
						{
						}
						{
							path := path.IndexString("ownership_declaration")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
						}
						{
						}
						{
						}
						{
							path := path.IndexString("address_kanji")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
						}
						{
						}
						{
							path := path.IndexString("verification")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("document")
								outerV := shared.Dig[any](outerV)
								{
								}
								{
								}
							}
						}
						{
							path := path.IndexString("address")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
						}
						{
						}
						{
						}
						{
						}
						if inEnum(shared.Dig[any](outerV), []string{"", "free_zone_establishment", "free_zone_llc", "government_instrumentality", "governmental_unit", "incorporated_non_profit", "limited_liability_partnership", "llc", "multi_member_llc", "private_company", "private_corporation", "private_partnership", "public_company", "public_corporation", "public_partnership", "single_member_llc", "sole_establishment", "sole_proprietorship", "tax_exempt_government_instrumentality", "unincorporated_association", "unincorporated_non_profit"}) {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       fmt.Sprintf("Field must be one of: , free_zone_establishment, free_zone_llc, government_instrumentality, governmental_unit, incorporated_non_profit, limited_liability_partnership, llc, multi_member_llc, private_company, private_corporation, private_partnership, public_company, public_corporation, public_partnership, single_member_llc, sole_establishment, sole_proprietorship, tax_exempt_government_instrumentality, unincorporated_association, unincorporated_non_profit"),
							})
						}
						{
						}
					}
					params["company"] = v
				}
				if d.HasChange("external_account") {
					v := d.Get("external_account")
					{
					}
					params["external_account"] = v
				}
				if d.HasChange("default_currency") {
					v := d.Get("default_currency")
					{
					}
					params["default_currency"] = v
				}
				_, err := f.Post(fmt.Sprintf("/v1/accounts/%v", d.Get("id")), params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to update account %s", d.Id(), err),
					})
					return out
				}
				return out
			},
		},
		"apple_pay_domains": &schema.Resource{
			CreateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				{
					v := d.Get("domain_name")
					{
					}
					params["domain_name"] = v
				}
				res, err := f.Post("/v1/apple_pay/domains", params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to create new apple_pay_domain", err),
					})
					return out
				}
				d.SetId(Dig(res, "id"))
				return out
			},
			DeleteWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				err := f.Delete(fmt.Sprintf("/v1/apple_pay/domains/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to delete apple_pay_domain %s", d.Id(), err),
					})
					return out
				}
				return out
			},
			ReadWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				res, err := f.Get(fmt.Sprintf("/v1/apple_pay/domains/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to look up apple_pay_domain %s", d.Id(), err),
					})
					return out
				}
				d.Set("domain_name", Dig[any](res, "domain_name"))
				return out
			},
			Schema: map[string]*schema.Schema{"domain_name": &schema.Schema{
				Description: "",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
			}},
			UpdateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				_, err := f.Post(fmt.Sprintf("/v1/apple_pay/domains/%v", d.Get("id")), params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to update apple_pay_domain %s", d.Id(), err),
					})
					return out
				}
				return out
			},
		},
		"coupons": &schema.Resource{
			CreateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if v, exists := d.GetOk("max_redemptions"); exists {
					{
					}
					params["max_redemptions"] = v
				}
				if v, exists := d.GetOk("percent_off"); exists {
					{
					}
					params["percent_off"] = v
				}
				if v, exists := d.GetOk("applies_to"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
							path := path.IndexString("products")
							outerV := shared.Dig[any](outerV)
							for i, v := range outerV.([]any) {
								{
								}
							}
						}
					}
					params["applies_to"] = v
				}
				if v, exists := d.GetOk("currency"); exists {
					{
					}
					params["currency"] = v
				}
				if v, exists := d.GetOk("duration"); exists {
					if inEnum(v, []string{"forever", "once", "repeating"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: forever, once, repeating"),
						})
					}
					params["duration"] = v
				}
				if v, exists := d.GetOk("duration_in_months"); exists {
					{
					}
					params["duration_in_months"] = v
				}
				if v, exists := d.GetOk("id"); exists {
					{
					}
					params["id"] = v
				}
				if v, exists := d.GetOk("amount_off"); exists {
					{
					}
					params["amount_off"] = v
				}
				if v, exists := d.GetOk("currency_options"); exists {
					{
						path := cty.Path{}
						outerV := v
					}
					params["currency_options"] = v
				}
				if v, exists := d.GetOk("name"); exists {
					{
					}
					params["name"] = v
				}
				if v, exists := d.GetOk("redeem_by"); exists {
					{
					}
					params["redeem_by"] = v
				}
				res, err := f.Post("/v1/coupons", params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to create new coupon", err),
					})
					return out
				}
				d.SetId(Dig(res, "id"))
				return out
			},
			DeleteWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				err := f.Delete(fmt.Sprintf("/v1/coupons/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to delete coupon %s", d.Id(), err),
					})
					return out
				}
				return out
			},
			ReadWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				res, err := f.Get(fmt.Sprintf("/v1/coupons/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to look up coupon %s", d.Id(), err),
					})
					return out
				}
				d.Set("max_redemptions", Dig[any](res, "max_redemptions"))
				d.Set("percent_off", Dig[any](res, "percent_off"))
				d.Set("applies_to", Dig[any](res, "applies_to"))
				d.Set("currency", Dig[any](res, "currency"))
				d.Set("duration", Dig[any](res, "duration"))
				d.Set("duration_in_months", Dig[any](res, "duration_in_months"))
				d.Set("id", Dig[any](res, "id"))
				d.Set("amount_off", Dig[any](res, "amount_off"))
				d.Set("currency_options", Dig[any](res, "currency_options"))
				d.Set("name", Dig[any](res, "name"))
				d.Set("redeem_by", Dig[any](res, "redeem_by"))
				return out
			},
			Schema: map[string]*schema.Schema{
				"amount_off": &schema.Schema{
					Description: "A positive integer representing the amount to subtract from an invoice total (required if `percent_off` is not passed).",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"applies_to": &schema.Schema{
					Description: "A hash containing directions for what this Coupon will apply discounts to.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{"products": &schema.Schema{
						Description: "",
						Elem:        &schema.Schema{Type: schema.TypeString},
						Required:    false,
						Type:        schema.TypeList,
					}}},
					ForceNew: true,
					Required: false,
					Type:     schema.TypeMap,
				},
				"currency": &schema.Schema{
					Description: "Three-letter [ISO code for the currency](https://stripe.com/docs/currencies) of the `amount_off` parameter (required if `amount_off` is passed).",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"currency_options": &schema.Schema{
					Description: "Coupons defined in each available currency option (only supported if `amount_off` is passed). Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).",
					Elem:        &schema.Resource{Schema: map[string]*schema.Schema{}},
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeMap,
				},
				"duration": &schema.Schema{
					Description: "Specifies how long the discount will be in effect if used on a subscription. Defaults to `once`.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"duration_in_months": &schema.Schema{
					Description: "Required only if `duration` is `repeating`, in which case it must be a positive integer that specifies the number of months the discount will be in effect.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"id": &schema.Schema{
					Description: "Unique string of your choice that will be used to identify this coupon when applying it to a customer. If you don't want to specify a particular code, you can leave the ID blank and we'll generate a random code for you.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"max_redemptions": &schema.Schema{
					Description: "A positive integer specifying the number of times the coupon can be redeemed before it's no longer valid. For example, you might have a 50% off coupon that the first 20 readers of your blog can use.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"name": &schema.Schema{
					Description: "Name of the coupon displayed to customers on, for instance invoices, or receipts. By default the `id` is shown if `name` is not set.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"percent_off": &schema.Schema{
					Description: "A positive float larger than 0, and smaller or equal to 100, that represents the discount the coupon will apply (required if `amount_off` is not passed).",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeFloat,
				},
				"redeem_by": &schema.Schema{
					Description: "Unix timestamp specifying the last time at which the coupon can be redeemed. After the redeem_by date, the coupon can no longer be applied to new customers.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeInt,
				},
			},
			UpdateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if d.HasChange("currency_options") {
					v := d.Get("currency_options")
					{
						path := cty.Path{}
						outerV := v
					}
					params["currency_options"] = v
				}
				if d.HasChange("name") {
					v := d.Get("name")
					{
					}
					params["name"] = v
				}
				_, err := f.Post(fmt.Sprintf("/v1/coupons/%v", d.Get("id")), params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to update coupon %s", d.Id(), err),
					})
					return out
				}
				return out
			},
		},
		"customers": &schema.Resource{
			CreateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if v, exists := d.GetOk("balance"); exists {
					{
					}
					params["balance"] = v
				}
				if v, exists := d.GetOk("description"); exists {
					{
					}
					params["description"] = v
				}
				if v, exists := d.GetOk("source"); exists {
					{
					}
					params["source"] = v
				}
				if v, exists := d.GetOk("name"); exists {
					{
					}
					params["name"] = v
				}
				if v, exists := d.GetOk("shipping"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
							path := path.IndexString("address")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
						}
						{
						}
						{
						}
					}
					params["shipping"] = v
				}
				if v, exists := d.GetOk("tax"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
						}
					}
					params["tax"] = v
				}
				if v, exists := d.GetOk("tax_id_data"); exists {
					{
						path := cty.Path{}
						outerV := v
						for i, v := range outerV.([]any) {
							{
								path := path.IndexInt(i)
								outerV := v
								if inEnum(shared.Dig[any](outerV), []string{"ae_trn", "au_abn", "au_arn", "bg_uic", "br_cnpj", "br_cpf", "ca_bn", "ca_gst_hst", "ca_pst_bc", "ca_pst_mb", "ca_pst_sk", "ca_qst", "ch_vat", "cl_tin", "eg_tin", "es_cif", "eu_oss_vat", "eu_vat", "gb_vat", "ge_vat", "hk_br", "hu_tin", "id_npwp", "il_vat", "in_gst", "is_vat", "jp_cn", "jp_rn", "jp_trn", "ke_pin", "kr_brn", "li_uid", "mx_rfc", "my_frp", "my_itn", "my_sst", "no_vat", "nz_gst", "ph_tin", "ru_inn", "ru_kpp", "sa_vat", "sg_gst", "sg_uen", "si_tin", "th_vat", "tr_tin", "tw_vat", "ua_vat", "us_ein", "za_vat"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: ae_trn, au_abn, au_arn, bg_uic, br_cnpj, br_cpf, ca_bn, ca_gst_hst, ca_pst_bc, ca_pst_mb, ca_pst_sk, ca_qst, ch_vat, cl_tin, eg_tin, es_cif, eu_oss_vat, eu_vat, gb_vat, ge_vat, hk_br, hu_tin, id_npwp, il_vat, in_gst, is_vat, jp_cn, jp_rn, jp_trn, ke_pin, kr_brn, li_uid, mx_rfc, my_frp, my_itn, my_sst, no_vat, nz_gst, ph_tin, ru_inn, ru_kpp, sa_vat, sg_gst, sg_uen, si_tin, th_vat, tr_tin, tw_vat, ua_vat, us_ein, za_vat"),
									})
								}
								{
								}
							}
						}
					}
					params["tax_id_data"] = v
				}
				if v, exists := d.GetOk("cash_balance"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
							path := path.IndexString("settings")
							outerV := shared.Dig[any](outerV)
							if inEnum(shared.Dig[any](outerV), []string{"automatic", "manual"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: automatic, manual"),
								})
							}
						}
					}
					params["cash_balance"] = v
				}
				if v, exists := d.GetOk("invoice_prefix"); exists {
					{
					}
					params["invoice_prefix"] = v
				}
				if v, exists := d.GetOk("invoice_settings"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
							path := path.IndexString("custom_fields")
							outerV := shared.Dig[any](outerV)
							for i, v := range outerV.([]any) {
								{
									path := path.IndexInt(i)
									outerV := v
									{
									}
									{
									}
								}
							}
						}
						{
						}
						{
						}
						{
							path := path.IndexString("rendering_options")
							outerV := shared.Dig[any](outerV)
							if inEnum(shared.Dig[any](outerV), []string{"", "exclude_tax", "include_inclusive_tax"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: , exclude_tax, include_inclusive_tax"),
								})
							}
						}
					}
					params["invoice_settings"] = v
				}
				if v, exists := d.GetOk("next_invoice_sequence"); exists {
					{
					}
					params["next_invoice_sequence"] = v
				}
				if v, exists := d.GetOk("phone"); exists {
					{
					}
					params["phone"] = v
				}
				if v, exists := d.GetOk("test_clock"); exists {
					{
					}
					params["test_clock"] = v
				}
				if v, exists := d.GetOk("promotion_code"); exists {
					{
					}
					params["promotion_code"] = v
				}
				if v, exists := d.GetOk("tax_exempt"); exists {
					if inEnum(v, []string{"", "exempt", "none", "reverse"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: , exempt, none, reverse"),
						})
					}
					params["tax_exempt"] = v
				}
				if v, exists := d.GetOk("address"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
						{
						}
						{
						}
						{
						}
						{
						}
					}
					params["address"] = v
				}
				if v, exists := d.GetOk("coupon"); exists {
					{
					}
					params["coupon"] = v
				}
				if v, exists := d.GetOk("email"); exists {
					{
					}
					params["email"] = v
				}
				if v, exists := d.GetOk("payment_method"); exists {
					{
					}
					params["payment_method"] = v
				}
				if v, exists := d.GetOk("preferred_locales"); exists {
					{
						path := cty.Path{}
						outerV := v
						for i, v := range outerV.([]any) {
							{
							}
						}
					}
					params["preferred_locales"] = v
				}
				res, err := f.Post("/v1/customers", params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to create new customer", err),
					})
					return out
				}
				d.SetId(Dig(res, "id"))
				return out
			},
			DeleteWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				err := f.Delete(fmt.Sprintf("/v1/customers/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to delete customer %s", d.Id(), err),
					})
					return out
				}
				return out
			},
			ReadWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				res, err := f.Get(fmt.Sprintf("/v1/customers/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to look up customer %s", d.Id(), err),
					})
					return out
				}
				d.Set("balance", Dig[any](res, "balance"))
				d.Set("description", Dig[any](res, "description"))
				d.Set("source", Dig[any](res, "source"))
				d.Set("name", Dig[any](res, "name"))
				d.Set("shipping", Dig[any](res, "shipping"))
				d.Set("tax", Dig[any](res, "tax"))
				d.Set("tax_id_data", Dig[any](res, "tax_id_data"))
				d.Set("cash_balance", Dig[any](res, "cash_balance"))
				d.Set("invoice_prefix", Dig[any](res, "invoice_prefix"))
				d.Set("invoice_settings", Dig[any](res, "invoice_settings"))
				d.Set("next_invoice_sequence", Dig[any](res, "next_invoice_sequence"))
				d.Set("phone", Dig[any](res, "phone"))
				d.Set("test_clock", Dig[any](res, "test_clock"))
				d.Set("promotion_code", Dig[any](res, "promotion_code"))
				d.Set("tax_exempt", Dig[any](res, "tax_exempt"))
				d.Set("address", Dig[any](res, "address"))
				d.Set("coupon", Dig[any](res, "coupon"))
				d.Set("email", Dig[any](res, "email"))
				d.Set("payment_method", Dig[any](res, "payment_method"))
				d.Set("preferred_locales", Dig[any](res, "preferred_locales"))
				return out
			},
			Schema: map[string]*schema.Schema{
				"address": &schema.Schema{
					Description: "The customer's address.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"city": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"country": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"line1": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"line2": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"postal_code": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"state": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
					}},
					ForceNew: false,
					Required: false,
					Type:     schema.TypeMap,
				},
				"balance": &schema.Schema{
					Description: "An integer amount in cents (or local equivalent) that represents the customer's current balance, which affect the customer's future invoices. A negative amount represents a credit that decreases the amount due on an invoice; a positive amount increases the amount due on an invoice.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"cash_balance": &schema.Schema{
					Description: "Balance information and default balance settings for this customer.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{"settings": &schema.Schema{
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"reconciliation_mode": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						}}},
						Required: false,
						Type:     schema.TypeMap,
					}}},
					ForceNew: false,
					Required: false,
					Type:     schema.TypeMap,
				},
				"coupon": &schema.Schema{
					Description: "",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"description": &schema.Schema{
					Description: "An arbitrary string that you can attach to a customer object. It is displayed alongside the customer in the dashboard.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"email": &schema.Schema{
					Description: "Customer's email address. It's displayed alongside the customer in your dashboard and can be useful for searching and tracking. This may be up to *512 characters*.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"invoice_prefix": &schema.Schema{
					Description: "The prefix for the customer used to generate unique invoice numbers. Must be 3–12 uppercase letters or numbers.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"invoice_settings": &schema.Schema{
					Description: "Default invoice settings for this customer.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"custom_fields": &schema.Schema{
							Description: "",
							Elem: &schema.Schema{
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Description: "",
										Required:    true,
										Type:        schema.TypeString,
									},
									"value": &schema.Schema{
										Description: "",
										Required:    true,
										Type:        schema.TypeString,
									},
								}},
								Type: schema.TypeMap,
							},
							Required: false,
							Type:     schema.TypeList,
						},
						"default_payment_method": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"footer": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"rendering_options": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{"amount_tax_display": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							}}},
							Required: false,
							Type:     schema.TypeMap,
						},
					}},
					ForceNew: false,
					Required: false,
					Type:     schema.TypeMap,
				},
				"name": &schema.Schema{
					Description: "The customer's full name or business name.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"next_invoice_sequence": &schema.Schema{
					Description: "The sequence to be used on the customer's next invoice. Defaults to 1.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"payment_method": &schema.Schema{
					Description: "",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"phone": &schema.Schema{
					Description: "The customer's phone number.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"preferred_locales": &schema.Schema{
					Description: "Customer's preferred languages, ordered by preference.",
					Elem:        &schema.Schema{Type: schema.TypeString},
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeList,
				},
				"promotion_code": &schema.Schema{
					Description: "The API ID of a promotion code to apply to the customer. The customer will have a discount applied on all recurring payments. Charges you create through the API will not have the discount.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"shipping": &schema.Schema{
					Description: "The customer's shipping information. Appears on invoices emailed to this customer.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"address": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"city": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"country": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"line1": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"line2": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"postal_code": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
								"state": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeString,
								},
							}},
							Required: true,
							Type:     schema.TypeMap,
						},
						"name": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeString,
						},
						"phone": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
					}},
					ForceNew: false,
					Required: false,
					Type:     schema.TypeMap,
				},
				"source": &schema.Schema{
					Description: "",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"tax": &schema.Schema{
					Description: "Tax details about the customer.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{"ip_address": &schema.Schema{
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					}}},
					ForceNew: false,
					Required: false,
					Type:     schema.TypeMap,
				},
				"tax_exempt": &schema.Schema{
					Description: "The customer's tax exemption. One of `none`, `exempt`, or `reverse`.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"tax_id_data": &schema.Schema{
					Description: "The customer's tax IDs.",
					Elem: &schema.Schema{
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"type": &schema.Schema{
								Description: "",
								Required:    true,
								Type:        schema.TypeString,
							},
							"value": &schema.Schema{
								Description: "",
								Required:    true,
								Type:        schema.TypeString,
							},
						}},
						Type: schema.TypeMap,
					},
					ForceNew: true,
					Required: false,
					Type:     schema.TypeList,
				},
				"test_clock": &schema.Schema{
					Description: "ID of the test clock to attach to the customer.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
			},
			UpdateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if d.HasChange("balance") {
					v := d.Get("balance")
					{
					}
					params["balance"] = v
				}
				if d.HasChange("description") {
					v := d.Get("description")
					{
					}
					params["description"] = v
				}
				if d.HasChange("source") {
					v := d.Get("source")
					{
					}
					params["source"] = v
				}
				if d.HasChange("name") {
					v := d.Get("name")
					{
					}
					params["name"] = v
				}
				if d.HasChange("shipping") {
					v := d.Get("shipping")
					{
						path := cty.Path{}
						outerV := v
						{
							path := path.IndexString("address")
							outerV := shared.Dig[any](outerV)
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
							{
							}
						}
						{
						}
						{
						}
					}
					params["shipping"] = v
				}
				if d.HasChange("tax") {
					v := d.Get("tax")
					{
						path := cty.Path{}
						outerV := v
						{
						}
					}
					params["tax"] = v
				}
				if d.HasChange("cash_balance") {
					v := d.Get("cash_balance")
					{
						path := cty.Path{}
						outerV := v
						{
							path := path.IndexString("settings")
							outerV := shared.Dig[any](outerV)
							if inEnum(shared.Dig[any](outerV), []string{"automatic", "manual"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: automatic, manual"),
								})
							}
						}
					}
					params["cash_balance"] = v
				}
				if d.HasChange("invoice_prefix") {
					v := d.Get("invoice_prefix")
					{
					}
					params["invoice_prefix"] = v
				}
				if d.HasChange("invoice_settings") {
					v := d.Get("invoice_settings")
					{
						path := cty.Path{}
						outerV := v
						{
							path := path.IndexString("custom_fields")
							outerV := shared.Dig[any](outerV)
							for i, v := range outerV.([]any) {
								{
									path := path.IndexInt(i)
									outerV := v
									{
									}
									{
									}
								}
							}
						}
						{
						}
						{
						}
						{
							path := path.IndexString("rendering_options")
							outerV := shared.Dig[any](outerV)
							if inEnum(shared.Dig[any](outerV), []string{"", "exclude_tax", "include_inclusive_tax"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: , exclude_tax, include_inclusive_tax"),
								})
							}
						}
					}
					params["invoice_settings"] = v
				}
				if d.HasChange("next_invoice_sequence") {
					v := d.Get("next_invoice_sequence")
					{
					}
					params["next_invoice_sequence"] = v
				}
				if d.HasChange("phone") {
					v := d.Get("phone")
					{
					}
					params["phone"] = v
				}
				if d.HasChange("promotion_code") {
					v := d.Get("promotion_code")
					{
					}
					params["promotion_code"] = v
				}
				if d.HasChange("tax_exempt") {
					v := d.Get("tax_exempt")
					if inEnum(v, []string{"", "exempt", "none", "reverse"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: , exempt, none, reverse"),
						})
					}
					params["tax_exempt"] = v
				}
				if d.HasChange("address") {
					v := d.Get("address")
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
						{
						}
						{
						}
						{
						}
						{
						}
					}
					params["address"] = v
				}
				if d.HasChange("coupon") {
					v := d.Get("coupon")
					{
					}
					params["coupon"] = v
				}
				if d.HasChange("email") {
					v := d.Get("email")
					{
					}
					params["email"] = v
				}
				if d.HasChange("preferred_locales") {
					v := d.Get("preferred_locales")
					{
						path := cty.Path{}
						outerV := v
						for i, v := range outerV.([]any) {
							{
							}
						}
					}
					params["preferred_locales"] = v
				}
				_, err := f.Post(fmt.Sprintf("/v1/customers/%v", d.Get("id")), params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to update customer %s", d.Id(), err),
					})
					return out
				}
				return out
			},
		},
		"ephemeral_keys": &schema.Resource{
			CreateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if v, exists := d.GetOk("issuing_card"); exists {
					{
					}
					params["issuing_card"] = v
				}
				if v, exists := d.GetOk("customer"); exists {
					{
					}
					params["customer"] = v
				}
				res, err := f.Post("/v1/ephemeral_keys", params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to create new ephemeral_key", err),
					})
					return out
				}
				d.SetId(Dig(res, "id"))
				return out
			},
			DeleteWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				err := f.Delete(fmt.Sprintf("/v1/ephemeral_keys/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to delete ephemeral_key %s", d.Id(), err),
					})
					return out
				}
				return out
			},
			ReadWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				res, err := f.Get(fmt.Sprintf("/v1/ephemeral_keys/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to look up ephemeral_key %s", d.Id(), err),
					})
					return out
				}
				d.Set("issuing_card", Dig[any](res, "issuing_card"))
				d.Set("customer", Dig[any](res, "customer"))
				return out
			},
			Schema: map[string]*schema.Schema{
				"customer": &schema.Schema{
					Description: "The ID of the Customer you'd like to modify using the resulting ephemeral key.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"issuing_card": &schema.Schema{
					Description: "The ID of the Issuing Card you'd like to access using the resulting ephemeral key.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
			},
			UpdateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				_, err := f.Post(fmt.Sprintf("/v1/ephemeral_keys/%v", d.Get("id")), params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to update ephemeral_key %s", d.Id(), err),
					})
					return out
				}
				return out
			},
		},
		"invoiceitems": &schema.Resource{
			CreateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				{
					v := d.Get("customer")
					{
					}
					params["customer"] = v
				}
				if v, exists := d.GetOk("discountable"); exists {
					{
					}
					params["discountable"] = v
				}
				if v, exists := d.GetOk("invoice"); exists {
					{
					}
					params["invoice"] = v
				}
				if v, exists := d.GetOk("period"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
					}
					params["period"] = v
				}
				if v, exists := d.GetOk("price"); exists {
					{
					}
					params["price"] = v
				}
				if v, exists := d.GetOk("unit_amount_decimal"); exists {
					{
					}
					params["unit_amount_decimal"] = v
				}
				if v, exists := d.GetOk("subscription"); exists {
					{
					}
					params["subscription"] = v
				}
				if v, exists := d.GetOk("tax_behavior"); exists {
					if inEnum(v, []string{"exclusive", "inclusive", "unspecified"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: exclusive, inclusive, unspecified"),
						})
					}
					params["tax_behavior"] = v
				}
				if v, exists := d.GetOk("amount"); exists {
					{
					}
					params["amount"] = v
				}
				if v, exists := d.GetOk("description"); exists {
					{
					}
					params["description"] = v
				}
				if v, exists := d.GetOk("discounts"); exists {
					{
						path := cty.Path{}
						outerV := v
						for i, v := range outerV.([]any) {
							{
								path := path.IndexInt(i)
								outerV := v
								{
								}
								{
								}
							}
						}
					}
					params["discounts"] = v
				}
				if v, exists := d.GetOk("quantity"); exists {
					{
					}
					params["quantity"] = v
				}
				if v, exists := d.GetOk("tax_rates"); exists {
					{
						path := cty.Path{}
						outerV := v
						for i, v := range outerV.([]any) {
							{
							}
						}
					}
					params["tax_rates"] = v
				}
				if v, exists := d.GetOk("unit_amount"); exists {
					{
					}
					params["unit_amount"] = v
				}
				if v, exists := d.GetOk("currency"); exists {
					{
					}
					params["currency"] = v
				}
				if v, exists := d.GetOk("price_data"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
						{
						}
						if inEnum(shared.Dig[any](outerV), []string{"exclusive", "inclusive", "unspecified"}) {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       fmt.Sprintf("Field must be one of: exclusive, inclusive, unspecified"),
							})
						}
						{
						}
					}
					params["price_data"] = v
				}
				if v, exists := d.GetOk("tax_code"); exists {
					{
					}
					params["tax_code"] = v
				}
				res, err := f.Post("/v1/invoiceitems", params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to create new invoiceitem", err),
					})
					return out
				}
				d.SetId(Dig(res, "id"))
				return out
			},
			DeleteWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				err := f.Delete(fmt.Sprintf("/v1/invoiceitems/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to delete invoiceitem %s", d.Id(), err),
					})
					return out
				}
				return out
			},
			ReadWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				res, err := f.Get(fmt.Sprintf("/v1/invoiceitems/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to look up invoiceitem %s", d.Id(), err),
					})
					return out
				}
				d.Set("customer", Dig[any](res, "customer"))
				d.Set("discountable", Dig[any](res, "discountable"))
				d.Set("invoice", Dig[any](res, "invoice"))
				d.Set("period", Dig[any](res, "period"))
				d.Set("price", Dig[any](res, "price"))
				d.Set("unit_amount_decimal", Dig[any](res, "unit_amount_decimal"))
				d.Set("subscription", Dig[any](res, "subscription"))
				d.Set("tax_behavior", Dig[any](res, "tax_behavior"))
				d.Set("amount", Dig[any](res, "amount"))
				d.Set("description", Dig[any](res, "description"))
				d.Set("discounts", Dig[any](res, "discounts"))
				d.Set("quantity", Dig[any](res, "quantity"))
				d.Set("tax_rates", Dig[any](res, "tax_rates"))
				d.Set("unit_amount", Dig[any](res, "unit_amount"))
				d.Set("currency", Dig[any](res, "currency"))
				d.Set("price_data", Dig[any](res, "price_data"))
				d.Set("tax_code", Dig[any](res, "tax_code"))
				return out
			},
			Schema: map[string]*schema.Schema{
				"amount": &schema.Schema{
					Description: "The integer amount in cents (or local equivalent) of the charge to be applied to the upcoming invoice. Passing in a negative `amount` will reduce the `amount_due` on the invoice.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"currency": &schema.Schema{
					Description: "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"customer": &schema.Schema{
					Description: "The ID of the customer who will be billed when this invoice item is billed.",
					ForceNew:    true,
					Required:    true,
					Type:        schema.TypeString,
				},
				"description": &schema.Schema{
					Description: "An arbitrary string which you can attach to the invoice item. The description is displayed in the invoice for easy tracking.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"discountable": &schema.Schema{
					Description: "Controls whether discounts apply to this invoice item. Defaults to false for prorations or negative invoice items, and true for all other invoice items.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeBool,
				},
				"discounts": &schema.Schema{
					Description: "The coupons to redeem into discounts for the invoice item or invoice line item.",
					Elem: &schema.Schema{
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"coupon": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"discount": &schema.Schema{
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
				"invoice": &schema.Schema{
					Description: "The ID of an existing invoice to add this invoice item to. When left blank, the invoice item will be added to the next upcoming scheduled invoice. This is useful when adding invoice items in response to an invoice.created webhook. You can only add invoice items to draft invoices and there is a maximum of 250 items per invoice.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"period": &schema.Schema{
					Description: "The period associated with this invoice item. When set to different values, the period will be rendered on the invoice.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"end": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeInt,
						},
						"start": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeInt,
						},
					}},
					ForceNew: false,
					Required: false,
					Type:     schema.TypeMap,
				},
				"price": &schema.Schema{
					Description: "The ID of the price object.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"price_data": &schema.Schema{
					Description: "Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"currency": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeString,
						},
						"product": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeString,
						},
						"tax_behavior": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"unit_amount": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeInt,
						},
						"unit_amount_decimal": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
					}},
					ForceNew: false,
					Required: false,
					Type:     schema.TypeMap,
				},
				"quantity": &schema.Schema{
					Description: "Non-negative integer. The quantity of units for the invoice item.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"subscription": &schema.Schema{
					Description: "The ID of a subscription to add this invoice item to. When left blank, the invoice item will be be added to the next upcoming scheduled invoice. When set, scheduled invoices for subscriptions other than the specified subscription will ignore the invoice item. Use this when you want to express that an invoice item has been accrued within the context of a particular subscription.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"tax_behavior": &schema.Schema{
					Description: "Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"tax_code": &schema.Schema{
					Description: "A [tax code](https://stripe.com/docs/tax/tax-categories) ID.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"tax_rates": &schema.Schema{
					Description: "The tax rates which apply to the invoice item. When set, the `default_tax_rates` on the invoice do not apply to this invoice item.",
					Elem:        &schema.Schema{Type: schema.TypeString},
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeList,
				},
				"unit_amount": &schema.Schema{
					Description: "The integer unit amount in cents (or local equivalent) of the charge to be applied to the upcoming invoice. This `unit_amount` will be multiplied by the quantity to get the full amount. Passing in a negative `unit_amount` will reduce the `amount_due` on the invoice.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"unit_amount_decimal": &schema.Schema{
					Description: "Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
			},
			UpdateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if d.HasChange("discountable") {
					v := d.Get("discountable")
					{
					}
					params["discountable"] = v
				}
				if d.HasChange("period") {
					v := d.Get("period")
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
					}
					params["period"] = v
				}
				if d.HasChange("price") {
					v := d.Get("price")
					{
					}
					params["price"] = v
				}
				if d.HasChange("unit_amount_decimal") {
					v := d.Get("unit_amount_decimal")
					{
					}
					params["unit_amount_decimal"] = v
				}
				if d.HasChange("tax_behavior") {
					v := d.Get("tax_behavior")
					if inEnum(v, []string{"exclusive", "inclusive", "unspecified"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: exclusive, inclusive, unspecified"),
						})
					}
					params["tax_behavior"] = v
				}
				if d.HasChange("amount") {
					v := d.Get("amount")
					{
					}
					params["amount"] = v
				}
				if d.HasChange("description") {
					v := d.Get("description")
					{
					}
					params["description"] = v
				}
				if d.HasChange("discounts") {
					v := d.Get("discounts")
					{
						path := cty.Path{}
						outerV := v
						for i, v := range outerV.([]any) {
							{
								path := path.IndexInt(i)
								outerV := v
								{
								}
								{
								}
							}
						}
					}
					params["discounts"] = v
				}
				if d.HasChange("quantity") {
					v := d.Get("quantity")
					{
					}
					params["quantity"] = v
				}
				if d.HasChange("tax_rates") {
					v := d.Get("tax_rates")
					{
						path := cty.Path{}
						outerV := v
						for i, v := range outerV.([]any) {
							{
							}
						}
					}
					params["tax_rates"] = v
				}
				if d.HasChange("unit_amount") {
					v := d.Get("unit_amount")
					{
					}
					params["unit_amount"] = v
				}
				if d.HasChange("price_data") {
					v := d.Get("price_data")
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
						{
						}
						if inEnum(shared.Dig[any](outerV), []string{"exclusive", "inclusive", "unspecified"}) {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       fmt.Sprintf("Field must be one of: exclusive, inclusive, unspecified"),
							})
						}
						{
						}
					}
					params["price_data"] = v
				}
				if d.HasChange("tax_code") {
					v := d.Get("tax_code")
					{
					}
					params["tax_code"] = v
				}
				_, err := f.Post(fmt.Sprintf("/v1/invoiceitems/%v", d.Get("id")), params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to update invoiceitem %s", d.Id(), err),
					})
					return out
				}
				return out
			},
		},
		"invoices": &schema.Resource{
			CreateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if v, exists := d.GetOk("account_tax_ids"); exists {
					{
						path := cty.Path{}
						outerV := v
						for i, v := range outerV.([]any) {
							{
							}
						}
					}
					params["account_tax_ids"] = v
				}
				if v, exists := d.GetOk("application_fee_amount"); exists {
					{
					}
					params["application_fee_amount"] = v
				}
				if v, exists := d.GetOk("currency"); exists {
					{
					}
					params["currency"] = v
				}
				if v, exists := d.GetOk("payment_settings"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
							path := path.IndexString("payment_method_options")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("acss_debit")
								outerV := shared.Dig[any](outerV)
								if inEnum(shared.Dig[any](outerV), []string{"automatic", "instant", "microdeposits"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: automatic, instant, microdeposits"),
									})
								}
								{
									path := path.IndexString("mandate_options")
									outerV := shared.Dig[any](outerV)
									if inEnum(shared.Dig[any](outerV), []string{"business", "personal"}) {
										out = append(out, diag.Diagnostic{
											AttributePath: path,
											Severity:      diag.Error,
											Summary:       fmt.Sprintf("Field must be one of: business, personal"),
										})
									}
								}
							}
							{
								path := path.IndexString("bancontact")
								outerV := shared.Dig[any](outerV)
								if inEnum(shared.Dig[any](outerV), []string{"de", "en", "fr", "nl"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: de, en, fr, nl"),
									})
								}
							}
							{
								path := path.IndexString("card")
								outerV := shared.Dig[any](outerV)
								{
									path := path.IndexString("installments")
									outerV := shared.Dig[any](outerV)
									{
									}
									{
										path := path.IndexString("plan")
										outerV := shared.Dig[any](outerV)
										{
										}
										if inEnum(shared.Dig[any](outerV), []string{"month"}) {
											out = append(out, diag.Diagnostic{
												AttributePath: path,
												Severity:      diag.Error,
												Summary:       fmt.Sprintf("Field must be one of: month"),
											})
										}
										if inEnum(shared.Dig[any](outerV), []string{"fixed_count"}) {
											out = append(out, diag.Diagnostic{
												AttributePath: path,
												Severity:      diag.Error,
												Summary:       fmt.Sprintf("Field must be one of: fixed_count"),
											})
										}
									}
								}
								if inEnum(shared.Dig[any](outerV), []string{"any", "automatic"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: any, automatic"),
									})
								}
							}
							{
								path := path.IndexString("customer_balance")
								outerV := shared.Dig[any](outerV)
								{
								}
								{
									path := path.IndexString("bank_transfer")
									outerV := shared.Dig[any](outerV)
									{
									}
									{
										path := path.IndexString("eu_bank_transfer")
										outerV := shared.Dig[any](outerV)
										{
										}
									}
								}
							}
							{
								path := path.IndexString("konbini")
								outerV := shared.Dig[any](outerV)
							}
							{
								path := path.IndexString("us_bank_account")
								outerV := shared.Dig[any](outerV)
								{
									path := path.IndexString("financial_connections")
									outerV := shared.Dig[any](outerV)
									{
										path := path.IndexString("permissions")
										outerV := shared.Dig[any](outerV)
										for i, v := range outerV.([]any) {
											if inEnum(v, []string{"balances", "ownership", "payment_method", "transactions"}) {
												out = append(out, diag.Diagnostic{
													AttributePath: path,
													Severity:      diag.Error,
													Summary:       fmt.Sprintf("Field must be one of: balances, ownership, payment_method, transactions"),
												})
											}
										}
									}
								}
								if inEnum(shared.Dig[any](outerV), []string{"automatic", "instant", "microdeposits"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: automatic, instant, microdeposits"),
									})
								}
							}
						}
						{
							path := path.IndexString("payment_method_types")
							outerV := shared.Dig[any](outerV)
							for i, v := range outerV.([]any) {
								if inEnum(v, []string{"ach_credit_transfer", "ach_debit", "acss_debit", "au_becs_debit", "bacs_debit", "bancontact", "boleto", "card", "customer_balance", "fpx", "giropay", "grabpay", "ideal", "konbini", "link", "paynow", "promptpay", "sepa_debit", "sofort", "us_bank_account", "wechat_pay"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: ach_credit_transfer, ach_debit, acss_debit, au_becs_debit, bacs_debit, bancontact, boleto, card, customer_balance, fpx, giropay, grabpay, ideal, konbini, link, paynow, promptpay, sepa_debit, sofort, us_bank_account, wechat_pay"),
									})
								}
							}
						}
					}
					params["payment_settings"] = v
				}
				if v, exists := d.GetOk("statement_descriptor"); exists {
					{
					}
					params["statement_descriptor"] = v
				}
				if v, exists := d.GetOk("subscription"); exists {
					{
					}
					params["subscription"] = v
				}
				if v, exists := d.GetOk("transfer_data"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
					}
					params["transfer_data"] = v
				}
				if v, exists := d.GetOk("automatic_tax"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
						}
					}
					params["automatic_tax"] = v
				}
				if v, exists := d.GetOk("default_source"); exists {
					{
					}
					params["default_source"] = v
				}
				if v, exists := d.GetOk("from_invoice"); exists {
					{
						path := cty.Path{}
						outerV := v
						if inEnum(shared.Dig[any](outerV), []string{"revision"}) {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       fmt.Sprintf("Field must be one of: revision"),
							})
						}
						{
						}
					}
					params["from_invoice"] = v
				}
				if v, exists := d.GetOk("pending_invoice_items_behavior"); exists {
					if inEnum(v, []string{"exclude", "include", "include_and_require"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: exclude, include, include_and_require"),
						})
					}
					params["pending_invoice_items_behavior"] = v
				}
				if v, exists := d.GetOk("rendering_options"); exists {
					{
						path := cty.Path{}
						outerV := v
						if inEnum(shared.Dig[any](outerV), []string{"", "exclude_tax", "include_inclusive_tax"}) {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       fmt.Sprintf("Field must be one of: , exclude_tax, include_inclusive_tax"),
							})
						}
					}
					params["rendering_options"] = v
				}
				if v, exists := d.GetOk("footer"); exists {
					{
					}
					params["footer"] = v
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
								}
								{
								}
							}
						}
					}
					params["custom_fields"] = v
				}
				if v, exists := d.GetOk("days_until_due"); exists {
					{
					}
					params["days_until_due"] = v
				}
				if v, exists := d.GetOk("default_payment_method"); exists {
					{
					}
					params["default_payment_method"] = v
				}
				if v, exists := d.GetOk("description"); exists {
					{
					}
					params["description"] = v
				}
				if v, exists := d.GetOk("due_date"); exists {
					{
					}
					params["due_date"] = v
				}
				if v, exists := d.GetOk("on_behalf_of"); exists {
					{
					}
					params["on_behalf_of"] = v
				}
				if v, exists := d.GetOk("auto_advance"); exists {
					{
					}
					params["auto_advance"] = v
				}
				if v, exists := d.GetOk("collection_method"); exists {
					if inEnum(v, []string{"charge_automatically", "send_invoice"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: charge_automatically, send_invoice"),
						})
					}
					params["collection_method"] = v
				}
				if v, exists := d.GetOk("customer"); exists {
					{
					}
					params["customer"] = v
				}
				if v, exists := d.GetOk("default_tax_rates"); exists {
					{
						path := cty.Path{}
						outerV := v
						for i, v := range outerV.([]any) {
							{
							}
						}
					}
					params["default_tax_rates"] = v
				}
				if v, exists := d.GetOk("discounts"); exists {
					{
						path := cty.Path{}
						outerV := v
						for i, v := range outerV.([]any) {
							{
								path := path.IndexInt(i)
								outerV := v
								{
								}
								{
								}
							}
						}
					}
					params["discounts"] = v
				}
				res, err := f.Post("/v1/invoices", params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to create new invoice", err),
					})
					return out
				}
				d.SetId(Dig(res, "id"))
				return out
			},
			DeleteWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				err := f.Delete(fmt.Sprintf("/v1/invoices/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to delete invoice %s", d.Id(), err),
					})
					return out
				}
				return out
			},
			ReadWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				res, err := f.Get(fmt.Sprintf("/v1/invoices/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to look up invoice %s", d.Id(), err),
					})
					return out
				}
				d.Set("account_tax_ids", Dig[any](res, "account_tax_ids"))
				d.Set("application_fee_amount", Dig[any](res, "application_fee_amount"))
				d.Set("currency", Dig[any](res, "currency"))
				d.Set("payment_settings", Dig[any](res, "payment_settings"))
				d.Set("statement_descriptor", Dig[any](res, "statement_descriptor"))
				d.Set("subscription", Dig[any](res, "subscription"))
				d.Set("transfer_data", Dig[any](res, "transfer_data"))
				d.Set("automatic_tax", Dig[any](res, "automatic_tax"))
				d.Set("default_source", Dig[any](res, "default_source"))
				d.Set("from_invoice", Dig[any](res, "from_invoice"))
				d.Set("pending_invoice_items_behavior", Dig[any](res, "pending_invoice_items_behavior"))
				d.Set("rendering_options", Dig[any](res, "rendering_options"))
				d.Set("footer", Dig[any](res, "footer"))
				d.Set("custom_fields", Dig[any](res, "custom_fields"))
				d.Set("days_until_due", Dig[any](res, "days_until_due"))
				d.Set("default_payment_method", Dig[any](res, "default_payment_method"))
				d.Set("description", Dig[any](res, "description"))
				d.Set("due_date", Dig[any](res, "due_date"))
				d.Set("on_behalf_of", Dig[any](res, "on_behalf_of"))
				d.Set("auto_advance", Dig[any](res, "auto_advance"))
				d.Set("collection_method", Dig[any](res, "collection_method"))
				d.Set("customer", Dig[any](res, "customer"))
				d.Set("default_tax_rates", Dig[any](res, "default_tax_rates"))
				d.Set("discounts", Dig[any](res, "discounts"))
				return out
			},
			Schema: map[string]*schema.Schema{
				"account_tax_ids": &schema.Schema{
					Description: "The account tax IDs associated with the invoice. Only editable when the invoice is a draft.",
					Elem:        &schema.Schema{Type: schema.TypeString},
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeList,
				},
				"application_fee_amount": &schema.Schema{
					Description: "A fee in cents (or local equivalent) that will be applied to the invoice and transferred to the application owner's Stripe account. The request must be made with an OAuth key or the Stripe-Account header in order to take an application fee. For more information, see the application fees [documentation](https://stripe.com/docs/billing/invoices/connect#collecting-fees).",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"auto_advance": &schema.Schema{
					Description: "Controls whether Stripe will perform [automatic collection](https://stripe.com/docs/billing/invoices/workflow/#auto_advance) of the invoice. When `false`, the invoice's state will not automatically advance without an explicit action.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeBool,
				},
				"automatic_tax": &schema.Schema{
					Description: "Settings for automatic tax lookup for this invoice.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{"enabled": &schema.Schema{
						Description: "",
						Required:    true,
						Type:        schema.TypeBool,
					}}},
					ForceNew: false,
					Required: false,
					Type:     schema.TypeMap,
				},
				"collection_method": &schema.Schema{
					Description: "Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay this invoice using the default source attached to the customer. When sending an invoice, Stripe will email this invoice to the customer with payment instructions. Defaults to `charge_automatically`.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"currency": &schema.Schema{
					Description: "The currency to create this invoice in. Defaults to that of `customer` if not specified.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"custom_fields": &schema.Schema{
					Description: "A list of up to 4 custom fields to be displayed on the invoice.",
					Elem: &schema.Schema{
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"name": &schema.Schema{
								Description: "",
								Required:    true,
								Type:        schema.TypeString,
							},
							"value": &schema.Schema{
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
				"customer": &schema.Schema{
					Description: "The ID of the customer who will be billed.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"days_until_due": &schema.Schema{
					Description: "The number of days from when the invoice is created until it is due. Valid only for invoices where `collection_method=send_invoice`.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"default_payment_method": &schema.Schema{
					Description: "ID of the default payment method for the invoice. It must belong to the customer associated with the invoice. If not set, defaults to the subscription's default payment method, if any, or to the default payment method in the customer's invoice settings.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"default_source": &schema.Schema{
					Description: "ID of the default payment source for the invoice. It must belong to the customer associated with the invoice and be in a chargeable state. If not set, defaults to the subscription's default source, if any, or to the customer's default source.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"default_tax_rates": &schema.Schema{
					Description: "The tax rates that will apply to any line item that does not have `tax_rates` set.",
					Elem:        &schema.Schema{Type: schema.TypeString},
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeList,
				},
				"description": &schema.Schema{
					Description: "An arbitrary string attached to the object. Often useful for displaying to users. Referenced as 'memo' in the Dashboard.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"discounts": &schema.Schema{
					Description: "The coupons to redeem into discounts for the invoice. If not specified, inherits the discount from the invoice's customer. Pass an empty string to avoid inheriting any discounts.",
					Elem: &schema.Schema{
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"coupon": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"discount": &schema.Schema{
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
				"due_date": &schema.Schema{
					Description: "The date on which payment for this invoice is due. Valid only for invoices where `collection_method=send_invoice`.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"footer": &schema.Schema{
					Description: "Footer to be displayed on the invoice.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"from_invoice": &schema.Schema{
					Description: "Revise an existing invoice. The new invoice will be created in `status=draft`. See the [revision documentation](https://stripe.com/docs/invoicing/invoice-revisions) for more details.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeString,
						},
						"invoice": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeString,
						},
					}},
					ForceNew: true,
					Required: false,
					Type:     schema.TypeMap,
				},
				"on_behalf_of": &schema.Schema{
					Description: "The account (if any) for which the funds of the invoice payment are intended. If set, the invoice will be presented with the branding and support information of the specified account. See the [Invoices with Connect](https://stripe.com/docs/billing/invoices/connect) documentation for details.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"payment_settings": &schema.Schema{
					Description: "Configuration settings for the PaymentIntent that is generated when the invoice is finalized.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"default_mandate": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"payment_method_options": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"acss_debit": &schema.Schema{
									Description: "",
									Elem: &schema.Resource{Schema: map[string]*schema.Schema{
										"mandate_options": &schema.Schema{
											Description: "",
											Elem: &schema.Resource{Schema: map[string]*schema.Schema{"transaction_type": &schema.Schema{
												Description: "",
												Required:    false,
												Type:        schema.TypeString,
											}}},
											Required: false,
											Type:     schema.TypeMap,
										},
										"verification_method": &schema.Schema{
											Description: "",
											Required:    false,
											Type:        schema.TypeString,
										},
									}},
									Required: false,
									Type:     schema.TypeMap,
								},
								"bancontact": &schema.Schema{
									Description: "",
									Elem: &schema.Resource{Schema: map[string]*schema.Schema{"preferred_language": &schema.Schema{
										Description: "",
										Required:    false,
										Type:        schema.TypeString,
									}}},
									Required: false,
									Type:     schema.TypeMap,
								},
								"card": &schema.Schema{
									Description: "",
									Elem: &schema.Resource{Schema: map[string]*schema.Schema{
										"installments": &schema.Schema{
											Description: "",
											Elem: &schema.Resource{Schema: map[string]*schema.Schema{
												"enabled": &schema.Schema{
													Description: "",
													Required:    false,
													Type:        schema.TypeBool,
												},
												"plan": &schema.Schema{
													Description: "",
													Elem: &schema.Resource{Schema: map[string]*schema.Schema{
														"count": &schema.Schema{
															Description: "",
															Required:    true,
															Type:        schema.TypeInt,
														},
														"interval": &schema.Schema{
															Description: "",
															Required:    true,
															Type:        schema.TypeString,
														},
														"type": &schema.Schema{
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
										"request_three_d_secure": &schema.Schema{
											Description: "",
											Required:    false,
											Type:        schema.TypeString,
										},
									}},
									Required: false,
									Type:     schema.TypeMap,
								},
								"customer_balance": &schema.Schema{
									Description: "",
									Elem: &schema.Resource{Schema: map[string]*schema.Schema{
										"bank_transfer": &schema.Schema{
											Description: "",
											Elem: &schema.Resource{Schema: map[string]*schema.Schema{
												"eu_bank_transfer": &schema.Schema{
													Description: "",
													Elem: &schema.Resource{Schema: map[string]*schema.Schema{"country": &schema.Schema{
														Description: "",
														Required:    true,
														Type:        schema.TypeString,
													}}},
													Required: false,
													Type:     schema.TypeMap,
												},
												"type": &schema.Schema{
													Description: "",
													Required:    false,
													Type:        schema.TypeString,
												},
											}},
											Required: false,
											Type:     schema.TypeMap,
										},
										"funding_type": &schema.Schema{
											Description: "",
											Required:    false,
											Type:        schema.TypeString,
										},
									}},
									Required: false,
									Type:     schema.TypeMap,
								},
								"konbini": &schema.Schema{
									Description: "",
									Elem:        &schema.Resource{Schema: map[string]*schema.Schema{}},
									Required:    false,
									Type:        schema.TypeMap,
								},
								"us_bank_account": &schema.Schema{
									Description: "",
									Elem: &schema.Resource{Schema: map[string]*schema.Schema{
										"financial_connections": &schema.Schema{
											Description: "",
											Elem: &schema.Resource{Schema: map[string]*schema.Schema{"permissions": &schema.Schema{
												Description: "",
												Elem:        &schema.Schema{Type: schema.TypeString},
												Required:    false,
												Type:        schema.TypeList,
											}}},
											Required: false,
											Type:     schema.TypeMap,
										},
										"verification_method": &schema.Schema{
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
						"payment_method_types": &schema.Schema{
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
				"pending_invoice_items_behavior": &schema.Schema{
					Description: "How to handle pending invoice items on invoice creation. One of `include` or `exclude`. `include` will include any pending invoice items, and will create an empty draft invoice if no pending invoice items exist. `exclude` will always create an empty invoice draft regardless if there are pending invoice items or not. Defaults to `exclude` if the parameter is omitted.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"rendering_options": &schema.Schema{
					Description: "Options for invoice PDF rendering.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{"amount_tax_display": &schema.Schema{
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					}}},
					ForceNew: false,
					Required: false,
					Type:     schema.TypeMap,
				},
				"statement_descriptor": &schema.Schema{
					Description: "Extra information about a charge for the customer's credit card statement. It must contain at least one letter. If not specified and this invoice is part of a subscription, the default `statement_descriptor` will be set to the first subscription item's product's `statement_descriptor`.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"subscription": &schema.Schema{
					Description: "The ID of the subscription to invoice, if any. If set, the created invoice will only include pending invoice items for that subscription and pending invoice items not associated with any subscription if `pending_invoice_items_behavior` is `include`. The subscription's billing cycle and regular subscription events won't be affected.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"transfer_data": &schema.Schema{
					Description: "If specified, the funds from the invoice will be transferred to the destination and the ID of the resulting transfer will be found on the invoice's charge.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"amount": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeInt,
						},
						"destination": &schema.Schema{
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
			UpdateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if d.HasChange("account_tax_ids") {
					v := d.Get("account_tax_ids")
					{
						path := cty.Path{}
						outerV := v
						for i, v := range outerV.([]any) {
							{
							}
						}
					}
					params["account_tax_ids"] = v
				}
				if d.HasChange("application_fee_amount") {
					v := d.Get("application_fee_amount")
					{
					}
					params["application_fee_amount"] = v
				}
				if d.HasChange("payment_settings") {
					v := d.Get("payment_settings")
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
							path := path.IndexString("payment_method_options")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("acss_debit")
								outerV := shared.Dig[any](outerV)
								if inEnum(shared.Dig[any](outerV), []string{"automatic", "instant", "microdeposits"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: automatic, instant, microdeposits"),
									})
								}
								{
									path := path.IndexString("mandate_options")
									outerV := shared.Dig[any](outerV)
									if inEnum(shared.Dig[any](outerV), []string{"business", "personal"}) {
										out = append(out, diag.Diagnostic{
											AttributePath: path,
											Severity:      diag.Error,
											Summary:       fmt.Sprintf("Field must be one of: business, personal"),
										})
									}
								}
							}
							{
								path := path.IndexString("bancontact")
								outerV := shared.Dig[any](outerV)
								if inEnum(shared.Dig[any](outerV), []string{"de", "en", "fr", "nl"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: de, en, fr, nl"),
									})
								}
							}
							{
								path := path.IndexString("card")
								outerV := shared.Dig[any](outerV)
								{
									path := path.IndexString("installments")
									outerV := shared.Dig[any](outerV)
									{
									}
									{
										path := path.IndexString("plan")
										outerV := shared.Dig[any](outerV)
										{
										}
										if inEnum(shared.Dig[any](outerV), []string{"month"}) {
											out = append(out, diag.Diagnostic{
												AttributePath: path,
												Severity:      diag.Error,
												Summary:       fmt.Sprintf("Field must be one of: month"),
											})
										}
										if inEnum(shared.Dig[any](outerV), []string{"fixed_count"}) {
											out = append(out, diag.Diagnostic{
												AttributePath: path,
												Severity:      diag.Error,
												Summary:       fmt.Sprintf("Field must be one of: fixed_count"),
											})
										}
									}
								}
								if inEnum(shared.Dig[any](outerV), []string{"any", "automatic"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: any, automatic"),
									})
								}
							}
							{
								path := path.IndexString("customer_balance")
								outerV := shared.Dig[any](outerV)
								{
								}
								{
									path := path.IndexString("bank_transfer")
									outerV := shared.Dig[any](outerV)
									{
									}
									{
										path := path.IndexString("eu_bank_transfer")
										outerV := shared.Dig[any](outerV)
										{
										}
									}
								}
							}
							{
								path := path.IndexString("konbini")
								outerV := shared.Dig[any](outerV)
							}
							{
								path := path.IndexString("us_bank_account")
								outerV := shared.Dig[any](outerV)
								{
									path := path.IndexString("financial_connections")
									outerV := shared.Dig[any](outerV)
									{
										path := path.IndexString("permissions")
										outerV := shared.Dig[any](outerV)
										for i, v := range outerV.([]any) {
											if inEnum(v, []string{"balances", "ownership", "payment_method", "transactions"}) {
												out = append(out, diag.Diagnostic{
													AttributePath: path,
													Severity:      diag.Error,
													Summary:       fmt.Sprintf("Field must be one of: balances, ownership, payment_method, transactions"),
												})
											}
										}
									}
								}
								if inEnum(shared.Dig[any](outerV), []string{"automatic", "instant", "microdeposits"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: automatic, instant, microdeposits"),
									})
								}
							}
						}
						{
							path := path.IndexString("payment_method_types")
							outerV := shared.Dig[any](outerV)
							for i, v := range outerV.([]any) {
								if inEnum(v, []string{"ach_credit_transfer", "ach_debit", "acss_debit", "au_becs_debit", "bacs_debit", "bancontact", "boleto", "card", "customer_balance", "fpx", "giropay", "grabpay", "ideal", "konbini", "link", "paynow", "promptpay", "sepa_debit", "sofort", "us_bank_account", "wechat_pay"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: ach_credit_transfer, ach_debit, acss_debit, au_becs_debit, bacs_debit, bancontact, boleto, card, customer_balance, fpx, giropay, grabpay, ideal, konbini, link, paynow, promptpay, sepa_debit, sofort, us_bank_account, wechat_pay"),
									})
								}
							}
						}
					}
					params["payment_settings"] = v
				}
				if d.HasChange("statement_descriptor") {
					v := d.Get("statement_descriptor")
					{
					}
					params["statement_descriptor"] = v
				}
				if d.HasChange("transfer_data") {
					v := d.Get("transfer_data")
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
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
						}
					}
					params["automatic_tax"] = v
				}
				if d.HasChange("default_source") {
					v := d.Get("default_source")
					{
					}
					params["default_source"] = v
				}
				if d.HasChange("rendering_options") {
					v := d.Get("rendering_options")
					{
						path := cty.Path{}
						outerV := v
						if inEnum(shared.Dig[any](outerV), []string{"", "exclude_tax", "include_inclusive_tax"}) {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       fmt.Sprintf("Field must be one of: , exclude_tax, include_inclusive_tax"),
							})
						}
					}
					params["rendering_options"] = v
				}
				if d.HasChange("footer") {
					v := d.Get("footer")
					{
					}
					params["footer"] = v
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
								}
								{
								}
							}
						}
					}
					params["custom_fields"] = v
				}
				if d.HasChange("days_until_due") {
					v := d.Get("days_until_due")
					{
					}
					params["days_until_due"] = v
				}
				if d.HasChange("default_payment_method") {
					v := d.Get("default_payment_method")
					{
					}
					params["default_payment_method"] = v
				}
				if d.HasChange("description") {
					v := d.Get("description")
					{
					}
					params["description"] = v
				}
				if d.HasChange("due_date") {
					v := d.Get("due_date")
					{
					}
					params["due_date"] = v
				}
				if d.HasChange("on_behalf_of") {
					v := d.Get("on_behalf_of")
					{
					}
					params["on_behalf_of"] = v
				}
				if d.HasChange("auto_advance") {
					v := d.Get("auto_advance")
					{
					}
					params["auto_advance"] = v
				}
				if d.HasChange("collection_method") {
					v := d.Get("collection_method")
					if inEnum(v, []string{"charge_automatically", "send_invoice"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: charge_automatically, send_invoice"),
						})
					}
					params["collection_method"] = v
				}
				if d.HasChange("default_tax_rates") {
					v := d.Get("default_tax_rates")
					{
						path := cty.Path{}
						outerV := v
						for i, v := range outerV.([]any) {
							{
							}
						}
					}
					params["default_tax_rates"] = v
				}
				if d.HasChange("discounts") {
					v := d.Get("discounts")
					{
						path := cty.Path{}
						outerV := v
						for i, v := range outerV.([]any) {
							{
								path := path.IndexInt(i)
								outerV := v
								{
								}
								{
								}
							}
						}
					}
					params["discounts"] = v
				}
				_, err := f.Post(fmt.Sprintf("/v1/invoices/%v", d.Get("id")), params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to update invoice %s", d.Id(), err),
					})
					return out
				}
				return out
			},
		},
		"plans": &schema.Resource{
			CreateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				{
					v := d.Get("currency")
					{
					}
					params["currency"] = v
				}
				if v, exists := d.GetOk("nickname"); exists {
					{
					}
					params["nickname"] = v
				}
				if v, exists := d.GetOk("transform_usage"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
						}
						if inEnum(shared.Dig[any](outerV), []string{"down", "up"}) {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       fmt.Sprintf("Field must be one of: down, up"),
							})
						}
					}
					params["transform_usage"] = v
				}
				if v, exists := d.GetOk("trial_period_days"); exists {
					{
					}
					params["trial_period_days"] = v
				}
				if v, exists := d.GetOk("amount"); exists {
					{
					}
					params["amount"] = v
				}
				if v, exists := d.GetOk("billing_scheme"); exists {
					if inEnum(v, []string{"per_unit", "tiered"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: per_unit, tiered"),
						})
					}
					params["billing_scheme"] = v
				}
				if v, exists := d.GetOk("usage_type"); exists {
					if inEnum(v, []string{"licensed", "metered"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: licensed, metered"),
						})
					}
					params["usage_type"] = v
				}
				if v, exists := d.GetOk("interval_count"); exists {
					{
					}
					params["interval_count"] = v
				}
				{
					v := d.Get("interval")
					if inEnum(v, []string{"day", "month", "week", "year"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: day, month, week, year"),
						})
					}
					params["interval"] = v
				}
				if v, exists := d.GetOk("product"); exists {
					{
					}
					params["product"] = v
				}
				if v, exists := d.GetOk("tiers_mode"); exists {
					if inEnum(v, []string{"graduated", "volume"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: graduated, volume"),
						})
					}
					params["tiers_mode"] = v
				}
				if v, exists := d.GetOk("active"); exists {
					{
					}
					params["active"] = v
				}
				if v, exists := d.GetOk("id"); exists {
					{
					}
					params["id"] = v
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
								}
								{
								}
								{
								}
								{
								}
								{
								}
							}
						}
					}
					params["tiers"] = v
				}
				if v, exists := d.GetOk("aggregate_usage"); exists {
					if inEnum(v, []string{"last_during_period", "last_ever", "max", "sum"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: last_during_period, last_ever, max, sum"),
						})
					}
					params["aggregate_usage"] = v
				}
				if v, exists := d.GetOk("amount_decimal"); exists {
					{
					}
					params["amount_decimal"] = v
				}
				res, err := f.Post("/v1/plans", params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to create new plan", err),
					})
					return out
				}
				d.SetId(Dig(res, "id"))
				return out
			},
			DeleteWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				err := f.Delete(fmt.Sprintf("/v1/plans/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to delete plan %s", d.Id(), err),
					})
					return out
				}
				return out
			},
			ReadWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				res, err := f.Get(fmt.Sprintf("/v1/plans/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to look up plan %s", d.Id(), err),
					})
					return out
				}
				d.Set("currency", Dig[any](res, "currency"))
				d.Set("nickname", Dig[any](res, "nickname"))
				d.Set("transform_usage", Dig[any](res, "transform_usage"))
				d.Set("trial_period_days", Dig[any](res, "trial_period_days"))
				d.Set("amount", Dig[any](res, "amount"))
				d.Set("billing_scheme", Dig[any](res, "billing_scheme"))
				d.Set("usage_type", Dig[any](res, "usage_type"))
				d.Set("interval_count", Dig[any](res, "interval_count"))
				d.Set("interval", Dig[any](res, "interval"))
				d.Set("product", Dig[any](res, "product"))
				d.Set("tiers_mode", Dig[any](res, "tiers_mode"))
				d.Set("active", Dig[any](res, "active"))
				d.Set("id", Dig[any](res, "id"))
				d.Set("tiers", Dig[any](res, "tiers"))
				d.Set("aggregate_usage", Dig[any](res, "aggregate_usage"))
				d.Set("amount_decimal", Dig[any](res, "amount_decimal"))
				return out
			},
			Schema: map[string]*schema.Schema{
				"active": &schema.Schema{
					Description: "Whether the plan is currently available for new subscriptions. Defaults to `true`.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeBool,
				},
				"aggregate_usage": &schema.Schema{
					Description: "Specifies a usage aggregation strategy for plans of `usage_type=metered`. Allowed values are `sum` for summing up all usage during a period, `last_during_period` for using the last usage record reported within a period, `last_ever` for using the last usage record ever (across period bounds) or `max` which uses the usage record with the maximum reported usage during a period. Defaults to `sum`.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"amount": &schema.Schema{
					Description: "A positive integer in cents (or local equivalent) (or 0 for a free plan) representing how much to charge on a recurring basis.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"amount_decimal": &schema.Schema{
					Description: "Same as `amount`, but accepts a decimal value with at most 12 decimal places. Only one of `amount` and `amount_decimal` can be set.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"billing_scheme": &schema.Schema{
					Description: "Describes how to compute the price per period. Either `per_unit` or `tiered`. `per_unit` indicates that the fixed amount (specified in `amount`) will be charged per unit in `quantity` (for plans with `usage_type=licensed`), or per unit of total usage (for plans with `usage_type=metered`). `tiered` indicates that the unit pricing will be computed using a tiering strategy as defined using the `tiers` and `tiers_mode` attributes.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"currency": &schema.Schema{
					Description: "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
					ForceNew:    true,
					Required:    true,
					Type:        schema.TypeString,
				},
				"id": &schema.Schema{
					Description: "An identifier randomly generated by Stripe. Used to identify this plan when subscribing a customer. You can optionally override this ID, but the ID must be unique across all plans in your Stripe account. You can, however, use the same plan ID in both live and test modes.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"interval": &schema.Schema{
					Description: "Specifies billing frequency. Either `day`, `week`, `month` or `year`.",
					ForceNew:    true,
					Required:    true,
					Type:        schema.TypeString,
				},
				"interval_count": &schema.Schema{
					Description: "The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of one year interval allowed (1 year, 12 months, or 52 weeks).",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"nickname": &schema.Schema{
					Description: "A brief description of the plan, hidden from customers.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"product": &schema.Schema{
					Description: "",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"tiers": &schema.Schema{
					Description: "Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.",
					Elem: &schema.Schema{
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"flat_amount": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeInt,
							},
							"flat_amount_decimal": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"unit_amount": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeInt,
							},
							"unit_amount_decimal": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"up_to": &schema.Schema{
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
				"tiers_mode": &schema.Schema{
					Description: "Defines if the tiering price should be `graduated` or `volume` based. In `volume`-based tiering, the maximum quantity within a period determines the per unit price, in `graduated` tiering pricing can successively change as the quantity grows.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"transform_usage": &schema.Schema{
					Description: "Apply a transformation to the reported usage or set quantity before computing the billed price. Cannot be combined with `tiers`.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"divide_by": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeInt,
						},
						"round": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeString,
						},
					}},
					ForceNew: true,
					Required: false,
					Type:     schema.TypeMap,
				},
				"trial_period_days": &schema.Schema{
					Description: "Default number of trial days when subscribing a customer to this plan using [`trial_from_plan=true`](https://stripe.com/docs/api#create_subscription-trial_from_plan).",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"usage_type": &schema.Schema{
					Description: "Configures how the quantity per period should be determined. Can be either `metered` or `licensed`. `licensed` automatically bills the `quantity` set when adding it to a subscription. `metered` aggregates the total usage based on usage records. Defaults to `licensed`.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
			},
			UpdateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if d.HasChange("nickname") {
					v := d.Get("nickname")
					{
					}
					params["nickname"] = v
				}
				if d.HasChange("trial_period_days") {
					v := d.Get("trial_period_days")
					{
					}
					params["trial_period_days"] = v
				}
				if d.HasChange("product") {
					v := d.Get("product")
					{
					}
					params["product"] = v
				}
				if d.HasChange("active") {
					v := d.Get("active")
					{
					}
					params["active"] = v
				}
				_, err := f.Post(fmt.Sprintf("/v1/plans/%v", d.Get("id")), params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to update plan %s", d.Id(), err),
					})
					return out
				}
				return out
			},
		},
		"prices": &schema.Resource{
			CreateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if v, exists := d.GetOk("tiers"); exists {
					{
						path := cty.Path{}
						outerV := v
						for i, v := range outerV.([]any) {
							{
								path := path.IndexInt(i)
								outerV := v
								{
								}
								{
								}
								{
								}
								{
								}
								{
								}
							}
						}
					}
					params["tiers"] = v
				}
				if v, exists := d.GetOk("active"); exists {
					{
					}
					params["active"] = v
				}
				if v, exists := d.GetOk("billing_scheme"); exists {
					if inEnum(v, []string{"per_unit", "tiered"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: per_unit, tiered"),
						})
					}
					params["billing_scheme"] = v
				}
				if v, exists := d.GetOk("product"); exists {
					{
					}
					params["product"] = v
				}
				if v, exists := d.GetOk("tax_behavior"); exists {
					if inEnum(v, []string{"exclusive", "inclusive", "unspecified"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: exclusive, inclusive, unspecified"),
						})
					}
					params["tax_behavior"] = v
				}
				if v, exists := d.GetOk("transfer_lookup_key"); exists {
					{
					}
					params["transfer_lookup_key"] = v
				}
				if v, exists := d.GetOk("unit_amount"); exists {
					{
					}
					params["unit_amount"] = v
				}
				if v, exists := d.GetOk("nickname"); exists {
					{
					}
					params["nickname"] = v
				}
				if v, exists := d.GetOk("product_data"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
						{
						}
						{
						}
						{
						}
						{
						}
						{
							path := path.IndexString("metadata")
							outerV := shared.Dig[any](outerV)
						}
					}
					params["product_data"] = v
				}
				if v, exists := d.GetOk("recurring"); exists {
					{
						path := cty.Path{}
						outerV := v
						if inEnum(shared.Dig[any](outerV), []string{"last_during_period", "last_ever", "max", "sum"}) {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       fmt.Sprintf("Field must be one of: last_during_period, last_ever, max, sum"),
							})
						}
						if inEnum(shared.Dig[any](outerV), []string{"day", "month", "week", "year"}) {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       fmt.Sprintf("Field must be one of: day, month, week, year"),
							})
						}
						{
						}
						if inEnum(shared.Dig[any](outerV), []string{"licensed", "metered"}) {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       fmt.Sprintf("Field must be one of: licensed, metered"),
							})
						}
					}
					params["recurring"] = v
				}
				if v, exists := d.GetOk("tiers_mode"); exists {
					if inEnum(v, []string{"graduated", "volume"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
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
						}
						if inEnum(shared.Dig[any](outerV), []string{"down", "up"}) {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       fmt.Sprintf("Field must be one of: down, up"),
							})
						}
					}
					params["transform_quantity"] = v
				}
				if v, exists := d.GetOk("currency_options"); exists {
					{
						path := cty.Path{}
						outerV := v
					}
					params["currency_options"] = v
				}
				if v, exists := d.GetOk("custom_unit_amount"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
						{
						}
						{
						}
					}
					params["custom_unit_amount"] = v
				}
				if v, exists := d.GetOk("lookup_key"); exists {
					{
					}
					params["lookup_key"] = v
				}
				if v, exists := d.GetOk("unit_amount_decimal"); exists {
					{
					}
					params["unit_amount_decimal"] = v
				}
				{
					v := d.Get("currency")
					{
					}
					params["currency"] = v
				}
				res, err := f.Post("/v1/prices", params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to create new price", err),
					})
					return out
				}
				d.SetId(Dig(res, "id"))
				return out
			},
			DeleteWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				_, err := f.Post(fmt.Sprintf("/v1/prices/%v", d.Get("id")), map[string]any{"active": false})
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to set active to false price %s", d.Id(), err),
					})
					return out
				}
				return out
			},
			ReadWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				res, err := f.Get(fmt.Sprintf("/v1/prices/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to look up price %s", d.Id(), err),
					})
					return out
				}
				d.Set("tiers", Dig[any](res, "tiers"))
				d.Set("active", Dig[any](res, "active"))
				d.Set("billing_scheme", Dig[any](res, "billing_scheme"))
				d.Set("product", Dig[any](res, "product"))
				d.Set("tax_behavior", Dig[any](res, "tax_behavior"))
				d.Set("transfer_lookup_key", Dig[any](res, "transfer_lookup_key"))
				d.Set("unit_amount", Dig[any](res, "unit_amount"))
				d.Set("nickname", Dig[any](res, "nickname"))
				d.Set("product_data", Dig[any](res, "product_data"))
				d.Set("recurring", Dig[any](res, "recurring"))
				d.Set("tiers_mode", Dig[any](res, "tiers_mode"))
				d.Set("transform_quantity", Dig[any](res, "transform_quantity"))
				d.Set("currency_options", Dig[any](res, "currency_options"))
				d.Set("custom_unit_amount", Dig[any](res, "custom_unit_amount"))
				d.Set("lookup_key", Dig[any](res, "lookup_key"))
				d.Set("unit_amount_decimal", Dig[any](res, "unit_amount_decimal"))
				d.Set("currency", Dig[any](res, "currency"))
				return out
			},
			Schema: map[string]*schema.Schema{
				"active": &schema.Schema{
					Description: "Whether the price can be used for new purchases. Defaults to `true`.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeBool,
				},
				"billing_scheme": &schema.Schema{
					Description: "Describes how to compute the price per period. Either `per_unit` or `tiered`. `per_unit` indicates that the fixed amount (specified in `unit_amount` or `unit_amount_decimal`) will be charged per unit in `quantity` (for prices with `usage_type=licensed`), or per unit of total usage (for prices with `usage_type=metered`). `tiered` indicates that the unit pricing will be computed using a tiering strategy as defined using the `tiers` and `tiers_mode` attributes.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"currency": &schema.Schema{
					Description: "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
					ForceNew:    true,
					Required:    true,
					Type:        schema.TypeString,
				},
				"currency_options": &schema.Schema{
					Description: "Prices defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).",
					Elem:        &schema.Resource{Schema: map[string]*schema.Schema{}},
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeMap,
				},
				"custom_unit_amount": &schema.Schema{
					Description: "When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"enabled": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeBool,
						},
						"maximum": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeInt,
						},
						"minimum": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeInt,
						},
						"preset": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeInt,
						},
					}},
					ForceNew: true,
					Required: false,
					Type:     schema.TypeMap,
				},
				"lookup_key": &schema.Schema{
					Description: "A lookup key used to retrieve prices dynamically from a static string. This may be up to 200 characters.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"nickname": &schema.Schema{
					Description: "A brief description of the price, hidden from customers.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"product": &schema.Schema{
					Description: "The ID of the product that this price will belong to.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"product_data": &schema.Schema{
					Description: "These fields can be used to create a new product that this price will belong to.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"active": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						},
						"id": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"metadata": &schema.Schema{
							Description: "",
							Elem:        &schema.Resource{Schema: map[string]*schema.Schema{}},
							Required:    false,
							Type:        schema.TypeMap,
						},
						"name": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeString,
						},
						"statement_descriptor": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"tax_code": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"unit_label": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
					}},
					ForceNew: true,
					Required: false,
					Type:     schema.TypeMap,
				},
				"recurring": &schema.Schema{
					Description: "The recurring components of a price such as `interval` and `usage_type`.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"aggregate_usage": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"interval": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeString,
						},
						"interval_count": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeInt,
						},
						"usage_type": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
					}},
					ForceNew: true,
					Required: false,
					Type:     schema.TypeMap,
				},
				"tax_behavior": &schema.Schema{
					Description: "Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"tiers": &schema.Schema{
					Description: "Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.",
					Elem: &schema.Schema{
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"flat_amount": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeInt,
							},
							"flat_amount_decimal": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"unit_amount": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeInt,
							},
							"unit_amount_decimal": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"up_to": &schema.Schema{
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
				"tiers_mode": &schema.Schema{
					Description: "Defines if the tiering price should be `graduated` or `volume` based. In `volume`-based tiering, the maximum quantity within a period determines the per unit price, in `graduated` tiering pricing can successively change as the quantity grows.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"transfer_lookup_key": &schema.Schema{
					Description: "If set to true, will atomically remove the lookup key from the existing price, and assign it to this price.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeBool,
				},
				"transform_quantity": &schema.Schema{
					Description: "Apply a transformation to the reported usage or set quantity before computing the billed price. Cannot be combined with `tiers`.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"divide_by": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeInt,
						},
						"round": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeString,
						},
					}},
					ForceNew: true,
					Required: false,
					Type:     schema.TypeMap,
				},
				"unit_amount": &schema.Schema{
					Description: "A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge. One of `unit_amount` or `custom_unit_amount` is required, unless `billing_scheme=tiered`.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"unit_amount_decimal": &schema.Schema{
					Description: "Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
			},
			UpdateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if d.HasChange("active") {
					v := d.Get("active")
					{
					}
					params["active"] = v
				}
				if d.HasChange("tax_behavior") {
					v := d.Get("tax_behavior")
					if inEnum(v, []string{"exclusive", "inclusive", "unspecified"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: exclusive, inclusive, unspecified"),
						})
					}
					params["tax_behavior"] = v
				}
				if d.HasChange("transfer_lookup_key") {
					v := d.Get("transfer_lookup_key")
					{
					}
					params["transfer_lookup_key"] = v
				}
				if d.HasChange("nickname") {
					v := d.Get("nickname")
					{
					}
					params["nickname"] = v
				}
				if d.HasChange("currency_options") {
					v := d.Get("currency_options")
					{
						path := cty.Path{}
						outerV := v
					}
					params["currency_options"] = v
				}
				if d.HasChange("lookup_key") {
					v := d.Get("lookup_key")
					{
					}
					params["lookup_key"] = v
				}
				_, err := f.Post(fmt.Sprintf("/v1/prices/%v", d.Get("id")), params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to update price %s", d.Id(), err),
					})
					return out
				}
				return out
			},
		},
		"products": &schema.Resource{
			CreateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if v, exists := d.GetOk("shippable"); exists {
					{
					}
					params["shippable"] = v
				}
				if v, exists := d.GetOk("statement_descriptor"); exists {
					{
					}
					params["statement_descriptor"] = v
				}
				if v, exists := d.GetOk("url"); exists {
					{
					}
					params["url"] = v
				}
				{
					v := d.Get("name")
					{
					}
					params["name"] = v
				}
				if v, exists := d.GetOk("active"); exists {
					{
					}
					params["active"] = v
				}
				if v, exists := d.GetOk("id"); exists {
					{
					}
					params["id"] = v
				}
				if v, exists := d.GetOk("images"); exists {
					{
						path := cty.Path{}
						outerV := v
						for i, v := range outerV.([]any) {
							{
							}
						}
					}
					params["images"] = v
				}
				if v, exists := d.GetOk("package_dimensions"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
						{
						}
						{
						}
					}
					params["package_dimensions"] = v
				}
				if v, exists := d.GetOk("unit_label"); exists {
					{
					}
					params["unit_label"] = v
				}
				if v, exists := d.GetOk("default_price_data"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
						{
						}
						{
							path := path.IndexString("currency_options")
							outerV := shared.Dig[any](outerV)
						}
						{
							path := path.IndexString("recurring")
							outerV := shared.Dig[any](outerV)
							if inEnum(shared.Dig[any](outerV), []string{"day", "month", "week", "year"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: day, month, week, year"),
								})
							}
							{
							}
						}
						if inEnum(shared.Dig[any](outerV), []string{"exclusive", "inclusive", "unspecified"}) {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       fmt.Sprintf("Field must be one of: exclusive, inclusive, unspecified"),
							})
						}
					}
					params["default_price_data"] = v
				}
				if v, exists := d.GetOk("tax_code"); exists {
					{
					}
					params["tax_code"] = v
				}
				if v, exists := d.GetOk("description"); exists {
					{
					}
					params["description"] = v
				}
				res, err := f.Post("/v1/products", params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to create new product", err),
					})
					return out
				}
				d.SetId(Dig(res, "id"))
				return out
			},
			DeleteWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				err := f.Delete(fmt.Sprintf("/v1/products/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to delete product %s", d.Id(), err),
					})
					return out
				}
				return out
			},
			ReadWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				res, err := f.Get(fmt.Sprintf("/v1/products/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to look up product %s", d.Id(), err),
					})
					return out
				}
				d.Set("shippable", Dig[any](res, "shippable"))
				d.Set("statement_descriptor", Dig[any](res, "statement_descriptor"))
				d.Set("url", Dig[any](res, "url"))
				d.Set("name", Dig[any](res, "name"))
				d.Set("active", Dig[any](res, "active"))
				d.Set("id", Dig[any](res, "id"))
				d.Set("images", Dig[any](res, "images"))
				d.Set("package_dimensions", Dig[any](res, "package_dimensions"))
				d.Set("unit_label", Dig[any](res, "unit_label"))
				d.Set("default_price_data", Dig[any](res, "default_price_data"))
				d.Set("tax_code", Dig[any](res, "tax_code"))
				d.Set("description", Dig[any](res, "description"))
				return out
			},
			Schema: map[string]*schema.Schema{
				"active": &schema.Schema{
					Description: "Whether the product is currently available for purchase. Defaults to `true`.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeBool,
				},
				"default_price_data": &schema.Schema{
					Description: "Data used to generate a new [Price](https://stripe.com/docs/api/prices) object. This Price will be set as the default price for this product.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"currency": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeString,
						},
						"currency_options": &schema.Schema{
							Description: "",
							Elem:        &schema.Resource{Schema: map[string]*schema.Schema{}},
							Required:    false,
							Type:        schema.TypeMap,
						},
						"recurring": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"interval": &schema.Schema{
									Description: "",
									Required:    true,
									Type:        schema.TypeString,
								},
								"interval_count": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeInt,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"tax_behavior": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"unit_amount": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeInt,
						},
						"unit_amount_decimal": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
					}},
					ForceNew: true,
					Required: false,
					Type:     schema.TypeMap,
				},
				"description": &schema.Schema{
					Description: "The product's description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"id": &schema.Schema{
					Description: "An identifier will be randomly generated by Stripe. You can optionally override this ID, but the ID must be unique across all products in your Stripe account.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"images": &schema.Schema{
					Description: "A list of up to 8 URLs of images for this product, meant to be displayable to the customer.",
					Elem:        &schema.Schema{Type: schema.TypeString},
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeList,
				},
				"name": &schema.Schema{
					Description: "The product's name, meant to be displayable to the customer.",
					ForceNew:    false,
					Required:    true,
					Type:        schema.TypeString,
				},
				"package_dimensions": &schema.Schema{
					Description: "The dimensions of this product for shipping purposes.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"height": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeFloat,
						},
						"length": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeFloat,
						},
						"weight": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeFloat,
						},
						"width": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeFloat,
						},
					}},
					ForceNew: false,
					Required: false,
					Type:     schema.TypeMap,
				},
				"shippable": &schema.Schema{
					Description: "Whether this product is shipped (i.e., physical goods).",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeBool,
				},
				"statement_descriptor": &schema.Schema{
					Description: "An arbitrary string to be displayed on your customer's credit card or bank statement. While most banks display this information consistently, some may display it incorrectly or not at all.\n\nThis may be up to 22 characters. The statement description may not include `<`, `>`, `\\`, `\"`, `'` characters, and will appear on your customer's statement in capital letters. Non-ASCII characters are automatically stripped.\n It must contain at least one letter.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"tax_code": &schema.Schema{
					Description: "A [tax code](https://stripe.com/docs/tax/tax-categories) ID.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"unit_label": &schema.Schema{
					Description: "A label that represents units of this product in Stripe and on customers’ receipts and invoices. When set, this will be included in associated invoice line item descriptions.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"url": &schema.Schema{
					Description: "A URL of a publicly-accessible webpage for this product.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
			},
			UpdateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if d.HasChange("shippable") {
					v := d.Get("shippable")
					{
					}
					params["shippable"] = v
				}
				if d.HasChange("statement_descriptor") {
					v := d.Get("statement_descriptor")
					{
					}
					params["statement_descriptor"] = v
				}
				if d.HasChange("url") {
					v := d.Get("url")
					{
					}
					params["url"] = v
				}
				if d.HasChange("name") {
					v := d.Get("name")
					{
					}
					params["name"] = v
				}
				if d.HasChange("active") {
					v := d.Get("active")
					{
					}
					params["active"] = v
				}
				if d.HasChange("images") {
					v := d.Get("images")
					{
						path := cty.Path{}
						outerV := v
						for i, v := range outerV.([]any) {
							{
							}
						}
					}
					params["images"] = v
				}
				if d.HasChange("package_dimensions") {
					v := d.Get("package_dimensions")
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
						{
						}
						{
						}
					}
					params["package_dimensions"] = v
				}
				if d.HasChange("unit_label") {
					v := d.Get("unit_label")
					{
					}
					params["unit_label"] = v
				}
				if d.HasChange("tax_code") {
					v := d.Get("tax_code")
					{
					}
					params["tax_code"] = v
				}
				if d.HasChange("description") {
					v := d.Get("description")
					{
					}
					params["description"] = v
				}
				_, err := f.Post(fmt.Sprintf("/v1/products/%v", d.Get("id")), params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to update product %s", d.Id(), err),
					})
					return out
				}
				return out
			},
		},
		"promotion_codes": &schema.Resource{
			CreateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if v, exists := d.GetOk("restrictions"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
							path := path.IndexString("currency_options")
							outerV := shared.Dig[any](outerV)
						}
						{
						}
						{
						}
					}
					params["restrictions"] = v
				}
				if v, exists := d.GetOk("active"); exists {
					{
					}
					params["active"] = v
				}
				{
					v := d.Get("coupon")
					{
					}
					params["coupon"] = v
				}
				if v, exists := d.GetOk("expires_at"); exists {
					{
					}
					params["expires_at"] = v
				}
				if v, exists := d.GetOk("max_redemptions"); exists {
					{
					}
					params["max_redemptions"] = v
				}
				if v, exists := d.GetOk("code"); exists {
					{
					}
					params["code"] = v
				}
				if v, exists := d.GetOk("customer"); exists {
					{
					}
					params["customer"] = v
				}
				res, err := f.Post("/v1/promotion_codes", params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to create new promotion_code", err),
					})
					return out
				}
				d.SetId(Dig(res, "id"))
				return out
			},
			DeleteWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				_, err := f.Post(fmt.Sprintf("/v1/promotion_codes/%v", d.Get("id")), map[string]any{"active": false})
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to set active to false promotion_code %s", d.Id(), err),
					})
					return out
				}
				return out
			},
			ReadWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				res, err := f.Get(fmt.Sprintf("/v1/promotion_codes/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to look up promotion_code %s", d.Id(), err),
					})
					return out
				}
				d.Set("restrictions", Dig[any](res, "restrictions"))
				d.Set("active", Dig[any](res, "active"))
				d.Set("coupon", Dig[any](res, "coupon"))
				d.Set("expires_at", Dig[any](res, "expires_at"))
				d.Set("max_redemptions", Dig[any](res, "max_redemptions"))
				d.Set("code", Dig[any](res, "code"))
				d.Set("customer", Dig[any](res, "customer"))
				return out
			},
			Schema: map[string]*schema.Schema{
				"active": &schema.Schema{
					Description: "Whether the promotion code is currently active.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeBool,
				},
				"code": &schema.Schema{
					Description: "The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for a specific customer. If left blank, we will generate one automatically.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"coupon": &schema.Schema{
					Description: "The coupon for this promotion code.",
					ForceNew:    true,
					Required:    true,
					Type:        schema.TypeString,
				},
				"customer": &schema.Schema{
					Description: "The customer that this promotion code can be used by. If not set, the promotion code can be used by all customers.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"expires_at": &schema.Schema{
					Description: "The timestamp at which this promotion code will expire. If the coupon has specified a `redeems_by`, then this value cannot be after the coupon's `redeems_by`.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"max_redemptions": &schema.Schema{
					Description: "A positive integer specifying the number of times the promotion code can be redeemed. If the coupon has specified a `max_redemptions`, then this value cannot be greater than the coupon's `max_redemptions`.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"restrictions": &schema.Schema{
					Description: "Settings that restrict the redemption of the promotion code.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"currency_options": &schema.Schema{
							Description: "",
							Elem:        &schema.Resource{Schema: map[string]*schema.Schema{}},
							Required:    false,
							Type:        schema.TypeMap,
						},
						"first_time_transaction": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						},
						"minimum_amount": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeInt,
						},
						"minimum_amount_currency": &schema.Schema{
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
			UpdateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if d.HasChange("restrictions") {
					v := d.Get("restrictions")
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
							path := path.IndexString("currency_options")
							outerV := shared.Dig[any](outerV)
						}
						{
						}
						{
						}
					}
					params["restrictions"] = v
				}
				if d.HasChange("active") {
					v := d.Get("active")
					{
					}
					params["active"] = v
				}
				_, err := f.Post(fmt.Sprintf("/v1/promotion_codes/%v", d.Get("id")), params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to update promotion_code %s", d.Id(), err),
					})
					return out
				}
				return out
			},
		},
		"radar_value_list_items": &schema.Resource{
			CreateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				{
					v := d.Get("value")
					{
					}
					params["value"] = v
				}
				{
					v := d.Get("value_list")
					{
					}
					params["value_list"] = v
				}
				res, err := f.Post("/v1/radar/value_list_items", params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to create new radar.value_list_item", err),
					})
					return out
				}
				d.SetId(Dig(res, "id"))
				return out
			},
			DeleteWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				err := f.Delete(fmt.Sprintf("/v1/radar/value_list_items/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to delete radar.value_list_item %s", d.Id(), err),
					})
					return out
				}
				return out
			},
			ReadWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				res, err := f.Get(fmt.Sprintf("/v1/radar/value_list_items/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to look up radar.value_list_item %s", d.Id(), err),
					})
					return out
				}
				d.Set("value", Dig[any](res, "value"))
				d.Set("value_list", Dig[any](res, "value_list"))
				return out
			},
			Schema: map[string]*schema.Schema{
				"value": &schema.Schema{
					Description: "The value of the item (whose type must match the type of the parent value list).",
					ForceNew:    true,
					Required:    true,
					Type:        schema.TypeString,
				},
				"value_list": &schema.Schema{
					Description: "The identifier of the value list which the created item will be added to.",
					ForceNew:    true,
					Required:    true,
					Type:        schema.TypeString,
				},
			},
			UpdateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				_, err := f.Post(fmt.Sprintf("/v1/radar/value_list_items/%v", d.Get("id")), params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to update radar.value_list_item %s", d.Id(), err),
					})
					return out
				}
				return out
			},
		},
		"radar_value_lists": &schema.Resource{
			CreateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				{
					v := d.Get("name")
					{
					}
					params["name"] = v
				}
				{
					v := d.Get("alias")
					{
					}
					params["alias"] = v
				}
				if v, exists := d.GetOk("item_type"); exists {
					if inEnum(v, []string{"card_bin", "card_fingerprint", "case_sensitive_string", "country", "customer_id", "email", "ip_address", "string"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: card_bin, card_fingerprint, case_sensitive_string, country, customer_id, email, ip_address, string"),
						})
					}
					params["item_type"] = v
				}
				res, err := f.Post("/v1/radar/value_lists", params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to create new radar.value_list", err),
					})
					return out
				}
				d.SetId(Dig(res, "id"))
				return out
			},
			DeleteWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				err := f.Delete(fmt.Sprintf("/v1/radar/value_lists/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to delete radar.value_list %s", d.Id(), err),
					})
					return out
				}
				return out
			},
			ReadWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				res, err := f.Get(fmt.Sprintf("/v1/radar/value_lists/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to look up radar.value_list %s", d.Id(), err),
					})
					return out
				}
				d.Set("name", Dig[any](res, "name"))
				d.Set("alias", Dig[any](res, "alias"))
				d.Set("item_type", Dig[any](res, "item_type"))
				return out
			},
			Schema: map[string]*schema.Schema{
				"alias": &schema.Schema{
					Description: "The name of the value list for use in rules.",
					ForceNew:    false,
					Required:    true,
					Type:        schema.TypeString,
				},
				"item_type": &schema.Schema{
					Description: "Type of the items in the value list. One of `card_fingerprint`, `card_bin`, `email`, `ip_address`, `country`, `string`, `case_sensitive_string`, or `customer_id`. Use `string` if the item type is unknown or mixed.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"name": &schema.Schema{
					Description: "The human-readable name of the value list.",
					ForceNew:    false,
					Required:    true,
					Type:        schema.TypeString,
				},
			},
			UpdateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if d.HasChange("name") {
					v := d.Get("name")
					{
					}
					params["name"] = v
				}
				if d.HasChange("alias") {
					v := d.Get("alias")
					{
					}
					params["alias"] = v
				}
				_, err := f.Post(fmt.Sprintf("/v1/radar/value_lists/%v", d.Get("id")), params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to update radar.value_list %s", d.Id(), err),
					})
					return out
				}
				return out
			},
		},
		"subscription_items": &schema.Resource{
			CreateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if v, exists := d.GetOk("price"); exists {
					{
					}
					params["price"] = v
				}
				if v, exists := d.GetOk("price_data"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
							path := path.IndexString("recurring")
							outerV := shared.Dig[any](outerV)
							if inEnum(shared.Dig[any](outerV), []string{"day", "month", "week", "year"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: day, month, week, year"),
								})
							}
							{
							}
						}
						if inEnum(shared.Dig[any](outerV), []string{"exclusive", "inclusive", "unspecified"}) {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       fmt.Sprintf("Field must be one of: exclusive, inclusive, unspecified"),
							})
						}
						{
						}
						{
						}
						{
						}
						{
						}
					}
					params["price_data"] = v
				}
				if v, exists := d.GetOk("proration_behavior"); exists {
					if inEnum(v, []string{"always_invoice", "create_prorations", "none"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: always_invoice, create_prorations, none"),
						})
					}
					params["proration_behavior"] = v
				}
				if v, exists := d.GetOk("quantity"); exists {
					{
					}
					params["quantity"] = v
				}
				{
					v := d.Get("subscription")
					{
					}
					params["subscription"] = v
				}
				if v, exists := d.GetOk("billing_thresholds"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
						}
					}
					params["billing_thresholds"] = v
				}
				if v, exists := d.GetOk("proration_date"); exists {
					{
					}
					params["proration_date"] = v
				}
				if v, exists := d.GetOk("tax_rates"); exists {
					{
						path := cty.Path{}
						outerV := v
						for i, v := range outerV.([]any) {
							{
							}
						}
					}
					params["tax_rates"] = v
				}
				if v, exists := d.GetOk("payment_behavior"); exists {
					if inEnum(v, []string{"allow_incomplete", "default_incomplete", "error_if_incomplete", "pending_if_incomplete"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: allow_incomplete, default_incomplete, error_if_incomplete, pending_if_incomplete"),
						})
					}
					params["payment_behavior"] = v
				}
				res, err := f.Post("/v1/subscription_items", params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to create new subscription_item", err),
					})
					return out
				}
				d.SetId(Dig(res, "id"))
				return out
			},
			DeleteWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				err := f.Delete(fmt.Sprintf("/v1/subscription_items/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to delete subscription_item %s", d.Id(), err),
					})
					return out
				}
				return out
			},
			ReadWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				res, err := f.Get(fmt.Sprintf("/v1/subscription_items/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to look up subscription_item %s", d.Id(), err),
					})
					return out
				}
				d.Set("price", Dig[any](res, "price"))
				d.Set("price_data", Dig[any](res, "price_data"))
				d.Set("proration_behavior", Dig[any](res, "proration_behavior"))
				d.Set("quantity", Dig[any](res, "quantity"))
				d.Set("subscription", Dig[any](res, "subscription"))
				d.Set("billing_thresholds", Dig[any](res, "billing_thresholds"))
				d.Set("proration_date", Dig[any](res, "proration_date"))
				d.Set("tax_rates", Dig[any](res, "tax_rates"))
				d.Set("payment_behavior", Dig[any](res, "payment_behavior"))
				return out
			},
			Schema: map[string]*schema.Schema{
				"billing_thresholds": &schema.Schema{
					Description: "Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. When updating, pass an empty string to remove previously-defined thresholds.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{"usage_gte": &schema.Schema{
						Description: "",
						Required:    true,
						Type:        schema.TypeInt,
					}}},
					ForceNew: false,
					Required: false,
					Type:     schema.TypeMap,
				},
				"payment_behavior": &schema.Schema{
					Description: "Use `allow_incomplete` to transition the subscription to `status=past_due` if a payment is required but cannot be paid. This allows you to manage scenarios where additional user actions are needed to pay a subscription's invoice. For example, SCA regulation may require 3DS authentication to complete payment. See the [SCA Migration Guide](https://stripe.com/docs/billing/migration/strong-customer-authentication) for Billing to learn more. This is the default behavior.\n\nUse `default_incomplete` to transition the subscription to `status=past_due` when payment is required and await explicit confirmation of the invoice's payment intent. This allows simpler management of scenarios where additional user actions are needed to pay a subscription’s invoice. Such as failed payments, [SCA regulation](https://stripe.com/docs/billing/migration/strong-customer-authentication), or collecting a mandate for a bank debit payment method.\n\nUse `pending_if_incomplete` to update the subscription using [pending updates](https://stripe.com/docs/billing/subscriptions/pending-updates). When you use `pending_if_incomplete` you can only pass the parameters [supported by pending updates](https://stripe.com/docs/billing/pending-updates-reference#supported-attributes).\n\nUse `error_if_incomplete` if you want Stripe to return an HTTP 402 status code if a subscription's invoice cannot be paid. For example, if a payment method requires 3DS authentication due to SCA regulation and further user action is needed, this parameter does not update the subscription and returns an error instead. This was the default behavior for API versions prior to 2019-03-14. See the [changelog](https://stripe.com/docs/upgrades#2019-03-14) to learn more.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"price": &schema.Schema{
					Description: "The ID of the price object.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"price_data": &schema.Schema{
					Description: "Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"currency": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeString,
						},
						"product": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeString,
						},
						"recurring": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"interval": &schema.Schema{
									Description: "",
									Required:    true,
									Type:        schema.TypeString,
								},
								"interval_count": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeInt,
								},
							}},
							Required: true,
							Type:     schema.TypeMap,
						},
						"tax_behavior": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"unit_amount": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeInt,
						},
						"unit_amount_decimal": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
					}},
					ForceNew: false,
					Required: false,
					Type:     schema.TypeMap,
				},
				"proration_behavior": &schema.Schema{
					Description: "Determines how to handle [prorations](https://stripe.com/docs/subscriptions/billing-cycle#prorations) when the billing cycle changes (e.g., when switching plans, resetting `billing_cycle_anchor=now`, or starting a trial), or if an item's `quantity` changes.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"proration_date": &schema.Schema{
					Description: "If set, the proration will be calculated as though the subscription was updated at the given time. This can be used to apply the same proration that was previewed with the [upcoming invoice](https://stripe.com/docs/api#retrieve_customer_invoice) endpoint.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"quantity": &schema.Schema{
					Description: "The quantity you'd like to apply to the subscription item you're creating.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"subscription": &schema.Schema{
					Description: "The identifier of the subscription to modify.",
					ForceNew:    true,
					Required:    true,
					Type:        schema.TypeString,
				},
				"tax_rates": &schema.Schema{
					Description: "A list of [Tax Rate](https://stripe.com/docs/api/tax_rates) ids. These Tax Rates will override the [`default_tax_rates`](https://stripe.com/docs/api/subscriptions/create#create_subscription-default_tax_rates) on the Subscription. When updating, pass an empty string to remove previously-defined tax rates.",
					Elem:        &schema.Schema{Type: schema.TypeString},
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeList,
				},
			},
			UpdateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if d.HasChange("price") {
					v := d.Get("price")
					{
					}
					params["price"] = v
				}
				if d.HasChange("price_data") {
					v := d.Get("price_data")
					{
						path := cty.Path{}
						outerV := v
						{
							path := path.IndexString("recurring")
							outerV := shared.Dig[any](outerV)
							if inEnum(shared.Dig[any](outerV), []string{"day", "month", "week", "year"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: day, month, week, year"),
								})
							}
							{
							}
						}
						if inEnum(shared.Dig[any](outerV), []string{"exclusive", "inclusive", "unspecified"}) {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       fmt.Sprintf("Field must be one of: exclusive, inclusive, unspecified"),
							})
						}
						{
						}
						{
						}
						{
						}
						{
						}
					}
					params["price_data"] = v
				}
				if d.HasChange("proration_behavior") {
					v := d.Get("proration_behavior")
					if inEnum(v, []string{"always_invoice", "create_prorations", "none"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: always_invoice, create_prorations, none"),
						})
					}
					params["proration_behavior"] = v
				}
				if d.HasChange("quantity") {
					v := d.Get("quantity")
					{
					}
					params["quantity"] = v
				}
				if d.HasChange("billing_thresholds") {
					v := d.Get("billing_thresholds")
					{
						path := cty.Path{}
						outerV := v
						{
						}
					}
					params["billing_thresholds"] = v
				}
				if d.HasChange("proration_date") {
					v := d.Get("proration_date")
					{
					}
					params["proration_date"] = v
				}
				if d.HasChange("tax_rates") {
					v := d.Get("tax_rates")
					{
						path := cty.Path{}
						outerV := v
						for i, v := range outerV.([]any) {
							{
							}
						}
					}
					params["tax_rates"] = v
				}
				if d.HasChange("payment_behavior") {
					v := d.Get("payment_behavior")
					if inEnum(v, []string{"allow_incomplete", "default_incomplete", "error_if_incomplete", "pending_if_incomplete"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: allow_incomplete, default_incomplete, error_if_incomplete, pending_if_incomplete"),
						})
					}
					params["payment_behavior"] = v
				}
				_, err := f.Post(fmt.Sprintf("/v1/subscription_items/%v", d.Get("id")), params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to update subscription_item %s", d.Id(), err),
					})
					return out
				}
				return out
			},
		},
		"subscriptions": &schema.Resource{
			CreateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if v, exists := d.GetOk("cancel_at"); exists {
					{
					}
					params["cancel_at"] = v
				}
				if v, exists := d.GetOk("currency"); exists {
					{
					}
					params["currency"] = v
				}
				if v, exists := d.GetOk("days_until_due"); exists {
					{
					}
					params["days_until_due"] = v
				}
				if v, exists := d.GetOk("payment_settings"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
							path := path.IndexString("payment_method_types")
							outerV := shared.Dig[any](outerV)
							for i, v := range outerV.([]any) {
								if inEnum(v, []string{"ach_credit_transfer", "ach_debit", "acss_debit", "au_becs_debit", "bacs_debit", "bancontact", "boleto", "card", "customer_balance", "fpx", "giropay", "grabpay", "ideal", "konbini", "link", "paynow", "promptpay", "sepa_debit", "sofort", "us_bank_account", "wechat_pay"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: ach_credit_transfer, ach_debit, acss_debit, au_becs_debit, bacs_debit, bancontact, boleto, card, customer_balance, fpx, giropay, grabpay, ideal, konbini, link, paynow, promptpay, sepa_debit, sofort, us_bank_account, wechat_pay"),
									})
								}
							}
						}
						if inEnum(shared.Dig[any](outerV), []string{"off", "on_subscription"}) {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       fmt.Sprintf("Field must be one of: off, on_subscription"),
							})
						}
						{
							path := path.IndexString("payment_method_options")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("customer_balance")
								outerV := shared.Dig[any](outerV)
								{
								}
								{
									path := path.IndexString("bank_transfer")
									outerV := shared.Dig[any](outerV)
									{
										path := path.IndexString("eu_bank_transfer")
										outerV := shared.Dig[any](outerV)
										{
										}
									}
									{
									}
								}
							}
							{
								path := path.IndexString("konbini")
								outerV := shared.Dig[any](outerV)
							}
							{
								path := path.IndexString("us_bank_account")
								outerV := shared.Dig[any](outerV)
								{
									path := path.IndexString("financial_connections")
									outerV := shared.Dig[any](outerV)
									{
										path := path.IndexString("permissions")
										outerV := shared.Dig[any](outerV)
										for i, v := range outerV.([]any) {
											if inEnum(v, []string{"balances", "ownership", "payment_method", "transactions"}) {
												out = append(out, diag.Diagnostic{
													AttributePath: path,
													Severity:      diag.Error,
													Summary:       fmt.Sprintf("Field must be one of: balances, ownership, payment_method, transactions"),
												})
											}
										}
									}
								}
								if inEnum(shared.Dig[any](outerV), []string{"automatic", "instant", "microdeposits"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: automatic, instant, microdeposits"),
									})
								}
							}
							{
								path := path.IndexString("acss_debit")
								outerV := shared.Dig[any](outerV)
								{
									path := path.IndexString("mandate_options")
									outerV := shared.Dig[any](outerV)
									if inEnum(shared.Dig[any](outerV), []string{"business", "personal"}) {
										out = append(out, diag.Diagnostic{
											AttributePath: path,
											Severity:      diag.Error,
											Summary:       fmt.Sprintf("Field must be one of: business, personal"),
										})
									}
								}
								if inEnum(shared.Dig[any](outerV), []string{"automatic", "instant", "microdeposits"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: automatic, instant, microdeposits"),
									})
								}
							}
							{
								path := path.IndexString("bancontact")
								outerV := shared.Dig[any](outerV)
								if inEnum(shared.Dig[any](outerV), []string{"de", "en", "fr", "nl"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: de, en, fr, nl"),
									})
								}
							}
							{
								path := path.IndexString("card")
								outerV := shared.Dig[any](outerV)
								{
									path := path.IndexString("mandate_options")
									outerV := shared.Dig[any](outerV)
									{
									}
									if inEnum(shared.Dig[any](outerV), []string{"fixed", "maximum"}) {
										out = append(out, diag.Diagnostic{
											AttributePath: path,
											Severity:      diag.Error,
											Summary:       fmt.Sprintf("Field must be one of: fixed, maximum"),
										})
									}
									{
									}
								}
								if inEnum(shared.Dig[any](outerV), []string{"amex", "cartes_bancaires", "diners", "discover", "interac", "jcb", "mastercard", "unionpay", "unknown", "visa"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: amex, cartes_bancaires, diners, discover, interac, jcb, mastercard, unionpay, unknown, visa"),
									})
								}
								if inEnum(shared.Dig[any](outerV), []string{"any", "automatic"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: any, automatic"),
									})
								}
							}
						}
					}
					params["payment_settings"] = v
				}
				if v, exists := d.GetOk("promotion_code"); exists {
					{
					}
					params["promotion_code"] = v
				}
				if v, exists := d.GetOk("default_source"); exists {
					{
					}
					params["default_source"] = v
				}
				if v, exists := d.GetOk("transfer_data"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
					}
					params["transfer_data"] = v
				}
				if v, exists := d.GetOk("backdate_start_date"); exists {
					{
					}
					params["backdate_start_date"] = v
				}
				if v, exists := d.GetOk("billing_thresholds"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
					}
					params["billing_thresholds"] = v
				}
				if v, exists := d.GetOk("cancel_at_period_end"); exists {
					{
					}
					params["cancel_at_period_end"] = v
				}
				if v, exists := d.GetOk("application_fee_percent"); exists {
					{
					}
					params["application_fee_percent"] = v
				}
				if v, exists := d.GetOk("automatic_tax"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
						}
					}
					params["automatic_tax"] = v
				}
				if v, exists := d.GetOk("off_session"); exists {
					{
					}
					params["off_session"] = v
				}
				if v, exists := d.GetOk("on_behalf_of"); exists {
					{
					}
					params["on_behalf_of"] = v
				}
				if v, exists := d.GetOk("trial_end"); exists {
					{
					}
					params["trial_end"] = v
				}
				if v, exists := d.GetOk("default_payment_method"); exists {
					{
					}
					params["default_payment_method"] = v
				}
				if v, exists := d.GetOk("description"); exists {
					{
					}
					params["description"] = v
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
								}
								{
									path := path.IndexString("price_data")
									outerV := shared.Dig[any](outerV)
									if inEnum(shared.Dig[any](outerV), []string{"exclusive", "inclusive", "unspecified"}) {
										out = append(out, diag.Diagnostic{
											AttributePath: path,
											Severity:      diag.Error,
											Summary:       fmt.Sprintf("Field must be one of: exclusive, inclusive, unspecified"),
										})
									}
									{
									}
									{
									}
									{
									}
									{
									}
									{
										path := path.IndexString("recurring")
										outerV := shared.Dig[any](outerV)
										if inEnum(shared.Dig[any](outerV), []string{"day", "month", "week", "year"}) {
											out = append(out, diag.Diagnostic{
												AttributePath: path,
												Severity:      diag.Error,
												Summary:       fmt.Sprintf("Field must be one of: day, month, week, year"),
											})
										}
										{
										}
									}
								}
								{
								}
								{
									path := path.IndexString("tax_rates")
									outerV := shared.Dig[any](outerV)
									for i, v := range outerV.([]any) {
										{
										}
									}
								}
								{
									path := path.IndexString("billing_thresholds")
									outerV := shared.Dig[any](outerV)
									{
									}
								}
								{
									path := path.IndexString("metadata")
									outerV := shared.Dig[any](outerV)
								}
							}
						}
					}
					params["items"] = v
				}
				if v, exists := d.GetOk("billing_cycle_anchor"); exists {
					{
					}
					params["billing_cycle_anchor"] = v
				}
				{
					v := d.Get("customer")
					{
					}
					params["customer"] = v
				}
				if v, exists := d.GetOk("proration_behavior"); exists {
					if inEnum(v, []string{"always_invoice", "create_prorations", "none"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: always_invoice, create_prorations, none"),
						})
					}
					params["proration_behavior"] = v
				}
				if v, exists := d.GetOk("default_tax_rates"); exists {
					{
						path := cty.Path{}
						outerV := v
						for i, v := range outerV.([]any) {
							{
							}
						}
					}
					params["default_tax_rates"] = v
				}
				if v, exists := d.GetOk("payment_behavior"); exists {
					if inEnum(v, []string{"allow_incomplete", "default_incomplete", "error_if_incomplete", "pending_if_incomplete"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: allow_incomplete, default_incomplete, error_if_incomplete, pending_if_incomplete"),
						})
					}
					params["payment_behavior"] = v
				}
				if v, exists := d.GetOk("trial_from_plan"); exists {
					{
					}
					params["trial_from_plan"] = v
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
								}
								{
									path := path.IndexString("price_data")
									outerV := shared.Dig[any](outerV)
									{
									}
									{
									}
									if inEnum(shared.Dig[any](outerV), []string{"exclusive", "inclusive", "unspecified"}) {
										out = append(out, diag.Diagnostic{
											AttributePath: path,
											Severity:      diag.Error,
											Summary:       fmt.Sprintf("Field must be one of: exclusive, inclusive, unspecified"),
										})
									}
									{
									}
									{
									}
								}
								{
								}
								{
									path := path.IndexString("tax_rates")
									outerV := shared.Dig[any](outerV)
									for i, v := range outerV.([]any) {
										{
										}
									}
								}
							}
						}
					}
					params["add_invoice_items"] = v
				}
				if v, exists := d.GetOk("collection_method"); exists {
					if inEnum(v, []string{"charge_automatically", "send_invoice"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: charge_automatically, send_invoice"),
						})
					}
					params["collection_method"] = v
				}
				if v, exists := d.GetOk("coupon"); exists {
					{
					}
					params["coupon"] = v
				}
				if v, exists := d.GetOk("pending_invoice_item_interval"); exists {
					{
						path := cty.Path{}
						outerV := v
						if inEnum(shared.Dig[any](outerV), []string{"day", "month", "week", "year"}) {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       fmt.Sprintf("Field must be one of: day, month, week, year"),
							})
						}
						{
						}
					}
					params["pending_invoice_item_interval"] = v
				}
				if v, exists := d.GetOk("trial_period_days"); exists {
					{
					}
					params["trial_period_days"] = v
				}
				res, err := f.Post("/v1/subscriptions", params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to create new subscription", err),
					})
					return out
				}
				d.SetId(Dig(res, "id"))
				return out
			},
			DeleteWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				err := f.Delete(fmt.Sprintf("/v1/subscriptions/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to delete subscription %s", d.Id(), err),
					})
					return out
				}
				return out
			},
			ReadWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				res, err := f.Get(fmt.Sprintf("/v1/subscriptions/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to look up subscription %s", d.Id(), err),
					})
					return out
				}
				d.Set("cancel_at", Dig[any](res, "cancel_at"))
				d.Set("currency", Dig[any](res, "currency"))
				d.Set("days_until_due", Dig[any](res, "days_until_due"))
				d.Set("payment_settings", Dig[any](res, "payment_settings"))
				d.Set("promotion_code", Dig[any](res, "promotion_code"))
				d.Set("default_source", Dig[any](res, "default_source"))
				d.Set("transfer_data", Dig[any](res, "transfer_data"))
				d.Set("backdate_start_date", Dig[any](res, "backdate_start_date"))
				d.Set("billing_thresholds", Dig[any](res, "billing_thresholds"))
				d.Set("cancel_at_period_end", Dig[any](res, "cancel_at_period_end"))
				d.Set("application_fee_percent", Dig[any](res, "application_fee_percent"))
				d.Set("automatic_tax", Dig[any](res, "automatic_tax"))
				d.Set("off_session", Dig[any](res, "off_session"))
				d.Set("on_behalf_of", Dig[any](res, "on_behalf_of"))
				d.Set("trial_end", Dig[any](res, "trial_end"))
				d.Set("default_payment_method", Dig[any](res, "default_payment_method"))
				d.Set("description", Dig[any](res, "description"))
				d.Set("items", Dig[any](res, "items"))
				d.Set("billing_cycle_anchor", Dig[any](res, "billing_cycle_anchor"))
				d.Set("customer", Dig[any](res, "customer"))
				d.Set("proration_behavior", Dig[any](res, "proration_behavior"))
				d.Set("default_tax_rates", Dig[any](res, "default_tax_rates"))
				d.Set("payment_behavior", Dig[any](res, "payment_behavior"))
				d.Set("trial_from_plan", Dig[any](res, "trial_from_plan"))
				d.Set("add_invoice_items", Dig[any](res, "add_invoice_items"))
				d.Set("collection_method", Dig[any](res, "collection_method"))
				d.Set("coupon", Dig[any](res, "coupon"))
				d.Set("pending_invoice_item_interval", Dig[any](res, "pending_invoice_item_interval"))
				d.Set("trial_period_days", Dig[any](res, "trial_period_days"))
				return out
			},
			Schema: map[string]*schema.Schema{
				"add_invoice_items": &schema.Schema{
					Description: "A list of prices and quantities that will generate invoice items appended to the next invoice for this subscription. You may pass up to 20 items.",
					Elem: &schema.Schema{
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"price": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"price_data": &schema.Schema{
								Description: "",
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"currency": &schema.Schema{
										Description: "",
										Required:    true,
										Type:        schema.TypeString,
									},
									"product": &schema.Schema{
										Description: "",
										Required:    true,
										Type:        schema.TypeString,
									},
									"tax_behavior": &schema.Schema{
										Description: "",
										Required:    false,
										Type:        schema.TypeString,
									},
									"unit_amount": &schema.Schema{
										Description: "",
										Required:    false,
										Type:        schema.TypeInt,
									},
									"unit_amount_decimal": &schema.Schema{
										Description: "",
										Required:    false,
										Type:        schema.TypeString,
									},
								}},
								Required: false,
								Type:     schema.TypeMap,
							},
							"quantity": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeInt,
							},
							"tax_rates": &schema.Schema{
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
				"application_fee_percent": &schema.Schema{
					Description: "A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice subtotal that will be transferred to the application owner's Stripe account. The request must be made by a platform account on a connected account in order to set an application fee percentage. For more information, see the application fees [documentation](https://stripe.com/docs/connect/subscriptions#collecting-fees-on-subscriptions).",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeFloat,
				},
				"automatic_tax": &schema.Schema{
					Description: "Automatic tax settings for this subscription. We recommend you only include this parameter when the existing value is being changed.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{"enabled": &schema.Schema{
						Description: "",
						Required:    true,
						Type:        schema.TypeBool,
					}}},
					ForceNew: false,
					Required: false,
					Type:     schema.TypeMap,
				},
				"backdate_start_date": &schema.Schema{
					Description: "For new subscriptions, a past timestamp to backdate the subscription's start date to. If set, the first invoice will contain a proration for the timespan between the start date and the current time. Can be combined with trials and the billing cycle anchor.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"billing_cycle_anchor": &schema.Schema{
					Description: "A future timestamp to anchor the subscription's [billing cycle](https://stripe.com/docs/subscriptions/billing-cycle). This is used to determine the date of the first full invoice, and, for plans with `month` or `year` intervals, the day of the month for subsequent invoices. The timestamp is in UTC format.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"billing_thresholds": &schema.Schema{
					Description: "Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"amount_gte": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeInt,
						},
						"reset_billing_cycle_anchor": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeBool,
						},
					}},
					ForceNew: false,
					Required: false,
					Type:     schema.TypeMap,
				},
				"cancel_at": &schema.Schema{
					Description: "A timestamp at which the subscription should cancel. If set to a date before the current period ends, this will cause a proration if prorations have been enabled using `proration_behavior`. If set during a future period, this will always cause a proration for that period.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"cancel_at_period_end": &schema.Schema{
					Description: "Boolean indicating whether this subscription should cancel at the end of the current period.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeBool,
				},
				"collection_method": &schema.Schema{
					Description: "Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay this subscription at the end of the cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`. Defaults to `charge_automatically`.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"coupon": &schema.Schema{
					Description: "The ID of the coupon to apply to this subscription. A coupon applied to a subscription will only affect invoices created for that particular subscription.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"currency": &schema.Schema{
					Description: "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"customer": &schema.Schema{
					Description: "The identifier of the customer to subscribe.",
					ForceNew:    true,
					Required:    true,
					Type:        schema.TypeString,
				},
				"days_until_due": &schema.Schema{
					Description: "Number of days a customer has to pay invoices generated by this subscription. Valid only for subscriptions where `collection_method` is set to `send_invoice`.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"default_payment_method": &schema.Schema{
					Description: "ID of the default payment method for the subscription. It must belong to the customer associated with the subscription. This takes precedence over `default_source`. If neither are set, invoices will use the customer's [invoice_settings.default_payment_method](https://stripe.com/docs/api/customers/object#customer_object-invoice_settings-default_payment_method) or [default_source](https://stripe.com/docs/api/customers/object#customer_object-default_source).",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"default_source": &schema.Schema{
					Description: "ID of the default payment source for the subscription. It must belong to the customer associated with the subscription and be in a chargeable state. If `default_payment_method` is also set, `default_payment_method` will take precedence. If neither are set, invoices will use the customer's [invoice_settings.default_payment_method](https://stripe.com/docs/api/customers/object#customer_object-invoice_settings-default_payment_method) or [default_source](https://stripe.com/docs/api/customers/object#customer_object-default_source).",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"default_tax_rates": &schema.Schema{
					Description: "The tax rates that will apply to any subscription item that does not have `tax_rates` set. Invoices created will have their `default_tax_rates` populated from the subscription.",
					Elem:        &schema.Schema{Type: schema.TypeString},
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeList,
				},
				"description": &schema.Schema{
					Description: "The subscription's description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"items": &schema.Schema{
					Description: "A list of up to 20 subscription items, each with an attached price.",
					Elem: &schema.Schema{
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"billing_thresholds": &schema.Schema{
								Description: "",
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{"usage_gte": &schema.Schema{
									Description: "",
									Required:    true,
									Type:        schema.TypeInt,
								}}},
								Required: false,
								Type:     schema.TypeMap,
							},
							"metadata": &schema.Schema{
								Description: "",
								Elem:        &schema.Resource{Schema: map[string]*schema.Schema{}},
								Required:    false,
								Type:        schema.TypeMap,
							},
							"price": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeString,
							},
							"price_data": &schema.Schema{
								Description: "",
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"currency": &schema.Schema{
										Description: "",
										Required:    true,
										Type:        schema.TypeString,
									},
									"product": &schema.Schema{
										Description: "",
										Required:    true,
										Type:        schema.TypeString,
									},
									"recurring": &schema.Schema{
										Description: "",
										Elem: &schema.Resource{Schema: map[string]*schema.Schema{
											"interval": &schema.Schema{
												Description: "",
												Required:    true,
												Type:        schema.TypeString,
											},
											"interval_count": &schema.Schema{
												Description: "",
												Required:    false,
												Type:        schema.TypeInt,
											},
										}},
										Required: true,
										Type:     schema.TypeMap,
									},
									"tax_behavior": &schema.Schema{
										Description: "",
										Required:    false,
										Type:        schema.TypeString,
									},
									"unit_amount": &schema.Schema{
										Description: "",
										Required:    false,
										Type:        schema.TypeInt,
									},
									"unit_amount_decimal": &schema.Schema{
										Description: "",
										Required:    false,
										Type:        schema.TypeString,
									},
								}},
								Required: false,
								Type:     schema.TypeMap,
							},
							"quantity": &schema.Schema{
								Description: "",
								Required:    false,
								Type:        schema.TypeInt,
							},
							"tax_rates": &schema.Schema{
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
				"off_session": &schema.Schema{
					Description: "Indicates if a customer is on or off-session while an invoice payment is attempted.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeBool,
				},
				"on_behalf_of": &schema.Schema{
					Description: "The account on behalf of which to charge, for each of the subscription's invoices.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"payment_behavior": &schema.Schema{
					Description: "Only applies to subscriptions with `collection_method=charge_automatically`.\n\nUse `allow_incomplete` to create subscriptions with `status=incomplete` if the first invoice cannot be paid. Creating subscriptions with this status allows you to manage scenarios where additional user actions are needed to pay a subscription's invoice. For example, SCA regulation may require 3DS authentication to complete payment. See the [SCA Migration Guide](https://stripe.com/docs/billing/migration/strong-customer-authentication) for Billing to learn more. This is the default behavior.\n\nUse `default_incomplete` to create Subscriptions with `status=incomplete` when the first invoice requires payment, otherwise start as active. Subscriptions transition to `status=active` when successfully confirming the payment intent on the first invoice. This allows simpler management of scenarios where additional user actions are needed to pay a subscription’s invoice. Such as failed payments, [SCA regulation](https://stripe.com/docs/billing/migration/strong-customer-authentication), or collecting a mandate for a bank debit payment method. If the payment intent is not confirmed within 23 hours subscriptions transition to `status=incomplete_expired`, which is a terminal state.\n\nUse `error_if_incomplete` if you want Stripe to return an HTTP 402 status code if a subscription's first invoice cannot be paid. For example, if a payment method requires 3DS authentication due to SCA regulation and further user action is needed, this parameter does not create a subscription and returns an error instead. This was the default behavior for API versions prior to 2019-03-14. See the [changelog](https://stripe.com/docs/upgrades#2019-03-14) to learn more.\n\n`pending_if_incomplete` is only used with updates and cannot be passed when creating a subscription.\n\nSubscriptions with `collection_method=send_invoice` are automatically activated regardless of the first invoice status.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"payment_settings": &schema.Schema{
					Description: "Payment settings to pass to invoices created by the subscription.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"payment_method_options": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"acss_debit": &schema.Schema{
									Description: "",
									Elem: &schema.Resource{Schema: map[string]*schema.Schema{
										"mandate_options": &schema.Schema{
											Description: "",
											Elem: &schema.Resource{Schema: map[string]*schema.Schema{"transaction_type": &schema.Schema{
												Description: "",
												Required:    false,
												Type:        schema.TypeString,
											}}},
											Required: false,
											Type:     schema.TypeMap,
										},
										"verification_method": &schema.Schema{
											Description: "",
											Required:    false,
											Type:        schema.TypeString,
										},
									}},
									Required: false,
									Type:     schema.TypeMap,
								},
								"bancontact": &schema.Schema{
									Description: "",
									Elem: &schema.Resource{Schema: map[string]*schema.Schema{"preferred_language": &schema.Schema{
										Description: "",
										Required:    false,
										Type:        schema.TypeString,
									}}},
									Required: false,
									Type:     schema.TypeMap,
								},
								"card": &schema.Schema{
									Description: "",
									Elem: &schema.Resource{Schema: map[string]*schema.Schema{
										"mandate_options": &schema.Schema{
											Description: "",
											Elem: &schema.Resource{Schema: map[string]*schema.Schema{
												"amount": &schema.Schema{
													Description: "",
													Required:    false,
													Type:        schema.TypeInt,
												},
												"amount_type": &schema.Schema{
													Description: "",
													Required:    false,
													Type:        schema.TypeString,
												},
												"description": &schema.Schema{
													Description: "",
													Required:    false,
													Type:        schema.TypeString,
												},
											}},
											Required: false,
											Type:     schema.TypeMap,
										},
										"network": &schema.Schema{
											Description: "",
											Required:    false,
											Type:        schema.TypeString,
										},
										"request_three_d_secure": &schema.Schema{
											Description: "",
											Required:    false,
											Type:        schema.TypeString,
										},
									}},
									Required: false,
									Type:     schema.TypeMap,
								},
								"customer_balance": &schema.Schema{
									Description: "",
									Elem: &schema.Resource{Schema: map[string]*schema.Schema{
										"bank_transfer": &schema.Schema{
											Description: "",
											Elem: &schema.Resource{Schema: map[string]*schema.Schema{
												"eu_bank_transfer": &schema.Schema{
													Description: "",
													Elem: &schema.Resource{Schema: map[string]*schema.Schema{"country": &schema.Schema{
														Description: "",
														Required:    true,
														Type:        schema.TypeString,
													}}},
													Required: false,
													Type:     schema.TypeMap,
												},
												"type": &schema.Schema{
													Description: "",
													Required:    false,
													Type:        schema.TypeString,
												},
											}},
											Required: false,
											Type:     schema.TypeMap,
										},
										"funding_type": &schema.Schema{
											Description: "",
											Required:    false,
											Type:        schema.TypeString,
										},
									}},
									Required: false,
									Type:     schema.TypeMap,
								},
								"konbini": &schema.Schema{
									Description: "",
									Elem:        &schema.Resource{Schema: map[string]*schema.Schema{}},
									Required:    false,
									Type:        schema.TypeMap,
								},
								"us_bank_account": &schema.Schema{
									Description: "",
									Elem: &schema.Resource{Schema: map[string]*schema.Schema{
										"financial_connections": &schema.Schema{
											Description: "",
											Elem: &schema.Resource{Schema: map[string]*schema.Schema{"permissions": &schema.Schema{
												Description: "",
												Elem:        &schema.Schema{Type: schema.TypeString},
												Required:    false,
												Type:        schema.TypeList,
											}}},
											Required: false,
											Type:     schema.TypeMap,
										},
										"verification_method": &schema.Schema{
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
						"payment_method_types": &schema.Schema{
							Description: "",
							Elem:        &schema.Schema{Type: schema.TypeString},
							Required:    false,
							Type:        schema.TypeList,
						},
						"save_default_payment_method": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
					}},
					ForceNew: false,
					Required: false,
					Type:     schema.TypeMap,
				},
				"pending_invoice_item_interval": &schema.Schema{
					Description: "Specifies an interval for how often to bill for any pending invoice items. It is analogous to calling [Create an invoice](https://stripe.com/docs/api#create_invoice) for the given subscription at the specified interval.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"interval": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeString,
						},
						"interval_count": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeInt,
						},
					}},
					ForceNew: false,
					Required: false,
					Type:     schema.TypeMap,
				},
				"promotion_code": &schema.Schema{
					Description: "The API ID of a promotion code to apply to this subscription. A promotion code applied to a subscription will only affect invoices created for that particular subscription.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"proration_behavior": &schema.Schema{
					Description: "Determines how to handle [prorations](https://stripe.com/docs/subscriptions/billing-cycle#prorations) resulting from the `billing_cycle_anchor`. If no value is passed, the default is `create_prorations`.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"transfer_data": &schema.Schema{
					Description: "If specified, the funds from the subscription's invoices will be transferred to the destination and the ID of the resulting transfers will be found on the resulting charges.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"amount_percent": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeFloat,
						},
						"destination": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeString,
						},
					}},
					ForceNew: false,
					Required: false,
					Type:     schema.TypeMap,
				},
				"trial_end": &schema.Schema{
					Description: "Unix timestamp representing the end of the trial period the customer will get before being charged for the first time. This will always overwrite any trials that might apply via a subscribed plan. If set, trial_end will override the default trial period of the plan the customer is being subscribed to. The special value `now` can be provided to end the customer's trial immediately. Can be at most two years from `billing_cycle_anchor`. See [Using trial periods on subscriptions](https://stripe.com/docs/billing/subscriptions/trials) to learn more.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeInt,
				},
				"trial_from_plan": &schema.Schema{
					Description: "Indicates if a plan's `trial_period_days` should be applied to the subscription. Setting `trial_end` per subscription is preferred, and this defaults to `false`. Setting this flag to `true` together with `trial_end` is not allowed. See [Using trial periods on subscriptions](https://stripe.com/docs/billing/subscriptions/trials) to learn more.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeBool,
				},
				"trial_period_days": &schema.Schema{
					Description: "Integer representing the number of trial period days before the customer is charged for the first time. This will always overwrite any trials that might apply via a subscribed plan. See [Using trial periods on subscriptions](https://stripe.com/docs/billing/subscriptions/trials) to learn more.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeInt,
				},
			},
			UpdateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if d.HasChange("cancel_at") {
					v := d.Get("cancel_at")
					{
					}
					params["cancel_at"] = v
				}
				if d.HasChange("days_until_due") {
					v := d.Get("days_until_due")
					{
					}
					params["days_until_due"] = v
				}
				if d.HasChange("payment_settings") {
					v := d.Get("payment_settings")
					{
						path := cty.Path{}
						outerV := v
						{
							path := path.IndexString("payment_method_types")
							outerV := shared.Dig[any](outerV)
							for i, v := range outerV.([]any) {
								if inEnum(v, []string{"ach_credit_transfer", "ach_debit", "acss_debit", "au_becs_debit", "bacs_debit", "bancontact", "boleto", "card", "customer_balance", "fpx", "giropay", "grabpay", "ideal", "konbini", "link", "paynow", "promptpay", "sepa_debit", "sofort", "us_bank_account", "wechat_pay"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: ach_credit_transfer, ach_debit, acss_debit, au_becs_debit, bacs_debit, bancontact, boleto, card, customer_balance, fpx, giropay, grabpay, ideal, konbini, link, paynow, promptpay, sepa_debit, sofort, us_bank_account, wechat_pay"),
									})
								}
							}
						}
						if inEnum(shared.Dig[any](outerV), []string{"off", "on_subscription"}) {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       fmt.Sprintf("Field must be one of: off, on_subscription"),
							})
						}
						{
							path := path.IndexString("payment_method_options")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("customer_balance")
								outerV := shared.Dig[any](outerV)
								{
								}
								{
									path := path.IndexString("bank_transfer")
									outerV := shared.Dig[any](outerV)
									{
										path := path.IndexString("eu_bank_transfer")
										outerV := shared.Dig[any](outerV)
										{
										}
									}
									{
									}
								}
							}
							{
								path := path.IndexString("konbini")
								outerV := shared.Dig[any](outerV)
							}
							{
								path := path.IndexString("us_bank_account")
								outerV := shared.Dig[any](outerV)
								{
									path := path.IndexString("financial_connections")
									outerV := shared.Dig[any](outerV)
									{
										path := path.IndexString("permissions")
										outerV := shared.Dig[any](outerV)
										for i, v := range outerV.([]any) {
											if inEnum(v, []string{"balances", "ownership", "payment_method", "transactions"}) {
												out = append(out, diag.Diagnostic{
													AttributePath: path,
													Severity:      diag.Error,
													Summary:       fmt.Sprintf("Field must be one of: balances, ownership, payment_method, transactions"),
												})
											}
										}
									}
								}
								if inEnum(shared.Dig[any](outerV), []string{"automatic", "instant", "microdeposits"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: automatic, instant, microdeposits"),
									})
								}
							}
							{
								path := path.IndexString("acss_debit")
								outerV := shared.Dig[any](outerV)
								{
									path := path.IndexString("mandate_options")
									outerV := shared.Dig[any](outerV)
									if inEnum(shared.Dig[any](outerV), []string{"business", "personal"}) {
										out = append(out, diag.Diagnostic{
											AttributePath: path,
											Severity:      diag.Error,
											Summary:       fmt.Sprintf("Field must be one of: business, personal"),
										})
									}
								}
								if inEnum(shared.Dig[any](outerV), []string{"automatic", "instant", "microdeposits"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: automatic, instant, microdeposits"),
									})
								}
							}
							{
								path := path.IndexString("bancontact")
								outerV := shared.Dig[any](outerV)
								if inEnum(shared.Dig[any](outerV), []string{"de", "en", "fr", "nl"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: de, en, fr, nl"),
									})
								}
							}
							{
								path := path.IndexString("card")
								outerV := shared.Dig[any](outerV)
								{
									path := path.IndexString("mandate_options")
									outerV := shared.Dig[any](outerV)
									{
									}
									if inEnum(shared.Dig[any](outerV), []string{"fixed", "maximum"}) {
										out = append(out, diag.Diagnostic{
											AttributePath: path,
											Severity:      diag.Error,
											Summary:       fmt.Sprintf("Field must be one of: fixed, maximum"),
										})
									}
									{
									}
								}
								if inEnum(shared.Dig[any](outerV), []string{"amex", "cartes_bancaires", "diners", "discover", "interac", "jcb", "mastercard", "unionpay", "unknown", "visa"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: amex, cartes_bancaires, diners, discover, interac, jcb, mastercard, unionpay, unknown, visa"),
									})
								}
								if inEnum(shared.Dig[any](outerV), []string{"any", "automatic"}) {
									out = append(out, diag.Diagnostic{
										AttributePath: path,
										Severity:      diag.Error,
										Summary:       fmt.Sprintf("Field must be one of: any, automatic"),
									})
								}
							}
						}
					}
					params["payment_settings"] = v
				}
				if d.HasChange("promotion_code") {
					v := d.Get("promotion_code")
					{
					}
					params["promotion_code"] = v
				}
				if d.HasChange("default_source") {
					v := d.Get("default_source")
					{
					}
					params["default_source"] = v
				}
				if d.HasChange("transfer_data") {
					v := d.Get("transfer_data")
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
					}
					params["transfer_data"] = v
				}
				if d.HasChange("billing_thresholds") {
					v := d.Get("billing_thresholds")
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
					}
					params["billing_thresholds"] = v
				}
				if d.HasChange("cancel_at_period_end") {
					v := d.Get("cancel_at_period_end")
					{
					}
					params["cancel_at_period_end"] = v
				}
				if d.HasChange("application_fee_percent") {
					v := d.Get("application_fee_percent")
					{
					}
					params["application_fee_percent"] = v
				}
				if d.HasChange("automatic_tax") {
					v := d.Get("automatic_tax")
					{
						path := cty.Path{}
						outerV := v
						{
						}
					}
					params["automatic_tax"] = v
				}
				if d.HasChange("off_session") {
					v := d.Get("off_session")
					{
					}
					params["off_session"] = v
				}
				if d.HasChange("on_behalf_of") {
					v := d.Get("on_behalf_of")
					{
					}
					params["on_behalf_of"] = v
				}
				if d.HasChange("trial_end") {
					v := d.Get("trial_end")
					{
					}
					params["trial_end"] = v
				}
				if d.HasChange("default_payment_method") {
					v := d.Get("default_payment_method")
					{
					}
					params["default_payment_method"] = v
				}
				if d.HasChange("description") {
					v := d.Get("description")
					{
					}
					params["description"] = v
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
								}
								{
									path := path.IndexString("price_data")
									outerV := shared.Dig[any](outerV)
									if inEnum(shared.Dig[any](outerV), []string{"exclusive", "inclusive", "unspecified"}) {
										out = append(out, diag.Diagnostic{
											AttributePath: path,
											Severity:      diag.Error,
											Summary:       fmt.Sprintf("Field must be one of: exclusive, inclusive, unspecified"),
										})
									}
									{
									}
									{
									}
									{
									}
									{
									}
									{
										path := path.IndexString("recurring")
										outerV := shared.Dig[any](outerV)
										if inEnum(shared.Dig[any](outerV), []string{"day", "month", "week", "year"}) {
											out = append(out, diag.Diagnostic{
												AttributePath: path,
												Severity:      diag.Error,
												Summary:       fmt.Sprintf("Field must be one of: day, month, week, year"),
											})
										}
										{
										}
									}
								}
								{
								}
								{
									path := path.IndexString("tax_rates")
									outerV := shared.Dig[any](outerV)
									for i, v := range outerV.([]any) {
										{
										}
									}
								}
								{
									path := path.IndexString("billing_thresholds")
									outerV := shared.Dig[any](outerV)
									{
									}
								}
								{
									path := path.IndexString("metadata")
									outerV := shared.Dig[any](outerV)
								}
							}
						}
					}
					params["items"] = v
				}
				if d.HasChange("billing_cycle_anchor") {
					v := d.Get("billing_cycle_anchor")
					{
					}
					params["billing_cycle_anchor"] = v
				}
				if d.HasChange("proration_behavior") {
					v := d.Get("proration_behavior")
					if inEnum(v, []string{"always_invoice", "create_prorations", "none"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: always_invoice, create_prorations, none"),
						})
					}
					params["proration_behavior"] = v
				}
				if d.HasChange("default_tax_rates") {
					v := d.Get("default_tax_rates")
					{
						path := cty.Path{}
						outerV := v
						for i, v := range outerV.([]any) {
							{
							}
						}
					}
					params["default_tax_rates"] = v
				}
				if d.HasChange("payment_behavior") {
					v := d.Get("payment_behavior")
					if inEnum(v, []string{"allow_incomplete", "default_incomplete", "error_if_incomplete", "pending_if_incomplete"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: allow_incomplete, default_incomplete, error_if_incomplete, pending_if_incomplete"),
						})
					}
					params["payment_behavior"] = v
				}
				if d.HasChange("trial_from_plan") {
					v := d.Get("trial_from_plan")
					{
					}
					params["trial_from_plan"] = v
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
								}
								{
									path := path.IndexString("price_data")
									outerV := shared.Dig[any](outerV)
									{
									}
									{
									}
									if inEnum(shared.Dig[any](outerV), []string{"exclusive", "inclusive", "unspecified"}) {
										out = append(out, diag.Diagnostic{
											AttributePath: path,
											Severity:      diag.Error,
											Summary:       fmt.Sprintf("Field must be one of: exclusive, inclusive, unspecified"),
										})
									}
									{
									}
									{
									}
								}
								{
								}
								{
									path := path.IndexString("tax_rates")
									outerV := shared.Dig[any](outerV)
									for i, v := range outerV.([]any) {
										{
										}
									}
								}
							}
						}
					}
					params["add_invoice_items"] = v
				}
				if d.HasChange("collection_method") {
					v := d.Get("collection_method")
					if inEnum(v, []string{"charge_automatically", "send_invoice"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: charge_automatically, send_invoice"),
						})
					}
					params["collection_method"] = v
				}
				if d.HasChange("coupon") {
					v := d.Get("coupon")
					{
					}
					params["coupon"] = v
				}
				if d.HasChange("pending_invoice_item_interval") {
					v := d.Get("pending_invoice_item_interval")
					{
						path := cty.Path{}
						outerV := v
						if inEnum(shared.Dig[any](outerV), []string{"day", "month", "week", "year"}) {
							out = append(out, diag.Diagnostic{
								AttributePath: path,
								Severity:      diag.Error,
								Summary:       fmt.Sprintf("Field must be one of: day, month, week, year"),
							})
						}
						{
						}
					}
					params["pending_invoice_item_interval"] = v
				}
				_, err := f.Post(fmt.Sprintf("/v1/subscriptions/%v", d.Get("id")), params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to update subscription %s", d.Id(), err),
					})
					return out
				}
				return out
			},
		},
		"tax_rates": &schema.Resource{
			CreateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if v, exists := d.GetOk("country"); exists {
					{
					}
					params["country"] = v
				}
				if v, exists := d.GetOk("description"); exists {
					{
					}
					params["description"] = v
				}
				{
					v := d.Get("inclusive")
					{
					}
					params["inclusive"] = v
				}
				if v, exists := d.GetOk("active"); exists {
					{
					}
					params["active"] = v
				}
				if v, exists := d.GetOk("jurisdiction"); exists {
					{
					}
					params["jurisdiction"] = v
				}
				{
					v := d.Get("percentage")
					{
					}
					params["percentage"] = v
				}
				if v, exists := d.GetOk("state"); exists {
					{
					}
					params["state"] = v
				}
				if v, exists := d.GetOk("tax_type"); exists {
					if inEnum(v, []string{"gst", "hst", "jct", "pst", "qst", "rst", "sales_tax", "vat"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: gst, hst, jct, pst, qst, rst, sales_tax, vat"),
						})
					}
					params["tax_type"] = v
				}
				{
					v := d.Get("display_name")
					{
					}
					params["display_name"] = v
				}
				res, err := f.Post("/v1/tax_rates", params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to create new tax_rate", err),
					})
					return out
				}
				d.SetId(Dig(res, "id"))
				return out
			},
			DeleteWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				_, err := f.Post(fmt.Sprintf("/v1/tax_rates/%v", d.Get("id")), map[string]any{"active": false})
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to set active to false tax_rate %s", d.Id(), err),
					})
					return out
				}
				return out
			},
			ReadWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				res, err := f.Get(fmt.Sprintf("/v1/tax_rates/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to look up tax_rate %s", d.Id(), err),
					})
					return out
				}
				d.Set("country", Dig[any](res, "country"))
				d.Set("description", Dig[any](res, "description"))
				d.Set("inclusive", Dig[any](res, "inclusive"))
				d.Set("active", Dig[any](res, "active"))
				d.Set("jurisdiction", Dig[any](res, "jurisdiction"))
				d.Set("percentage", Dig[any](res, "percentage"))
				d.Set("state", Dig[any](res, "state"))
				d.Set("tax_type", Dig[any](res, "tax_type"))
				d.Set("display_name", Dig[any](res, "display_name"))
				return out
			},
			Schema: map[string]*schema.Schema{
				"active": &schema.Schema{
					Description: "Flag determining whether the tax rate is active or inactive (archived). Inactive tax rates cannot be used with new applications or Checkout Sessions, but will still work for subscriptions and invoices that already have it set.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeBool,
				},
				"country": &schema.Schema{
					Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"description": &schema.Schema{
					Description: "An arbitrary string attached to the tax rate for your internal use only. It will not be visible to your customers.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"display_name": &schema.Schema{
					Description: "The display name of the tax rate, which will be shown to users.",
					ForceNew:    false,
					Required:    true,
					Type:        schema.TypeString,
				},
				"inclusive": &schema.Schema{
					Description: "This specifies if the tax rate is inclusive or exclusive.",
					ForceNew:    true,
					Required:    true,
					Type:        schema.TypeBool,
				},
				"jurisdiction": &schema.Schema{
					Description: "The jurisdiction for the tax rate. You can use this label field for tax reporting purposes. It also appears on your customer’s invoice.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"percentage": &schema.Schema{
					Description: "This represents the tax rate percent out of 100.",
					ForceNew:    true,
					Required:    true,
					Type:        schema.TypeFloat,
				},
				"state": &schema.Schema{
					Description: "[ISO 3166-2 subdivision code](https://en.wikipedia.org/wiki/ISO_3166-2:US), without country prefix. For example, \"NY\" for New York, United States.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"tax_type": &schema.Schema{
					Description: "The high-level tax type, such as `vat` or `sales_tax`.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
			},
			UpdateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if d.HasChange("country") {
					v := d.Get("country")
					{
					}
					params["country"] = v
				}
				if d.HasChange("description") {
					v := d.Get("description")
					{
					}
					params["description"] = v
				}
				if d.HasChange("active") {
					v := d.Get("active")
					{
					}
					params["active"] = v
				}
				if d.HasChange("jurisdiction") {
					v := d.Get("jurisdiction")
					{
					}
					params["jurisdiction"] = v
				}
				if d.HasChange("state") {
					v := d.Get("state")
					{
					}
					params["state"] = v
				}
				if d.HasChange("tax_type") {
					v := d.Get("tax_type")
					if inEnum(v, []string{"gst", "hst", "jct", "pst", "qst", "rst", "sales_tax", "vat"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: gst, hst, jct, pst, qst, rst, sales_tax, vat"),
						})
					}
					params["tax_type"] = v
				}
				if d.HasChange("display_name") {
					v := d.Get("display_name")
					{
					}
					params["display_name"] = v
				}
				_, err := f.Post(fmt.Sprintf("/v1/tax_rates/%v", d.Get("id")), params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to update tax_rate %s", d.Id(), err),
					})
					return out
				}
				return out
			},
		},
		"terminal_configurations": &schema.Resource{
			CreateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if v, exists := d.GetOk("bbpos_wisepos_e"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
						}
					}
					params["bbpos_wisepos_e"] = v
				}
				if v, exists := d.GetOk("tipping"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
							path := path.IndexString("usd")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("cad")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("chf")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("czk")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("dkk")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("gbp")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("aud")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("nok")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("nzd")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("sek")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
						}
						{
							path := path.IndexString("hkd")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("sgd")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("eur")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("myr")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
					}
					params["tipping"] = v
				}
				if v, exists := d.GetOk("verifone_p400"); exists {
					{
						path := cty.Path{}
						outerV := v
						{
						}
					}
					params["verifone_p400"] = v
				}
				res, err := f.Post("/v1/terminal/configurations", params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to create new terminal.configuration", err),
					})
					return out
				}
				d.SetId(Dig(res, "id"))
				return out
			},
			DeleteWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				err := f.Delete(fmt.Sprintf("/v1/terminal/configurations/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to delete terminal.configuration %s", d.Id(), err),
					})
					return out
				}
				return out
			},
			ReadWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				res, err := f.Get(fmt.Sprintf("/v1/terminal/configurations/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to look up terminal.configuration %s", d.Id(), err),
					})
					return out
				}
				d.Set("bbpos_wisepos_e", Dig[any](res, "bbpos_wisepos_e"))
				d.Set("tipping", Dig[any](res, "tipping"))
				d.Set("verifone_p400", Dig[any](res, "verifone_p400"))
				return out
			},
			Schema: map[string]*schema.Schema{
				"bbpos_wisepos_e": &schema.Schema{
					Description: "An object containing device type specific settings for BBPOS WisePOS E readers",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{"splashscreen": &schema.Schema{
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					}}},
					ForceNew: false,
					Required: false,
					Type:     schema.TypeMap,
				},
				"tipping": &schema.Schema{
					Description: "Tipping configurations for readers supporting on-reader tips",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"aud": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"fixed_amounts": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"percentages": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"smart_tip_threshold": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeInt,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"cad": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"fixed_amounts": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"percentages": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"smart_tip_threshold": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeInt,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"chf": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"fixed_amounts": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"percentages": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"smart_tip_threshold": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeInt,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"czk": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"fixed_amounts": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"percentages": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"smart_tip_threshold": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeInt,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"dkk": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"fixed_amounts": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"percentages": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"smart_tip_threshold": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeInt,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"eur": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"fixed_amounts": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"percentages": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"smart_tip_threshold": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeInt,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"gbp": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"fixed_amounts": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"percentages": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"smart_tip_threshold": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeInt,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"hkd": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"fixed_amounts": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"percentages": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"smart_tip_threshold": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeInt,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"myr": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"fixed_amounts": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"percentages": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"smart_tip_threshold": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeInt,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"nok": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"fixed_amounts": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"percentages": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"smart_tip_threshold": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeInt,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"nzd": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"fixed_amounts": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"percentages": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"smart_tip_threshold": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeInt,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"sek": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"fixed_amounts": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"percentages": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"smart_tip_threshold": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeInt,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"sgd": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"fixed_amounts": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"percentages": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"smart_tip_threshold": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeInt,
								},
							}},
							Required: false,
							Type:     schema.TypeMap,
						},
						"usd": &schema.Schema{
							Description: "",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"fixed_amounts": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"percentages": &schema.Schema{
									Description: "",
									Elem:        &schema.Schema{Type: schema.TypeInt},
									Required:    false,
									Type:        schema.TypeList,
								},
								"smart_tip_threshold": &schema.Schema{
									Description: "",
									Required:    false,
									Type:        schema.TypeInt,
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
				"verifone_p400": &schema.Schema{
					Description: "An object containing device type specific settings for Verifone P400 readers",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{"splashscreen": &schema.Schema{
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					}}},
					ForceNew: false,
					Required: false,
					Type:     schema.TypeMap,
				},
			},
			UpdateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if d.HasChange("bbpos_wisepos_e") {
					v := d.Get("bbpos_wisepos_e")
					{
						path := cty.Path{}
						outerV := v
						{
						}
					}
					params["bbpos_wisepos_e"] = v
				}
				if d.HasChange("tipping") {
					v := d.Get("tipping")
					{
						path := cty.Path{}
						outerV := v
						{
							path := path.IndexString("usd")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("cad")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("chf")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("czk")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("dkk")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("gbp")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("aud")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("nok")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("nzd")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("sek")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
						}
						{
							path := path.IndexString("hkd")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("sgd")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("eur")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
						{
							path := path.IndexString("myr")
							outerV := shared.Dig[any](outerV)
							{
								path := path.IndexString("fixed_amounts")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
								path := path.IndexString("percentages")
								outerV := shared.Dig[any](outerV)
								for i, v := range outerV.([]any) {
									{
									}
								}
							}
							{
							}
						}
					}
					params["tipping"] = v
				}
				if d.HasChange("verifone_p400") {
					v := d.Get("verifone_p400")
					{
						path := cty.Path{}
						outerV := v
						{
						}
					}
					params["verifone_p400"] = v
				}
				_, err := f.Post(fmt.Sprintf("/v1/terminal/configurations/%v", d.Get("id")), params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to update terminal.configuration %s", d.Id(), err),
					})
					return out
				}
				return out
			},
		},
		"terminal_locations": &schema.Resource{
			CreateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				{
					v := d.Get("address")
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
						{
						}
						{
						}
						{
						}
						{
						}
					}
					params["address"] = v
				}
				if v, exists := d.GetOk("configuration_overrides"); exists {
					{
					}
					params["configuration_overrides"] = v
				}
				{
					v := d.Get("display_name")
					{
					}
					params["display_name"] = v
				}
				res, err := f.Post("/v1/terminal/locations", params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to create new terminal.location", err),
					})
					return out
				}
				d.SetId(Dig(res, "id"))
				return out
			},
			DeleteWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				err := f.Delete(fmt.Sprintf("/v1/terminal/locations/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to delete terminal.location %s", d.Id(), err),
					})
					return out
				}
				return out
			},
			ReadWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				res, err := f.Get(fmt.Sprintf("/v1/terminal/locations/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to look up terminal.location %s", d.Id(), err),
					})
					return out
				}
				d.Set("address", Dig[any](res, "address"))
				d.Set("configuration_overrides", Dig[any](res, "configuration_overrides"))
				d.Set("display_name", Dig[any](res, "display_name"))
				return out
			},
			Schema: map[string]*schema.Schema{
				"address": &schema.Schema{
					Description: "The full address of the location.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"city": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"country": &schema.Schema{
							Description: "",
							Required:    true,
							Type:        schema.TypeString,
						},
						"line1": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"line2": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"postal_code": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
						"state": &schema.Schema{
							Description: "",
							Required:    false,
							Type:        schema.TypeString,
						},
					}},
					ForceNew: false,
					Required: true,
					Type:     schema.TypeMap,
				},
				"configuration_overrides": &schema.Schema{
					Description: "The ID of a configuration that will be used to customize all readers in this location.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"display_name": &schema.Schema{
					Description: "A name for the location.",
					ForceNew:    false,
					Required:    true,
					Type:        schema.TypeString,
				},
			},
			UpdateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if d.HasChange("address") {
					v := d.Get("address")
					{
						path := cty.Path{}
						outerV := v
						{
						}
						{
						}
						{
						}
						{
						}
						{
						}
						{
						}
					}
					params["address"] = v
				}
				if d.HasChange("configuration_overrides") {
					v := d.Get("configuration_overrides")
					{
					}
					params["configuration_overrides"] = v
				}
				if d.HasChange("display_name") {
					v := d.Get("display_name")
					{
					}
					params["display_name"] = v
				}
				_, err := f.Post(fmt.Sprintf("/v1/terminal/locations/%v", d.Get("id")), params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to update terminal.location %s", d.Id(), err),
					})
					return out
				}
				return out
			},
		},
		"terminal_readers": &schema.Resource{
			CreateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				{
					v := d.Get("registration_code")
					{
					}
					params["registration_code"] = v
				}
				if v, exists := d.GetOk("label"); exists {
					{
					}
					params["label"] = v
				}
				if v, exists := d.GetOk("location"); exists {
					{
					}
					params["location"] = v
				}
				res, err := f.Post("/v1/terminal/readers", params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to create new terminal.reader", err),
					})
					return out
				}
				d.SetId(Dig(res, "id"))
				return out
			},
			DeleteWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				err := f.Delete(fmt.Sprintf("/v1/terminal/readers/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to delete terminal.reader %s", d.Id(), err),
					})
					return out
				}
				return out
			},
			ReadWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				res, err := f.Get(fmt.Sprintf("/v1/terminal/readers/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to look up terminal.reader %s", d.Id(), err),
					})
					return out
				}
				d.Set("registration_code", Dig[any](res, "registration_code"))
				d.Set("label", Dig[any](res, "label"))
				d.Set("location", Dig[any](res, "location"))
				return out
			},
			Schema: map[string]*schema.Schema{
				"label": &schema.Schema{
					Description: "Custom label given to the reader for easier identification. If no label is specified, the registration code will be used.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"location": &schema.Schema{
					Description: "The location to assign the reader to.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"registration_code": &schema.Schema{
					Description: "A code generated by the reader used for registering to an account.",
					ForceNew:    true,
					Required:    true,
					Type:        schema.TypeString,
				},
			},
			UpdateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if d.HasChange("label") {
					v := d.Get("label")
					{
					}
					params["label"] = v
				}
				_, err := f.Post(fmt.Sprintf("/v1/terminal/readers/%v", d.Get("id")), params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to update terminal.reader %s", d.Id(), err),
					})
					return out
				}
				return out
			},
		},
		"test_helpers_test_clocks": &schema.Resource{
			CreateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				{
					v := d.Get("frozen_time")
					{
					}
					params["frozen_time"] = v
				}
				if v, exists := d.GetOk("name"); exists {
					{
					}
					params["name"] = v
				}
				res, err := f.Post("/v1/test_helpers/test_clocks", params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to create new test_helpers.test_clock", err),
					})
					return out
				}
				d.SetId(Dig(res, "id"))
				return out
			},
			DeleteWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				err := f.Delete(fmt.Sprintf("/v1/test_helpers/test_clocks/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to delete test_helpers.test_clock %s", d.Id(), err),
					})
					return out
				}
				return out
			},
			ReadWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				res, err := f.Get(fmt.Sprintf("/v1/test_helpers/test_clocks/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to look up test_helpers.test_clock %s", d.Id(), err),
					})
					return out
				}
				d.Set("frozen_time", Dig[any](res, "frozen_time"))
				d.Set("name", Dig[any](res, "name"))
				return out
			},
			Schema: map[string]*schema.Schema{
				"frozen_time": &schema.Schema{
					Description: "The initial frozen time for this test clock.",
					ForceNew:    true,
					Required:    true,
					Type:        schema.TypeInt,
				},
				"name": &schema.Schema{
					Description: "The name for this test clock.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
			},
			UpdateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				_, err := f.Post(fmt.Sprintf("/v1/test_helpers/test_clocks/%v", d.Get("id")), params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to update test_helpers.test_clock %s", d.Id(), err),
					})
					return out
				}
				return out
			},
		},
		"webhook_endpoints": &schema.Resource{
			CreateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				{
					v := d.Get("url")
					{
					}
					params["url"] = v
				}
				if v, exists := d.GetOk("api_version"); exists {
					if inEnum(v, []string{"2011-01-01", "2011-06-21", "2011-06-28", "2011-08-01", "2011-09-15", "2011-11-17", "2012-02-23", "2012-03-25", "2012-06-18", "2012-06-28", "2012-07-09", "2012-09-24", "2012-10-26", "2012-11-07", "2013-02-11", "2013-02-13", "2013-07-05", "2013-08-12", "2013-08-13", "2013-10-29", "2013-12-03", "2014-01-31", "2014-03-13", "2014-03-28", "2014-05-19", "2014-06-13", "2014-06-17", "2014-07-22", "2014-07-26", "2014-08-04", "2014-08-20", "2014-09-08", "2014-10-07", "2014-11-05", "2014-11-20", "2014-12-08", "2014-12-17", "2014-12-22", "2015-01-11", "2015-01-26", "2015-02-10", "2015-02-16", "2015-02-18", "2015-03-24", "2015-04-07", "2015-06-15", "2015-07-07", "2015-07-13", "2015-07-28", "2015-08-07", "2015-08-19", "2015-09-03", "2015-09-08", "2015-09-23", "2015-10-01", "2015-10-12", "2015-10-16", "2016-02-03", "2016-02-19", "2016-02-22", "2016-02-23", "2016-02-29", "2016-03-07", "2016-06-15", "2016-07-06", "2016-10-19", "2017-01-27", "2017-02-14", "2017-04-06", "2017-05-25", "2017-06-05", "2017-08-15", "2017-12-14", "2018-01-23", "2018-02-05", "2018-02-06", "2018-02-28", "2018-05-21", "2018-07-27", "2018-08-23", "2018-09-06", "2018-09-24", "2018-10-31", "2018-11-08", "2019-02-11", "2019-02-19", "2019-03-14", "2019-05-16", "2019-08-14", "2019-09-09", "2019-10-08", "2019-10-17", "2019-11-05", "2019-12-03", "2020-03-02", "2020-08-27", "2022-08-01", "2022-11-15"}) {
						out = append(out, diag.Diagnostic{
							AttributePath: path,
							Severity:      diag.Error,
							Summary:       fmt.Sprintf("Field must be one of: 2011-01-01, 2011-06-21, 2011-06-28, 2011-08-01, 2011-09-15, 2011-11-17, 2012-02-23, 2012-03-25, 2012-06-18, 2012-06-28, 2012-07-09, 2012-09-24, 2012-10-26, 2012-11-07, 2013-02-11, 2013-02-13, 2013-07-05, 2013-08-12, 2013-08-13, 2013-10-29, 2013-12-03, 2014-01-31, 2014-03-13, 2014-03-28, 2014-05-19, 2014-06-13, 2014-06-17, 2014-07-22, 2014-07-26, 2014-08-04, 2014-08-20, 2014-09-08, 2014-10-07, 2014-11-05, 2014-11-20, 2014-12-08, 2014-12-17, 2014-12-22, 2015-01-11, 2015-01-26, 2015-02-10, 2015-02-16, 2015-02-18, 2015-03-24, 2015-04-07, 2015-06-15, 2015-07-07, 2015-07-13, 2015-07-28, 2015-08-07, 2015-08-19, 2015-09-03, 2015-09-08, 2015-09-23, 2015-10-01, 2015-10-12, 2015-10-16, 2016-02-03, 2016-02-19, 2016-02-22, 2016-02-23, 2016-02-29, 2016-03-07, 2016-06-15, 2016-07-06, 2016-10-19, 2017-01-27, 2017-02-14, 2017-04-06, 2017-05-25, 2017-06-05, 2017-08-15, 2017-12-14, 2018-01-23, 2018-02-05, 2018-02-06, 2018-02-28, 2018-05-21, 2018-07-27, 2018-08-23, 2018-09-06, 2018-09-24, 2018-10-31, 2018-11-08, 2019-02-11, 2019-02-19, 2019-03-14, 2019-05-16, 2019-08-14, 2019-09-09, 2019-10-08, 2019-10-17, 2019-11-05, 2019-12-03, 2020-03-02, 2020-08-27, 2022-08-01, 2022-11-15"),
						})
					}
					params["api_version"] = v
				}
				if v, exists := d.GetOk("connect"); exists {
					{
					}
					params["connect"] = v
				}
				if v, exists := d.GetOk("description"); exists {
					{
					}
					params["description"] = v
				}
				{
					v := d.Get("enabled_events")
					{
						path := cty.Path{}
						outerV := v
						for i, v := range outerV.([]any) {
							if inEnum(v, []string{"*", "account.application.authorized", "account.application.deauthorized", "account.external_account.created", "account.external_account.deleted", "account.external_account.updated", "account.updated", "application_fee.created", "application_fee.refund.updated", "application_fee.refunded", "balance.available", "billing_portal.configuration.created", "billing_portal.configuration.updated", "billing_portal.session.created", "capability.updated", "cash_balance.funds_available", "charge.captured", "charge.dispute.closed", "charge.dispute.created", "charge.dispute.funds_reinstated", "charge.dispute.funds_withdrawn", "charge.dispute.updated", "charge.expired", "charge.failed", "charge.pending", "charge.refund.updated", "charge.refunded", "charge.succeeded", "charge.updated", "checkout.session.async_payment_failed", "checkout.session.async_payment_succeeded", "checkout.session.completed", "checkout.session.expired", "coupon.created", "coupon.deleted", "coupon.updated", "credit_note.created", "credit_note.updated", "credit_note.voided", "customer.created", "customer.deleted", "customer.discount.created", "customer.discount.deleted", "customer.discount.updated", "customer.source.created", "customer.source.deleted", "customer.source.expiring", "customer.source.updated", "customer.subscription.created", "customer.subscription.deleted", "customer.subscription.pending_update_applied", "customer.subscription.pending_update_expired", "customer.subscription.trial_will_end", "customer.subscription.updated", "customer.tax_id.created", "customer.tax_id.deleted", "customer.tax_id.updated", "customer.updated", "customer_cash_balance_transaction.created", "file.created", "financial_connections.account.created", "financial_connections.account.deactivated", "financial_connections.account.disconnected", "financial_connections.account.reactivated", "financial_connections.account.refreshed_balance", "identity.verification_session.canceled", "identity.verification_session.created", "identity.verification_session.processing", "identity.verification_session.redacted", "identity.verification_session.requires_input", "identity.verification_session.verified", "invoice.created", "invoice.deleted", "invoice.finalization_failed", "invoice.finalized", "invoice.marked_uncollectible", "invoice.paid", "invoice.payment_action_required", "invoice.payment_failed", "invoice.payment_succeeded", "invoice.sent", "invoice.upcoming", "invoice.updated", "invoice.voided", "invoiceitem.created", "invoiceitem.deleted", "invoiceitem.updated", "issuing_authorization.created", "issuing_authorization.request", "issuing_authorization.updated", "issuing_card.created", "issuing_card.updated", "issuing_cardholder.created", "issuing_cardholder.updated", "issuing_dispute.closed", "issuing_dispute.created", "issuing_dispute.funds_reinstated", "issuing_dispute.submitted", "issuing_dispute.updated", "issuing_transaction.created", "issuing_transaction.updated", "mandate.updated", "order.created", "payment_intent.amount_capturable_updated", "payment_intent.canceled", "payment_intent.created", "payment_intent.partially_funded", "payment_intent.payment_failed", "payment_intent.processing", "payment_intent.requires_action", "payment_intent.succeeded", "payment_link.created", "payment_link.updated", "payment_method.attached", "payment_method.automatically_updated", "payment_method.detached", "payment_method.updated", "payout.canceled", "payout.created", "payout.failed", "payout.paid", "payout.updated", "person.created", "person.deleted", "person.updated", "plan.created", "plan.deleted", "plan.updated", "price.created", "price.deleted", "price.updated", "product.created", "product.deleted", "product.updated", "promotion_code.created", "promotion_code.updated", "quote.accepted", "quote.canceled", "quote.created", "quote.finalized", "radar.early_fraud_warning.created", "radar.early_fraud_warning.updated", "recipient.created", "recipient.deleted", "recipient.updated", "reporting.report_run.failed", "reporting.report_run.succeeded", "reporting.report_type.updated", "review.closed", "review.opened", "setup_intent.canceled", "setup_intent.created", "setup_intent.requires_action", "setup_intent.setup_failed", "setup_intent.succeeded", "sigma.scheduled_query_run.created", "sku.created", "sku.deleted", "sku.updated", "source.canceled", "source.chargeable", "source.failed", "source.mandate_notification", "source.refund_attributes_required", "source.transaction.created", "source.transaction.updated", "subscription_schedule.aborted", "subscription_schedule.canceled", "subscription_schedule.completed", "subscription_schedule.created", "subscription_schedule.expiring", "subscription_schedule.released", "subscription_schedule.updated", "tax_rate.created", "tax_rate.updated", "terminal.reader.action_failed", "terminal.reader.action_succeeded", "test_helpers.test_clock.advancing", "test_helpers.test_clock.created", "test_helpers.test_clock.deleted", "test_helpers.test_clock.internal_failure", "test_helpers.test_clock.ready", "topup.canceled", "topup.created", "topup.failed", "topup.reversed", "topup.succeeded", "transfer.created", "transfer.reversed", "transfer.updated", "treasury.credit_reversal.created", "treasury.credit_reversal.posted", "treasury.debit_reversal.completed", "treasury.debit_reversal.created", "treasury.debit_reversal.initial_credit_granted", "treasury.financial_account.closed", "treasury.financial_account.created", "treasury.financial_account.features_status_updated", "treasury.inbound_transfer.canceled", "treasury.inbound_transfer.created", "treasury.inbound_transfer.failed", "treasury.inbound_transfer.succeeded", "treasury.outbound_payment.canceled", "treasury.outbound_payment.created", "treasury.outbound_payment.expected_arrival_date_updated", "treasury.outbound_payment.failed", "treasury.outbound_payment.posted", "treasury.outbound_payment.returned", "treasury.outbound_transfer.canceled", "treasury.outbound_transfer.created", "treasury.outbound_transfer.expected_arrival_date_updated", "treasury.outbound_transfer.failed", "treasury.outbound_transfer.posted", "treasury.outbound_transfer.returned", "treasury.received_credit.created", "treasury.received_credit.failed", "treasury.received_credit.succeeded", "treasury.received_debit.created"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: *, account.application.authorized, account.application.deauthorized, account.external_account.created, account.external_account.deleted, account.external_account.updated, account.updated, application_fee.created, application_fee.refund.updated, application_fee.refunded, balance.available, billing_portal.configuration.created, billing_portal.configuration.updated, billing_portal.session.created, capability.updated, cash_balance.funds_available, charge.captured, charge.dispute.closed, charge.dispute.created, charge.dispute.funds_reinstated, charge.dispute.funds_withdrawn, charge.dispute.updated, charge.expired, charge.failed, charge.pending, charge.refund.updated, charge.refunded, charge.succeeded, charge.updated, checkout.session.async_payment_failed, checkout.session.async_payment_succeeded, checkout.session.completed, checkout.session.expired, coupon.created, coupon.deleted, coupon.updated, credit_note.created, credit_note.updated, credit_note.voided, customer.created, customer.deleted, customer.discount.created, customer.discount.deleted, customer.discount.updated, customer.source.created, customer.source.deleted, customer.source.expiring, customer.source.updated, customer.subscription.created, customer.subscription.deleted, customer.subscription.pending_update_applied, customer.subscription.pending_update_expired, customer.subscription.trial_will_end, customer.subscription.updated, customer.tax_id.created, customer.tax_id.deleted, customer.tax_id.updated, customer.updated, customer_cash_balance_transaction.created, file.created, financial_connections.account.created, financial_connections.account.deactivated, financial_connections.account.disconnected, financial_connections.account.reactivated, financial_connections.account.refreshed_balance, identity.verification_session.canceled, identity.verification_session.created, identity.verification_session.processing, identity.verification_session.redacted, identity.verification_session.requires_input, identity.verification_session.verified, invoice.created, invoice.deleted, invoice.finalization_failed, invoice.finalized, invoice.marked_uncollectible, invoice.paid, invoice.payment_action_required, invoice.payment_failed, invoice.payment_succeeded, invoice.sent, invoice.upcoming, invoice.updated, invoice.voided, invoiceitem.created, invoiceitem.deleted, invoiceitem.updated, issuing_authorization.created, issuing_authorization.request, issuing_authorization.updated, issuing_card.created, issuing_card.updated, issuing_cardholder.created, issuing_cardholder.updated, issuing_dispute.closed, issuing_dispute.created, issuing_dispute.funds_reinstated, issuing_dispute.submitted, issuing_dispute.updated, issuing_transaction.created, issuing_transaction.updated, mandate.updated, order.created, payment_intent.amount_capturable_updated, payment_intent.canceled, payment_intent.created, payment_intent.partially_funded, payment_intent.payment_failed, payment_intent.processing, payment_intent.requires_action, payment_intent.succeeded, payment_link.created, payment_link.updated, payment_method.attached, payment_method.automatically_updated, payment_method.detached, payment_method.updated, payout.canceled, payout.created, payout.failed, payout.paid, payout.updated, person.created, person.deleted, person.updated, plan.created, plan.deleted, plan.updated, price.created, price.deleted, price.updated, product.created, product.deleted, product.updated, promotion_code.created, promotion_code.updated, quote.accepted, quote.canceled, quote.created, quote.finalized, radar.early_fraud_warning.created, radar.early_fraud_warning.updated, recipient.created, recipient.deleted, recipient.updated, reporting.report_run.failed, reporting.report_run.succeeded, reporting.report_type.updated, review.closed, review.opened, setup_intent.canceled, setup_intent.created, setup_intent.requires_action, setup_intent.setup_failed, setup_intent.succeeded, sigma.scheduled_query_run.created, sku.created, sku.deleted, sku.updated, source.canceled, source.chargeable, source.failed, source.mandate_notification, source.refund_attributes_required, source.transaction.created, source.transaction.updated, subscription_schedule.aborted, subscription_schedule.canceled, subscription_schedule.completed, subscription_schedule.created, subscription_schedule.expiring, subscription_schedule.released, subscription_schedule.updated, tax_rate.created, tax_rate.updated, terminal.reader.action_failed, terminal.reader.action_succeeded, test_helpers.test_clock.advancing, test_helpers.test_clock.created, test_helpers.test_clock.deleted, test_helpers.test_clock.internal_failure, test_helpers.test_clock.ready, topup.canceled, topup.created, topup.failed, topup.reversed, topup.succeeded, transfer.created, transfer.reversed, transfer.updated, treasury.credit_reversal.created, treasury.credit_reversal.posted, treasury.debit_reversal.completed, treasury.debit_reversal.created, treasury.debit_reversal.initial_credit_granted, treasury.financial_account.closed, treasury.financial_account.created, treasury.financial_account.features_status_updated, treasury.inbound_transfer.canceled, treasury.inbound_transfer.created, treasury.inbound_transfer.failed, treasury.inbound_transfer.succeeded, treasury.outbound_payment.canceled, treasury.outbound_payment.created, treasury.outbound_payment.expected_arrival_date_updated, treasury.outbound_payment.failed, treasury.outbound_payment.posted, treasury.outbound_payment.returned, treasury.outbound_transfer.canceled, treasury.outbound_transfer.created, treasury.outbound_transfer.expected_arrival_date_updated, treasury.outbound_transfer.failed, treasury.outbound_transfer.posted, treasury.outbound_transfer.returned, treasury.received_credit.created, treasury.received_credit.failed, treasury.received_credit.succeeded, treasury.received_debit.created"),
								})
							}
						}
					}
					params["enabled_events"] = v
				}
				res, err := f.Post("/v1/webhook_endpoints", params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to create new webhook_endpoint", err),
					})
					return out
				}
				d.SetId(Dig(res, "id"))
				return out
			},
			DeleteWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				err := f.Delete(fmt.Sprintf("/v1/webhook_endpoints/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to delete webhook_endpoint %s", d.Id(), err),
					})
					return out
				}
				return out
			},
			ReadWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				res, err := f.Get(fmt.Sprintf("/v1/webhook_endpoints/%v", d.Get("id")))
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to look up webhook_endpoint %s", d.Id(), err),
					})
					return out
				}
				d.Set("url", Dig[any](res, "url"))
				d.Set("api_version", Dig[any](res, "api_version"))
				d.Set("connect", Dig[any](res, "connect"))
				d.Set("description", Dig[any](res, "description"))
				d.Set("enabled_events", Dig[any](res, "enabled_events"))
				return out
			},
			Schema: map[string]*schema.Schema{
				"api_version": &schema.Schema{
					Description: "Events sent to this endpoint will be generated with this Stripe Version instead of your account's default Stripe Version.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeString,
				},
				"connect": &schema.Schema{
					Description: "Whether this endpoint should receive events from connected accounts (`true`), or from your account (`false`). Defaults to `false`.",
					ForceNew:    true,
					Required:    false,
					Type:        schema.TypeBool,
				},
				"description": &schema.Schema{
					Description: "An optional description of what the webhook is used for.",
					ForceNew:    false,
					Required:    false,
					Type:        schema.TypeString,
				},
				"enabled_events": &schema.Schema{
					Description: "The list of events to enable for this endpoint. You may specify `['*']` to enable all events, except those that require explicit selection.",
					Elem:        &schema.Schema{Type: schema.TypeString},
					ForceNew:    false,
					Required:    true,
					Type:        schema.TypeList,
				},
				"url": &schema.Schema{
					Description: "The URL of the webhook endpoint.",
					ForceNew:    false,
					Required:    true,
					Type:        schema.TypeString,
				},
			},
			UpdateWithoutTimeout: func(_ context.Context, d *schema.ResourceData, f any) diag.Diagnostics {
				f := m.(*Facilitator)
				out := diag.Diagnostics{}
				params := map[string]any{}
				if d.HasChange("url") {
					v := d.Get("url")
					{
					}
					params["url"] = v
				}
				if d.HasChange("description") {
					v := d.Get("description")
					{
					}
					params["description"] = v
				}
				if d.HasChange("enabled_events") {
					v := d.Get("enabled_events")
					{
						path := cty.Path{}
						outerV := v
						for i, v := range outerV.([]any) {
							if inEnum(v, []string{"*", "account.application.authorized", "account.application.deauthorized", "account.external_account.created", "account.external_account.deleted", "account.external_account.updated", "account.updated", "application_fee.created", "application_fee.refund.updated", "application_fee.refunded", "balance.available", "billing_portal.configuration.created", "billing_portal.configuration.updated", "billing_portal.session.created", "capability.updated", "cash_balance.funds_available", "charge.captured", "charge.dispute.closed", "charge.dispute.created", "charge.dispute.funds_reinstated", "charge.dispute.funds_withdrawn", "charge.dispute.updated", "charge.expired", "charge.failed", "charge.pending", "charge.refund.updated", "charge.refunded", "charge.succeeded", "charge.updated", "checkout.session.async_payment_failed", "checkout.session.async_payment_succeeded", "checkout.session.completed", "checkout.session.expired", "coupon.created", "coupon.deleted", "coupon.updated", "credit_note.created", "credit_note.updated", "credit_note.voided", "customer.created", "customer.deleted", "customer.discount.created", "customer.discount.deleted", "customer.discount.updated", "customer.source.created", "customer.source.deleted", "customer.source.expiring", "customer.source.updated", "customer.subscription.created", "customer.subscription.deleted", "customer.subscription.pending_update_applied", "customer.subscription.pending_update_expired", "customer.subscription.trial_will_end", "customer.subscription.updated", "customer.tax_id.created", "customer.tax_id.deleted", "customer.tax_id.updated", "customer.updated", "customer_cash_balance_transaction.created", "file.created", "financial_connections.account.created", "financial_connections.account.deactivated", "financial_connections.account.disconnected", "financial_connections.account.reactivated", "financial_connections.account.refreshed_balance", "identity.verification_session.canceled", "identity.verification_session.created", "identity.verification_session.processing", "identity.verification_session.redacted", "identity.verification_session.requires_input", "identity.verification_session.verified", "invoice.created", "invoice.deleted", "invoice.finalization_failed", "invoice.finalized", "invoice.marked_uncollectible", "invoice.paid", "invoice.payment_action_required", "invoice.payment_failed", "invoice.payment_succeeded", "invoice.sent", "invoice.upcoming", "invoice.updated", "invoice.voided", "invoiceitem.created", "invoiceitem.deleted", "invoiceitem.updated", "issuing_authorization.created", "issuing_authorization.request", "issuing_authorization.updated", "issuing_card.created", "issuing_card.updated", "issuing_cardholder.created", "issuing_cardholder.updated", "issuing_dispute.closed", "issuing_dispute.created", "issuing_dispute.funds_reinstated", "issuing_dispute.submitted", "issuing_dispute.updated", "issuing_transaction.created", "issuing_transaction.updated", "mandate.updated", "order.created", "payment_intent.amount_capturable_updated", "payment_intent.canceled", "payment_intent.created", "payment_intent.partially_funded", "payment_intent.payment_failed", "payment_intent.processing", "payment_intent.requires_action", "payment_intent.succeeded", "payment_link.created", "payment_link.updated", "payment_method.attached", "payment_method.automatically_updated", "payment_method.detached", "payment_method.updated", "payout.canceled", "payout.created", "payout.failed", "payout.paid", "payout.updated", "person.created", "person.deleted", "person.updated", "plan.created", "plan.deleted", "plan.updated", "price.created", "price.deleted", "price.updated", "product.created", "product.deleted", "product.updated", "promotion_code.created", "promotion_code.updated", "quote.accepted", "quote.canceled", "quote.created", "quote.finalized", "radar.early_fraud_warning.created", "radar.early_fraud_warning.updated", "recipient.created", "recipient.deleted", "recipient.updated", "reporting.report_run.failed", "reporting.report_run.succeeded", "reporting.report_type.updated", "review.closed", "review.opened", "setup_intent.canceled", "setup_intent.created", "setup_intent.requires_action", "setup_intent.setup_failed", "setup_intent.succeeded", "sigma.scheduled_query_run.created", "sku.created", "sku.deleted", "sku.updated", "source.canceled", "source.chargeable", "source.failed", "source.mandate_notification", "source.refund_attributes_required", "source.transaction.created", "source.transaction.updated", "subscription_schedule.aborted", "subscription_schedule.canceled", "subscription_schedule.completed", "subscription_schedule.created", "subscription_schedule.expiring", "subscription_schedule.released", "subscription_schedule.updated", "tax_rate.created", "tax_rate.updated", "terminal.reader.action_failed", "terminal.reader.action_succeeded", "test_helpers.test_clock.advancing", "test_helpers.test_clock.created", "test_helpers.test_clock.deleted", "test_helpers.test_clock.internal_failure", "test_helpers.test_clock.ready", "topup.canceled", "topup.created", "topup.failed", "topup.reversed", "topup.succeeded", "transfer.created", "transfer.reversed", "transfer.updated", "treasury.credit_reversal.created", "treasury.credit_reversal.posted", "treasury.debit_reversal.completed", "treasury.debit_reversal.created", "treasury.debit_reversal.initial_credit_granted", "treasury.financial_account.closed", "treasury.financial_account.created", "treasury.financial_account.features_status_updated", "treasury.inbound_transfer.canceled", "treasury.inbound_transfer.created", "treasury.inbound_transfer.failed", "treasury.inbound_transfer.succeeded", "treasury.outbound_payment.canceled", "treasury.outbound_payment.created", "treasury.outbound_payment.expected_arrival_date_updated", "treasury.outbound_payment.failed", "treasury.outbound_payment.posted", "treasury.outbound_payment.returned", "treasury.outbound_transfer.canceled", "treasury.outbound_transfer.created", "treasury.outbound_transfer.expected_arrival_date_updated", "treasury.outbound_transfer.failed", "treasury.outbound_transfer.posted", "treasury.outbound_transfer.returned", "treasury.received_credit.created", "treasury.received_credit.failed", "treasury.received_credit.succeeded", "treasury.received_debit.created"}) {
								out = append(out, diag.Diagnostic{
									AttributePath: path,
									Severity:      diag.Error,
									Summary:       fmt.Sprintf("Field must be one of: *, account.application.authorized, account.application.deauthorized, account.external_account.created, account.external_account.deleted, account.external_account.updated, account.updated, application_fee.created, application_fee.refund.updated, application_fee.refunded, balance.available, billing_portal.configuration.created, billing_portal.configuration.updated, billing_portal.session.created, capability.updated, cash_balance.funds_available, charge.captured, charge.dispute.closed, charge.dispute.created, charge.dispute.funds_reinstated, charge.dispute.funds_withdrawn, charge.dispute.updated, charge.expired, charge.failed, charge.pending, charge.refund.updated, charge.refunded, charge.succeeded, charge.updated, checkout.session.async_payment_failed, checkout.session.async_payment_succeeded, checkout.session.completed, checkout.session.expired, coupon.created, coupon.deleted, coupon.updated, credit_note.created, credit_note.updated, credit_note.voided, customer.created, customer.deleted, customer.discount.created, customer.discount.deleted, customer.discount.updated, customer.source.created, customer.source.deleted, customer.source.expiring, customer.source.updated, customer.subscription.created, customer.subscription.deleted, customer.subscription.pending_update_applied, customer.subscription.pending_update_expired, customer.subscription.trial_will_end, customer.subscription.updated, customer.tax_id.created, customer.tax_id.deleted, customer.tax_id.updated, customer.updated, customer_cash_balance_transaction.created, file.created, financial_connections.account.created, financial_connections.account.deactivated, financial_connections.account.disconnected, financial_connections.account.reactivated, financial_connections.account.refreshed_balance, identity.verification_session.canceled, identity.verification_session.created, identity.verification_session.processing, identity.verification_session.redacted, identity.verification_session.requires_input, identity.verification_session.verified, invoice.created, invoice.deleted, invoice.finalization_failed, invoice.finalized, invoice.marked_uncollectible, invoice.paid, invoice.payment_action_required, invoice.payment_failed, invoice.payment_succeeded, invoice.sent, invoice.upcoming, invoice.updated, invoice.voided, invoiceitem.created, invoiceitem.deleted, invoiceitem.updated, issuing_authorization.created, issuing_authorization.request, issuing_authorization.updated, issuing_card.created, issuing_card.updated, issuing_cardholder.created, issuing_cardholder.updated, issuing_dispute.closed, issuing_dispute.created, issuing_dispute.funds_reinstated, issuing_dispute.submitted, issuing_dispute.updated, issuing_transaction.created, issuing_transaction.updated, mandate.updated, order.created, payment_intent.amount_capturable_updated, payment_intent.canceled, payment_intent.created, payment_intent.partially_funded, payment_intent.payment_failed, payment_intent.processing, payment_intent.requires_action, payment_intent.succeeded, payment_link.created, payment_link.updated, payment_method.attached, payment_method.automatically_updated, payment_method.detached, payment_method.updated, payout.canceled, payout.created, payout.failed, payout.paid, payout.updated, person.created, person.deleted, person.updated, plan.created, plan.deleted, plan.updated, price.created, price.deleted, price.updated, product.created, product.deleted, product.updated, promotion_code.created, promotion_code.updated, quote.accepted, quote.canceled, quote.created, quote.finalized, radar.early_fraud_warning.created, radar.early_fraud_warning.updated, recipient.created, recipient.deleted, recipient.updated, reporting.report_run.failed, reporting.report_run.succeeded, reporting.report_type.updated, review.closed, review.opened, setup_intent.canceled, setup_intent.created, setup_intent.requires_action, setup_intent.setup_failed, setup_intent.succeeded, sigma.scheduled_query_run.created, sku.created, sku.deleted, sku.updated, source.canceled, source.chargeable, source.failed, source.mandate_notification, source.refund_attributes_required, source.transaction.created, source.transaction.updated, subscription_schedule.aborted, subscription_schedule.canceled, subscription_schedule.completed, subscription_schedule.created, subscription_schedule.expiring, subscription_schedule.released, subscription_schedule.updated, tax_rate.created, tax_rate.updated, terminal.reader.action_failed, terminal.reader.action_succeeded, test_helpers.test_clock.advancing, test_helpers.test_clock.created, test_helpers.test_clock.deleted, test_helpers.test_clock.internal_failure, test_helpers.test_clock.ready, topup.canceled, topup.created, topup.failed, topup.reversed, topup.succeeded, transfer.created, transfer.reversed, transfer.updated, treasury.credit_reversal.created, treasury.credit_reversal.posted, treasury.debit_reversal.completed, treasury.debit_reversal.created, treasury.debit_reversal.initial_credit_granted, treasury.financial_account.closed, treasury.financial_account.created, treasury.financial_account.features_status_updated, treasury.inbound_transfer.canceled, treasury.inbound_transfer.created, treasury.inbound_transfer.failed, treasury.inbound_transfer.succeeded, treasury.outbound_payment.canceled, treasury.outbound_payment.created, treasury.outbound_payment.expected_arrival_date_updated, treasury.outbound_payment.failed, treasury.outbound_payment.posted, treasury.outbound_payment.returned, treasury.outbound_transfer.canceled, treasury.outbound_transfer.created, treasury.outbound_transfer.expected_arrival_date_updated, treasury.outbound_transfer.failed, treasury.outbound_transfer.posted, treasury.outbound_transfer.returned, treasury.received_credit.created, treasury.received_credit.failed, treasury.received_credit.succeeded, treasury.received_debit.created"),
								})
							}
						}
					}
					params["enabled_events"] = v
				}
				_, err := f.Post(fmt.Sprintf("/v1/webhook_endpoints/%v", d.Get("id")), params)
				if err != nil {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("%s: %w", "failed to update webhook_endpoint %s", d.Id(), err),
					})
					return out
				}
				return out
			},
		},
	}
}
func gitHash() string {
	return "907ce9d"
}
