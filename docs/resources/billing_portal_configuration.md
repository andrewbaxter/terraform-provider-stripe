---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "stripe_billing_portal_configuration Resource - stripe"
subcategory: ""
description: |-
  
---

# stripe_billing_portal_configuration (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `business_profile_headline` (String) The messaging shown to customers in the portal.
- `business_profile_privacy_policy_url` (String) A link to the business’s publicly available privacy policy.
- `business_profile_terms_of_service_url` (String) A link to the business’s publicly available terms of service.
- `default_return_url` (String) The default URL to redirect customers to when they click on the portal's link to return to your website. This can be [overriden](https://stripe.com/docs/api/customer_portal/sessions/create#create_portal_session-return_url) when creating the session.
- `features_customer_update_allowed_updates` (List of String) The types of customer updates that are supported. When empty, customers are not updateable.
- `features_customer_update_enabled` (Boolean) Whether the feature is enabled.
- `features_invoice_history_enabled` (Boolean) Whether the feature is enabled.
- `features_payment_method_update_enabled` (Boolean) Whether the feature is enabled.
- `features_subscription_cancel_cancellation_reason_enabled` (Boolean) Whether the feature is enabled.
- `features_subscription_cancel_cancellation_reason_options` (List of String) Which cancellation reasons will be given as options to the customer.
- `features_subscription_cancel_enabled` (Boolean) Whether the feature is enabled.
- `features_subscription_cancel_mode` (String) Whether to cancel subscriptions immediately or at the end of the billing period.
- `features_subscription_cancel_proration_behavior` (String) Whether to create prorations when canceling subscriptions. Possible values are `none` and `create_prorations`.
- `features_subscription_pause_enabled` (Boolean) Whether the feature is enabled.
- `features_subscription_update_default_allowed_updates` (List of String) The types of subscription updates that are supported for items listed in the `products` attribute. When empty, subscriptions are not updateable.
- `features_subscription_update_enabled` (Boolean) Whether the feature is enabled.
- `features_subscription_update_products` (Block List) The list of products that support subscription updates. (see [below for nested schema](#nestedblock--features_subscription_update_products))
- `features_subscription_update_proration_behavior` (String) Determines how to handle prorations resulting from subscription updates. Valid values are `none`, `create_prorations`, and `always_invoice`.
- `login_page_enabled` (Boolean) If `true`, a shareable `url` will be generated that will take your customers to a hosted login page for the customer portal.

If `false`, the previously generated `url`, if any, will be deactivated.

### Read-Only

- `application` (String) ID of the Connect Application that created the configuration.
- `created` (Number) Time at which the object was created. Measured in seconds since the Unix epoch.
- `id` (String) Unique identifier for the object.
- `is_default` (Boolean) Whether the configuration is the default. If `true`, this configuration can be managed in the Dashboard and portal sessions will use this configuration unless it is overriden when creating the session.
- `livemode` (Boolean) Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
- `login_page_url` (String) A shareable URL to the hosted portal login page. Your customers will be able to log in with their [email](https://stripe.com/docs/api/customers/object#customer_object-email) and receive a link to their customer portal.
- `object` (String) String representing the object's type. Objects of the same type share the same value.
- `updated` (Number) Time at which the object was last updated. Measured in seconds since the Unix epoch.

<a id="nestedblock--features_subscription_update_products"></a>
### Nested Schema for `features_subscription_update_products`

Required:

- `prices` (List of String) The list of price IDs which, when subscribed to, a subscription can be updated.
- `product` (String) The product ID.


