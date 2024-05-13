variable "ns_name" {
  type        = string
  description = "name of the minikube namespace to be created"
}

variable "config_path" {
  type        = string
  description = "path of the k8s config for minikube"
  default     = "~/.kube/config"
}

variable "config_context" {
  type        = string
  description = "minikube context"
  default     = "minikube"
}