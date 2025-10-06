package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubOrganizationCustomProperty(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates a single_select custom property without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_organization_custom_property" "test" {
				name               = "environment_%s"
				value_type         = "single_select"
				required           = true
				default_value      = "production"
				description        = "Prod or dev environment"
				allowed_values     = ["production", "development"]
				values_editable_by = "org_actors"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "name", fmt.Sprintf("environment_%s", randomID)),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "value_type", "single_select"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "required", "true"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "default_value", "production"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "description", "Prod or dev environment"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "allowed_values.#", "2"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "allowed_values.0", "production"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "allowed_values.1", "development"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "values_editable_by", "org_actors"),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("creates a multi_select custom property without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_organization_custom_property" "test" {
				name               = "tags_%s"
				value_type         = "multi_select"
				required           = false
				description        = "Project tags"
				allowed_values     = ["frontend", "backend", "database", "api"]
				values_editable_by = "org_and_repo_actors"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "name", fmt.Sprintf("tags_%s", randomID)),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "value_type", "multi_select"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "required", "false"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "description", "Project tags"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "allowed_values.#", "4"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "values_editable_by", "org_and_repo_actors"),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("creates a string custom property without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_organization_custom_property" "test" {
				name               = "owner_%s"
				value_type         = "string"
				required           = false
				description        = "Repository owner"
				values_editable_by = "org_actors"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "name", fmt.Sprintf("owner_%s", randomID)),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "value_type", "string"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "required", "false"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "description", "Repository owner"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "values_editable_by", "org_actors"),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("creates a true_false custom property without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_organization_custom_property" "test" {
				name               = "archived_%s"
				value_type         = "true_false"
				required           = false
				description        = "Is this repository archived"
				values_editable_by = "org_actors"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "name", fmt.Sprintf("archived_%s", randomID)),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "value_type", "true_false"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "required", "false"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "description", "Is this repository archived"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "values_editable_by", "org_actors"),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("updates a custom property without error", func(t *testing.T) {
		configBefore := fmt.Sprintf(`
			resource "github_organization_custom_property" "test" {
				name               = "status_%s"
				value_type         = "single_select"
				required           = false
				default_value      = "active"
				description        = "Repository status"
				allowed_values     = ["active", "inactive"]
				values_editable_by = "org_actors"
			}
		`, randomID)

		configAfter := fmt.Sprintf(`
			resource "github_organization_custom_property" "test" {
				name               = "status_%s"
				value_type         = "single_select"
				required           = true
				default_value      = "archived"
				description        = "Updated repository status"
				allowed_values     = ["active", "inactive", "archived"]
				values_editable_by = "org_and_repo_actors"
			}
		`, randomID)

		checkBefore := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "name", fmt.Sprintf("status_%s", randomID)),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "required", "false"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "default_value", "active"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "description", "Repository status"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "allowed_values.#", "2"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "values_editable_by", "org_actors"),
		)

		checkAfter := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "name", fmt.Sprintf("status_%s", randomID)),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "required", "true"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "default_value", "archived"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "description", "Updated repository status"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "allowed_values.#", "3"),
			resource.TestCheckResourceAttr("github_organization_custom_property.test", "values_editable_by", "org_and_repo_actors"),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: configBefore,
						Check:  checkBefore,
					},
					{
						Config: configAfter,
						Check:  checkAfter,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
