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

variable "wordpress_user" {
  description = "The WordPress admin user"
  type        = string
  sensitive   = true
}

variable "wordpress_password" {
  description = "The WordPress admin password"
  type        = string
  sensitive   = true
}