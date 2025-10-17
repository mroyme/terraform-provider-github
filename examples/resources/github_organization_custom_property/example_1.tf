resource "github_organization_custom_property" "environment" {
  name               = "environment"
  value_type         = "single_select"
  required           = true
  default_value      = "production"
  description        = "Prod or dev environment"
  allowed_values     = ["production", "development"]
  values_editable_by = "org_actors"
}
