
locals {
  enabled = module.this.enabled
  cluster_name = module.this.name

  controller_container_name = format("%s-control-plane", local.cluster_name)
}

resource "kind_cluster" "default" {
    count = local.enabled? 1: 0
    name = local.cluster_name
    
    wait_for_ready = true
}

data "docker-utils_inspect" "default" {
    count = local.enabled? 1: 0

    container_name = local.controller_container_name

    depends_on = [
      kind_cluster.default
    ]
}

locals {
  ipAddress = join("", data.docker-utils_inspect.default.*.networks[0].ip_address)
  toBeMerged = {
    clusters = [
        {
            name = format("kind-%s", local.cluster_name)
            cluster = {
                server = format("https://%s:6443", local.ipAddress)
            }
        }
    ]
  }
}

data "utils_deep_merge_yaml" "default" {
    count = local.enabled? 1: 0

    append_list = false
    deep_copy_list = true

    input = [
       join("", kind_cluster.default.*.kubeconfig),
       yamlencode(local.toBeMerged)
    ]
}