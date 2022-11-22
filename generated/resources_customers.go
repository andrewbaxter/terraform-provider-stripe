package main

import (
	"context"
	"fmt"
	shared "github.com/andrewbaxter/terraform-provider-stripe/shared"
	cty "github.com/hashicorp/go-cty/cty"
	diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resources_customers() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if v, exists := d.GetOk("address"); exists {
				params["address"] = v
			}
			if v, exists := d.GetOk("balance"); exists {
				params["balance"] = v
			}
			if v, exists := d.GetOk("cash_balance"); exists {
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("settings")
						v := shared.Dig[any](outerV, "settings")
						if v != nil {
							{
								path := path
								outerV := v
								{
									path := path.IndexString("reconciliation_mode")
									v := shared.Dig[any](outerV, "reconciliation_mode")
									if v != nil {
										if inEnum(v.(string), []string{"automatic", "manual"}) {
											out = append(out, diag.Diagnostic{
												AttributePath: path,
												Severity:      diag.Error,
												Summary:       fmt.Sprintf("Field must be one of: automatic, manual"),
											})
										}
									}
								}
							}
						}
					}
				}
				params["cash_balance"] = v
			}
			if v, exists := d.GetOk("name"); exists {
				params["name"] = v
			}
			if v, exists := d.GetOk("email"); exists {
				params["email"] = v
			}
			if v, exists := d.GetOk("preferred_locales"); exists {
				params["preferred_locales"] = v
			}
			if v, exists := d.GetOk("tax"); exists {
				params["tax"] = v
			}
			if v, exists := d.GetOk("coupon"); exists {
				params["coupon"] = v
			}
			if v, exists := d.GetOk("description"); exists {
				params["description"] = v
			}
			if v, exists := d.GetOk("invoice_prefix"); exists {
				params["invoice_prefix"] = v
			}
			if v, exists := d.GetOk("next_invoice_sequence"); exists {
				params["next_invoice_sequence"] = v
			}
			if v, exists := d.GetOk("payment_method"); exists {
				params["payment_method"] = v
			}
			if v, exists := d.GetOk("promotion_code"); exists {
				params["promotion_code"] = v
			}
			if v, exists := d.GetOk("tax_exempt"); exists {
				if inEnum(v.(string), []string{"", "exempt", "none", "reverse"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: , exempt, none, reverse"),
					})
				}
				params["tax_exempt"] = v
			}
			if v, exists := d.GetOk("invoice_settings"); exists {
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("custom_fields")
						v := shared.Dig[any](outerV, "custom_fields")
						if v != nil {
							{
								path := path
								outerV := v
								for i, v := range outerV.([]any) {
									{
										path := path.IndexInt(i)
										outerV := v
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
								}
							}
						}
					}
					{
						path := path.IndexString("rendering_options")
						v := shared.Dig[any](outerV, "rendering_options")
						if v != nil {
							{
								path := path
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
						}
					}
				}
				params["invoice_settings"] = v
			}
			if v, exists := d.GetOk("phone"); exists {
				params["phone"] = v
			}
			if v, exists := d.GetOk("shipping"); exists {
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("address")
						v := shared.Dig[any](outerV, "address")
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
				params["shipping"] = v
			}
			if v, exists := d.GetOk("source"); exists {
				params["source"] = v
			}
			if v, exists := d.GetOk("tax_id_data"); exists {
				{
					path := cty.Path{}
					outerV := v
					for i, v := range outerV.([]any) {
						{
							path := path.IndexInt(i)
							outerV := v
							{
								path := path.IndexString("type")
								v := shared.Dig[any](outerV, "type")
								if v != nil {
									if inEnum(v.(string), []string{"ae_trn", "au_abn", "au_arn", "bg_uic", "br_cnpj", "br_cpf", "ca_bn", "ca_gst_hst", "ca_pst_bc", "ca_pst_mb", "ca_pst_sk", "ca_qst", "ch_vat", "cl_tin", "eg_tin", "es_cif", "eu_oss_vat", "eu_vat", "gb_vat", "ge_vat", "hk_br", "hu_tin", "id_npwp", "il_vat", "in_gst", "is_vat", "jp_cn", "jp_rn", "jp_trn", "ke_pin", "kr_brn", "li_uid", "mx_rfc", "my_frp", "my_itn", "my_sst", "no_vat", "nz_gst", "ph_tin", "ru_inn", "ru_kpp", "sa_vat", "sg_gst", "sg_uen", "si_tin", "th_vat", "tr_tin", "tw_vat", "ua_vat", "us_ein", "za_vat"}) {
										out = append(out, diag.Diagnostic{
											AttributePath: path,
											Severity:      diag.Error,
											Summary:       fmt.Sprintf("Field must be one of: ae_trn, au_abn, au_arn, bg_uic, br_cnpj, br_cpf, ca_bn, ca_gst_hst, ca_pst_bc, ca_pst_mb, ca_pst_sk, ca_qst, ch_vat, cl_tin, eg_tin, es_cif, eu_oss_vat, eu_vat, gb_vat, ge_vat, hk_br, hu_tin, id_npwp, il_vat, in_gst, is_vat, jp_cn, jp_rn, jp_trn, ke_pin, kr_brn, li_uid, mx_rfc, my_frp, my_itn, my_sst, no_vat, nz_gst, ph_tin, ru_inn, ru_kpp, sa_vat, sg_gst, sg_uen, si_tin, th_vat, tr_tin, tw_vat, ua_vat, us_ein, za_vat"),
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
				params["tax_id_data"] = v
			}
			if v, exists := d.GetOk("test_clock"); exists {
				params["test_clock"] = v
			}
			res, err := f.Post(ctx, "/v1/customers", params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to create new customer: %s", err),
				})
				return out
			}
			d.SetId(shared.Dig[string](res, "id"))
			return out
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			err := f.Delete(ctx, fmt.Sprintf("/v1/customers/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to delete customer %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			res, err := f.Get(ctx, fmt.Sprintf("/v1/customers/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to look up customer %s: %s", d.Id(), err),
				})
				return out
			}
			d.Set("address", shared.Dig[any](res, "address"))
			d.Set("balance", shared.Dig[any](res, "balance"))
			d.Set("cash_balance", shared.Dig[any](res, "cash_balance"))
			d.Set("name", shared.Dig[any](res, "name"))
			d.Set("email", shared.Dig[any](res, "email"))
			d.Set("preferred_locales", shared.Dig[any](res, "preferred_locales"))
			d.Set("tax", shared.Dig[any](res, "tax"))
			d.Set("coupon", shared.Dig[any](res, "coupon"))
			d.Set("description", shared.Dig[any](res, "description"))
			d.Set("invoice_prefix", shared.Dig[any](res, "invoice_prefix"))
			d.Set("next_invoice_sequence", shared.Dig[any](res, "next_invoice_sequence"))
			d.Set("payment_method", shared.Dig[any](res, "payment_method"))
			d.Set("promotion_code", shared.Dig[any](res, "promotion_code"))
			d.Set("tax_exempt", shared.Dig[any](res, "tax_exempt"))
			d.Set("invoice_settings", shared.Dig[any](res, "invoice_settings"))
			d.Set("phone", shared.Dig[any](res, "phone"))
			d.Set("shipping", shared.Dig[any](res, "shipping"))
			d.Set("source", shared.Dig[any](res, "source"))
			d.Set("tax_id_data", shared.Dig[any](res, "tax_id_data"))
			d.Set("test_clock", shared.Dig[any](res, "test_clock"))
			return out
		},
		Schema: map[string]*schema.Schema{
			"address": {
				Description: "The customer's address.",
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
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"balance": {
				Description: "An integer amount in cents (or local equivalent) that represents the customer's current balance, which affect the customer's future invoices. A negative amount represents a credit that decreases the amount due on an invoice; a positive amount increases the amount due on an invoice.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"cash_balance": {
				Description: "Balance information and default balance settings for this customer.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"settings": {
					Description: "",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{"reconciliation_mode": {
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
			"coupon": {
				Description: "",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"description": {
				Description: "An arbitrary string that you can attach to a customer object. It is displayed alongside the customer in the dashboard.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"email": {
				Description: "Customer's email address. It's displayed alongside the customer in your dashboard and can be useful for searching and tracking. This may be up to *512 characters*.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"invoice_prefix": {
				Description: "The prefix for the customer used to generate unique invoice numbers. Must be 3–12 uppercase letters or numbers.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"invoice_settings": {
				Description: "Default invoice settings for this customer.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"custom_fields": {
						Description: "",
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
						Required: false,
						Type:     schema.TypeList,
					},
					"default_payment_method": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"footer": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
					"rendering_options": {
						Description: "",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"amount_tax_display": {
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
			"name": {
				Description: "The customer's full name or business name.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"next_invoice_sequence": {
				Description: "The sequence to be used on the customer's next invoice. Defaults to 1.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeInt,
			},
			"payment_method": {
				Description: "",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"phone": {
				Description: "The customer's phone number.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"preferred_locales": {
				Description: "Customer's preferred languages, ordered by preference.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeList,
			},
			"promotion_code": {
				Description: "The API ID of a promotion code to apply to the customer. The customer will have a discount applied on all recurring payments. Charges you create through the API will not have the discount.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"shipping": {
				Description: "The customer's shipping information. Appears on invoices emailed to this customer.",
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
						Required: true,
						Type:     schema.TypeMap,
					},
					"name": {
						Description: "",
						Required:    true,
						Type:        schema.TypeString,
					},
					"phone": {
						Description: "",
						Required:    false,
						Type:        schema.TypeString,
					},
				}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"source": {
				Description: "",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"tax": {
				Description: "Tax details about the customer.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"ip_address": {
					Description: "",
					Required:    false,
					Type:        schema.TypeString,
				}}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"tax_exempt": {
				Description: "The customer's tax exemption. One of `none`, `exempt`, or `reverse`.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"tax_id_data": {
				Description: "The customer's tax IDs.",
				Elem: &schema.Schema{
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"type": {
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
				ForceNew: true,
				Required: false,
				Type:     schema.TypeList,
			},
			"test_clock": {
				Description: "ID of the test clock to attach to the customer.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if d.HasChange("address") {
				v := d.Get("address")
				params["address"] = v
			}
			if d.HasChange("balance") {
				v := d.Get("balance")
				params["balance"] = v
			}
			if d.HasChange("cash_balance") {
				v := d.Get("cash_balance")
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("settings")
						v := shared.Dig[any](outerV, "settings")
						if v != nil {
							{
								path := path
								outerV := v
								{
									path := path.IndexString("reconciliation_mode")
									v := shared.Dig[any](outerV, "reconciliation_mode")
									if v != nil {
										if inEnum(v.(string), []string{"automatic", "manual"}) {
											out = append(out, diag.Diagnostic{
												AttributePath: path,
												Severity:      diag.Error,
												Summary:       fmt.Sprintf("Field must be one of: automatic, manual"),
											})
										}
									}
								}
							}
						}
					}
				}
				params["cash_balance"] = v
			}
			if d.HasChange("name") {
				v := d.Get("name")
				params["name"] = v
			}
			if d.HasChange("email") {
				v := d.Get("email")
				params["email"] = v
			}
			if d.HasChange("preferred_locales") {
				v := d.Get("preferred_locales")
				params["preferred_locales"] = v
			}
			if d.HasChange("tax") {
				v := d.Get("tax")
				params["tax"] = v
			}
			if d.HasChange("coupon") {
				v := d.Get("coupon")
				params["coupon"] = v
			}
			if d.HasChange("description") {
				v := d.Get("description")
				params["description"] = v
			}
			if d.HasChange("invoice_prefix") {
				v := d.Get("invoice_prefix")
				params["invoice_prefix"] = v
			}
			if d.HasChange("next_invoice_sequence") {
				v := d.Get("next_invoice_sequence")
				params["next_invoice_sequence"] = v
			}
			if d.HasChange("promotion_code") {
				v := d.Get("promotion_code")
				params["promotion_code"] = v
			}
			if d.HasChange("tax_exempt") {
				v := d.Get("tax_exempt")
				if inEnum(v.(string), []string{"", "exempt", "none", "reverse"}) {
					out = append(out, diag.Diagnostic{
						AttributePath: cty.Path{},
						Severity:      diag.Error,
						Summary:       fmt.Sprintf("Field must be one of: , exempt, none, reverse"),
					})
				}
				params["tax_exempt"] = v
			}
			if d.HasChange("invoice_settings") {
				v := d.Get("invoice_settings")
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("custom_fields")
						v := shared.Dig[any](outerV, "custom_fields")
						if v != nil {
							{
								path := path
								outerV := v
								for i, v := range outerV.([]any) {
									{
										path := path.IndexInt(i)
										outerV := v
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
								}
							}
						}
					}
					{
						path := path.IndexString("rendering_options")
						v := shared.Dig[any](outerV, "rendering_options")
						if v != nil {
							{
								path := path
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
						}
					}
				}
				params["invoice_settings"] = v
			}
			if d.HasChange("phone") {
				v := d.Get("phone")
				params["phone"] = v
			}
			if d.HasChange("shipping") {
				v := d.Get("shipping")
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("address")
						v := shared.Dig[any](outerV, "address")
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
				params["shipping"] = v
			}
			if d.HasChange("source") {
				v := d.Get("source")
				params["source"] = v
			}
			_, err := f.Post(ctx, fmt.Sprintf("/v1/customers/%v", d.Get("id")), params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to update customer %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
	}
}
