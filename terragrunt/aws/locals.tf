locals {
  common_tags = {
    Terraform  = "true"
    CostCentre = var.billing_code
  }
  superset_docs_cloudwatch_log_group_name = "/aws/lambda/${module.superset_docs.function_name}"
  superset_docs_function_url              = split("/", aws_lambda_function_url.superset_docs.function_url)[2]
}