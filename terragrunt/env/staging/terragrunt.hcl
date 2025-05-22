terraform {
  source = "../..//aws"
}

include {
  path = find_in_parent_folders("root.hcl")
}

inputs = {
  google_analytics_id = ""
  menu_id_en          = "9"
  menu_id_fr          = "10"
  site_name_en        = "Superset Docs"
  site_name_fr        = "Docs de Superset"
  wordpress_url       = "https://articles.alpha.canada.ca/pcs-superset"
}
