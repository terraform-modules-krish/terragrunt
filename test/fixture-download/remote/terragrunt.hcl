inputs = {
  name = "World"
}

terraform {
  source = "git::https://github.com/terraform-modules-krish/terragrunt.git//test/fixture-download/hello-world?ref=v0.9.9"
}
