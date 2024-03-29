---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "stripe_product Resource - stripe"
subcategory: ""
description: |-
  
---

# stripe_product (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The product's name, meant to be displayable to the customer.

### Optional

- `default_price_data_currency` (String)
- `default_price_data_currency_options` (Block List) (see [below for nested schema](#nestedblock--default_price_data_currency_options))
- `default_price_data_recurring_interval` (String)
- `default_price_data_recurring_interval_count` (Number)
- `default_price_data_tax_behavior` (String)
- `default_price_data_unit_amount` (Number)
- `default_price_data_unit_amount_decimal` (String)
- `description` (String) The product's description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.
- `images` (List of String) A list of up to 8 URLs of images for this product, meant to be displayable to the customer.
- `package_dimensions_height` (Number) Height, in inches.
- `package_dimensions_length` (Number) Length, in inches.
- `package_dimensions_weight` (Number) Weight, in ounces.
- `package_dimensions_width` (Number) Width, in inches.
- `shippable` (Boolean) Whether this product is shipped (i.e., physical goods).
- `statement_descriptor` (String) An arbitrary string to be displayed on your customer's credit card or bank statement. While most banks display this information consistently, some may display it incorrectly or not at all.

This may be up to 22 characters. The statement description may not include `<`, `>`, `\`, `"`, `'` characters, and will appear on your customer's statement in capital letters. Non-ASCII characters are automatically stripped.
 It must contain at least one letter.
- `tax_code` (String) A [tax code](https://stripe.com/docs/tax/tax-categories) ID.
- `unit_label` (String) A label that represents units of this product in Stripe and on customers’ receipts and invoices. When set, this will be included in associated invoice line item descriptions.
- `url` (String) A URL of a publicly-accessible webpage for this product.

### Read-Only

- `created` (Number) Time at which the object was created. Measured in seconds since the Unix epoch.
- `default_price` (String) The ID of the [Price](https://stripe.com/docs/api/prices) object that is the default price for this product.
- `id` (String) An identifier will be randomly generated by Stripe. You can optionally override this ID, but the ID must be unique across all products in your Stripe account.
- `livemode` (Boolean) Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
- `object` (String) String representing the object's type. Objects of the same type share the same value.
- `updated` (Number) Time at which the object was last updated. Measured in seconds since the Unix epoch.

<a id="nestedblock--default_price_data_currency_options"></a>
### Nested Schema for `default_price_data_currency_options`

Required:

- `key` (String) Key for this field in parent map (synthetic to work around Terraform limitations)

Optional:

- `custom_unit_amount_enabled` (Boolean)
- `custom_unit_amount_maximum` (Number)
- `custom_unit_amount_minimum` (Number)
- `custom_unit_amount_preset` (Number)
- `tax_behavior` (String)
- `tiers` (Block List) (see [below for nested schema](#nestedblock--default_price_data_currency_options--tiers))
- `unit_amount` (Number)
- `unit_amount_decimal` (String)

<a id="nestedblock--default_price_data_currency_options--tiers"></a>
### Nested Schema for `default_price_data_currency_options.tiers`

Required:

- `up_to` (String)

Optional:

- `flat_amount` (Number)
- `flat_amount_decimal` (String)
- `unit_amount` (Number)
- `unit_amount_decimal` (String)


