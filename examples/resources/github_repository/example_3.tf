# This example matches the user's requested format with native HCL lists
resource "github_repository" "example" {
  name        = "example"
  description = "My awesome codebase"

  custom_property {
    property_name = "foo"
    value         = ["bar"]
  }

  custom_property {
    property_name = "boolean"
    value         = ["false"]
  }

  custom_property {
    property_name = "multiselect"
    value         = ["goo", "zoo"]
  }

  custom_property {
    property_name = "singleselect"
    value         = ["acme"]
  }
}
