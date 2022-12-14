---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "stripe_webhook_endpoint Resource - stripe"
subcategory: ""
description: |-
  
---

# stripe_webhook_endpoint (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enabled_events` (List of String) The list of events to enable for this endpoint. You may specify `['*']` to enable all events, except those that require explicit selection.
- `url` (String) The URL of the webhook endpoint.

### Optional

- `api_version` (String) Events sent to this endpoint will be generated with this Stripe Version instead of your account's default Stripe Version.
- `connect` (Boolean) Whether this endpoint should receive events from connected accounts (`true`), or from your account (`false`). Defaults to `false`.
- `description` (String) An optional description of what the webhook is used for.

### Read-Only

- `application` (String) The ID of the associated Connect application.
- `created` (Number) Time at which the object was created. Measured in seconds since the Unix epoch.
- `id` (String) Unique identifier for the object.
- `livemode` (Boolean) Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
- `object` (String) String representing the object's type. Objects of the same type share the same value.
- `secret` (String) The endpoint's secret, used to generate [webhook signatures](https://stripe.com/docs/webhooks/signatures). Only returned at creation.
- `status` (String) The status of the webhook. It can be `enabled` or `disabled`.


