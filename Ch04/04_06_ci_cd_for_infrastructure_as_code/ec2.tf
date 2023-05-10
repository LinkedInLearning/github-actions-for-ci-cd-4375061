resource "aws_security_group" "ec2" {
  name_prefix = substr("${var.name}-${var.environment}-ec2-", 0, 32)
  description = "EC2 security group for ${var.name}-${var.environment}"
  tags        = merge(var.tags, local.tags, { Resource = "EC2" })
}

resource "aws_security_group_rule" "ec2-http" {
  security_group_id = aws_security_group.ec2.id
  type              = "ingress"
  to_port           = 80
  from_port         = 80
  protocol          = "6"
  cidr_blocks       = ["0.0.0.0/0"]
  description       = "HTTP Ingress for ${var.name}-${var.environment}"
  ipv6_cidr_blocks  = []
}

resource "aws_security_group_rule" "ec2-egress" {
  security_group_id = aws_security_group.ec2.id
  type              = "egress"
  to_port           = 0
  from_port         = 0
  protocol          = "-1"
  cidr_blocks       = ["0.0.0.0/0"]
  description       = "Egress for ${var.name}-${var.environment}"
  ipv6_cidr_blocks  = []
  prefix_list_ids   = []
}

resource "aws_instance" "ec2" {
  for_each                    = local.host
  ami                         = data.aws_ami.ubuntu.id
  instance_type               = var.instance_type
  user_data                   = local.user_data
  vpc_security_group_ids      = [aws_security_group.ec2.id]
  associate_public_ip_address = true
  tags                        = merge(local.tags, { Name = each.value.name })

  root_block_device {
    volume_size = var.root_volume_size
    encrypted   = true
    tags        = merge(local.tags, { Name = each.key })
  }

  lifecycle {
    create_before_destroy = "true"
  }
}
