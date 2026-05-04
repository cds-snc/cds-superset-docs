locals {
  rate_limit_all_requests      = 1000
  rate_limit_mutating_requests = 200
}

resource "aws_wafv2_web_acl" "superset_docs" {
  provider = aws.us-east-1

  name        = "cds-superset-docs"
  description = "Superset Docs CloudFront distribution"
  scope       = "CLOUDFRONT"

  default_action {
    allow {}
  }

  rule {
    name     = "CanadaOnlyGeoRestriction"
    priority = 1

    action {
      block {
        custom_response {
          response_code = 403
          response_header {
            name  = "waf-block"
            value = "CanadaOnlyGeoRestriction"
          }
        }
      }
    }

    statement {
      not_statement {
        statement {
          or_statement {
            statement {
              geo_match_statement {
                country_codes = ["CA"]
              }
            }
            statement {
              byte_match_statement {
                positional_constraint = "EXACTLY"
                field_to_match {
                  single_header {
                    name = "upptime"
                  }
                }
                search_string = var.upptime_status_header
                text_transformation {
                  priority = 1
                  type     = "NONE"
                }
              }
            }
          }
        }
      }
    }

    visibility_config {
      cloudwatch_metrics_enabled = true
      metric_name                = "CanadaOnlyGeoRestriction"
      sampled_requests_enabled   = true
    }
  }

  rule {
    name     = "RateLimitAllRequestsIp"
    priority = 10

    action {
      block {
        custom_response {
          response_code = 429
          response_header {
            name  = "waf-block"
            value = "RateLimitAllRequestsIp"
          }
        }
      }
    }

    statement {
      rate_based_statement {
        limit              = local.rate_limit_all_requests
        aggregate_key_type = "IP"
      }
    }

    visibility_config {
      cloudwatch_metrics_enabled = true
      metric_name                = "RateLimitAllRequestsIp"
      sampled_requests_enabled   = true
    }
  }

  rule {
    name     = "RateLimitMutatingRequestsIp"
    priority = 20

    action {
      block {}
    }

    statement {
      rate_based_statement {
        limit              = local.rate_limit_mutating_requests
        aggregate_key_type = "IP"
        scope_down_statement {
          regex_match_statement {
            field_to_match {
              method {}
            }
            regex_string = "^(delete|patch|post|put)$"
            text_transformation {
              priority = 1
              type     = "LOWERCASE"
            }
          }
        }
      }
    }

    visibility_config {
      cloudwatch_metrics_enabled = true
      metric_name                = "RateLimitMutatingRequestsIp"
      sampled_requests_enabled   = true
    }
  }

  rule {
    name     = "RateLimitAllRequestsJA4"
    priority = 30

    action {
      block {}
    }

    statement {
      rate_based_statement {
        limit              = local.rate_limit_all_requests
        aggregate_key_type = "CUSTOM_KEYS"

        custom_key {
          ja4_fingerprint {
            fallback_behavior = "MATCH"
          }
        }
      }
    }

    visibility_config {
      cloudwatch_metrics_enabled = true
      metric_name                = "RateLimitAllRequestsJA4"
      sampled_requests_enabled   = true
    }
  }

  rule {
    name     = "RateLimitMutatingRequestsJA4"
    priority = 40

    action {
      block {}
    }

    statement {
      rate_based_statement {
        limit              = local.rate_limit_mutating_requests
        aggregate_key_type = "CUSTOM_KEYS"

        custom_key {
          ja4_fingerprint {
            fallback_behavior = "MATCH"
          }
        }

        scope_down_statement {
          regex_match_statement {
            field_to_match {
              method {}
            }
            regex_string = "^(delete|patch|post|put)$"
            text_transformation {
              priority = 1
              type     = "LOWERCASE"
            }
          }
        }
      }
    }

    visibility_config {
      cloudwatch_metrics_enabled = true
      metric_name                = "RateLimitMutatingRequestsJA4"
      sampled_requests_enabled   = true
    }
  }

  rule {
    name     = "BlockLargeRequests"
    priority = 50

    action {
      block {}
    }

    statement {
      or_statement {
        statement {
          size_constraint_statement {
            field_to_match {
              body {
                oversize_handling = "MATCH"
              }
            }
            comparison_operator = "GT"
            size                = 8192
            text_transformation {
              priority = 0
              type     = "NONE"
            }
          }
        }
        statement {
          size_constraint_statement {
            field_to_match {
              cookies {
                match_pattern {
                  all {}
                }
                match_scope       = "ALL"
                oversize_handling = "MATCH"
              }
            }
            comparison_operator = "GT"
            size                = 8192
            text_transformation {
              priority = 0
              type     = "NONE"
            }
          }
        }
        statement {
          size_constraint_statement {
            field_to_match {
              headers {
                match_pattern {
                  all {}
                }
                match_scope       = "ALL"
                oversize_handling = "MATCH"
              }
            }
            comparison_operator = "GT"
            size                = 8192
            text_transformation {
              priority = 0
              type     = "NONE"
            }
          }
        }
      }
    }

    visibility_config {
      cloudwatch_metrics_enabled = true
      metric_name                = "BlockLargeRequests"
      sampled_requests_enabled   = true
    }
  }

  rule {
    name     = "AWSManagedRulesAmazonIpReputationList"
    priority = 60

    override_action {
      none {}
    }

    statement {
      managed_rule_group_statement {
        name        = "AWSManagedRulesAmazonIpReputationList"
        vendor_name = "AWS"
      }
    }

    visibility_config {
      cloudwatch_metrics_enabled = true
      metric_name                = "AWSManagedRulesAmazonIpReputationList"
      sampled_requests_enabled   = true
    }
  }

  rule {
    name     = "AWSManagedRulesKnownBadInputsRuleSet"
    priority = 70

    override_action {
      none {}
    }

    statement {
      managed_rule_group_statement {
        name        = "AWSManagedRulesKnownBadInputsRuleSet"
        vendor_name = "AWS"
      }
    }

    visibility_config {
      cloudwatch_metrics_enabled = true
      metric_name                = "AWSManagedRulesKnownBadInputsRuleSet"
      sampled_requests_enabled   = true
    }
  }

  rule {
    name     = "AWSManagedRulesCommonRuleSet"
    priority = 80

    override_action {
      none {}
    }

    statement {
      managed_rule_group_statement {
        name        = "AWSManagedRulesCommonRuleSet"
        vendor_name = "AWS"
      }
    }

    visibility_config {
      cloudwatch_metrics_enabled = true
      metric_name                = "AWSManagedRulesCommonRuleSet"
      sampled_requests_enabled   = true
    }
  }

  visibility_config {
    cloudwatch_metrics_enabled = true
    metric_name                = "cds-superset-docs"
    sampled_requests_enabled   = false
  }

  tags = local.common_tags
}

#
# AWS Shield Advanced
#
resource "aws_shield_protection" "superset_docs_cloudfront" {
  name         = "superset-docs-cloudfront"
  resource_arn = aws_cloudfront_distribution.superset_docs.arn
  tags         = local.common_tags
}

resource "aws_shield_application_layer_automatic_response" "superset_docs_cloudfront" {
  resource_arn = aws_cloudfront_distribution.superset_docs.arn
  action       = "BLOCK"
}
