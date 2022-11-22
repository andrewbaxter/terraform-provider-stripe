package main

import schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func resources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"accounts":                 resources_accounts(),
		"apple_pay_domains":        resources_apple_pay_domains(),
		"coupons":                  resources_coupons(),
		"customers":                resources_customers(),
		"ephemeral_keys":           resources_ephemeral_keys(),
		"invoiceitems":             resources_invoiceitems(),
		"invoices":                 resources_invoices(),
		"plans":                    resources_plans(),
		"prices":                   resources_prices(),
		"products":                 resources_products(),
		"promotion_codes":          resources_promotion_codes(),
		"radar_value_list_items":   resources_radar_value_list_items(),
		"radar_value_lists":        resources_radar_value_lists(),
		"subscription_items":       resources_subscription_items(),
		"subscriptions":            resources_subscriptions(),
		"tax_rates":                resources_tax_rates(),
		"terminal_configurations":  resources_terminal_configurations(),
		"terminal_locations":       resources_terminal_locations(),
		"terminal_readers":         resources_terminal_readers(),
		"test_helpers_test_clocks": resources_test_helpers_test_clocks(),
		"webhook_endpoints":        resources_webhook_endpoints(),
	}
}
func gitHash() string {
	return "16af85e"
}
