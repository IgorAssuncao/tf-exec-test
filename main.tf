resource "github_repository" "project_repository" {
  name        = "testing"
  description = "test description"

  visibility = "public"

  template {
    owner      = "IgorAssuncao"
    repository = "igor-nodejs-boilerplate"
  }
}

