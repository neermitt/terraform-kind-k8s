output "name" {
  value = join("", kind_cluster.default.*.name)
  description = "the name of the kind_cluster"
}

output "kubeconfig" {
  value = join("", kind_cluster.default.*.kubeconfig)
  description = "the kubeconfig of the kind_cluster"
}

output "ipAddress" {
    value = local.ipAddress
    description = "the IP address of the kind_cluster controller node"
}

output "kubeconfig_with_ip" {
    value = data.utils_deep_merge_yaml.default[0].output
    description = "the kubeconfig of the kind_cluster with the IP address of the controller node"
}