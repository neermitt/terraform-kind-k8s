

module "kind_cluster" {
  source = "../../"

  nodes = var.nodes

  context = module.this.context
}