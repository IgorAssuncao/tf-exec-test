terraform {
  required_providers {
    github = {
      source  = "integrations/github"
      version = "~> 4.31.0"
    }
  }
}

provider "github" {
  # token = "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"
}