resource "github_organization_custom_property" "archived" {
  name               = "archived"
  value_type         = "true_false"
  required           = false
  default_value      = "false"
  description        = "Is this repository archived"
  values_editable_by = "org_actors"
}
