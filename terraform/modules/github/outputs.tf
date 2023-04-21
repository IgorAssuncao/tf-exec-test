output "repository_html_url" {
  value = github_repository.project_repository.html_url
}

output "repository_git_url" {
  value = github_repository.project_repository.git_clone_url
}

output "repository_ssh_url" {
  value = github_repository.project_repository.ssh_clone_url
}
