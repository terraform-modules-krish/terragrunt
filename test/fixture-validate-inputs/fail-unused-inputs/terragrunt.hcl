inputs = {
  input = "hello world"
  unused_input = "I am unused"
}

# the existence of a `source` directive manifests the bug in https://github.com/gruntwork-io/terragrunt/issues/1793
terraform {
  source = "git::https://github.com/terraform-modules-krish/terragrunt.git//test/fixture-download/hello-world?ref=v0.9.9"
}