module "superset_docs" {
  source    = "github.com/cds-snc/terraform-modules//lambda?ref=v10.5.3"
  name      = var.product_name
  ecr_arn   = aws_ecr_repository.superset_docs.arn
  image_uri = "${aws_ecr_repository.superset_docs.repository_url}:latest"

  architectures          = ["arm64"]
  memory                 = 1024
  timeout                = 10
  enable_lambda_insights = true

  environment_variables = {
    GOOGLE_ANALYTICS_ID  = var.google_analytics_id
    SITE_NAME_EN         = var.site_name_en
    SITE_NAME_FR         = var.site_name_fr
    WORDPRESS_MENU_ID_EN = var.menu_id_en
    WORDPRESS_MENU_ID_FR = var.menu_id_fr
    WORDPRESS_URL        = var.wordpress_url
    WORDPRESS_USERNAME   = var.wordpress_user
    WORDPRESS_PASSWORD   = var.wordpress_password
  }

  billing_tag_value = var.billing_code
}

resource "aws_lambda_function_url" "superset_docs" {
  function_name      = module.superset_docs.function_name
  authorization_type = "NONE"
}

#
# Function warmer
#
resource "aws_cloudwatch_event_rule" "superset_docs" {
  name                = "invoke-superset-docs"
  description         = "Keep the function toasty warm"
  schedule_expression = "rate(5 minutes)"
}

resource "aws_cloudwatch_event_target" "superset_docs" {
  target_id = "invoke-lambda"
  rule      = aws_cloudwatch_event_rule.superset_docs.name
  arn       = module.superset_docs.function_arn
  input     = jsonencode({})
}

resource "aws_lambda_permission" "superset_docs" {
  statement_id  = "AllowExecutionFromEventBridge"
  action        = "lambda:InvokeFunction"
  function_name = module.superset_docs.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.superset_docs.arn
}