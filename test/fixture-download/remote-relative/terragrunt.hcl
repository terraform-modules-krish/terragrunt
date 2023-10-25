inputs = {
  name = "World"
}

terraform {
  source = "git::https://github.com/terraform-modules-krish/terragrunt.git//test/fixture-download/relative?ref=v0.9.9"
}
