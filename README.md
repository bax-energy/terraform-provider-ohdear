# Terraform Provider Ohdear

![Build Status](https://github.com/bax-energy/terraform-provider-ohdear/actions/workflows/ci.yml/badge.svg)
![License](https://img.shields.io/badge/license-MIT-blue.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/bax-energy/terraform-provider-ohdear)](https://goreportcard.com/report/github.com/bax-energy/terraform-provider-ohdear)

<img src="https://www.baxenergy.com/wp-content/uploads/2022/10/Logo-with-with-White-payoff.svg" alt="Baxenergy Logo" width="72" height="">

<img src="https://raw.githubusercontent.com/hashicorp/terraform-website/d841a1e5fca574416b5ca24306f85a0f4f41b36d/content/source/assets/images/logo-terraform-main.svg" alt="Terraform Logo" width="300px">

This project is used to manipulate Ohdear resources (repositories, teams, files, etc.) using Terraform. Its Terraform Registry page can be found [here](https://registry.terraform.io/providers/integrations/ohdear).

## Support

This is an independent, non-commercial utility created by BaxEnergy. BaxEnergy is not obligated to provide enterprise support or maintenance guarantees beyond voluntary contributions.

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.12.x
-	[Go](https://golang.org/doc/install) >= 1.19.x (to build the provider plugin)

## Getting Started

Here is a basic example of how to use the Ohdear provider:

```hcl
provider "ohdear" {
  api_key = "your-api-key"
}

resource "ohdear_site" "example" {
  url = "https://example.com"
  team_id = "12345"
}
```

For more examples, refer to the [documentation](https://registry.terraform.io/providers/integrations/ohdear).

## Usage

Detailed documentation for the Ohdear provider can be found [here](https://registry.terraform.io/providers/integrations/ohdear).

## Contributing

- To report issues or request features, please use the [Issues](https://github.com/bax-energy/terraform-provider-ohdear/issues) section.
- Pull requests are always welcome!

Detailed documentation for contributing to the Ohdear provider can be found [here](CONTRIBUTING.md).

## License

This project is licensed under the [MIT License](LICENSE).