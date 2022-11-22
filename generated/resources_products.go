package main

import (
	"context"
	"fmt"
	shared "github.com/andrewbaxter/terraform-provider-stripe/shared"
	cty "github.com/hashicorp/go-cty/cty"
	diag "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resources_products() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if v, exists := d.GetOk("id"); exists {
				params["id"] = v
			}
			{
				v := d.Get("name")
				params["name"] = v
			}
			if v, exists := d.GetOk("package_dimensions"); exists {
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("height")
						v := shared.Dig[any](outerV, "height")
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
						path := path.IndexString("length")
						v := shared.Dig[any](outerV, "length")
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
						path := path.IndexString("weight")
						v := shared.Dig[any](outerV, "weight")
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
						path := path.IndexString("width")
						v := shared.Dig[any](outerV, "width")
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
				params["package_dimensions"] = v
			}
			if v, exists := d.GetOk("images"); exists {
				params["images"] = v
			}
			if v, exists := d.GetOk("default_price_data"); exists {
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
				params["default_price_data"] = v
			}
			if v, exists := d.GetOk("url"); exists {
				params["url"] = v
			}
			if v, exists := d.GetOk("tax_code"); exists {
				params["tax_code"] = v
			}
			if v, exists := d.GetOk("unit_label"); exists {
				params["unit_label"] = v
			}
			if v, exists := d.GetOk("description"); exists {
				params["description"] = v
			}
			if v, exists := d.GetOk("shippable"); exists {
				params["shippable"] = v
			}
			if v, exists := d.GetOk("statement_descriptor"); exists {
				params["statement_descriptor"] = v
			}
			res, err := f.Post(ctx, "/v1/products", params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to create new product: %s", err),
				})
				return out
			}
			d.SetId(shared.Dig[string](res, "id"))
			return out
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			err := f.Delete(ctx, fmt.Sprintf("/v1/products/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to delete product %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			res, err := f.Get(ctx, fmt.Sprintf("/v1/products/%v", d.Get("id")))
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to look up product %s: %s", d.Id(), err),
				})
				return out
			}
			d.Set("id", shared.Dig[any](res, "id"))
			d.Set("name", shared.Dig[any](res, "name"))
			d.Set("package_dimensions", shared.Dig[any](res, "package_dimensions"))
			d.Set("images", shared.Dig[any](res, "images"))
			d.Set("default_price_data", shared.Dig[any](res, "default_price_data"))
			d.Set("url", shared.Dig[any](res, "url"))
			d.Set("tax_code", shared.Dig[any](res, "tax_code"))
			d.Set("unit_label", shared.Dig[any](res, "unit_label"))
			d.Set("description", shared.Dig[any](res, "description"))
			d.Set("shippable", shared.Dig[any](res, "shippable"))
			d.Set("statement_descriptor", shared.Dig[any](res, "statement_descriptor"))
			return out
		},
		Schema: map[string]*schema.Schema{
			"default_price_data": {
				Description: "Data used to generate a new [Price](https://stripe.com/docs/api/prices) object. This Price will be set as the default price for this product.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"currency": {
						Description: "",
						Required:    true,
						Type:        schema.TypeString,
					},
					"currency_options": {
						Description: "",
						Elem:        &schema.Resource{Schema: map[string]*schema.Schema{}},
						Required:    false,
						Type:        schema.TypeMap,
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
						Required: false,
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
				ForceNew: true,
				Required: false,
				Type:     schema.TypeMap,
			},
			"description": {
				Description: "The product's description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"id": {
				Description: "An identifier will be randomly generated by Stripe. You can optionally override this ID, but the ID must be unique across all products in your Stripe account.",
				ForceNew:    true,
				Required:    false,
				Type:        schema.TypeString,
			},
			"images": {
				Description: "A list of up to 8 URLs of images for this product, meant to be displayable to the customer.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeList,
			},
			"name": {
				Description: "The product's name, meant to be displayable to the customer.",
				ForceNew:    false,
				Required:    true,
				Type:        schema.TypeString,
			},
			"package_dimensions": {
				Description: "The dimensions of this product for shipping purposes.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"height": {
						Description: "",
						Required:    true,
						Type:        schema.TypeFloat,
					},
					"length": {
						Description: "",
						Required:    true,
						Type:        schema.TypeFloat,
					},
					"weight": {
						Description: "",
						Required:    true,
						Type:        schema.TypeFloat,
					},
					"width": {
						Description: "",
						Required:    true,
						Type:        schema.TypeFloat,
					},
				}},
				ForceNew: false,
				Required: false,
				Type:     schema.TypeMap,
			},
			"shippable": {
				Description: "Whether this product is shipped (i.e., physical goods).",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeBool,
			},
			"statement_descriptor": {
				Description: "An arbitrary string to be displayed on your customer's credit card or bank statement. While most banks display this information consistently, some may display it incorrectly or not at all.\n\nThis may be up to 22 characters. The statement description may not include `<`, `>`, `\\`, `\"`, `'` characters, and will appear on your customer's statement in capital letters. Non-ASCII characters are automatically stripped.\n It must contain at least one letter.",
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
			"unit_label": {
				Description: "A label that represents units of this product in Stripe and on customers’ receipts and invoices. When set, this will be included in associated invoice line item descriptions.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
			"url": {
				Description: "A URL of a publicly-accessible webpage for this product.",
				ForceNew:    false,
				Required:    false,
				Type:        schema.TypeString,
			},
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, f0 any) diag.Diagnostics {
			f := f0.(*Facilitator)
			out := diag.Diagnostics{}
			params := map[string]any{}
			if d.HasChange("name") {
				v := d.Get("name")
				params["name"] = v
			}
			if d.HasChange("package_dimensions") {
				v := d.Get("package_dimensions")
				{
					path := cty.Path{}
					outerV := v
					{
						path := path.IndexString("height")
						v := shared.Dig[any](outerV, "height")
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
						path := path.IndexString("length")
						v := shared.Dig[any](outerV, "length")
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
						path := path.IndexString("weight")
						v := shared.Dig[any](outerV, "weight")
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
						path := path.IndexString("width")
						v := shared.Dig[any](outerV, "width")
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
				params["package_dimensions"] = v
			}
			if d.HasChange("images") {
				v := d.Get("images")
				params["images"] = v
			}
			if d.HasChange("url") {
				v := d.Get("url")
				params["url"] = v
			}
			if d.HasChange("tax_code") {
				v := d.Get("tax_code")
				params["tax_code"] = v
			}
			if d.HasChange("unit_label") {
				v := d.Get("unit_label")
				params["unit_label"] = v
			}
			if d.HasChange("description") {
				v := d.Get("description")
				params["description"] = v
			}
			if d.HasChange("shippable") {
				v := d.Get("shippable")
				params["shippable"] = v
			}
			if d.HasChange("statement_descriptor") {
				v := d.Get("statement_descriptor")
				params["statement_descriptor"] = v
			}
			_, err := f.Post(ctx, fmt.Sprintf("/v1/products/%v", d.Get("id")), params)
			if err != nil {
				out = append(out, diag.Diagnostic{
					AttributePath: cty.Path{},
					Severity:      diag.Error,
					Summary:       fmt.Sprintf("failed to update product %s: %s", d.Id(), err),
				})
				return out
			}
			return out
		},
	}
}
