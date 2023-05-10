
variable "server_count" {
  type        = number
  default     = 3
  description = "The total number of VMs to create"
}

variable "name" {
  type        = string
  description = "The name to assign to resources in this module"
  default     = "super-webserver"
}

variable "environment" {
  type        = string
  description = "The environment to assign to resource in this module"
  default     = "prod"
}

variable "tags" {
  type        = map(any)
  default     = {}
  description = "Tags to asscociate to taggable resources in this module"
}

variable "instance_type" {
  type    = string
  default = "t2.micro"
}

variable "root_volume_size" {
  type    = string
  default = 50
}

variable "production" {
  type    = bool
  default = false
}

locals {
  host = { for i in range(var.server_count) : join("-", [var.name, var.environment, i]) => {
    index = i
    name  = "${join("-", [var.name, var.environment, i])}"
    }
  }

  user_data = templatefile("${path.module}/user_data.template",
    {
      name        = var.name
      environment = var.environment
  })

}
