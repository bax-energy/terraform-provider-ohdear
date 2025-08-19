package provider

import "terraform-provider-ohdear/ohdear"

type Config struct {
	client *ohdear.Client
	teamID string
}
