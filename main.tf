terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
    random = {
      source = "hashicorp/random"
    }
  }

  cloud {
    organization = "City-Hotels"

    workspaces {
      name = "ch-backend"
    }
  }
}
