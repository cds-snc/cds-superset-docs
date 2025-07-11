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

resource "aws_route53_record" "superset_docs_HTTPS" {
  zone_id = var.hosted_zone_id
  name    = var.domain
  type    = "HTTPS"

  alias {
    name                   = aws_cloudfront_distribution.superset_docs.domain_name
    zone_id                = aws_cloudfront_distribution.superset_docs.hosted_zone_id
    evaluate_target_health = false
  }
}
