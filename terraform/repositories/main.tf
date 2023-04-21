module "github_repository_no_template" {
  source = "../modules/github"

  project_name        = "example"
  project_description = "example"
}

module "github_repository_template" {
  source = "../modules/github"

  project_name        = "example_public"
  project_description = "example_public"

  project_visibility = "public"

  template_repository_info = [{
    owner      = "IgorAssuncao"
    repository = "igor-nodejs-boilerplate"
  }]
}
