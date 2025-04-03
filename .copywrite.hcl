schema_version = 1

project {
  license        = "MIT"
  copyright_year = 2025
  copyright_holder = "BAXENERGY ITALIA SRL"

  header_ignore = [
    # examples used within documentation (prose)
    "examples/**",

    # GitHub issue template configuration
    ".github/ISSUE_TEMPLATE/*.yml",

    # golangci-lint tooling configuration
    ".golangci.yml",

    # GoReleaser tooling configuration
    ".goreleaser.yml",
  ]
}