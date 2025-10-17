resource "github_organization_custom_property" "owner" {
  name               = "owner"
  value_type         = "string"
  required           = false
  description        = "Repository owner"
  values_editable_by = "org_actors"
}
