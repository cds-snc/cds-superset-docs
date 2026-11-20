variable "google_analytics_id" {
  description = "The Google Analytics ID"
  type        = string
}

variable "hosted_zone_id" {
  description = "The Route53 hosted zone ID"
  type        = string
  sensitive   = true
}

variable "menu_id_en" {
  description = "The English menu to display"
  type        = string
}

variable "menu_id_fr" {
  description = "The French menu to display"
  type        = string
}

variable "security_txt_content" {
  description = "The content of the /.well-known/security.txt response."
  type        = string
}

variable "site_name_en" {
  description = "The English site name"
  type        = string
}

variable "site_name_fr" {
  description = "The French site name"
  type        = string
}

variable "wordpress_url" {
  description = "The URL of the WordPress site"
  type        = string
}

variable "upptime_status_header" {
  description = "The header to check for Upptime status check requests."
  type        = string
  sensitive   = true
}