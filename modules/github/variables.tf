variable "project_name" {
  type        = string
  description = "Project name"
  default     = ""
}

variable "project_description" {
  type        = string
  description = "Project description"
  default     = ""
}

variable "project_visibility" {
  type        = string
  description = "Project visibility"
  default     = "public"
}

variable "template_repository_info" {
  type = list(object({
    owner      = string
    repository = string
  }))
  description = "Template repository information if willing to use a template repository."
  default     = []
}
