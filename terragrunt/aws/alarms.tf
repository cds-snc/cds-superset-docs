#
# SNS topics created by the cds-snc/cds-superset Terraform 
#
data "aws_sns_topic" "cloudwatch_alert_warning" {
  name = "cloudwatch-alert-warning"
}

data "aws_sns_topic" "cloudwatch_alert_ok" {
  name = "cloudwatch-alert-ok"
}

data "aws_sns_topic" "cloudwatch_alert_warning_us_east_1" {
  provider = aws.us-east-1
  name     = "cloudwatch-alert-warning"
}

data "aws_sns_topic" "cloudwatch_alert_ok_us_east_1" {
  provider = aws.us-east-1
  name     = "cloudwatch-alert-ok"
}

#
# Errors logged
#
resource "aws_cloudwatch_log_metric_filter" "superset_docs_errors" {
  name           = "error-superset-docs"
  pattern        = "ERROR? Error? error?"
  log_group_name = local.superset_docs_cloudwatch_log_group_name

  metric_transformation {
    name          = "error-superset-docs"
    namespace     = "superset-docs"
    value         = "1"
    default_value = "0"
  }
}

resource "aws_cloudwatch_metric_alarm" "superset_docs_errors" {
  alarm_name          = "errors-superset-docs"
  alarm_description   = "`superset-docs` errors logged over 1 minute."
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "1"
  metric_name         = aws_cloudwatch_log_metric_filter.superset_docs_errors.metric_transformation[0].name
  namespace           = aws_cloudwatch_log_metric_filter.superset_docs_errors.metric_transformation[0].namespace
  period              = "60"
  statistic           = "Sum"
  threshold           = "0"
  treat_missing_data  = "notBreaching"

  alarm_actions = [data.aws_sns_topic.cloudwatch_alert_warning.arn]
  ok_actions    = [data.aws_sns_topic.cloudwatch_alert_ok.arn]

  tags = local.common_tags
}

resource "aws_cloudwatch_metric_alarm" "superset_docs_invocation_errors" {
  alarm_name          = "invokcation-errors-superset-docs"
  alarm_description   = "`superset-docs` invocation errors logged over 1 minute."
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "1"
  metric_name         = "Errors"
  namespace           = "AWS/Lambda"
  period              = "60"
  statistic           = "Sum"
  threshold           = "0"
  treat_missing_data  = "notBreaching"

  alarm_actions = [data.aws_sns_topic.cloudwatch_alert_warning.arn]
  ok_actions    = [data.aws_sns_topic.cloudwatch_alert_ok.arn]

  tags = local.common_tags
}

resource "aws_cloudwatch_metric_alarm" "superset_docs_health_check" {
  provider = aws.us-east-1

  alarm_name          = "health-check-superset-docs"
  alarm_description   = "`superset-docs` heath check has failed in a 2 minute period."
  comparison_operator = "LessThanThreshold"
  evaluation_periods  = "2"
  metric_name         = "HealthCheckStatus"
  namespace           = "AWS/Route53"
  period              = "60"
  statistic           = "Minimum"
  threshold           = "1"
  treat_missing_data  = "notBreaching"

  alarm_actions = [data.aws_sns_topic.cloudwatch_alert_warning_us_east_1.arn]
  ok_actions    = [data.aws_sns_topic.cloudwatch_alert_ok_us_east_1.arn]

  dimensions = {
    HealthCheckId = aws_route53_health_check.superset_docs.id
  }

  tags = local.common_tags
}

#
# Log Insight queries
#
resource "aws_cloudwatch_query_definition" "superset_ecs_errors" {
  name            = "Superset Docs - errors"
  log_group_names = [local.superset_docs_cloudwatch_log_group_name]

  query_string = <<-QUERY
    fields @timestamp, @message, @logStream
    | filter @message like /(?i)error/
    | sort @timestamp desc
    | limit 100
  QUERY
}