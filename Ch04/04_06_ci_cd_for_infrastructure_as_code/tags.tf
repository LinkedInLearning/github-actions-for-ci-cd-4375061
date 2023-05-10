data "aws_caller_identity" "id" {}

locals {
  tags = {
    Name        = var.name
    Environment = var.environment
    Designation          = "${var.name}-${var.environment}"
    Terraform   = true
  }
}
