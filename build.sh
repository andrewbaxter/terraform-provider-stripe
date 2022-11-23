#!/usr/bin/bash -xeu
rm generated/resources*
go run ./generator
#cd generated
#rm -rf docs
#go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name stripe --rendered-provider-name Stripe --rendered-website-dir ../docs
