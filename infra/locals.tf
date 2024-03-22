locals {
  function_name = "goterraform-demo"
  owner         = "Massimo Biagioli"

  app_id = "${lower(local.function_name)}-${lower(var.app_env)}"

  runtime = "provided.al2"
  handler = "main"

  common_tags = {
    Service = local.function_name
    Owner   = local.owner
    Destroy = "false"
  }
}
