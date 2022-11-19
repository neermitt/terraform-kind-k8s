

module "kind_cluster" {
  source = "../../"

  nodes = var.nodes

  networking = var.networking

  context = module.this.context
}