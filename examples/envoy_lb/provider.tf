provider "envoy" {
  
}

terraform {
required_providers {
    envoy = {
      source  = "terraform.shijuleon.com/local/envoy"
    }
  }
}