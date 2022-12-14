---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "stripe_customer Resource - stripe"
subcategory: ""
description: |-
  
---

# stripe_customer (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `address_city` (String) City, district, suburb, town, or village.
- `address_country` (String) Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
- `address_line1` (String) Address line 1 (e.g., street, PO Box, or company name).
- `address_line2` (String) Address line 2 (e.g., apartment, suite, unit, or building).
- `address_postal_code` (String) ZIP or postal code.
- `address_state` (String) State, county, province, or region.
- `balance` (Number) An integer amount in cents (or local equivalent) that represents the customer's current balance, which affect the customer's future invoices. A negative amount represents a credit that decreases the amount due on an invoice; a positive amount increases the amount due on an invoice.
- `cash_balance_settings_reconciliation_mode` (String) The configuration for how funds that land in the customer cash balance are reconciled.
- `coupon` (String)
- `description` (String) An arbitrary string that you can attach to a customer object. It is displayed alongside the customer in the dashboard.
- `email` (String) Customer's email address. It's displayed alongside the customer in your dashboard and can be useful for searching and tracking. This may be up to *512 characters*.
- `invoice_prefix` (String) The prefix for the customer used to generate unique invoice numbers. Must be 3–12 uppercase letters or numbers.
- `invoice_settings_custom_fields` (Block List) Default custom fields to be displayed on invoices for this customer. (see [below for nested schema](#nestedblock--invoice_settings_custom_fields))
- `invoice_settings_default_payment_method` (String) ID of a payment method that's attached to the customer, to be used as the customer's default payment method for subscriptions and invoices.
- `invoice_settings_footer` (String) Default footer to be displayed on invoices for this customer.
- `invoice_settings_rendering_options_amount_tax_display` (String) How line-item prices and amounts will be displayed with respect to tax on invoice PDFs.
- `name` (String) The customer's full name or business name.
- `next_invoice_sequence` (Number) The sequence to be used on the customer's next invoice. Defaults to 1.
- `payment_method` (String)
- `phone` (String) The customer's phone number.
- `preferred_locales` (List of String) Customer's preferred languages, ordered by preference.
- `promotion_code` (String) The API ID of a promotion code to apply to the customer. The customer will have a discount applied on all recurring payments. Charges you create through the API will not have the discount.
- `shipping_address_city` (String) City, district, suburb, town, or village.
- `shipping_address_country` (String) Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
- `shipping_address_line1` (String) Address line 1 (e.g., street, PO Box, or company name).
- `shipping_address_line2` (String) Address line 2 (e.g., apartment, suite, unit, or building).
- `shipping_address_postal_code` (String) ZIP or postal code.
- `shipping_address_state` (String) State, county, province, or region.
- `shipping_name` (String) Recipient name.
- `shipping_phone` (String) Recipient phone (including extension).
- `source` (String)
- `tax_exempt` (String) The customer's tax exemption. One of `none`, `exempt`, or `reverse`.
- `tax_id_data` (Block List) The customer's tax IDs. (see [below for nested schema](#nestedblock--tax_id_data))
- `tax_ip_address` (String) A recent IP address of the customer used for tax reporting and tax location inference.
- `test_clock` (String) ID of the test clock to attach to the customer.

### Read-Only

- `cash_balance_available` (Map of Number) A hash of all cash balances available to this customer. You cannot delete a customer with any cash balances, even if the balance is 0. Amounts are represented in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
- `cash_balance_customer` (String) The ID of the customer whose cash balance this object represents.
- `cash_balance_livemode` (Boolean) Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
- `cash_balance_object` (String) String representing the object's type. Objects of the same type share the same value.
- `created` (Number) Time at which the object was created. Measured in seconds since the Unix epoch.
- `currency` (String) Three-letter [ISO code for the currency](https://stripe.com/docs/currencies) the customer can be charged in for recurring billing purposes.
- `default_source` (String) ID of the default payment source for the customer.

If you are using payment methods created via the PaymentMethods API, see the [invoice_settings.default_payment_method](https://stripe.com/docs/api/customers/object#customer_object-invoice_settings-default_payment_method) field instead.
- `delinquent` (Boolean) When the customer's latest invoice is billed by charging automatically, `delinquent` is `true` if the invoice's latest charge failed. When the customer's latest invoice is billed by sending an invoice, `delinquent` is `true` if the invoice isn't paid by its due date.

If an invoice is marked uncollectible by [dunning](https://stripe.com/docs/billing/automatic-collection), `delinquent` doesn't get reset to `false`.
- `discount_checkout_session` (String) The Checkout session that this coupon is applied to, if it is applied to a particular session in payment mode. Will not be present for subscription mode.
- `discount_coupon_amount_off` (Number) Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.
- `discount_coupon_applies_to_products` (List of String) A list of product IDs this coupon applies to
- `discount_coupon_created` (Number) Time at which the object was created. Measured in seconds since the Unix epoch.
- `discount_coupon_currency` (String) If `amount_off` has been set, the three-letter [ISO code for the currency](https://stripe.com/docs/currencies) of the amount to take off.
- `discount_coupon_currency_options` (List of Object) Coupons defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies). (see [below for nested schema](#nestedatt--discount_coupon_currency_options))
- `discount_coupon_duration` (String) One of `forever`, `once`, and `repeating`. Describes how long a customer who applies this coupon will get the discount.
- `discount_coupon_duration_in_months` (Number) If `duration` is `repeating`, the number of months the coupon applies. Null if coupon `duration` is `forever` or `once`.
- `discount_coupon_id` (String) Unique identifier for the object.
- `discount_coupon_livemode` (Boolean) Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
- `discount_coupon_max_redemptions` (Number) Maximum number of times this coupon can be redeemed, in total, across all customers, before it is no longer valid.
- `discount_coupon_metadata` (Map of String) Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
- `discount_coupon_name` (String) Name of the coupon displayed to customers on for instance invoices or receipts.
- `discount_coupon_object` (String) String representing the object's type. Objects of the same type share the same value.
- `discount_coupon_percent_off` (Number) Percent that will be taken off the subtotal of any invoices for this customer for the duration of the coupon. For example, a coupon with percent_off of 50 will make a %s100 invoice %s50 instead.
- `discount_coupon_redeem_by` (Number) Date after which the coupon can no longer be redeemed.
- `discount_coupon_times_redeemed` (Number) Number of times this coupon has been applied to a customer.
- `discount_coupon_valid` (Boolean) Taking account of the above properties, whether this coupon can still be applied to a customer.
- `discount_customer` (String) The ID of the customer associated with this discount.
- `discount_end` (Number) If the coupon has a duration of `repeating`, the date that this discount will end. If the coupon has a duration of `once` or `forever`, this attribute will be null.
- `discount_id` (String) The ID of the discount object. Discounts cannot be fetched by ID. Use `expand[]=discounts` in API calls to expand discount IDs in an array.
- `discount_invoice` (String) The invoice that the discount's coupon was applied to, if it was applied directly to a particular invoice.
- `discount_invoice_item` (String) The invoice item `id` (or invoice line item `id` for invoice line items of type='subscription') that the discount's coupon was applied to, if it was applied directly to a particular invoice item or invoice line item.
- `discount_object` (String) String representing the object's type. Objects of the same type share the same value.
- `discount_promotion_code` (String) The promotion code applied to create this discount.
- `discount_start` (Number) Date that the coupon was applied.
- `discount_subscription` (String) The subscription that this coupon is applied to, if it is applied to a particular subscription.
- `id` (String) Unique identifier for the object.
- `invoice_credit_balance` (Map of Number) The current multi-currency balances, if any, being stored on the customer. If positive in a currency, the customer has a credit to apply to their next invoice denominated in that currency. If negative, the customer has an amount owed that will be added to their next invoice denominated in that currency. These balances do not refer to any unpaid invoices. They solely track amounts that have yet to be successfully applied to any invoice. A balance in a particular currency is only applied to any invoice as an invoice in that currency is finalized.
- `livemode` (Boolean) Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
- `object` (String) String representing the object's type. Objects of the same type share the same value.
- `shipping_carrier` (String) The delivery service that shipped a physical product, such as Fedex, UPS, USPS, etc.
- `shipping_tracking_number` (String) The tracking number for a physical product, obtained from the delivery service. If multiple tracking numbers were generated for this purchase, please separate them with commas.
- `sources_has_more` (Boolean) True if this list has another page of items after this one that can be fetched.
- `sources_object` (String) String representing the object's type. Objects of the same type share the same value. Always has the value `list`.
- `sources_url` (String) The URL where this list can be accessed.
- `subscriptions_data` (List of Object) Details about each object. (see [below for nested schema](#nestedatt--subscriptions_data))
- `subscriptions_has_more` (Boolean) True if this list has another page of items after this one that can be fetched.
- `subscriptions_object` (String) String representing the object's type. Objects of the same type share the same value. Always has the value `list`.
- `subscriptions_url` (String) The URL where this list can be accessed.
- `tax_automatic_tax` (String) Surfaces if automatic tax computation is possible given the current customer location information.
- `tax_ids_data` (List of Object) Details about each object. (see [below for nested schema](#nestedatt--tax_ids_data))
- `tax_ids_has_more` (Boolean) True if this list has another page of items after this one that can be fetched.
- `tax_ids_object` (String) String representing the object's type. Objects of the same type share the same value. Always has the value `list`.
- `tax_ids_url` (String) The URL where this list can be accessed.
- `tax_location_country` (String) The customer's country as identified by Stripe Tax.
- `tax_location_source` (String) The data source used to infer the customer's location.
- `tax_location_state` (String) The customer's state, county, province, or region as identified by Stripe Tax.

<a id="nestedblock--invoice_settings_custom_fields"></a>
### Nested Schema for `invoice_settings_custom_fields`

Required:

- `name` (String) The name of the custom field.
- `value` (String) The value of the custom field.


<a id="nestedblock--tax_id_data"></a>
### Nested Schema for `tax_id_data`

Required:

- `type` (String)
- `value` (String)


<a id="nestedatt--discount_coupon_currency_options"></a>
### Nested Schema for `discount_coupon_currency_options`

Read-Only:

- `amount_off` (Number)
- `key` (String)


<a id="nestedatt--subscriptions_data"></a>
### Nested Schema for `subscriptions_data`

Read-Only:

- `application` (String)
- `application_fee_percent` (Number)
- `automatic_tax_enabled` (Boolean)
- `billing_cycle_anchor` (Number)
- `billing_thresholds_amount_gte` (Number)
- `billing_thresholds_reset_billing_cycle_anchor` (Boolean)
- `cancel_at` (Number)
- `cancel_at_period_end` (Boolean)
- `canceled_at` (Number)
- `collection_method` (String)
- `created` (Number)
- `currency` (String)
- `current_period_end` (Number)
- `current_period_start` (Number)
- `customer` (String)
- `days_until_due` (Number)
- `default_payment_method` (String)
- `default_source` (String)
- `default_tax_rates` (List of Object) (see [below for nested schema](#nestedobjatt--subscriptions_data--default_tax_rates))
- `description` (String)
- `discount_checkout_session` (String)
- `discount_coupon_amount_off` (Number)
- `discount_coupon_applies_to_products` (List of String)
- `discount_coupon_created` (Number)
- `discount_coupon_currency` (String)
- `discount_coupon_currency_options` (List of Object) (see [below for nested schema](#nestedobjatt--subscriptions_data--discount_coupon_currency_options))
- `discount_coupon_duration` (String)
- `discount_coupon_duration_in_months` (Number)
- `discount_coupon_id` (String)
- `discount_coupon_livemode` (Boolean)
- `discount_coupon_max_redemptions` (Number)
- `discount_coupon_metadata` (Map of String)
- `discount_coupon_name` (String)
- `discount_coupon_object` (String)
- `discount_coupon_percent_off` (Number)
- `discount_coupon_redeem_by` (Number)
- `discount_coupon_times_redeemed` (Number)
- `discount_coupon_valid` (Boolean)
- `discount_customer` (String)
- `discount_end` (Number)
- `discount_id` (String)
- `discount_invoice` (String)
- `discount_invoice_item` (String)
- `discount_object` (String)
- `discount_promotion_code` (String)
- `discount_start` (Number)
- `discount_subscription` (String)
- `ended_at` (Number)
- `id` (String)
- `items_data` (List of Object) (see [below for nested schema](#nestedobjatt--subscriptions_data--items_data))
- `items_has_more` (Boolean)
- `items_object` (String)
- `items_url` (String)
- `latest_invoice` (String)
- `livemode` (Boolean)
- `metadata` (Map of String)
- `next_pending_invoice_item_invoice` (Number)
- `object` (String)
- `on_behalf_of` (String)
- `pause_collection_behavior` (String)
- `pause_collection_resumes_at` (Number)
- `payment_settings_payment_method_options_acss_debit_mandate_options_transaction_type` (String)
- `payment_settings_payment_method_options_acss_debit_verification_method` (String)
- `payment_settings_payment_method_options_bancontact_preferred_language` (String)
- `payment_settings_payment_method_options_card_mandate_options_amount` (Number)
- `payment_settings_payment_method_options_card_mandate_options_amount_type` (String)
- `payment_settings_payment_method_options_card_mandate_options_description` (String)
- `payment_settings_payment_method_options_card_network` (String)
- `payment_settings_payment_method_options_card_request_three_d_secure` (String)
- `payment_settings_payment_method_options_customer_balance_bank_transfer_eu_bank_transfer_country` (String)
- `payment_settings_payment_method_options_customer_balance_bank_transfer_type` (String)
- `payment_settings_payment_method_options_customer_balance_funding_type` (String)
- `payment_settings_payment_method_options_us_bank_account_financial_connections_permissions` (List of String)
- `payment_settings_payment_method_options_us_bank_account_verification_method` (String)
- `payment_settings_payment_method_types` (List of String)
- `payment_settings_save_default_payment_method` (String)
- `pending_invoice_item_interval_interval` (String)
- `pending_invoice_item_interval_interval_count` (Number)
- `pending_setup_intent` (String)
- `pending_update_billing_cycle_anchor` (Number)
- `pending_update_expires_at` (Number)
- `pending_update_subscription_items` (List of Object) (see [below for nested schema](#nestedobjatt--subscriptions_data--pending_update_subscription_items))
- `pending_update_trial_end` (Number)
- `pending_update_trial_from_plan` (Boolean)
- `schedule` (String)
- `start_date` (Number)
- `status` (String)
- `test_clock` (String)
- `transfer_data_amount_percent` (Number)
- `transfer_data_destination` (String)
- `trial_end` (Number)
- `trial_start` (Number)

<a id="nestedobjatt--subscriptions_data--default_tax_rates"></a>
### Nested Schema for `subscriptions_data.default_tax_rates`

Read-Only:

- `active` (Boolean)
- `country` (String)
- `created` (Number)
- `description` (String)
- `display_name` (String)
- `id` (String)
- `inclusive` (Boolean)
- `jurisdiction` (String)
- `livemode` (Boolean)
- `metadata` (Map of String)
- `object` (String)
- `percentage` (Number)
- `state` (String)
- `tax_type` (String)


<a id="nestedobjatt--subscriptions_data--discount_coupon_currency_options"></a>
### Nested Schema for `subscriptions_data.discount_coupon_currency_options`

Read-Only:

- `amount_off` (Number)
- `key` (String)


<a id="nestedobjatt--subscriptions_data--items_data"></a>
### Nested Schema for `subscriptions_data.items_data`

Read-Only:

- `billing_thresholds_usage_gte` (Number)
- `created` (Number)
- `id` (String)
- `metadata` (Map of String)
- `object` (String)
- `price_active` (Boolean)
- `price_billing_scheme` (String)
- `price_created` (Number)
- `price_currency` (String)
- `price_currency_options` (List of Object) (see [below for nested schema](#nestedobjatt--subscriptions_data--items_data--price_currency_options))
- `price_custom_unit_amount_maximum` (Number)
- `price_custom_unit_amount_minimum` (Number)
- `price_custom_unit_amount_preset` (Number)
- `price_id` (String)
- `price_livemode` (Boolean)
- `price_lookup_key` (String)
- `price_metadata` (Map of String)
- `price_nickname` (String)
- `price_object` (String)
- `price_product` (String)
- `price_recurring_aggregate_usage` (String)
- `price_recurring_interval` (String)
- `price_recurring_interval_count` (Number)
- `price_recurring_usage_type` (String)
- `price_tax_behavior` (String)
- `price_tiers` (List of Object) (see [below for nested schema](#nestedobjatt--subscriptions_data--items_data--price_tiers))
- `price_tiers_mode` (String)
- `price_transform_quantity_divide_by` (Number)
- `price_transform_quantity_round` (String)
- `price_type` (String)
- `price_unit_amount` (Number)
- `price_unit_amount_decimal` (String)
- `quantity` (Number)
- `subscription` (String)
- `tax_rates` (List of Object) (see [below for nested schema](#nestedobjatt--subscriptions_data--items_data--tax_rates))

<a id="nestedobjatt--subscriptions_data--items_data--price_currency_options"></a>
### Nested Schema for `subscriptions_data.items_data.price_currency_options`

Read-Only:

- `custom_unit_amount_maximum` (Number)
- `custom_unit_amount_minimum` (Number)
- `custom_unit_amount_preset` (Number)
- `key` (String)
- `tax_behavior` (String)
- `tiers` (List of Object) (see [below for nested schema](#nestedobjatt--subscriptions_data--items_data--price_currency_options--tiers))
- `unit_amount` (Number)
- `unit_amount_decimal` (String)

<a id="nestedobjatt--subscriptions_data--items_data--price_currency_options--tiers"></a>
### Nested Schema for `subscriptions_data.items_data.price_currency_options.unit_amount_decimal`

Read-Only:

- `flat_amount` (Number)
- `flat_amount_decimal` (String)
- `unit_amount` (Number)
- `unit_amount_decimal` (String)
- `up_to` (Number)



<a id="nestedobjatt--subscriptions_data--items_data--price_tiers"></a>
### Nested Schema for `subscriptions_data.items_data.price_tiers`

Read-Only:

- `flat_amount` (Number)
- `flat_amount_decimal` (String)
- `unit_amount` (Number)
- `unit_amount_decimal` (String)
- `up_to` (Number)


<a id="nestedobjatt--subscriptions_data--items_data--tax_rates"></a>
### Nested Schema for `subscriptions_data.items_data.tax_rates`

Read-Only:

- `active` (Boolean)
- `country` (String)
- `created` (Number)
- `description` (String)
- `display_name` (String)
- `id` (String)
- `inclusive` (Boolean)
- `jurisdiction` (String)
- `livemode` (Boolean)
- `metadata` (Map of String)
- `object` (String)
- `percentage` (Number)
- `state` (String)
- `tax_type` (String)



<a id="nestedobjatt--subscriptions_data--pending_update_subscription_items"></a>
### Nested Schema for `subscriptions_data.pending_update_subscription_items`

Read-Only:

- `billing_thresholds_usage_gte` (Number)
- `created` (Number)
- `id` (String)
- `metadata` (Map of String)
- `object` (String)
- `price_active` (Boolean)
- `price_billing_scheme` (String)
- `price_created` (Number)
- `price_currency` (String)
- `price_currency_options` (List of Object) (see [below for nested schema](#nestedobjatt--subscriptions_data--pending_update_subscription_items--price_currency_options))
- `price_custom_unit_amount_maximum` (Number)
- `price_custom_unit_amount_minimum` (Number)
- `price_custom_unit_amount_preset` (Number)
- `price_id` (String)
- `price_livemode` (Boolean)
- `price_lookup_key` (String)
- `price_metadata` (Map of String)
- `price_nickname` (String)
- `price_object` (String)
- `price_product` (String)
- `price_recurring_aggregate_usage` (String)
- `price_recurring_interval` (String)
- `price_recurring_interval_count` (Number)
- `price_recurring_usage_type` (String)
- `price_tax_behavior` (String)
- `price_tiers` (List of Object) (see [below for nested schema](#nestedobjatt--subscriptions_data--pending_update_subscription_items--price_tiers))
- `price_tiers_mode` (String)
- `price_transform_quantity_divide_by` (Number)
- `price_transform_quantity_round` (String)
- `price_type` (String)
- `price_unit_amount` (Number)
- `price_unit_amount_decimal` (String)
- `quantity` (Number)
- `subscription` (String)
- `tax_rates` (List of Object) (see [below for nested schema](#nestedobjatt--subscriptions_data--pending_update_subscription_items--tax_rates))

<a id="nestedobjatt--subscriptions_data--pending_update_subscription_items--price_currency_options"></a>
### Nested Schema for `subscriptions_data.pending_update_subscription_items.price_currency_options`

Read-Only:

- `custom_unit_amount_maximum` (Number)
- `custom_unit_amount_minimum` (Number)
- `custom_unit_amount_preset` (Number)
- `key` (String)
- `tax_behavior` (String)
- `tiers` (List of Object) (see [below for nested schema](#nestedobjatt--subscriptions_data--pending_update_subscription_items--price_currency_options--tiers))
- `unit_amount` (Number)
- `unit_amount_decimal` (String)

<a id="nestedobjatt--subscriptions_data--pending_update_subscription_items--price_currency_options--tiers"></a>
### Nested Schema for `subscriptions_data.pending_update_subscription_items.price_currency_options.unit_amount_decimal`

Read-Only:

- `flat_amount` (Number)
- `flat_amount_decimal` (String)
- `unit_amount` (Number)
- `unit_amount_decimal` (String)
- `up_to` (Number)



<a id="nestedobjatt--subscriptions_data--pending_update_subscription_items--price_tiers"></a>
### Nested Schema for `subscriptions_data.pending_update_subscription_items.price_tiers`

Read-Only:

- `flat_amount` (Number)
- `flat_amount_decimal` (String)
- `unit_amount` (Number)
- `unit_amount_decimal` (String)
- `up_to` (Number)


<a id="nestedobjatt--subscriptions_data--pending_update_subscription_items--tax_rates"></a>
### Nested Schema for `subscriptions_data.pending_update_subscription_items.tax_rates`

Read-Only:

- `active` (Boolean)
- `country` (String)
- `created` (Number)
- `description` (String)
- `display_name` (String)
- `id` (String)
- `inclusive` (Boolean)
- `jurisdiction` (String)
- `livemode` (Boolean)
- `metadata` (Map of String)
- `object` (String)
- `percentage` (Number)
- `state` (String)
- `tax_type` (String)




<a id="nestedatt--tax_ids_data"></a>
### Nested Schema for `tax_ids_data`

Read-Only:

- `country` (String)
- `created` (Number)
- `customer` (String)
- `id` (String)
- `livemode` (Boolean)
- `object` (String)
- `type` (String)
- `value` (String)
- `verification_status` (String)
- `verification_verified_address` (String)
- `verification_verified_name` (String)


