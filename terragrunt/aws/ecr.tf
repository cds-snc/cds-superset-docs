resource "aws_ecr_repository" "superset_docs_test" {
  name                 = var.product_name
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = local.common_tags
}

resource "aws_ecr_lifecycle_policy" "superset_docs_test" {
  repository = aws_ecr_repository.superset_docs_test.name
  policy = jsonencode({
    "rules" : [
      {
        "rulePriority" : 1,
        "description" : "Keep last 30 release tagged images",
        "selection" : {
          "tagStatus" : "tagged",
          "tagPrefixList" : ["v"],
          "countType" : "imageCountMoreThan",
          "countNumber" : 30
        },
        "action" : {
          "type" : "expire"
        }
      },
      {
        "rulePriority" : 10,
        "description" : "Keep last 10 git SHA tagged images",
        "selection" : {
          "tagStatus" : "tagged",
          "tagPrefixList" : ["sha-"],
          "countType" : "imageCountMoreThan",
          "countNumber" : 10
        },
        "action" : {
          "type" : "expire"
        }
      },
      {
        "rulePriority" : 20,
        "description" : "Expire untagged images older than 7 days",
        "selection" : {
          "tagStatus" : "untagged",
          "countType" : "sinceImagePushed",
          "countUnit" : "days",
          "countNumber" : 7
        },
        "action" : {
          "type" : "expire"
        }
      }
    ]
  })
}

moved {
  from = aws_ecr_repository.superset_docs
  to   = aws_ecr_repository.superset_doc_test
}

moved {
  from = aws_ecr_lifecycle_policy.superset_docs
  to   = aws_ecr_lifecycle_policy.superset_doc_test
}
