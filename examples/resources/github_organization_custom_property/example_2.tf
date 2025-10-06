resource "github_organization_custom_property" "tags" {
  name               = "tags"
  value_type         = "multi_select"
  required           = false
  description        = "Project tags"
  allowed_values     = ["frontend", "backend", "database", "api"]
  values_editable_by = "org_and_repo_actors"
}
