terraform {
  required_version = "~> 1.4"

  backend "s3" {
    key    = "github-actions-cicd/terraform.tfstate" # the directory/file.tfstate
    bucket = "ADD_YOUR_BUCKET_NAME_HERE"             # the bucket
    region = "ADD_YOUR_REGION_NAME_HERE"             # the region
  }
}
