package provider

import (
	"terraform-provider-ohdear/pkg/ohdear"
)

type Config struct {
	client *ohdear.Client
	teamID string
}
