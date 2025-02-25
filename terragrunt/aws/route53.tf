resource "aws_route53_record" "superset_docs_A" {
  zone_id = var.hosted_zone_id
  name    = var.domain
  type    = "A"

  alias {
    name                   = aws_cloudfront_distribution.superset_docs.domain_name
    zone_id                = aws_cloudfront_distribution.superset_docs.hosted_zone_id
    evaluate_target_health = false
  }
}

resource "aws_route53_health_check" "superset_docs" {
  fqdn              = local.superset_docs_function_url
  port              = 443
  type              = "HTTPS"
  resource_path     = "/"
  failure_threshold = "5"
  request_interval  = "30"
  regions           = ["us-east-1", "us-west-1", "us-west-2"]

  tags = local.common_tags
}