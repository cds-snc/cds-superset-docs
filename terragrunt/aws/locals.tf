locals {
  common_tags = {
    Terraform  = "true"
    CostCentre = var.billing_code
  }
  superset_docs_function_url = split("/", aws_lambda_function_url.superset_docs.function_url)[2]
}