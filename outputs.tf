output "name" {
  value       = join("", kind_cluster.default.*.name)
  description = "the name of the kind_cluster"
}

output "kubeconfig" {
  value       = join("", kind_cluster.default.*.kubeconfig)
  description = "the kubeconfig of the kind_cluster"
}

output "ipAddress" {
  value       = local.ipAddress
  description = "the IP address of the kind_cluster controller node"
}

output "kubeconfig_with_ip" {
  value       = data.utils_deep_merge_yaml.default[0].output
  description = "the kubeconfig of the kind_cluster with the IP address of the controller node"
}

output "context_name" {
  value       = local.context_name
  description = "context name for the kind cluster in kube_config file"
}

output "loadbalancer_container_name" {
  value       = local.loadbalancer_container_name
  description = "loadbalancer container name for the kind cluster"
}