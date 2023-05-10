# Create an output for the public_dns of each EC2 instance
output "ec2" {
  value = {
    for name, this in aws_instance.ec2 : name => this.public_dns
  }
}
