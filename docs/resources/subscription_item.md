---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "stripe_subscription_item Resource - stripe"
subcategory: ""
description: |-
  
---

# stripe_subscription_item (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `subscription` (String) The identifier of the subscription to modify.

### Optional

- `billing_thresholds_usage_gte` (Number) Usage threshold that triggers the subscription to create an invoice
- `payment_behavior` (String) Use `allow_incomplete` to transition the subscription to `status=past_due` if a payment is required but cannot be paid. This allows you to manage scenarios where additional user actions are needed to pay a subscription's invoice. For example, SCA regulation may require 3DS authentication to complete payment. See the [SCA Migration Guide](https://stripe.com/docs/billing/migration/strong-customer-authentication) for Billing to learn more. This is the default behavior.

Use `default_incomplete` to transition the subscription to `status=past_due` when payment is required and await explicit confirmation of the invoice's payment intent. This allows simpler management of scenarios where additional user actions are needed to pay a subscription’s invoice. Such as failed payments, [SCA regulation](https://stripe.com/docs/billing/migration/strong-customer-authentication), or collecting a mandate for a bank debit payment method.

Use `pending_if_incomplete` to update the subscription using [pending updates](https://stripe.com/docs/billing/subscriptions/pending-updates). When you use `pending_if_incomplete` you can only pass the parameters [supported by pending updates](https://stripe.com/docs/billing/pending-updates-reference#supported-attributes).

Use `error_if_incomplete` if you want Stripe to return an HTTP 402 status code if a subscription's invoice cannot be paid. For example, if a payment method requires 3DS authentication due to SCA regulation and further user action is needed, this parameter does not update the subscription and returns an error instead. This was the default behavior for API versions prior to 2019-03-14. See the [changelog](https://stripe.com/docs/upgrades#2019-03-14) to learn more.
- `price` (String) The ID of the price object.
- `price_data_currency` (String)
- `price_data_product` (String)
- `price_data_recurring_interval` (String)
- `price_data_recurring_interval_count` (Number)
- `price_data_tax_behavior` (String)
- `price_data_unit_amount` (Number)
- `price_data_unit_amount_decimal` (String)
- `proration_behavior` (String) Determines how to handle [prorations](https://stripe.com/docs/subscriptions/billing-cycle#prorations) when the billing cycle changes (e.g., when switching plans, resetting `billing_cycle_anchor=now`, or starting a trial), or if an item's `quantity` changes.
- `proration_date` (Number) If set, the proration will be calculated as though the subscription was updated at the given time. This can be used to apply the same proration that was previewed with the [upcoming invoice](https://stripe.com/docs/api#retrieve_customer_invoice) endpoint.
- `quantity` (Number) The quantity you'd like to apply to the subscription item you're creating.
- `tax_rates` (List of String) A list of [Tax Rate](https://stripe.com/docs/api/tax_rates) ids. These Tax Rates will override the [`default_tax_rates`](https://stripe.com/docs/api/subscriptions/create#create_subscription-default_tax_rates) on the Subscription. When updating, pass an empty string to remove previously-defined tax rates.

### Read-Only

- `created` (Number) Time at which the object was created. Measured in seconds since the Unix epoch.
- `id` (String) Unique identifier for the object.
- `object` (String) String representing the object's type. Objects of the same type share the same value.


