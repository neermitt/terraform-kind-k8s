variable "nodes" {
    type = list(
        object({
            role = string
        })
    )
    default = []
    description = "values for the kind_config node block"
}