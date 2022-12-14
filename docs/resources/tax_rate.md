---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "stripe_tax_rate Resource - stripe"
subcategory: ""
description: |-
  
---

# stripe_tax_rate (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `display_name` (String) The display name of the tax rate, which will be shown to users.
- `inclusive` (Boolean) This specifies if the tax rate is inclusive or exclusive.
- `percentage` (Number) This represents the tax rate percent out of 100.

### Optional

- `country` (String) Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
- `description` (String) An arbitrary string attached to the tax rate for your internal use only. It will not be visible to your customers.
- `jurisdiction` (String) The jurisdiction for the tax rate. You can use this label field for tax reporting purposes. It also appears on your customer’s invoice.
- `state` (String) [ISO 3166-2 subdivision code](https://en.wikipedia.org/wiki/ISO_3166-2:US), without country prefix. For example, "NY" for New York, United States.
- `tax_type` (String) The high-level tax type, such as `vat` or `sales_tax`.

### Read-Only

- `created` (Number) Time at which the object was created. Measured in seconds since the Unix epoch.
- `id` (String) Unique identifier for the object.
- `livemode` (Boolean) Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
- `object` (String) String representing the object's type. Objects of the same type share the same value.


