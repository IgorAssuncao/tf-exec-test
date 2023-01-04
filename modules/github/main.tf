resource "github_repository" "project_repository" {
  name        = var.project_name
  description = var.project_description

  visibility = var.project_visibility

  dynamic "template" {
    for_each = var.template_repository_info
    content {
      owner      = template.value["owner"]
      repository = template.value["repository"]
    }
  }
}

