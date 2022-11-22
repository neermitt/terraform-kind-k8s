variable "nodes" {
  type = list(
    object({
      role  = string
      image = optional(string)
    })
  )
  default     = []
  description = "values for the kind_config node block"
}

variable "networking" {
  type = object({
    ip_family           = optional(string)
    api_server_address  = optional(string)
    api_server_port     = optional(number)
    pod_subnet          = optional(string)
    service_subnet      = optional(string)
    disable_default_cni = optional(bool)
    kube_proxy_mode     = optional(string)
  })
  default     = null
  description = "Multiple details of the cluster's networking can be customized under the networking field. See https://kind.sigs.k8s.io/docs/user/configuration/#networking"
}

variable "kubeconfig_path" {
  type    = string
  default = null

  description = "file path, where kubeconfig should be stored, directory path should be provided as alternative"
}

variable "kubeconfig_base_path" {
  type        = string
  default     = null
  description = "directory, where kubeconfig should be generated, kubeconfig will have <cluster-name>.config pattern"
}
