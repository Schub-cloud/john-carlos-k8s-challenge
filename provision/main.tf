terraform {
  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "2.11.0"
    }
  }
}

provider "kubernetes" {
  config_path    = var.config_path
  config_context = var.config_context
}

resource "kubernetes_namespace" "k8s_namespace" {
  metadata {
    name = var.ns_name
  }
}