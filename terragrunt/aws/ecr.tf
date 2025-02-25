resource "aws_ecr_repository" "superset_docs" {
  name                 = var.product_name
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = local.common_tags
}