resource "aws_ssm_parameter" "wordpress_username" {
  name             = "/${var.product_name}/wordpress/username"
  description      = "WordPress API username"
  type             = "SecureString"
  value_wo         = "UPDATE_OUT_OF_BAND"
  value_wo_version = 1
  tags             = local.common_tags

  lifecycle {
    prevent_destroy = true
  }
}

resource "aws_ssm_parameter" "wordpress_password" {
  name             = "/${var.product_name}/wordpress/password"
  description      = "WordPress API password"
  type             = "SecureString"
  value_wo         = "UPDATE_OUT_OF_BAND"
  value_wo_version = 1
  tags             = local.common_tags

  lifecycle {
    prevent_destroy = true
  }
}
