module "github_repository" {
  source = "../modules/github"

  project_name        = "example"
  project_description = "example"
}

module "github_repository2" {
  source = "../modules/github"

  project_name        = "example2"
  project_description = "example2"

  project_visibility = "public"

  template_repository_info = [{
    owner      = "IgorAssuncao"
    repository = "igor-nodejs-boilerplate"
  }]
}
