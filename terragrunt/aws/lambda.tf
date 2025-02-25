module "superset_docs" {
  source    = "github.com/cds-snc/terraform-modules//lambda?ref=v10.3.0"
  name      = var.product_name
  ecr_arn   = aws_ecr_repository.superset_docs.arn
  image_uri = "${aws_ecr_repository.superset_docs.repository_url}:latest"

  architectures          = ["arm64"]
  memory                 = 4096
  timeout                = 10
  enable_lambda_insights = true

  environment_variables = {
    MENU_ID_EN         = var.menu_id_en
    MENU_ID_FR         = var.menu_id_fr
    SITE_NAME_EN       = var.site_name_en
    SITE_NAME_FR       = var.site_name_fr
    WORDPRESS_URL      = var.wordpress_url
    WORDPRESS_USER     = var.wordpress_user
    WORDPRESS_PASSWORD = var.wordpress_password
  }

  billing_tag_value = var.billing_code
}

resource "aws_lambda_function_url" "superset_docs" {
  function_name      = module.superset_docs.function_name
  authorization_type = "NONE"
}