---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "stripe_terminal_reader Resource - stripe"
subcategory: ""
description: |-
  
---

# stripe_terminal_reader (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `registration_code` (String) A code generated by the reader used for registering to an account.

### Optional

- `label` (String) Custom label given to the reader for easier identification. If no label is specified, the registration code will be used.
- `location` (String) The location to assign the reader to.

### Read-Only

- `action_failure_code` (String) Failure code, only set if status is `failed`.
- `action_failure_message` (String) Detailed failure message, only set if status is `failed`.
- `action_process_payment_intent_payment_intent` (String) Most recent PaymentIntent processed by the reader.
- `action_process_payment_intent_process_config_skip_tipping` (Boolean) Override showing a tipping selection screen on this transaction.
- `action_process_payment_intent_process_config_tipping_amount_eligible` (Number) Amount used to calculate tip suggestions on tipping selection screen for this transaction. Must be a positive integer in the smallest currency unit (e.g., 100 cents to represent $1.00 or 100 to represent ¥100, a zero-decimal currency).
- `action_process_setup_intent_generated_card` (String) ID of a card PaymentMethod generated from the card_present PaymentMethod that may be attached to a Customer for future transactions. Only present if it was possible to generate a card PaymentMethod.
- `action_process_setup_intent_setup_intent` (String) Most recent SetupIntent processed by the reader.
- `action_set_reader_display_cart_currency` (String) Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
- `action_set_reader_display_cart_line_items` (List of Object) List of line items in the cart. (see [below for nested schema](#nestedatt--action_set_reader_display_cart_line_items))
- `action_set_reader_display_cart_tax` (Number) Tax amount for the entire cart. A positive integer in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
- `action_set_reader_display_cart_total` (Number) Total amount for the entire cart, including tax. A positive integer in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
- `action_set_reader_display_type` (String) Type of information to be displayed by the reader.
- `action_status` (String) Status of the action performed by the reader.
- `action_type` (String) Type of action performed by the reader.
- `device_sw_version` (String) The current software version of the reader.
- `device_type` (String) Type of reader, one of `bbpos_wisepad3`, `stripe_m2`, `bbpos_chipper2x`, `bbpos_wisepos_e`, `verifone_P400`, or `simulated_wisepos_e`.
- `id` (String) Unique identifier for the object.
- `ip_address` (String) The local IP address of the reader.
- `livemode` (Boolean) Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
- `object` (String) String representing the object's type. Objects of the same type share the same value.
- `serial_number` (String) Serial number of the reader.
- `status` (String) The networking status of the reader.

<a id="nestedatt--action_set_reader_display_cart_line_items"></a>
### Nested Schema for `action_set_reader_display_cart_line_items`

Read-Only:

- `amount` (Number)
- `description` (String)
- `quantity` (Number)


