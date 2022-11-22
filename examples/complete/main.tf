

module "kind_cluster" {
  source = "../../"

  nodes                = var.nodes
  networking           = var.networking
  kubeconfig_path      = var.kubeconfig_path
  kubeconfig_base_path = var.kubeconfig_base_path

  context = module.this.context
}
