# Changelog


# v0.9.9
_Released Feb 4, 2017_
- When Terragrunt downloads source code, it will now download the entire repository (i.e. the part before the double-slash `//` in the URL). This way, relative paths between modules in that repository should work correctly.

<br>

# v0.9.8
_Released Feb 2, 2017_
- Terragrunt now supports a `--terragrunt-source-update` flag. If set, Terragrunt will delete the contents of the temporary folder to clear out any old, cached source code before downloading new source code into it.

<br>

# v0.9.7
_Released Feb 1, 2017_
- When downloading remote Terraform configurations, Terragrunt will now cache the download in a temporary folder, and only re-download the files when necessary. This way, you don't have to download source code, download modules, and configure remote state for every single Terragrunt command.

<br>

# v0.9.6
_Released Jan 28, 2017_
- When you run `terragrunt foo`, Terragrunt will now only send the output of the `foo` command to `stdout`. The output from any other commands Terragrunt might run, such as `terraform remote config` or `terraform get` will now be redirected to `stderr`. This allows you to parse the `stdout` of your Terraform commands without worrying about them being polluted by unrelated logging.

<br>

# v0.9.5
_Released Jan 28, 2017_
- Terragrunt now allows you to specify a login profile to use by specifying a `profile` key in the `config` block of the DynamoDB lock and the `config` block of remote state (the latter only applies when using S3).

<br>

# v0.9.4
_Released Jan 26, 2017_
- Terragrunt can now download Terraform templates from a specified URL into a temporary folder and run Terraform in that temporary folder. This can dramatically reduce the amount of copy/paste in Terraform projects, allows you to version all of your code, and to promote immutable versions of that code through each of your environments (e.g. qa -> stage -> prod). See the [Terraform remote configurations documentation](https://github.com/gruntwork-io/terragrunt#remote-terraform-configurations) for more info.

<br>

# v0.9.3
_Released Jan 22, 2017_
- Terragrunt now displays clear error messages when you try to release a lock on a table or item that does not exist in DynamoDB

<br>

# v0.9.2
_Released Jan 15, 2017_
- Fix a bug where Terragrunt was not respecting the `--working-dir` parameter when configuring remote state.

<br>

# v0.9.1
_Released Jan 14, 2017_
- Terragrunt now supports a `get_env()` helper that you can use to read in environment variables in your `.terragrunt` files.

<br>

# v0.9.0
_Released Jan 13, 2017_
- Fix a ton of bugs on Windows. Terragrunt was not properly handling file paths, which use backslashes on Windows, which caused bugs when using the `find_in_parent_folders()` and `path_relative_to_include()` helpers. Thanks to @cstavro for fixing these bugs, as well as fixing all the automated tests so they pass on Windows too!

<br>

# v0.8.1
_Released Jan 11, 2017_
- Terragrunt should now properly propagate the exit code from Terraform.

<br>

# v0.8.0
_Released Jan 8, 2017_
- All of Terragrunt's log messages now go to stderr instead of stdout. Only Terraform's output goes to stdout. That way, any code you write that depends on Terraform's output in stdout (e.g. the `graph` command work) will now work as expected, without lots of Terragrunt logging polluting things.

<br>

# v0.7.4
_Released Jan 7, 2017_
- Update to the latest version of the AWS Go SDK. This should make the `AWS_SHARED_CREDENTIALS_FILE` env var work correctly when used with Terragrunt.

<br>

# v0.7.3
_Released Jan 3, 2017_
- Terragrunt should now forward the SIGTERM signal to Terraform

<br>

# v0.7.2
_Released Dec 21, 2016_
- Fixed a bug where Terragrunt’s default path for the config file is `.terragrunt` in the current working directory rather than the working directory specified by the `--terragrunt-working-dir` parameter.

<br>

# v0.7.1
_Released Dec 21, 2016_
- You can now specify a custom path to the `terraform` binary using the `TERRAGRUNT_TFPATH` environment variable or the `--terragrunt-tfpath` command-line option. This way, you can remove `terraform` from your `PATH`, which will help you avoid running it directly by accident and thereby bypassing the protections offered by Terragrunt.

<br>

# v0.7.0
_Released Dec 20, 2016_
- We've updated our CI build job to use Go 1.7.3. Before, we were using Go 1.6.x, which apparently [does not work with the latest version of OS X](https://golang.org/doc/go1.7#ports). We hope that this update will fix the crashes we were seeing in #41.

<br>

# v0.6.4
_Released Dec 20, 2016_
- Fixes #81, where `terragrunt spin-up` would only search in immediate subfolders, but not recursive subfolders.

<br>

# v0.6.3
_Released Dec 14, 2016_
- Fix a bug where the `spin-up` and `tear-down` commands did not properly call `terraform get` for modules that depended on other modules.

<br>

# v0.6.2
_Released Dec 9, 2016_
- Fix a bug in the `find_in_parent_folders` helper where it would not handle relative paths correctly.
- Fix documentation in the Terragrunt `--help` command.

<br>

# v0.6.1
_Released Dec 8, 2016_
- Fix typo in Terragrunt log output

<br>

# v0.6.0
_Released Dec 7, 2016_
- Terragrunt now supports `spin-up` and `tear-down` commands that automatically find all the Terraform modules in the current working directory and runs `terraform apply` or `terraform destroy` in each one, in the proper order, based on their dependencies. See the docs here for details: https://github.com/gruntwork-io/terragrunt#the-spin-up-and-tear-down-commands
- Terragrunt now supports defining `dependencies` in the `.terragrunt` file to specify which other modules should be deployed before the current one. See the docs here for details: https://github.com/gruntwork-io/terragrunt#dependencies-between-modules
- You can now specify the current working directory using the `--terragrunt-working-dir` param. See the docs here for details: https://github.com/gruntwork-io/terragrunt#cli-options

<br>

# v0.52.5
_Released Oct 24, 2023_
## Updated CLI args, config attributes and blocks

* `read_tfvars_file`

## Description

- Added `read_tfvars_file` function for parsing tfvar.tf or tfvar.json, returning JSON string.
 
## Special thanks

Special thanks to @alikhil for their contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2757

<br>

# v0.52.4
_Released Oct 20, 2023_
## Description

- CICD improvements: Added Go code linting and resolved code issues.
 
## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2743

<br>

# v0.52.3
_Released Oct 13, 2023_
## Description
Updated `golang.org/x/net` dependency to 0.17.0

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2755

<br>

# v0.52.2
_Released Oct 12, 2023_
## Updated CLI args, config attributes and blocks

* `remote_state`

## Description

- Added checking of `skip_credentials_validation` field in the remote state to skip AWS session validation.
 
## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2753

<br>

# v0.52.1
_Released Oct 8, 2023_
## Description

- Enhanced dependent module scanning to eliminate unnecessary user input prompts.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2749

<br>

# v0.52.0
_Released Oct 6, 2023_
## Description

- Added support for OpenTofu in Terragrunt, by default, will be wrapped `terraform` command with a fallback to `tofu`.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2745

## Migration guide

Switching to tofu:
1. Install OpenTofu cli https://github.com/opentofu/opentofu
2. Apply one of:
- Option 1: Remove `terraform` binary from `PATH`
- Option 2: Define env variable `TERRAGRUNT_TFPATH=tofu`
- Option 3: When launching `terragrunt`, specify `--terragrunt-tfpath tofu`


<br>

# v0.51.9
_Released Oct 4, 2023_
## Updated CLI args, config attributes and blocks

* `render-json`

## Description

- Added `dependent_modules` field to `render-json` output with a list of paths to dependent modules
- Replaced deprecated calls to io/ioutil package
 
## Special thanks

Special thanks to @maunzCache for their contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2740
* https://github.com/gruntwork-io/terragrunt/pull/2742

<br>

# v0.51.8
_Released Oct 3, 2023_
## Updated CLI args, config attributes and blocks

* `remote_state`

## Description

- Added AWS session status validation before S3 bucket status check
 
## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2739

<br>

# v0.51.7
_Released Sep 29, 2023_
## Updated CLI args, config attributes and blocks

* `terraform`

## Description

- Enhanced error handling during reading Terraform state.
 
## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2735


<br>

# v0.51.6
_Released Sep 28, 2023_
## Description

- Updated dependencies to fix CVE-2023-27561 and CVE-2022-27191.
 
## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2733


<br>

# v0.51.5
_Released Sep 26, 2023_
## Updated CLI args, config attributes and blocks

* `generate`

## Description

- Updated validation of generate blocks to identify blocks with the same name
 
## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2281


<br>

# v0.51.4
_Released Sep 21, 2023_
## Description

- Make `auto-init` feature respect `-no-color` flag
- Add a new built-in function `get_default_retryable_errors()` that returns a list of default retryable errors.
 
## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2725
* https://github.com/gruntwork-io/terragrunt/pull/2722


<br>

# v0.51.3
_Released Sep 19, 2023_
## Updated CLI args, config attributes and blocks

* `terragrunt-debug`

## Description

- Updated `terragrunt-debug` command to print valid executed `terraform` commands
 
## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2724


<br>

# v0.51.2
_Released Sep 19, 2023_
## Updated CLI args, config attributes and blocks

* `remote_state`

## Description

- Improved handling of S3 bucket configuration errors
 
## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2721


<br>

# v0.51.1
_Released Sep 18, 2023_
## Updated CLI args, config attributes and blocks

* `get_path_to_repo_root()`

## Description

- Removes trailing slash appended to the end of `get_path_to_repo_root()`

## Migration Guide

  * `source = "${get_path_to_repo_root()}/modules` -> `source = "${get_path_to_repo_root()}//modules`
  * `source = "${get_path_to_repo_root()}modules` -> `source = "${get_path_to_repo_root()}/modules`
 
## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2719


<br>

# v0.51.0
_Released Sep 16, 2023_
## Description

Upgraded CICD jobs to build and test Terragrunt using Go 1.21

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2715
<br>

# v0.50.9
_Released Aug 28, 2023_
## Updated CLI args, config attributes and blocks

* `remote_state`

## Description

Improved handling of AWS S3 errors: 
* Handling of errors when the bucket is in a different region
* Logging of underlying S3 API errors

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2684


<br>

# v0.50.8
_Released Aug 25, 2023_
## Updated CLI args, config attributes and blocks

* `terraform`

## Description

- Updated `tflint` hook to handle the passing of custom configuration file through `--config` argument.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2683


<br>

# v0.50.7
_Released Aug 25, 2023_
## Updated CLI args, config attributes and blocks

* `terragrunt console`

## Description

- Fixed the issue when using the `terragrunt console` command in non-interactive mode would result in an error.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2686


<br>

# v0.50.6
_Released Aug 21, 2023_
## Updated CLI args, config attributes and blocks

* `terraform`

## Description

- Addressed the issue related to GCP impersonation during state bucket creation.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2679


<br>

# v0.50.5
_Released Aug 21, 2023_
## Description

- Implementation of validation for provided Terraform commands.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2676

<br>

# v0.50.4
_Released Aug 17, 2023_
## Updated CLI args, config attributes and blocks

* `terraform`

## Description

- Eliminated generation of unnecessary `.terragrunt-null-vars.auto.tfvars.json` files.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2672




<br>

# v0.50.3
_Released Aug 16, 2023_
## Description

- Updated logrus dependency to `v1.9.3`

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2665

<br>

# v0.50.2
_Released Aug 16, 2023_
## Description

- Fixed updating S3 bucket policy when the previous `statement` contains `NotPrincipal` field.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2667

<br>

# v0.50.17
_Released Sep 14, 2023_
## Updated CLI args, config attributes and blocks

* `dependency`

## Description

- Added `enabled` property on `dependency` block, once set to `false` - dependency will be skipped.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2712


<br>

# v0.50.16
_Released Sep 13, 2023_
## Description

- Reduced "The handle is invalid" log noise for Terragrunt executions in Windows.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2711

<br>

# v0.50.15
_Released Sep 11, 2023_
## Updated CLI args, config attributes and blocks

* `--terragrunt-config`

## Description

- Fixed searching for config with the specified `--terragrunt-config` flag when running `run-all` commands.
- Added the ability to explicitly specify the name of dependency configuration files with a name other than the default `terragrunt.conf`.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2704


<br>

# v0.50.14
_Released Sep 6, 2023_
## Description

- Added passing of Terragrunt version in user agent for AWS API calls.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2701


<br>

# v0.50.13
_Released Sep 5, 2023_
## Updated CLI args, config attributes and blocks

* `terraform`

## Description

- Improved checking for module path in terraform source URL.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2700


<br>

# v0.50.12
_Released Sep 2, 2023_
## Updated CLI args, config attributes and blocks

* `--terragrunt-fetch-dependency-output-from-state`

## Description

- Fixed handling of `--terragrunt-fetch-dependency-output-from-state` option.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2699


<br>

# v0.50.11
_Released Sep 1, 2023_
## Updated CLI args, config attributes and blocks

* `--terragrunt-disable-command-validation`

## Description

- Added a new flag `--terragrunt-disable-command-validation` to disable terraform command validation.
- Prevent `terraform init` command from parallel running if the terraform plugin cache is used.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2697
* https://github.com/gruntwork-io/terragrunt/pull/2698


<br>

# v0.50.10
_Released Sep 1, 2023_
## Updated CLI args, config attributes and blocks

* `remote_state`

## Description

- Enhanced GCP storage configuration to include optional prefix support.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2695


<br>

# v0.50.1
_Released Aug 15, 2023_
## Updated CLI args, config attributes and blocks

* `--terragrunt-exclude-dir`

## Description

- Prevent parsing of excluded modules.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2664

<br>

# v0.50.0
_Released Aug 14, 2023_
## Updated CLI args, config attributes and blocks

* `terraform` [BACKWARD INCOMPATIBLE]

## Description

- Updated the process of passing inputs to Terraform to support inputs with `null` values.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2663

## Migration guide

To keep passing `null` values as strings, wrap values in quotes:

```
inputs = {
  input_1 = "null"
}
```


<br>

# v0.5.1
_Released Dec 5, 2016_
- Fix panic when an S3 bucket exists but doesn't have versioning enabled

<br>

# v0.5.0
_Released Dec 4, 2016_
- Terragrunt now stores the user ARN, rather than the user ID, in DynamoDB. The ARN typically includes a human-readable username, which will make it easier to tell who actually has a lock.

<br>

# v0.49.1
_Released Aug 11, 2023_
## Description

Update CircleCI config to sign MacOS binaries! 🎉 

## Related links

https://github.com/gruntwork-io/terragrunt/pull/2661
<br>

# v0.49.0
_Released Aug 11, 2023_
## Updated CLI args, config attributes and blocks

* `terraform` [BACKWARD INCOMPATIBLE]

## Description

- Added support for the `--terragrunt-external-tflint` parameter into the `tflint` hook. This modification enables Terragrunt to execute `tflint` directly from the operating system.
- Upgraded internal `tflint` to version `0.47.0`, introducing backward incompatibility.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2645

## Migration guide

Follow `tflint` recommendations to upgrade to `0.47.0` version:

https://github.com/terraform-linters/tflint/blob/v0.47.0/CHANGELOG.md#breaking-changes

<br>

# v0.48.7-test-signing-binaries
_Released Aug 11, 2023_
Testing changes from https://github.com/gruntwork-io/terragrunt/pull/2661.
<br>

# v0.48.7
_Released Aug 9, 2023_
## Updated CLI args, config attributes and blocks

* `--terragrunt-no-color`

## Description

* Added a new flag `--terragrunt-no-color` to disable colors in logs.
* Replaced environment variables (backward compatibility preserved):
`TERRAGRUNT_AUTO_INIT` with `TERRAGRUNT_NO_AUTO_INIT`
`TERRAGRUNT_AUTO_RETRY` with `TERRAGRUNT_NO_AUTO_RETRY`
`TERRAGRUNT_AUTO_APPROVE` with `TERRAGRUNT_NO_AUTO_APPROVE`
`TF_INPUT` with `TERRAGRUNT_NON_INTERACTIVE`

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2608
<br>

# v0.48.6
_Released Aug 2, 2023_
## Updated CLI args, config attributes and blocks

* `run-all`

## Description

- Improved readability of the `run-all` confirmation message.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2653
<br>

# v0.48.5
_Released Jul 28, 2023_

## Updated CLI args, config attributes and blocks

* `output-module-groups`

## Description

- Added `output-module-groups` command which outputs as JSON the groups of modules

## Special thanks

Special thanks to @smaftoul for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2130
<br>

# v0.48.4
_Released Jul 19, 2023_
## Updated CLI args, config attributes and blocks

* terraform cli

## Description

- Added `--terragrunt-fail-on-state-bucket-creation` which will fail Terragrunt execution if state bucket creation is required
- Added `--terragrunt-disable-bucket-update` which will disable updating of state bucket

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2644
<br>

# v0.48.3
_Released Jul 18, 2023_
## Updated CLI args, config attributes and blocks

* `render-json`

## Description

- Updated `render-json` to handle output values from not applied dependencies.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2635
<br>

# v0.48.2
_Released Jul 17, 2023_

## Updated CLI args, config attributes and blocks

* `--terragrunt-use-partial-parse-config-cache`

## Description

- Updated partial cache key construction to include filename

## Special thanks

Special thanks to @untcha for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2518
<br>

# v0.48.1
_Released Jul 5, 2023_
## Updated CLI args, config attributes and blocks

* `terraform` block

## Description

* Improved terraform source hash calculation by excluding files from `.terragrunt-cache` and `.terraform.lock.hcl`. These files are automatically generated and modifying their content will no longer trigger a `terraform init` execution.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2630

<br>

# v0.48.0
_Released Jun 26, 2023_
## Description

**Terraform 1.5 support**: We are now testing Terragrunt against Terraform 1.5 and is confirmed to be working.

NOTE: Although this release is marked as backward incompatible, it is functionally compatible as nothing has been changed in Terragrunt internals. The minor version release is useful to mark the change in Terraform version that is being tested.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2616

<br>

# v0.47.0
_Released Jun 15, 2023_
## Description

Upgraded CICD jobs to build and test Terragrunt using Go 1.20

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2605
<br>

# v0.46.3
_Released Jun 7, 2023_
## Updated CLI args, config attributes and blocks

* `terraform`

## Description

* Updated handling of `?ref=` in `terraform` block to read correctly value with slashes.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2597

<br>

# v0.46.2
_Released Jun 5, 2023_
## Updated CLI args, config attributes and blocks

* `--terragrunt-source-map`

## Description

* Added support git tag in URL for --terragrunt-source-map option.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2586

<br>

# v0.46.1
_Released Jun 1, 2023_
## Updated CLI args, config attributes and blocks

* `dependency`

## Description

* Fixed dependencies output reading during destroy execution.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2589

<br>

# v0.46.0
_Released Jun 1, 2023_
## Updated CLI args, config attributes and blocks

* `before_hook` [**BACKWARD INCOMPATIBLE**]
* `after_hook` [**BACKWARD INCOMPATIBLE**]
* `error_hook` [**BACKWARD INCOMPATIBLE**]

## Description

* Fixed redirection of stdout to stderr in hooks execution.

# Migration guide

To disable stdout redirection should be defined attribute `suppress_stdout = true` in the hook declaration:
```
terraform {
    ... 
  after_hook "after_init_from_module" {
    ...
    suppress_stdout = true
  }
  after_hook "after_init" {
    ...
    suppress_stdout = true
  }
  error_hook "handle_error" {
    ...
    suppress_stdout = true
  }
}
```

## Related Links
- https://github.com/gruntwork-io/terragrunt/pull/2587

<br>

# v0.45.9
_Released May 5, 2023_
## Description
Fixed the output parsing when AWS CSM is enabled.

## Special thanks

* Special thanks to @levkoburburas for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2554

<br>

# v0.45.8
_Released May 1, 2023_
## Updated CLI args, config attributes and blocks

* `--terragrunt-iam-role`

## Description

Handling of `--terragrunt-iam-role` flag has been updated to avoid evaluating `iam_role` attribute from HCL files once the `--terragrunt-iam-role flag` is specified

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2507

<br>

# v0.45.7
_Released May 1, 2023_
## Updated CLI args, config attributes and blocks

* `sops_decrypt_file`

## Description

Updated sops version to v3.7.3.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2549

<br>

# v0.45.6
_Released Apr 28, 2023_
## Updated CLI args, config attributes and blocks

* `--terragrunt-include-module-prefix`

## Description

Updated handling of `--terragrunt-include-module-prefix` to not include module prefix in case of `-json` argument

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2548

<br>

# v0.45.5
_Released Apr 27, 2023_
## Description
Updated Terragrunt to print explanations on Terraform error messages.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2547

<br>

# v0.45.4
_Released Apr 19, 2023_
## Description
Updated ownership setting on Terragrunt created S3 bucket.
Improved AWS Auth documentation.

## Special thanks

* Special thanks to @Kiran01bm for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2540
- https://github.com/gruntwork-io/terragrunt/pull/2533

<br>

# v0.45.3
_Released Apr 17, 2023_
## Updated CLI args, config attributes and blocks

* `--terragrunt-include-module-prefix`

## Description

Updated module prefix printing to show module name for non *-run-all commands

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2527

<br>

# v0.45.2
_Released Apr 5, 2023_
## Updated CLI args, config attributes and blocks

* `--terragrunt-include-module-prefix`

## Description

Updated module prefix printing to not affect fetching of module outputs

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2516

<br>

# v0.45.18
_Released May 31, 2023_
## Updated CLI args, config attributes and blocks

* `hclfmt`

## Description

Added flag `--terragrunt-diff` to `hclfmt` which will lead to printing of differences to output.

## Special thanks

* Special thanks to @okgolove for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2570

<br>

# v0.45.17
_Released May 30, 2023_
## Updated CLI args, config attributes and blocks

* `destroy`

## Description

Improved checking dependent modules during `destroy` execution

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2581

<br>

# v0.45.16
_Released May 24, 2023_
## Updated CLI args, config attributes and blocks

* `generate`

## Description

* Added `disable` field to `generate` blocks

## Special thanks

* Special thanks to @szesch for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2497
<br>

# v0.45.15
_Released May 19, 2023_
## Updated CLI args, config attributes and blocks

* `terraform`

## Description

Updated generation blocks locks to synchronize only tflint hooks execution and improve performance of "non-tflint" hooks.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2576

<br>

# v0.45.14
_Released May 18, 2023_
## Description

Updated Terragrunt to include module path in the error message.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2574

<br>

# v0.45.13
_Released May 17, 2023_
## Description

Added support for `timecmp` function in HCL files

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2571
<br>

# v0.45.12
_Released May 17, 2023_
## Updated CLI args, config attributes and blocks

* `terraform`

## Description

Fixed provider lock file handling when retrieving outputs from dependencies via remote state.

## Special thanks

* Special thanks to @geekofalltrades for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2568

<br>

# v0.45.11
_Released May 10, 2023_
## Description

Fixed sending second interrupt signal to Terraform

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2559

<br>

# v0.45.10
_Released May 8, 2023_
## Description

Fixed double-rendering of `terraform init` output

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2558

<br>

# v0.45.1
_Released Apr 3, 2023_
## Updated CLI args, config attributes and blocks

* `--terragrunt-include-module-prefix`

## Description

Added `--terragrunt-include-module-prefix` flag / `TERRAGRUNT_INCLUDE_MODULE_PREFIX` environment variable to include module dir prefix in Terraform output

## Special thanks

* Special thanks to @maciasello for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2493

<br>

# v0.45.0
_Released Mar 17, 2023_
## Description

**Terraform 1.4 support**: We are now testing Terragrunt against Terraform 1.4 and is confirmed to be working.

NOTE: Although this release is marked as backward incompatible, it is functionally compatible as nothing has been changed in Terragrunt internals. The minor version release is useful to mark the change in Terraform version that is being tested.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2489

<br>

# v0.44.5
_Released Mar 8, 2023_
## Description

- Updated `terragrunt` to copy in the download directory tflint configuration files.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2474

<br>

# v0.44.4
_Released Mar 2, 2023_
## Updated CLI args, config attributes and blocks

- `remote_state` [gcs]

## Description

- The `gcs` configuration for `remote_state` now supports `access_token` with OAuth 2.0 access token.

## Special thanks

Thank you to @marcportabellaclotet-mt for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2463
<br>

# v0.44.3
_Released Mar 2, 2023_
## Description
Updated `golang.org/x/net` dependency to 0.7.0

## Special thanks

Special thanks to @dependabot for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2455

<br>

# v0.44.2
_Released Mar 1, 2023_
## Description
Updated `go-getter` dependency to 1.7.0

## Special thanks

Special thanks to @dependabot for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2454

<br>

# v0.44.1
_Released Feb 28, 2023_
## Description 
* tflint update: Parallel TFLint execution errors by @denis256

## Related links
- https://github.com/gruntwork-io/terragrunt/pull/2469

<br>

# v0.44.0
_Released Feb 22, 2023_
## Description

Adds support for parsing retryable errors out of json output when the -json flag is included in commands passed to terraform. Fixes #2462 

##  Special Thanks
* @leighpascoe made their first contribution in https://github.com/gruntwork-io/terragrunt/pull/2464

## Related Links

* #2464 

## Migration Guide

Previously, when JSON output was used retryable errors would not be detected.  Now std output will be checked for retryable errors.  Regular Expressions will now match stdout as well as stderr when determining retry logic.
<br>

# v0.43.3
_Released Feb 22, 2023_
## Description

Fixed bug in loading of TFLint configuration files.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2456
<br>

# v0.43.2
_Released Feb 2, 2023_
## Updated CLI args, config attributes and blocks

* `remote_state`

## Description

Improved handling of default value for `remote_state` `bucket_sse_algorithm`.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2438

<br>

# v0.43.1
_Released Jan 31, 2023_
## Updated CLI args, config attributes and blocks

* `remote_state`

## Description

Improved handling of encryption setting for `remote_state` to not log misleading "out of date" messages in case of AES256 algorithm.

## Special thanks

* Special thanks to @kevcube for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2337

<br>

# v0.43.0
_Released Jan 19, 2023_
## Description 
* tflint updates:
  * Remove `GITHUB_OAUTH_TOKEN` environment variable based on `GITHUB_TOKEN`
  * Remove `--module` argument from default configuration 

## Special thanks
* Special thanks to @theurichde for their contribution!

## Related links
- https://github.com/gruntwork-io/terragrunt/pull/2424
- https://github.com/gruntwork-io/terragrunt/pull/2422

<br>

# v0.42.8
_Released Jan 13, 2023_
## Description

Updated `golang.org/x/crypto` dependency to fix CVE-2020-9283.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2412
<br>

# v0.42.7
_Released Jan 6, 2023_
## Updated CLI args, config attributes and blocks

* `dependency`

## Description

Updated evaluation of `mock_outputs_merge_strategy_with_state` in the `dependency` block to avoid crashes on empty mocks.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2379

<br>

# v0.42.6
_Released Jan 5, 2023_
## Updated CLI args, config attributes and blocks

* `locals`

## Description

Improved error message to include the file name in which locals evaluation failed.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2404

<br>

# v0.42.5
_Released Dec 13, 2022_
## Updated CLI args, config attributes and blocks

 * Fix debug log to be printed only whe `GITHUB_TOKEN` is exported by tflint

## Description
Previously, the logging for the `tflint` hook was issuing a message that the `GITHUB_TOKEN` was exported even if it wasn't. So this release fixes that.

## Related links

-  https://github.com/gruntwork-io/terragrunt/pull/2395
<br>

# v0.42.4
_Released Dec 13, 2022_
## Updated CLI args, config attributes and blocks

 * Added `tflint` as a built-in hook

## Description

Added `tflint` as a built-in hook. `Tflint` will now work as a built-in function of terragrunt when configured to do so, without needing to install and run it separately.

## Special thanks

Special thanks to @everops-miked for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2362
- https://github.com/gruntwork-io/terragrunt/pull/2383
- https://github.com/gruntwork-io/terragrunt/pull/2387
<br>

# v0.42.3
_Released Dec 2, 2022_
## Updated CLI args, config attributes and blocks

 * `run_cmd`

## Description

Added `--terragrunt-global-cache` option for `run_cmd` to cache output globally instead of per directory.

## Special thanks

Special thanks to @tjstansell for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2348
<br>

# v0.42.2
_Released Nov 29, 2022_
# Updated CLI args, config attributes and blocks

* `remote_state`

# Description

Improved setting of encryption to access logs bucket to be decoupled from encryption of state bucket - will be set `AES256` if encryption is enabled.

# Related links
 
- https://github.com/gruntwork-io/terragrunt/pull/2375
<br>

# v0.42.1
_Released Nov 29, 2022_
# Updated CLI args, config attributes and blocks

* `remote_state`

# Description

Added to `remote_state` configuration block attribute `accesslogging_bucket_tags` containing a map of tags which will be set on access logging bucket,

# Special thanks

Special thanks to @edgarsandi for their contribution!

# Related links
 
- https://github.com/gruntwork-io/terragrunt/pull/2355
<br>

# v0.42.0
_Released Nov 29, 2022_
## Updated CLI args, config attributes and blocks

* `remote_state` [**BACKWARD INCOMPATIBLE**]

## Description

Updated creation of S3 access logging bucket to configure TLS and SSE.
Since _only_ supported encryption algorithm is [AES256 (SSE-S3)](https://docs.aws.amazon.com/AmazonS3/latest/userguide/enable-server-access-logging.html) Terragrunt will stop with an error if will be attempted to create access logging bucket with a different algorithm.

# Migration guide

Define `bucket_sse_algorithm = AES256` for S3 remote state backends:
```
remote_state {
  backend  = "s3"
  ...
  config = {
    ...
    accesslogging_bucket_name      = "access-log-bucket-123"
    bucket_sse_algorithm           = "AES256"
  }
}


```


## Related Links

- https://github.com/gruntwork-io/terragrunt/pull/2367

<br>

# v0.41.0
_Released Nov 24, 2022_
## Description

Upgraded CICD jobs to build and test Terragrunt using Go 1.18

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2365
<br>

# v0.40.2
_Released Nov 10, 2022_
## Description
Upgraded `golang.org/x/text` and `golang.org/x/net` dependencies  to fix [CVE-2022-32149](https://nvd.nist.gov/vuln/detail/CVE-2022-32149) and [CVE-2022-27664](https://avd.aquasec.com/nvd/2022/cve-2022-27664/)

## Special thanks

Special thanks to @peterdeme for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2341

<br>

# v0.40.1
_Released Nov 7, 2022_
## Updated CLI args, config attributes and blocks

* `sops_decrypt_file`

## Description

Fixed Terragrunt crash when using SOPS secrets in parallel.

## Special thanks

Special thanks to @adongy for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2352

<br>

# v0.40.0
_Released Nov 4, 2022_
## Description

**Terraform 1.3 support**: We are now testing Terragrunt against Terraform 1.3 and is confirmed to be working.

NOTE: Although this release is marked as backward incompatible, it is functionally compatible as nothing has been changed in Terragrunt internals. The minor version release is useful to mark the change in Terraform version that is being tested.

## Related links

https://github.com/gruntwork-io/terragrunt/pull/2308

<br>

# v0.4.0
_Released Nov 29, 2016_
- ENHANCEMENT: Terragrunt now supports an `include` block and several helpers that you can use in the `.terragrunt` file to reuse settings from another `.terragrunt` file. This significantly reduces the amount of copy/paste and manual maintenance you have to do if you are maintaining multiple Terraform folders (and thereby, multiple `.terragrunt` files). For usage instructions, see: https://github.com/gruntwork-io/terragrunt/tree/inherit#managing-multiple-terragrunt-files

<br>

# v0.39.2
_Released Oct 18, 2022_
## What's Changed
* fix work with multiple aws accounts typo by @senmm in https://github.com/gruntwork-io/terragrunt/pull/2307
* Fix a mistake in markdown by @aazon in https://github.com/gruntwork-io/terragrunt/pull/2303
* Updated navbar - added "we're hiring" by @klijakub in https://github.com/gruntwork-io/terragrunt/pull/2313
* Navbar we are hiring fix by @klijakub in https://github.com/gruntwork-io/terragrunt/pull/2321

## New Contributors
* @senmm made their first contribution in https://github.com/gruntwork-io/terragrunt/pull/2307
* @aazon made their first contribution in https://github.com/gruntwork-io/terragrunt/pull/2303

**Full Changelog**: https://github.com/gruntwork-io/terragrunt/compare/v0.39.1...v0.39.2
<br>

# v0.39.1
_Released Oct 5, 2022_
## Updated CLI args, config attributes and blocks

* `render-json`
* `aws-provider-patch`
* `validate-inputs`

## Description

Added handling of `--help` option for Terragrunt commands

## Related Links

- https://github.com/gruntwork-io/terragrunt/pull/2297
<br>

# v0.39.0
_Released Sep 29, 2022_
## Updated CLI args, config attributes and blocks

* `render-json` [**BACKWARD INCOMPATIBLE**]

## Description

Improved `render-json` execution to use mock outputs of dependency if read of terraform outputs fails.
Updated docs to reflect run-all syntax.

## Related Links

- https://github.com/gruntwork-io/terragrunt/pull/2288
- https://github.com/gruntwork-io/terragrunt/pull/2266
<br>

# v0.38.9
_Released Aug 29, 2022_

## Updated CLI args, config attributes and blocks

* `--terragrunt-use-partial-parse-config-cache`

## Description

Improved parsing speed of HCL files by introducing an in-memory cache

## Special thanks

Special thanks to @maunzCache for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2204
<br>

# v0.38.8
_Released Aug 26, 2022_
## Updated CLI args, config attributes and blocks

* `remote_state`

## Description

* This introduces the `skip_bucket_public_access_blocking` configuration to the `remote_state` block for the S3 backend, when set to `true`, created S3 bucket will not have public access blocking enabled. This will flag may be useful when Terragunt defaults conflict with AWS organization policies.
* Improve render JSON tests to be more stable
* Documentation fixes: broken link, and spelling correction

## Special thanks

Special thanks to the following users for their contribution:
- @gpdenny
- @rufusnufus
- @conorevans

## Links

- https://github.com/gruntwork-io/terragrunt/pull/2246
- https://github.com/gruntwork-io/terragrunt/pull/2248
- https://github.com/gruntwork-io/terragrunt/pull/2250
- https://github.com/gruntwork-io/terragrunt/pull/2238

<br>

# v0.38.7
_Released Aug 8, 2022_
## Updated CLI args, config attributes and blocks

- `render-json` [CLI cmd]

## Description

- Fixed a bug in `render-json` where when running with `run-all`, it reuses the same rendered json file for all runs, causing a race condition where the last module to run always wins. Now the rendered json output is created adjacent to each `terragrunt.hcl` config `terragrunt` finds.

## Links 

- https://github.com/gruntwork-io/terragrunt/pull/2230
<br>

# v0.38.6
_Released Jul 22, 2022_
## Updated CLI args, config attributes and blocks

- `render-json` [CLI cmd]

## Description

- Updated `render-json`  command to support `--with-metadata`  option that will add to each JSON field metadata information

## Links 

- https://github.com/gruntwork-io/terragrunt/pull/2199

<br>

# v0.38.5
_Released Jul 13, 2022_
## Updated CLI args, config attributes and blocks

- `render-json` [CLI cmd]
- `locals` [Block]

## Description

- Fixed a bug in `render-json` where `dependency` blocks and `locals` information were lost in the resulting rendered config.
- Improved error logging in Terragrunt:
    - Terragrunt will now show the paths of the files where `locals` parsing failed.
    - Terragrunt will now show the paths of the modules where Terraform calls had errors.
- Updated dependencies that security patches.

## Links 

- https://github.com/gruntwork-io/terragrunt/pull/2175
- https://github.com/gruntwork-io/terragrunt/pull/2191
- https://github.com/gruntwork-io/terragrunt/pull/2192
<br>

# v0.38.4
_Released Jul 5, 2022_
## Updated CLI args, config attributes and blocks

- `terraform`

## Description

 * Faster approach to identify that local terraform sources have changes

## Special thanks

Special thanks to the following users for their contribution:
- @Sinjo

## Links 

- https://github.com/gruntwork-io/terragrunt/pull/2168
<br>

# v0.38.3
_Released Jun 30, 2022_
## Updated CLI args, config attributes and blocks

- `remote_state`

## Description

 * Fixed parsing of existing IAM policy response for S3 buckets
 * Added handling of empty boolean fields in S3 IAM policy
 * Improved stability of integration tests

## Links 

- https://github.com/gruntwork-io/terragrunt/pull/2111
- https://github.com/gruntwork-io/terragrunt/pull/2118
- https://github.com/gruntwork-io/terragrunt/pull/2160
<br>

# v0.38.2
_Released Jun 28, 2022_
## Updated CLI args, config attributes and blocks

- `terraform`

## Description

 * Updated logic in the handling of `error_hook` to match terraform stdout and stderr messages

## Links 

- https://github.com/gruntwork-io/terragrunt/pull/2101
<br>

# v0.38.12
_Released Sep 14, 2022_
## Description

Updated Terragrunt to retry in case of "Could not download module The requested URL returned error: 429" errors.

## Special thanks

Special thanks to @lorengordon for their contribution!

## Related Links

- https://github.com/gruntwork-io/terragrunt/pull/2276

<br>

# v0.38.11
_Released Sep 14, 2022_
## Description

Implemented usage of AWS partition from the current session when constructing KMS key ARN

## Special thanks

Special thanks to @lorengordon for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2257
<br>

# v0.38.10
_Released Sep 13, 2022_
## Description

Added passing of environment variables to dependencies when preparing list of modules to confirm destroy action

## Special thanks

Special thanks to @jlepere-everlaw for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2274
<br>

# v0.38.1
_Released Jun 22, 2022_
## Updated CLI args, config attributes and blocks

- `--terragrunt-no-auto-approve` [**NEW CLI FLAG**]
- `--terragrunt-fetch-dependency-output-from-state` [**NEW CLI FLAG**]
- `run-all` (cmd)
- `dependency` (block)

## Description

- Added new flag `--terragrunt-no-auto-approve` which will prevent Terragrunt from automatically including the `-auto-approve` flag to `apply` and `destroy` calls for `run-all`. In order to make the prompts work correctly, Terragrunt will also automatically set parallelism to `1`.
- Added new flag `--terragrunt-fetch-dependency-output-from-state` which will instruct Terragrunt to directly lookup the outputs from the state object in S3, by passing `terraform`. This could drastically speed up the dependency fetching routine. Note that this feature is **experimental** and **only supports the S3 backend**.
- Added the ability to configure the Server Side Encryption settings for the S3 Bucket, using two new parameters on the config:
    - `bucket_sse_algorithm`
    - `bucket_sse_kms_key_id`

## Special thanks

Special thanks to the following users for their contribution:
- @Ido-DY

## Links

- https://github.com/gruntwork-io/terragrunt/pull/2156
- https://github.com/gruntwork-io/terragrunt/pull/2123
- https://github.com/gruntwork-io/terragrunt/pull/2157
<br>

# v0.38.0
_Released Jun 17, 2022_
## Description

**Terraform 1.2 support**: We are now testing Terragrunt against Terraform 1.2 and is confirmed to be working.

NOTE: Although this release is marked as backward incompatible, it is functionally compatible as nothing has been changed in Terragrunt internals. The minor version release is useful to mark the change in Terraform version that is being tested.

<br>

# v0.37.4
_Released Jun 14, 2022_
## Updated CLI args, config attributes and blocks

* `terraform`

## Description

* Fixed handling of `include_in_copy` patterns to include files that aren't in the root of terraform module

## Special thanks

Special thanks to the following users for their contribution!
- @slawekzachcial

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2112


<br>

# v0.37.3
_Released Jun 10, 2022_
## Description

Improved local source code download behavior to generate a hash of local directory and copy only changed files.

## Special thanks

Special thanks to @BlackDark for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2006
<br>

# v0.37.2
_Released Jun 8, 2022_
## Description

Updated dependency `go-getter` to `1.6.1`.
<br>

# v0.37.1
_Released May 13, 2022_
## Description

**This release is functionally equivalent to v0.37.0**

Update the documentation to suggest `generate` blocks for managing the remote state backend configuration for users who do not want automated state management by Terragrunt.
<br>

# v0.37.0
_Released May 11, 2022_
## Updated CLI args, config attributes and blocks

* `remote_state` [**BACKWARD INCOMPATIBLE**]

## Description

* Updated logic for handling `remote_state` to enforce what is defined in the block of config, the change affects default encryption settings, public access blocking, bucket policy (enforce SSL only), access logging, and versioning

* Fixed behavior for enforcing `EnforcedTLS` to not overwrite already configured `RootAccess` policy

## Migration guide

If you do not want `terragrunt` to update the S3 bucket based on the configurations, you can define the config attribute `disable_bucket_update = true` in the `remote_state` block.

If you have an environment where it is difficult to update `terragrunt` and your configuration simultaneously, you can use [v0.36.12](https://github.com/gruntwork-io/terragrunt/releases/tag/v0.36.12) which supports the new `disable_bucket_update` configuration, but does not implement the bucket updating behavior.

To perform a safe upgrade:

- Upgrade your Terragrunt environment to `v0.36.12`
- Add `disable_bucket_update = true` in the `remote_state` block `config` map.
- Run `terragrunt` to confirm Terragrunt doesn't update the state buckets.
- Upgrade your Terragrunt environment to `v0.37.0` and confirm Terragrunt still does not update the state buckets.

## Special thanks

Special thanks to the following users for their contribution!
- @leonardobiffi 

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2063

<br>

# v0.36.9
_Released Apr 29, 2022_
## Description

* Fixed bug in the invocation of `git rev-parse --show-toplevel` to pass environment variables

## Special Thanks

* @lexton 

## References

https://github.com/gruntwork-io/terragrunt/pull/2086
<br>

# v0.36.8
_Released Apr 27, 2022_
## Description
* Updated `sops` version to v3.7.2 
* Clarified behavior of `terragrunt-tfpath` override with respect to dependencies
* Failing tests fixes

## Special Thanks
* @aymericDD
* @yorinasub17 

## References
* https://github.com/gruntwork-io/terragrunt/pull/2083
* https://github.com/gruntwork-io/terragrunt/pull/2077
* https://github.com/gruntwork-io/terragrunt/pull/2083
<br>

# v0.36.7
_Released Apr 15, 2022_
## Description

Updated Terragrunt to retry in case of "Timeout exceeded while awaiting headers" errors.

## Special thanks

Special thanks to @MFAshby for their contribution!

## Related Links

https://github.com/gruntwork-io/terragrunt/pull/2066

<br>

# v0.36.6
_Released Mar 17, 2022_
## Description

* Updated repo root functions to use platform agnostic path separators

## Special thanks

Special thanks to @henworth for their contribution!

## References

https://github.com/gruntwork-io/terragrunt/pull/2042

<br>

# v0.36.5
_Released Mar 16, 2022_
## Description

Updated Terragrunt to configure blocking of public access to the access logs S3 bucket when access logging of the state bucket is configured.

## Related Links

https://github.com/gruntwork-io/terragrunt/pull/2040
<br>

# v0.36.4
_Released Mar 16, 2022_
## Description

Added caching of IAM roles to improve parsing speed of HCL files.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2010
<br>

# v0.36.3
_Released Mar 2, 2022_
## Updated CLI args, config attributes and blocks

- `get_repo_root` [**NEW**]

## Description

* Introduced new function `get_repo_root`, that can be used to get the absolute path to the root of Git repository

## Special thanks

Special thanks to @andreykaipov for their contribution!

## References

https://github.com/gruntwork-io/terragrunt/pull/2027

<br>

# v0.36.2
_Released Feb 24, 2022_
## Description

Implemented support for merge strategies to control how to merge outputs with mocks.

## Special thanks

Special thanks to @jon-walton for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1936
<br>

# v0.36.12
_Released May 13, 2022_
## Updated CLI args, config attributes and blocks

* `remote_state`

## Description

* This introduces the `disable_bucket_update` configuration to the `remote_state` block for the S3 backend. This flag has **no functional effect**, but is useful to ease the transition to `v0.37.0`. Refer to [the release notes for v0.37.0](https://github.com/gruntwork-io/terragrunt/releases/tag/v0.37.0) for more details.


## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2105

<br>

# v0.36.11
_Released May 9, 2022_
## Description
Upgraded `go-getter` dependency to 1.5.11 with CVE-2022-29810 fixes

## Special thanks

Special thanks to @dependabot for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2095

<br>

# v0.36.10
_Released May 4, 2022_
## Description

- Fixed bug where `--terragrunt-strict-validate` ignored unused input variables
- Added missing links in documentation for CLI options

## Special thanks

Special thanks to the following users for their contribution!
- @raidancampbell
- @duboiss

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/2056
* https://github.com/gruntwork-io/terragrunt/pull/2088 

<br>

# v0.36.1
_Released Jan 31, 2022_
## Description

Fixed a bug in `tfr` source where relative paths returned from third party registries was not handled correctly.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/2005
<br>

# v0.36.0
_Released Jan 20, 2022_
## Description

**Terraform 1.1 support**: We are now testing Terragrunt against Terraform 1.1 and is confirmed to be working.

NOTE: Although this release is marked as backward incompatible, it is functionally compatible as nothing has been changed in Terragrunt internals. The minor version release is useful to mark the change in Terraform version that is being tested.


## Related links

https://github.com/gruntwork-io/terragrunt/pull/1990
<br>

# v0.35.9
_Released Nov 11, 2021_
## Updated CLI args, config attributes and blocks

- `run-all`
- `--terragrunt-modules-that-include` [**NEW OPTION**]

## Description

- Added new option `--terragrunt-modules-that-include` for `run-all` command which will restrict the `run-all` stack set to only those modules that include the given configuration file. This is useful for driving CI/CD workloads based on updates to common files that are included in child configurations.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1893
<br>

# v0.35.8
_Released Nov 9, 2021_
## Updated CLI args, config attributes and blocks

- `dependency`
- `dependencies`
- `include`

## Description

- Fixed bug where deep dependency merge did not properly merge the `config_path` when the `dependency` block was redefined in the child.

## Special thanks

Special thanks to @denis256 for their contribution!


## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1906
<br>

# v0.35.7
_Released Nov 9, 2021_
## Updated CLI args, config attributes and blocks

- `terraform.source`
- auto-init

## Description

- Fixed bug where `tfr` source did not handle registries that returned absolute URLs.
- Fixed regression bug where auto-init no longer handled updates to the `source` attribute of the `terraform` block.

## Special thanks

Special thanks to @denis256 for their contribution!


## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1907
- https://github.com/gruntwork-io/terragrunt/pull/1903
<br>

# v0.35.6
_Released Nov 1, 2021_
## Updated CLI args, config attributes and blocks

- `include`

## Description

Warning logs about partial parsing of included configurations have been converted to debug level.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1886
<br>

# v0.35.5
_Released Oct 28, 2021_
## Updated CLI args, config attributes and blocks

- `destroy`

## Description

Terragrunt will now log parsing errors during the dependency detection phase for `destroy` at the debug level.

## Special thanks

Special thanks to @denis256 for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1889
<br>

# v0.35.4
_Released Oct 19, 2021_
## Updated CLI args, config attributes and blocks

- `run-all`

## Description

Terragrunt will now log the order in which modules are deployed when using `run-all`, instead of all the modules in the graph including those that are excluded. You can get the old format logs if you use `--terragrunt-log-level debug`.

### Example module

```
%~> tree .
.
├── account-baseline
│   ├── main.tf
│   └── terragrunt.hcl
└── services
    ├── myapp
    │   ├── main.tf
    │   └── terragrunt.hcl
    ├── mysql
    │   ├── main.tf
    │   └── terragrunt.hcl
    ├── redis
    │   ├── main.tf
    │   └── terragrunt.hcl
    └── vpc
        ├── main.tf
        └── terragrunt.hcl
```

The following runs are done from the `services` folder, so we expect the `account-baseline` to be skipped.

#### Graph output before

```
%~> terragrunt run-all apply --terragrunt-non-interactive
INFO[0000] Stack at ~/gruntwork/support/terragrunt-run-all-destroy/proj/services:
  => Module ~/gruntwork/support/terragrunt-run-all-destroy/proj/account-baseline (excluded: false, dependencies: [])
  => Module ~/gruntwork/support/terragrunt-run-all-destroy/proj/services/myapp (excluded: false, dependencies: [~/gruntwork/support/terragrunt-run-all-destroy/proj/services/mysql, ~/gruntwork/support/terragrunt-run-all-destroy/proj/services/redis])
  => Module ~/gruntwork/support/terragrunt-run-all-destroy/proj/services/mysql (excluded: false, dependencies: [~/gruntwork/support/terragrunt-run-all-destroy/proj/services/vpc])
  => Module ~/gruntwork/support/terragrunt-run-all-destroy/proj/services/redis (excluded: false, dependencies: [~/gruntwork/support/terragrunt-run-all-destroy/proj/services/vpc])
  => Module ~/gruntwork/support/terragrunt-run-all-destroy/proj/services/vpc (excluded: false, dependencies: [~/gruntwork/support/terragrunt-run-all-destroy/proj/account-baseline])
```

#### Graph output with this release (apply)

```
%~> ~/gruntwork/tools/terragrunt/terragrunt run-all apply --terragrunt-non-interactive
INFO[0000] The stack at ~/gruntwork/support/terragrunt-run-all-destroy/proj/services will be processed in the following order for command apply:
Group 1
- Module ~/gruntwork/support/terragrunt-run-all-destroy/proj/services/vpc

Group 2
- Module ~/gruntwork/support/terragrunt-run-all-destroy/proj/services/mysql
- Module ~/gruntwork/support/terragrunt-run-all-destroy/proj/services/redis

Group 3
- Module ~/gruntwork/support/terragrunt-run-all-destroy/proj/services/myapp
```

#### Graph output with this release (destroy)

```
%~> ~/gruntwork/tools/terragrunt/terragrunt run-all destroy --terragrunt-non-interactive
INFO[0000] The stack at ~/gruntwork/support/terragrunt-run-all-destroy/proj/services will be processed in the following order for command destroy:
Group 1
- Module ~/gruntwork/support/terragrunt-run-all-destroy/proj/services/myapp

Group 2
- Module ~/gruntwork/support/terragrunt-run-all-destroy/proj/services/mysql
- Module ~/gruntwork/support/terragrunt-run-all-destroy/proj/services/redis

Group 3
- Module ~/gruntwork/support/terragrunt-run-all-destroy/proj/services/vpc
```


## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1861
<br>

# v0.35.3
_Released Oct 14, 2021_
## Description

Fixed bug where Terragrunt panicked with a nil pointer error for certain configurations.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1872
<br>

# v0.35.20
_Released Jan 14, 2022_
## Updated CLI args, config attributes, and blocks

- `error_hook` [**NEW**]

## Description
* Added 1 new block type, `error_hook`, which can be used to run arbitrary commands when an error is thrown that matches the `on_errors` list attribute.

## Special thanks

Special thanks to @smitthakkar96 for their contribution!

## References
* https://github.com/gruntwork-io/terragrunt/pull/1967
* https://github.com/gruntwork-io/terragrunt/pull/1980
* https://github.com/gruntwork-io/terragrunt/pull/1982
<br>

# v0.35.2
_Released Oct 13, 2021_
## Description

Fixed a regression bug introduced in `v0.35.1` where the `iam_role` config was ignored when managing the s3 state bucket in auto init.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1867
<br>

# v0.35.19
_Released Jan 13, 2022_
## Updated CLI args, config attributes and blocks

- `get_path_to_repo_root` [**NEW**]
- `get_path_from_repo_root` [**NEW**]

## Description
* Introduced two new functions, `get_path_to_repo_root` and `get_path_from_repo_root`, that can be used to construct relative paths in relation to the root of the git repository, if the `terragrunt.hcl` configuration lives in a git repository and `terragrunt` is being run from a clone of the repository.

## Special thanks

Special thanks to @tbell83 for their contribution!

## References
* https://github.com/gruntwork-io/terragrunt/pull/1954

<br>

# v0.35.18
_Released Jan 12, 2022_
## Description
* Fixed bug where terragrunt executed auto-init if no module is used 

## References
* #1921

<br>

# v0.35.17
_Released Jan 12, 2022_
## Description
* Canonical path usage for dependencies exclusion
* Fix typo in config-blocks-and-attributes docs
* fix some 404 links

## Special Thanks

Special thanks to the following users for their contribution!

* @brad-alexander
* @stratusjerry
* @denis256 

## References

 * https://github.com/gruntwork-io/terragrunt/pull/1968
 * https://github.com/gruntwork-io/terragrunt/pull/1958
 * https://github.com/gruntwork-io/terragrunt/pull/1959

<br>

# v0.35.16
_Released Dec 17, 2021_
## Description

- Fixed bug where `--terragrunt-modules-that-include` didn't account for includes that are relative path.
- Added the ability to use `sops_decrypt_file` with `ini`, `env`, and raw text files.

## Special thanks

Special thanks to the following users for their contribution!

- @denis256


## Reference

- https://github.com/gruntwork-io/terragrunt/pull/1948
- https://github.com/gruntwork-io/terragrunt/pull/1956
<br>

# v0.35.15
_Released Dec 17, 2021_
## Description

- Removed unused functions in code base
- Updated behavior of `terraform_binary`. Now `terragrunt` will always prefer the setting in the configuration. This works around the issue where you need mixed terraform binaries in your project, and `terragrunt` doesn't know which one to use when fetching dependencies.

## Special thanks

Special thanks to the following users for their contribution!

- @denis256
- @leonardobiffi


## Reference

- https://github.com/gruntwork-io/terragrunt/pull/1943
- https://github.com/gruntwork-io/terragrunt/pull/1946
<br>

# v0.35.14
_Released Dec 10, 2021_
## Description

- Added the ability to specify files that are always included in the working `terraform` module copy (via the new `include_in_copy` attribute in the `terraform` block).


## Reference

- https://github.com/gruntwork-io/terragrunt/pull/1935
<br>

# v0.35.13
_Released Nov 24, 2021_
## Description

- Updated `aws-sdk-go` to v1.41.7
- Fixed bug where Terragrunt would prefer local state files even if Terraform was configured with remote state.

## Special thanks

Special thanks to the following users for their contribution

- @adrien-f 
- @andreykaipov

## Reference

- https://github.com/gruntwork-io/terragrunt/pull/1884
- https://github.com/gruntwork-io/terragrunt/pull/1885
<br>

# v0.35.12
_Released Nov 19, 2021_
## Description

- Fixed bug where terragrunt would log a "Failed to detect where module is used" warning unnecessarily.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1915 
<br>

# v0.35.11
_Released Nov 18, 2021_
## Description

- Fixed bug where terragrunt could not handle args passed in as `--terragrunt-option=value`.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1922
<br>

# v0.35.10
_Released Nov 12, 2021_
## Updated CLI args, config attributes and blocks

- `prevent_destroy`

## Description

- Fixed bug where terragrunt was not treating `apply -destroy` as a destroy operation, so it was not taking into account the `prevent_destroy` flag.

## Special thanks

Special thanks to @denis256 for their contribution!


## Related links

- https://github.com/gruntwork-io/terragrunt/issues/1904
<br>

# v0.35.1
_Released Oct 13, 2021_
## Updated CLI args, config attributes and blocks

* `--terragrunt-iam-role`
* `--terragrunt-iam-assume-role-duration`
* `--terragrunt-iam-assume-role-session-name`
* `iam_role`
* `iam_assume_role_duration`
* `iam_assume_role_session_name`

## Description

**NOTE: There is a regression bug introduced in this version. We recommend using [v0.35.2](https://github.com/gruntwork-io/terragrunt/releases/tag/v0.35.2).**

The following bugs related to the assume IAM role features of terragrunt has been fixed:

- There was a bug in the logic for `--terragrunt-assume-role-session-name`, where the default session name was always used when the CLI option was not passed in due to the session name being set on the options struct with the default. Now the default is only used if it is not explicitly set by CLI flag or config.
- There was a bug in the logic for assuming IAM roles in the remote state initialization, where it did not use the session duration or the session name that was set on the CLI. Now the remote state initialization will use the session duration and session name passed through the CLI, if set.
- There was a bug in the logic for assuming IAM roles for the terragrunt AWS internal helper functions, where it did not use the session duration or the session name from the config or CLI. Now the AWS internal helpers will use the session duration and session name if set.
- There was a bug in the logic for dependency, where it did not use the session duration or the session name from the config or CLI. Now dependency fetching will use the session duration and the session name if set.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1858
<br>

# v0.35.0
_Released Oct 13, 2021_
## Updated CLI args, config attributes and blocks

* `dependencies` [**BACKWARD INCOMPATIBLE**]
* `dependency`
* `include`

## Description

* Fixed bug where Terragrunt only took the last `include` block into consideration for the dependency run graph. Now all `dependency` blocks defined across all `include` configurations will be taken into consideration.

## Migration guide

As a part of this change, the behavior of how `dependencies` blocks are merged together in the `shallow` merge strategy has been updated to be a deep merge - now all the paths defined in `dependencies` blocks across the included modules are always concatenated together rather than replaced. If you have a configuration that depended on the old behavior, you will need to update your configuration to take advantage of multiple include blocks to selectively include the parent `dependencies` block.

E.g., if you had the following configurations:

_parent terragrunt.hcl_
```hcl
dependencies {
  paths = ["../vpc", "../mysql"]
}

# ... other blocks ...
```

_child terragrunt.hcl_
```hcl
include "root" {
  path = find_in_parent_folders()
}

dependencies {
  paths = ["../new_vpc"] # intended to replace dependencies block in parent
}

# ... other blocks ...
```

You will want to update to the following:

_parent terragrunt.hcl_
```hcl
# ... other blocks ...
```

_dependencies parent terragrunt.hcl_
```
dependencies {
  paths = ["../vpc", "../mysql"]
}
```

_child terragrunt.hcl_
```hcl
include "root" {
  path = find_in_parent_folders()
}

dependencies {
  paths = ["../new_vpc"] # intended to replace dependencies block in parent
}

# ... other blocks ...
```

_child who wants dependencies_
```hcl
include "root" {
  path = find_in_parent_folders()
}

include "dependencies" {
  path = find_in_parent_folders("dependencies_parent_terragrunt.hcl")
}

# ... other blocks ...
```


## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1864
<br>

# v0.34.3
_Released Oct 11, 2021_
## Updated CLI args, config attributes and blocks

* `--terragrunt-iam-assume-role-session-name` [**NEW CLI OPTION**]
* `iam_assume_role_session_name` [**NEW CONFIG ATTR**]
* `iam_role`
* `dependency`

## Description

* Fixed bug where `iam_role` configuration functionality broke for dependencies.
* Updated `hclfmt` to only log files that were updated by the command by default. You can get the original output if you set `--terragrunt-log-level debug`.
* Fixed bug where Terragrunt prompts for `run-all plan` were not shown to the console.
* Added new configuration option that can be used to set the assume role session name when `iam_role` is configured.

## Special thanks

Special thanks to the following users for their contribution!

- @denis256 
- @kartikay101
- @pjuanda

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1856
- https://github.com/gruntwork-io/terragrunt/pull/1857
- https://github.com/gruntwork-io/terragrunt/pull/1846
- https://github.com/gruntwork-io/terragrunt/pull/1843
<br>

# v0.34.2
_Released Oct 11, 2021_
## Description

* Updated error messages related to `locals` parsing.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1819
<br>

# v0.34.1
_Released Oct 6, 2021_
## Description

* Fixed bug where the `include` related functions were not being correctly parsed when used in `locals` blocks.

## Special thanks

Special thanks to @denis256 for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1847
<br>

# v0.34.0
_Released Oct 1, 2021_
## Updated CLI args, config attributes and blocks

* `iam_role` [**BEHAVIOR CHANGE**]
* `get_aws_account_id` [**BEHAVIOR CHANGE**]
* `get_aws_caller_identity_arn` [**BEHAVIOR CHANGE**]
* `get_aws_caller_identity_user_id` [**BEHAVIOR CHANGE**]
* `locals` [**BEHAVIOR CHANGE**]
 
## Description

* Fixed bug where the `iam_role` config attribute was ignored when creating the S3 bucket during Auto-Init.
* Terragrunt now has an additional parsing stage to parse `iam_role` attribute early, so that it can use it to resolve each of the `get_aws_*` functions. This means that the `get_aws_*` functions now return answers after the IAM role has been assumed, whereas before it was always based on the . Note that this additional parsing stage means that `locals` are parsed an additional time, which may cause side effects if you are using `run_cmd` in `locals`, and the args are dynamic so as to bust the cache.


## Special thanks

Special thanks to @denis256 for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1807
<br>

# v0.33.2
_Released Sep 30, 2021_
## Updated CLI args, config attributes and blocks

* `render-json` [**NEW CMD**]
* `dependency`

## Description

* Terragrunt now has a new command `render-json` which can be used to render the json representation of the fully interpreted `terragrunt.hcl` config. This can be used for debugging purposes, as well as for enforcing policies using [OPA](https://www.openpolicyagent.org/).
* Reverted the change from [v0.32.2](https://github.com/gruntwork-io/terragrunt/releases/tag/v0.32.2), as it is intentional for terragrunt to error out on dependency blocks that have no outputs. If it is desired to allow empty outputs on dependency blocks, you should configure `mock_outputs` and `mock_outputs_allowed_terraform_commands` or `skip_outputs`.


## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1821
- https://github.com/gruntwork-io/terragrunt/pull/1837
<br>

# v0.33.1
_Released Sep 29, 2021_
## Updated CLI args, config attributes and blocks

* `run-all`

## Description

* `run-all` will now error with a more sensible error message if you are missing the Terraform command that Terragrunt should run on all modules.

## Special thanks

Special thanks to @denis256 for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1836
<br>

# v0.33.0
_Released Sep 28, 2021_
## Updated CLI args, config attributes and blocks

* `destroy` [**BEHAVIOR CHANGE**]
 
## Description

* Running `destroy` on a single module will now check for `dependency` references that point to that module. If there are any modules that reference the module being destroyed, `terragrunt` will warn you with the list of modules that point to the module being destroyed and prompt for a confirmation that you are ok destroying the module. **To avoid the prompt and restore previous behavior, you can pass in `--terragrunt-non-interactive`.**

## Special thanks

Special thanks to @denis256 for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1823
<br>

# v0.32.6
_Released Sep 28, 2021_
## Updated CLI args, config attributes and blocks

* `include`
 
## Description

* You can now access later stage configuration elements in exposed includes (e.g., `inputs`) if certain conditions are met. Refer to the [updated documentation](https://terragrunt.gruntwork.io/docs/reference/config-blocks-and-attributes/#include) for more info (specifically, section "Limitations on accessing exposed config").


## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1817
<br>

# v0.32.5
_Released Sep 28, 2021_
## Description

* Fixed bug where Terragrunt did not correctly handle generated remote state backend configurations in `run-all` commands.


## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1827
<br>

# v0.32.4
_Released Sep 23, 2021_
## Description

* Fixed bug where Terragrunt did not honor `terraform_binary` when fetching outputs for `dependency`.

## Special thanks

* Special thanks to @pulchart  for the bug fix!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1813
<br>

# v0.32.3
_Released Sep 20, 2021_
## Description

* Fixed bug where Terragrunt would break on Windows paths in dependencies

## Special thanks

* Special thanks to @denis256  for the bug fix!
## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1816
- https://github.com/gruntwork-io/terragrunt/pull/1775
<br>

# v0.32.2
_Released Sep 16, 2021_
## Description

* Fixed bug where Terragrunt would error out if no outputs were defined

## Special thanks

* Special thanks to @denis256  for the bug fix!
## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1811
<br>

# v0.32.1
_Released Sep 14, 2021_
## Updated CLI args, config attributes and blocks

* `dependency`

## Description

* Fixed bug where `mock_outputs_merge_with_state` ignored `mock_outputs_allowed_terraform_commands`

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1800
<br>

# v0.32.0
_Released Sep 14, 2021_
## Updated CLI args, config attributes and blocks

* `include`

## Description

**NOTE: This release is marked as backward incompatible, but there is no backward incompatible configuration changes. We expect all existing configuration to work as is. However, given the amount of changes internally to the `include` mechanism, we are marking this release as backward incompatible out of caution. If you encounter any issues using this release with existing configurations, please file [an issue on this repo](https://github.com/gruntwork-io/terragrunt/issues) as the intention is to limit backward incompatibilities.** 

* Added support for using multiple `include` blocks in a single configuration. Note that with this change, usage of a bare `include` block with no labels (e.g., `include {}`) is deprecated. It is recommended to update your configuration starting this release to attach a label to all `include` blocks (e.g., `include "root" {}`). You can learn more about multiple include blocks in the [updated documentation](https://terragrunt.gruntwork.io/docs/reference/config-blocks-and-attributes/#include).

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1804
<br>

# v0.31.9
_Released Sep 9, 2021_
## Updated CLI args, config attributes and blocks

* `source`
* `generate`


## Description

* Terragrunt will now support Terraform registries, such as GitLab, that return relative paths.
* Terragrunt will now detect and show an error for `generate` blocks with duplicate names.

## Special thanks

* Special thanks to @hown3d for the registry URL fix!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1787
- https://github.com/gruntwork-io/terragrunt/pull/1795
<br>

# v0.31.8
_Released Aug 30, 2021_
## Description

Fix a bug where `terragrunt` would error out if no `terraform` files were found because they were being generated using `generate` blocks.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1781
<br>

# v0.31.7
_Released Aug 24, 2021_
## Description

Update `go-getter` to `v1.5.7`.


## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1777
<br>

# v0.31.6
_Released Aug 23, 2021_
## Updated CLI args, config attributes and blocks

- `run_cmd` [helper function]

## Description

- Terragrunt's parser ends up parsing `terragrunt.hcl` files multiple times, which is normally harmless, but in the case of the `run_cmd` helper function, which can execute arbitrary code with side effects, this could cause problems. Therefore, we now cache `run_cmd` calls based on the arguments passed to them so that they are only executed once. 


## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1763
<br>

# v0.31.5
_Released Aug 18, 2021_
## Updated CLI args, config attributes and blocks

- `terraform` [block]

## Description

- Added support for fetching modules from any Terraform Registry using the new `tfr://` protocol syntax for the `source` attribute. See the updated [docs on `source`](https://terragrunt.gruntwork.io/docs/reference/config-blocks-and-attributes/#terraform) for more details.


## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1767
<br>

# v0.31.4
_Released Aug 13, 2021_
## Updated CLI args, config attributes and blocks

- `dependency` [block]

## Description

- Added `mock_outputs_merge_with_state` option. When `true`, `dependency` fetching will always merge in `mock_outputs` into  the dependency output.

## Special thanks

Special thanks to @gregorycuellar for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1765
<br>

# v0.31.3
_Released Jul 30, 2021_
## Updated CLI args, config attributes and blocks

- `include` [block]

## Description

- Added `deep` merge strategy for `include`. Refer to [the updated documentation](https://terragrunt.gruntwork.io/docs/reference/config-blocks-and-attributes/#include) for more information on what deep merge means.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1759
<br>

# v0.31.2
_Released Jul 29, 2021_
## Updated CLI args, config attributes and blocks

- `include` [block]

## Description

- Added new attribute to `include`: `merge_strategy`. `merge_strategy` indicates how the included parent config should be merged with the child config. Currently, this only supports `no_merge` and `shallow`. When omitted, the merge strategy defaults to `shallow` (the same strategy as previous versions).

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1724
<br>

# v0.31.11
_Released Sep 13, 2021_
## Updated CLI args, config attributes and blocks

* `--terragrunt-log-level`


## Description

* Terragrunt will now honor the log level set using the environment variable `TERRAGRUNT_LOG_LEVEL`. This log level will also be used on the global fallback logger, which will log out stack traces on error.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1802
<br>

# v0.31.10
_Released Sep 9, 2021_
## Updated CLI args, config attributes and blocks

* `source`
* `dependencies`


## Description

* Terragrunt will now show an error if `source` or `dependencies` are referencing a folder that doesn't exist.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1796
- https://github.com/gruntwork-io/terragrunt/pull/1797
<br>

# v0.31.1
_Released Jul 20, 2021_
## Updated CLI args, config attributes and blocks

- CLI args

## Description

- Addressed bug where plan file args were not always passed to the end of the arg list when calling `terraform`. Now `terragrunt` will check all the args and determine if an arg is a plan file (a filename that exists on disk and ends with extension `tfplan`), and if it is, feed it to the end of the args list.

## Special thanks

- A special thank you to @PertsevRoman for the fix!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1740
<br>

# v0.31.0
_Released Jun 22, 2021_
## Description

- **Terraform 1.0 support**: We are now testing Terragrunt against Terraform 1.0 and is confirmed to be working.
- Terraform functions have been updated to the versions shipped with Terraform 0.15.3 (previously the functions were pulled in from Terraform 0.12.24). These may include backward incompatibilities. Refer to the terraform release notes for more information.


## Related links

https://github.com/gruntwork-io/terragrunt/pull/1726
<br>

# v0.30.7
_Released Jun 19, 2021_
## Updated CLI args, config attributes and blocks

- `include` [block]

## Description

Fix bug where using an exposed `include` with `local` in the same expression did not work when referencing in `locals` blocks.

## Related links

https://github.com/gruntwork-io/terragrunt/issues/1727
https://github.com/gruntwork-io/terragrunt/pull/1728
<br>

# v0.30.6
_Released Jun 18, 2021_
## Updated CLI args, config attributes and blocks

- `remote_state` [block]

## Description

- Fix the way the `remote_state` block handles the AWS partition settings so that it works correctly with GovCloud.

## Special thanks

- A special thank you to @ryno75 for the fix!

## Related links

- https://github.com/gruntwork-io/terragrunt/issues/1718
<br>

# v0.30.5
_Released Jun 18, 2021_
## Updated CLI args, config attributes and blocks

- `include` [block]

## Description

Fix bug where exposing `include` did not work when referencing in `locals` blocks.

## Related links

https://github.com/gruntwork-io/terragrunt/issues/1721
https://github.com/gruntwork-io/terragrunt/pull/1723
<br>

# v0.30.4
_Released Jun 17, 2021_
## Updated CLI args, config attributes and blocks

- `include` [block]

## Description

You can now access values from included config. E.g., if you want to access a local var `region` defined in the parent terragrunt config, you can reference `include.locals.region` in the child config.

Note that there are a few limitations/differences with `read_terragrunt_config`:

- `include` references do not include fetched dependencies. This will change in the future.
- `include` references are not automatically available. You must set the new `expose` attribute to `true` to access the included references.
- At the moment, you can only have a single `include` block in the child, and you can only `include` one level deep (no nested `include`s). This will change in the future.

(This is the first of several features that implement [the Imports RFC](https://terragrunt.gruntwork.io/docs/rfc/imports/))


## Related links

https://github.com/gruntwork-io/terragrunt/issues/1566
https://github.com/gruntwork-io/terragrunt/pull/1716
<br>

# v0.30.3
_Released Jun 11, 2021_
## Updated CLI args, config attributes and blocks

- `run-all` [command]

## Description

Improved error messaging when multiple errors are returned.

## Special thanks

Special thanks to @derom  for their contribution!

## Related links

https://github.com/gruntwork-io/terragrunt/pull/1703
<br>

# v0.30.2
_Released Jun 11, 2021_
## Updated CLI args, config attributes and blocks

- `aws-provider-patch` (command)

## Description

Improve error messages in `aws-provider-patch` when the json input is malformed.

## Related links

https://github.com/gruntwork-io/terragrunt/pull/1715
<br>

# v0.30.1
_Released Jun 11, 2021_
## Updated CLI args, config attributes and blocks

- `hclfmt` (command)

## Description

Updated documentation and help text in `hclfmt` command to clarify that it works on all files with `hcl` extension, not just `terragrunt.hcl`.

## Special thanks

Special thanks to @edgarsandi  for their contribution!

## Related links

https://github.com/gruntwork-io/terragrunt/pull/1713
<br>

# v0.30.0
_Released Jun 11, 2021_
## Updated CLI args, config attributes and blocks

- `aws-provider-patch` (command)

## Description

`aws-provider-patch` now supports additional data types. Previously `aws-provider-patch` only supported patching strings, which made it impossible to patch provider attributes that are not strings (e.g., the `allowed_account_ids` attribute of the `aws` provider, which is `list(string)` type).

Note that to support this, the `aws-provider-patch` now expects attribute values to be json encoded when passed in. That means that you need to quote the values in order for it to work. For example, if you previously ran:

```
terragrunt aws-provider-patch --terragrunt-override-attr region=us-east-2
```

you need to update the call to:

```
terragrunt aws-provider-patch --terragrunt-override-attr 'region="us-east-2"'
```


## Related links

https://github.com/gruntwork-io/terragrunt/pull/1714
https://github.com/gruntwork-io/terragrunt/issues/1709
<br>

# v0.3.0
_Released Nov 23, 2016_
- Terragrunt now supports an `acquire-lock` command that can be used to create long-term locks.

<br>

# v0.29.9
_Released Jun 7, 2021_
## Updated CLI args, config attributes and blocks

- `--terragrunt-include-external-dependencies`

## Description

You can now configure the `--terragrunt-include-external-dependencies` setting via the environment variable `TERRAGRUNT_INCLUDE_EXTERNAL_DEPENDENCIES`.


## Special thanks

Special thanks to @elebertus for the contribution!

## Related links

https://github.com/gruntwork-io/terragrunt/pull/1548
<br>

# v0.29.8
_Released Jun 2, 2021_
## Updated CLI args, config attributes and blocks

- `--terragrunt-debug`

## Description

You can now control the `--terragrunt-debug` flag using the `TERRAGRUNT_DEBUG` environment variable.


## Related links

https://github.com/gruntwork-io/terragrunt/pull/1698
<br>

# v0.29.7
_Released May 28, 2021_
## Updated CLI args, config attributes and blocks

- `iam_role`

## Description

Fix a bug where Terragrunt would not properly assume the IAM role specified via the `iam_role` parameter if you were using AWS SSO.

## Special thanks

Thank you to @stevie- for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1689
<br>

# v0.29.6
_Released May 25, 2021_
## Updated CLI args, config attributes and blocks

- `--help`


## Description

Update the usage text for Terragrunt to reflect that options should go after the command. There should be no impact on Terragrunt's behavior in this release.

## Special thanks

Thank you to @Tarasovych for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1687
<br>

# v0.29.5
_Released May 24, 2021_
## Updated CLI args, config attributes and blocks

- `--terragrunt-source-map` [CLI Arg]


## Description

You can now configure the Terragrunt source map option using the environment variable `TERRAGRUNT_SOURCE_MAP`. You can configure multiple mappings using comma separated value encoding. For example, the following configures three mappings:

```
github.com/org/modules.git=/local/path/to/modules,gitlab.com/org-lab/modules.git=/local/path/to/modules,bitbucket.org/bitorg/modules.git=/local/path/to/modules,gitlab.com/org-lab/modules.git=/local/path/to/modules
```

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1676
<br>

# v0.29.4
_Released May 20, 2021_
## Updated CLI args, config attributes and blocks

- `--terragrunt-iam-assume-role-duration` [new CLI Arg]
- `iam_assume_role_duration` [new config]

## Description

You can now use the new CLI arg and config setting to configure the duration for the IAM role from `--terragrunt-iam-role`.

## Special thanks

Thank you to @thehunt33r for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1673
<br>

# v0.29.3
_Released May 15, 2021_
## Updated CLI args, config attributes and blocks

- `--terragrunt-source-map` [CLI Arg]


## Description

This release introduces `--terragrunt-source-map`, which can be used to provide multiple mappings to translate terragrunt source URLs in the config with another path. See the [documentation](https://terragrunt.gruntwork.io/docs/reference/cli-options/#terragrunt-source-map) for more information.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1674
<br>

# v0.29.2
_Released May 3, 2021_
## Updated CLI args, config attributes and blocks

- `get_aws_account_id` [func]
- `get_aws_caller_identity_arn` [func]
- `get_aws_caller_identity_user_id` [func]


## Description

- Fixes bug where the default credentials chain for the `get_aws_**` functions ignored the config file (`~/.aws/config`).

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1663
<br>

# v0.29.10
_Released Jun 7, 2021_
## Updated CLI args, config attributes and blocks

- `yamldecode` (helper function)

## Description

Fix bug where `yamldecode` is unable decode certain forms of yaml.


## Special thanks

Special thanks to @andreykaipov for the contribution!

## Related links

https://github.com/gruntwork-io/terragrunt/pull/1706
<br>

# v0.29.1
_Released Apr 28, 2021_
## Updated CLI args, config attributes and blocks

- remote_state ; s3

## Description

- Fixes bug where AWS profile based authentication was not fully honored when automatic s3 bucket creation was happening.

## Related links

* https://github.com/gruntwork-io/terragrunt/issues/1650
* https://github.com/gruntwork-io/terragrunt/pull/1654

<br>

# v0.29.0
_Released Apr 23, 2021_
## Updated CLI args, config attributes and blocks

- (none)

## Description

- Update Terragrunt to work with Terraform 0.15. There are no code or behavior changes, but marking this release as incompatible because from this release onwards, we are only testing with Terraform 0.15 and up.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1651

<br>

# v0.28.9
_Released Mar 11, 2021_
## Updated CLI args, config attributes and blocks

- AWS Go SDK version

## Description

- Updated the version of the AWS Go SDK used in Terragrunt. This should allow Terragrunt to work with AWS SSO / AWS CLI v2.

## Special thanks

- Thank you to @z0mbix for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1537
<br>

# v0.28.8
_Released Mar 8, 2021_
## Updated CLI args, config attributes and blocks

- Logging for `run-all`

## Description

- Terragrunt will now remove extraneous newlines from the log output of the `run-all` command.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1541
<br>

# v0.28.7
_Released Feb 18, 2021_
## Updated CLI args, config attributes and blocks

- `run_cmd` [func]

## Description

- `run_cmd` will now automatically chomp trailing whitespace (including newlines) from the returned output.

## Special thanks

Special thanks to @twang817 and @tbell83 for their contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/728
* https://github.com/gruntwork-io/terragrunt/pull/1553
<br>

# v0.28.6
_Released Feb 17, 2021_
## Updated CLI args, config attributes and blocks

- Logging

## Description

- In [`v0.28.0`](https://github.com/gruntwork-io/terragrunt/releases/tag/v0.28.0), when we switched Terragrunt to use a proper logger that supported different log levels, we set the default log level to `WARN`. The intention was to make Terragrunt's logging less noisy, but as an unintended side effect, we ended up suppressing some log messages that are important to see. This release changes the default log level to `INFO` and should fix most of the logging issues.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1541
<br>

# v0.28.5
_Released Feb 16, 2021_
## Updated CLI args, config attributes and blocks

- `auto-init`

## Description

- Terragrunt will no longer error and continue to run commands when you disable AutoInit but init is needed.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1549
<br>

# v0.28.4
_Released Feb 12, 2021_
## Updated CLI args, config attributes and blocks

- `run-all` [CLI]

## Description

- Fix bug where `run-all` no longer allowed positional args.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1544
<br>

# v0.28.3
_Released Feb 10, 2021_
## Updated CLI args, config attributes and blocks

- `dependency` [block]

## Description

- This release introduces a further optimization to `dependency` output fetching where it will skip the `init` phase if it detects that the module has previously been locally initialized.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1533
<br>

# v0.28.24
_Released Apr 16, 2021_
## Updated CLI args, config attributes and blocks

- `skip_bucket_versioning`

## Description

- If `skip_bucket_versioning` is set to `true`, and you are using GCS as a backend, Terragrunt will not only not enable versioning automatically, but now it will also no longer try to check if versioning is enabled either. 

## Special thanks

- Thank you to @davidalger for the fix!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1610

<br>

# v0.28.23
_Released Apr 16, 2021_
## Updated CLI args, config attributes and blocks

- `get_terraform_commands_that_need_locking()`

## Description

- Remove `init` from the list of commands returned by `get_terraform_commands_that_need_locking()`, as `init` does not support locking, and as of Terraform 0.15, will exit with an error if you try to use the lock parameters with it.

## Special thanks

- Thank you to @grimm26 for the fix!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1642

<br>

# v0.28.22
_Released Apr 15, 2021_
## Updated CLI args, config attributes and blocks

- `sops_decrypt_file()`

## Description

- Updated the versions of `sops`, `aws-sdk-go`, and `vault` libraries that we depend on. As a result, the `sops_decrypt_file()` function should now work with data encrypted via HashiCorp Vault. 

## Special thanks

- Thank you to @teamfighter for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1564

<br>

# v0.28.21
_Released Apr 13, 2021_
## Updated CLI args, config attributes and blocks

- `get_original_terragrunt_dir()` **[NEW]**
- `generate`
- `remote_state`

## Description

- Added a new `get_original_terragrunt_dir()` helper, which returns the directory where the original Terragrunt configuration file (by default `terragrunt.hcl`) lives. This is primarily useful when one Terragrunt config is being read from another: e.g., if `/terraform-code/terragrunt.hcl` calls `read_terragrunt_config("/foo/bar.hcl")`, and within `bar.hcl`, you call `get_original_terragrunt_dir()`, you'll get back `/terraform-code`.
- Updated the `generate` and `remote_state` settings so that they can be set either as blocks or attributes. This makes it possible to, for example, read these settings from `common.hcl` using `read_terragrunt_config` and set them dynamically.


## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1637
* https://github.com/gruntwork-io/terragrunt/pull/1639

<br>

# v0.28.20
_Released Apr 12, 2021_
## Updated CLI args, config attributes and blocks

- `--terragrunt-strict-include`

## Description

- When you pass in `--terragrunt-strict-include`, Terragrunt will now _only_ execute within the directories passed in via `--terragrunt-include-dir`. If you set `--terragrunt-strict-include`, but don't pass in any directories via `--terragrunt-include-dir`, then Terragrunt will exit without doing anything. This is arguably a backwards incompatible change, but this is the behavior the `--terragrunt-strict-include` flag was intended to have originally, and is less surprising, so we're treating this as a bug fix. 

## Special thanks

- Thank you to @celestialorb for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1631

<br>

# v0.28.2
_Released Feb 4, 2021_
## Updated CLI args, config attributes and blocks

- (interactive prompts)

## Description

- Fix a bug introduced via the new logging functionality in `v0.28.0`, where, depending on your log level, Terragrunt's interactive prompts may have been hidden. Interactive prompts should now be visible regardless of log level.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1526
<br>

# v0.28.19
_Released Apr 8, 2021_
## Updated CLI args, config attributes and blocks

- `retry_max_attempts` **[NEW]**
- `retry_sleep_interval_sec` **[NEW]**

## Description

- Updated Terragrunt's [auto retry functionality](https://terragrunt.gruntwork.io/docs/features/auto-retry/) so that you can now configure the number of retry attempts and the time between retries using the new config attributes `retry_max_attempts` and `retry_sleep_interval_sec`, respectively.
- Switch log level of terragrunt dependency fetching informational messages to `debug`.

## Special thanks

- Thank you to @andreykaipov for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1632
* https://github.com/gruntwork-io/terragrunt/pull/1628

<br>

# v0.28.18
_Released Mar 29, 2021_
## Updated CLI args, config attributes and blocks

- `validate-inputs` [CLI Command]

## Description

- Fixes bug where `default = null` was viewed as a required input variable by `validate-inputs`.


## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1613

<br>

# v0.28.17
_Released Mar 29, 2021_
## Updated CLI args, config attributes and blocks

- Auto Init 

## Description

- When automatically calling `terraform init`, Terragrunt will no longer pass the deprecated `-get-plugins` flag.

## Special thanks

Special thanks to @davidalger for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1618

<br>

# v0.28.16
_Released Mar 24, 2021_
## Updated CLI args, config attributes and blocks

- `get_terragrunt_source_cli_flag` [func] [**NEW**]
- `terraform.source` [attr]

## Description

- Add a new helper function `get_terragrunt_source_cli_flag` for getting the value passed for the `--terragrunt-source` arg. This is useful for implementing various logic that depends on whether terragrunt is running in local dev mode or not.

- The `source` getter for terragrunt now supports multiple forced getters, such as `git-remote-codecommit` URLs.

## Special thanks

Special thanks to @suhussai for their contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1575
* https://github.com/gruntwork-io/terragrunt/pull/1594

<br>

# v0.28.15
_Released Mar 17, 2021_
## Updated CLI args, config attributes and blocks

- `remote_state` [S3 config] [block]

## Description

You can now specify a custom DynamoDB endpoint on the remote state configuration for S3 using the `dynamodb_endpoint` attribute in the config.

## Special thanks

Special thanks to @alxrem for their contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1586
* https://github.com/gruntwork-io/terragrunt/issues/875
* https://github.com/gruntwork-io/terragrunt/pull/876
<br>

# v0.28.14
_Released Mar 16, 2021_
## Updated CLI args, config attributes and blocks

- `terraform.before_hook` [subblock]
- `terraform.after_hook` [subblock]

## Description

Hook configurations now accept a `working_dir` attribute to specify where the command should run.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1588
<br>

# v0.28.13
_Released Mar 16, 2021_
## Description

We have updated the `creack/pty` dependency to to version 1.1.11. This will hopefully fix the "Setctty set but Ctty not valid in child" error when using `terragrunt console`.

## Special thanks

Thank you to @abohne for the fix!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1582
<br>

# v0.28.12
_Released Mar 15, 2021_
## Description

Starting this release, we will be publishing binaries for the `darwin/arm64` (compatible with Mac M1 chips) and `linux/arm64` platforms.

## Related links
* https://github.com/gruntwork-io/terragrunt/pull/1585
<br>

# v0.28.11
_Released Mar 12, 2021_
## Updated CLI args, config attributes and blocks

- `validate-inputs` [**new command**] 

## Description

This release introduces a new command `validate-inputs`. You can use this to sanity check the inputs you are passing through terragrunt against the available variables in the terraform module. Refer [to the command docs](https://terragrunt.gruntwork.io/docs/reference/cli-options/#validate-inputs) for more info.

## Related links
* https://github.com/gruntwork-io/terragrunt/pull/1572
<br>

# v0.28.10
_Released Mar 11, 2021_
## Updated CLI args, config attributes and blocks

- `remote_state` 

## Description

Added a new property `accesslogging_target_prefix` to enable control over the `TargetPrefix` setting when Server Access Logging is enabled for Remote TF State S3 buckets. 

Note: This attribute won’t take effect if the `accesslogging_bucket_name` attribute is not present.

More details can be found further in the [official Terragrunt docs](https://terragrunt.gruntwork.io/docs/reference/config-blocks-and-attributes/#remote_state)

```
remote_state {
  backend = "s3"
  config = {
    bucket         = "my-terraform-state"
    ...
    ...
    accesslogging_bucket_name = "valid_string_for_S3_bucket_bame"
    accesslogging_target_prefix = "valid_string" # NEW PROPERTY
  }
}
```

## Special thanks
- @dsecik for raising this issue!

## Related links
* https://github.com/gruntwork-io/terragrunt/pull/1507
* Issue linked: gruntwork-io/terragrunt#1437
* [Official AWS docs for TargetPrefix](https://docs.aws.amazon.com/AmazonS3/latest/API/API_LoggingEnabled.html)
* [Official AWS docs for Server Access Logging](https://docs.aws.amazon.com/AmazonS3/latest/userguide/enable-server-access-logging.html)
<br>

# v0.28.1
_Released Feb 2, 2021_
## Updated CLI args, config attributes and blocks

- `run-all` [**new CLI command**]
- `apply-all` [**deprecated**]
- `plan-all` [**deprecated**]
- `output-all` [**deprecated**]
- `destroy-all` [**deprecated**]
- `validate-all` [**deprecated**]

## Description

- Terragrunt now has a generic `run-all` command that allows you to run arbitrary terraform commands against a terragrunt stack. For example, you can run `terraform init` against a stack using `terragrunt run-all init`.
- `xxx-all` commands have now been deprecated in favor of `run-all`. Each command still works for backwards compatibility, but may be removed in the future.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1520
<br>

# v0.28.0
_Released Feb 2, 2021_
## Updated CLI args, config attributes and blocks

- `--terragrunt-log-level` [new CLI argument]

## Description

- Terragrunt now has improved logging! Each log message is now set with a different log level, and the default log level is set to `warn`, so by default, you'll get much less noise from Terragrunt's logging. You can configure the log level using the new `--terragrunt-log-level` argument. There should be no behavior difference, but since it's such a drastic difference in how logs look, marking this as backwards incompatible just in case. 

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1510
<br>

# v0.27.4
_Released Feb 1, 2021_
## Updated CLI args, config attributes and blocks

- `aws-provider-patch` [CLI command]

## Description

- Fix a bug in `aws-provider-patch` where it outputted invalid hcl in certain cases.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1519
<br>

# v0.27.3
_Released Jan 26, 2021_
## Updated CLI args, config attributes and blocks

- `providers lock` [CLI command]

## Description

- Terragrunt will now copy the `.terraform.lock.hcl` file not only when you run the `init` command, but also on the `providers lock` command.

## Special thanks

- Thank you to @ZymoticB for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1504
<br>

# v0.27.2
_Released Jan 25, 2021_
## Updated CLI args, config attributes and blocks

- `--terragurnt-debug` [CLI arg]

## Description

- This release fixes a bug where `--terragrunt-debug` was not always interpreted correctly.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1503
<br>

# v0.27.1
_Released Jan 12, 2021_
## Updated CLI args, config attributes and blocks

- `remote_state` [block]

## Description

- This release fixes a bug where the generated code from `remote_state` blocks in generate mode was nondeterministic. The backend configuration is now deterministically generated in sorted key order.

## Special thanks

Special thanks to the following users for their contribution:

- @Jerazol
- @jonathansp
- @SCTommyG

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1165
<br>

# v0.27.0
_Released Jan 8, 2021_
## Updated CLI args, config attributes and blocks

- Lock file handling **[NEW FUNCTIONALITY]**

## Description

- This release updates Terragrunt to work with Terraform 0.14!
    - Terragrunt now supports handling Terraform lock files. Check out the [lock file handling documentation](https://terragrunt.gruntwork.io/docs/features/lock-file-handling/) for details.
    - Fix a bug with Terragrunt always re-running `init`.

## Migration guide

1. Follow the [Terraform 0.14 upgrade guide](https://www.terraform.io/upgrade-guides/0-14.html) to update your Terraform code.
1. The only other thing you will need to do is start checking in the Terraform lock files (`.terraform.lock.hcl`), which is new functionality added in Terraform 0.14. Terragrunt will automatically copy the lock file next to your `terragrunt.hcl` file, so you should be able to commit it to version control just as with any other Terraform project. Check out the [lock file handling documentation](https://terragrunt.gruntwork.io/docs/features/lock-file-handling/) for details. 


## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1459
* https://github.com/gruntwork-io/terragrunt/pull/1452
<br>

# v0.26.7
_Released Nov 25, 2020_
## Updated CLI args, config attributes and blocks

- `dependency`

## Description

- Fix bug where `mock_outputs_allowed_terraform_commands` did not work correctly for nested dependencies.


## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1439
<br>

# v0.26.6
_Released Nov 25, 2020_
## Updated CLI args, config attributes and blocks

- `remote_state` [gcs]

## Description

- The `gcs` configuration for `remote_state` now supports `impersonate_service_account` when creating the bucket.

## Special thanks

Thank you to @ksvladimir for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1434
<br>

# v0.26.5
_Released Nov 25, 2020_
## Updated CLI args, config attributes and blocks

(none)

## Description

- The `xxx-all` commands (e.g., `apply-all`) now properly skip the `.terragrunt-cache` on Windows.

## Special thanks

Thank you to @xmclark for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1440
<br>

# v0.26.4
_Released Nov 18, 2020_
## Updated CLI args, config attributes and blocks

(none)

## Description

- The `xxx-all` commands (e.g., `apply-all`) now skip the `.terragrunt-cache` entirely when scanning for Terragrunt modules, which should make these commands run faster.

## Special thanks

Thank you to @awiddersheim for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1424
<br>

# v0.26.3
_Released Nov 17, 2020_
## Updated CLI args, config attributes and blocks

(none)

## Description

- Fix a bug where Terragrunt would pass through the `TF_CLI_ARGS` environment variable when calling `terraform --version` for an internal version check, which could lead to errors. Terragrunt will now omit this environment variable during this version check. 

## Special thanks

Thank you to @awiddersheim for the fix!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1428
<br>

# v0.26.2
_Released Nov 2, 2020_
## Updated CLI args, config attributes and blocks

(none)

## Description

- **Fixing version tags**: 
    - We previously released `v0.26.1` too quick - when the correct version number was `v0.26.0`. However, that's caused problems on other third-party tools using `terragrunt` - e.g. `brew install terragrunt` was pointing to the non-existing tag `v0.26.1`.

## Related links

* https://github.com/Homebrew/homebrew-core/pull/64096
<br>

# v0.26.0
_Released Nov 2, 2020_
## Updated CLI args, config attributes and blocks

- `remote_state_s3_test.go`
- `skip_bucket_accesslogging` config option
- `accesslogging_bucket_name` config option
- `gostructToCty` function

## Description

- Add commercial support subpages
- Rename the `gostructToCty` (in `config_as_cty.go`) -> `goTypeToCty` (fixes issue: #1384)
- Deprecate `skip_bucket_accesslogging` (fixes issue: #1333)
- Replace it with `accesslogging_bucket_name` (fixes issue: #1333)
- Updated docs for relevant configuration option updates & missing asset paths

## Migration guide 

**You only need to follow this guide if you used [Terragrunt v0.18.0](https://github.com/gruntwork-io/terragrunt/releases/tag/v0.18.0) or newer to automatically create S3 buckets as a backend _(Auto Init)_ with access logging. If you created the S3 buckets yourself using other mechanisms, you can skip this guide.**

_If you used Terragrunt v0.18.0 or newer with _Auto Init_ enabled to create S3 buckets as a backend, then you might be seeing the following bug, described in more details here: https://github.com/gruntwork-io/terragrunt/issues/1333_


### Fixing the bug on existing S3 buckets
   1. In the AWS Web Console, locate the S3 remote state buckets that were created with `terragrunt`.
   2. For all the relevant S3 buckets, find & navigate to their logging properties.
   3. Edit the properties and set the `Target Bucket` to point to a new S3 bucket dedicated for logs _(you may need to create this new bucket yourself, if it doesn't exist already)_.
   
### Skipping access logging if you had `skip_bucket_accesslogging` set
   1. Find all terragrunt configuration files from your code base that have the `skip_bucket_acesslogging` config option set.
   2. Remove the `skip_bucket_accesslogging` config option.
   3. Run `terragrunt plan` - no planned changes should be output & the access logging should still be disabled for the relevant S3 buckets.

### Enabling access logging with this release and going forward
_**Caveat:** This only applies to new S3 buckets for storing terraform state created by Terragrunt Auto Init. If the bucket already exists, Terragrunt will NOT update the existing S3 bucket._
   1. Find all `terragrunt.hcl` files that set up a remote state using S3, and you want to enable server access logging for the S3 bucket the state will be stored in.
   2. Add in the `remote_state.config` section of the configuration file the new config option and replace the value "A VALID S3 BUCKET NAME STRING" with a `string` fir the name of your logs `Target Bucket`.
```
# terragrunt.hcl
remote_state {
...
  config = {
    ...
    accesslogging_bucket_name = "A VALID S3 BUCKET NAME STRING"
    ...
  }
}
```
   3. Run one of the following commands `terragrunt plan | apply | validate` so that the S3 state bucket new configuration changes can take effect 

#### Docs
Check out the [docs](https://terragrunt.gruntwork.io/docs/reference/config-blocks-and-attributes/#remote_state) to read in more detail how to use the new attribute.

## Special thanks

- Thank you to @VishalCR7, @ArkuVonSymfon & @klijakub for the contribution!
- Thank you to @tony-harverson-moonpig for raising the issue for S3 bucket logging

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1400
* https://github.com/gruntwork-io/terragrunt/pull/1396
* https://github.com/gruntwork-io/terragrunt/pull/1405
* https://github.com/gruntwork-io/terragrunt/pull/1411
* https://github.com/gruntwork-io/terragrunt/pull/1412


<br>

# v0.25.5
_Released Oct 22, 2020_
## Updated CLI args, config attributes and blocks

- `remote_state` ; `s3` (config block)

## Description

This release fixes a bug in the auto S3 bucket creation routine, where it will ignore the `skip_` flags under certain circumstances where the backend configuration is modified (e.g., modifying the bucket name or object key location).

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1399
<br>

# v0.25.4
_Released Oct 15, 2020_
## Updated CLI args, config attributes and blocks

- `aws-provider-patch` (CLI arg)

## Description

The `aws-provider-patch` command now allows you to override attributes in nested blocks: e.g., override the `role_arn` attribute in an `assume_role { ... }` block. This is to work around [yet another Terraform bug with `import`](https://github.com/hashicorp/terraform/issues/26211).

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1391
<br>

# v0.25.3
_Released Oct 9, 2020_
## Updated CLI args, config attributes and blocks

New config attribute:`retryable_errors`

## Description

Retryable errors are now configurable. For more information see the following documentation pages:
- [Auto-retry](https://terragrunt.gruntwork.io/docs/features/auto-retry/)
- [Configuration Blocks and Attributes](https://terragrunt.gruntwork.io/docs/reference/config-blocks-and-attributes/#retryable_errors)

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1232
<br>

# v0.25.2
_Released Sep 30, 2020_
## Updated CLI args, config attributes and blocks

- `remote_state`

## Description

Previously, Terragrunt logged detailed information about changes in the config for the `remote_state` block, which could include secrets. Now these detailed change logs have been demoted to `debug` level and will only be logged if debug logging is enabled (environment variable `TG_LOG=debug`).

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1361
<br>

# v0.25.1
_Released Sep 20, 2020_
## Updated CLI args, config attributes and blocks

- `sops_decrypt_file`

## Description

- The `sops_decrypt_file` will now cache decrypted files in memory while Terragrunt runs. This can lead to a significant speed up, as it ensures that each unique file path is decrypted at most once per run.

## Special thanks

- Thank you to @abeluck for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1345
<br>

# v0.25.0
_Released Sep 18, 2020_
## Updated CLI args, config attributes and blocks

(none)

## Description

- **Terraform 0.13 upgrade**: We have verified that this repo is compatible with Terraform `0.13.x`! 
    - From this release onward, we will only be running tests with Terraform `0.13.x` against this repo, so we recommend updating to `0.13.x` soon! 

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1357
<br>

# v0.24.4
_Released Sep 16, 2020_
## Updated CLI args, config attributes and blocks

* `Auto Init`

## Description

When creating an S3 bucket for state storage, Terragrunt will now add an IAM policy that only allows the S3 bucket to be accessed via TLS. You can disable this policy (NOT recommended!) via the `skip_bucket_enforced_tls` setting.

## Special thanks

Thank you to @lazzurs for the PR!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1337
<br>

# v0.24.3
_Released Sep 15, 2020_
## Updated CLI args, config attributes and blocks

* `dependency`

## Description

Terragrunt will now use the terragrunt download directory for setting up shallow dependency fetching as described in [the dependency optimization documentation](https://terragrunt.gruntwork.io/docs/reference/config-blocks-and-attributes/#dependency).

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1346
<br>

# v0.24.2
_Released Sep 15, 2020_
## Updated CLI args, config attributes and blocks

* `Auto Retry`

## Description

Terragrunt will now automatically retry on several different flavors of the "429 Too Many Requests" error from app.terraform.io.

## Special thanks

Special thanks to @robbruce for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1342
<br>

# v0.24.1
_Released Sep 10, 2020_
## Updated CLI args, config attributes and blocks

* `Auto Retry`

## Description

Terragrunt will now automatically retry commands if it gets a "429 Too Many Requests" error from app.terraform.io.

## Special thanks

Special thanks to @robbruce for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1332
<br>

# v0.24.0
_Released Sep 8, 2020_
## Updated CLI args, config attributes and blocks

* `remote_state`

## Description

The `lock_table` attribute for the s3 remote state backend is now marked as deprecated and terragrunt will automatically convert it to the preferred `dynamodb_table` attribute. This means that existing configurations using the `lock_table` attribute will need to be reinitialized.

## Special thanks

Special thanks to @kwilczynski for their contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1327
<br>

# v0.23.9
_Released Apr 15, 2020_
## Updated CLI args, config attributes and blocks

- `terragrunt-source` [CLI Arg]
- `dependency` [block]

## Description

This fixes a bug where setting `--terragrunt-source` on a module with dependencies updated the source of both the current module and the dependency. Starting this version, when using `--terragrunt-source` on a module with a dependency the dependency's terraform source will be updated to the combined path using `//` as an anchor.

For example, if you had the following configs:

*live/vpc/terragrunt.hcl*
```hcl
terraform {
  source = "/path/to/my/modules//vpc
}
```

*live/app/terragrunt.hcl*
```hcl
terraform {
  source = "/path/to/my/modules//app
}

dependency "vpc" {
  config_path = "../vpc"
}
```

and you ran `terragrunt plan --terragrunt-source /alternative//app` in the app module folder (`live/app`), the source of the app will be updated to `/alternative//app` and the source of the vpc module when reading the dependency will be updated to `/alternative//vpc`. 

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1134
<br>

# v0.23.8
_Released Apr 8, 2020_
## Updated CLI args, config attributes and blocks

- `read_terragrunt_config` [function]
- `dependency` [block]

## Description

This fixes a bug where the relative paths used for `read_terragrunt_config` were broken when used in a module that was pulled in with `dependency` blocks.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1125
<br>

# v0.23.7
_Released Apr 8, 2020_
## Updated CLI args, config attributes and blocks

`graph-dependencies` [**NEW**]

## Description

This introduces a new command `graph-dependencies` which can be used to get a graph representation of the dependency relations between all the modules.

## Special thanks

Thanks to @mauriciopoppe for the fix!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1117
<br>

# v0.23.6
_Released Apr 6, 2020_
## Updated CLI args, config attributes and blocks

(none)

## Description

Fixed formatting in `--help` text for Terragrunt.

## Special thanks

Thanks to @KyMidd for the fix!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1115
<br>

# v0.23.5
_Released Apr 4, 2020_
## Updated CLI args, config attributes and blocks

- `dependency` [Block]

## Description

This release introduces an optimization for retrieving dependencies, where each dependency will be retrieved concurrently.

## Special thanks

Special thanks to @dmattia for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1114
<br>

# v0.23.40
_Released Sep 2, 2020_
## Updated CLI args, config attributes and blocks

* `aws-provider-patch` **[NEW CLI COMMAND]**

## Description

We've added a new `aws-provider-patch` command that can be used to override attributes in nested `provider` blocks. This is an attempt at a hacky workaround for a [Terraform bug](https://github.com/hashicorp/terraform/issues/13018) where `import` does not work if you are using any modules that have `provider` blocks with dynamic variables nested within them. With this release, you can run:

```bash
terragrunt aws-provider-patch --terragrunt-override-attr region=eu-west-1
```

And Terragrunt will:

1. Run `init` to download the code for all your modules into `.terraform/modules`.
1. Scan all the Terraform code in `.terraform/modules`, find AWS `provider` blocks, and hard-code the `region` param to `eu-west-1` for each one.

Once you do this, you'll hopefully be able to run `import` on that module. After that, you can delete the modified `.terraform/modules` and go back to normal.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1326
<br>

# v0.23.4
_Released Mar 29, 2020_
## Updated CLI args, config attributes and blocks

- `hclfmt` [command]
- Terraform functions

## Description

This release upgraded the internal terraform functions to the latest version (`v0.12.24`). Previously unavailable functions such as `trim`, `try`, and `can` are now available. `hclfmt` has also been upgraded to handle formatting files with heredoc syntax correctly.

Internally, terragrunt now uses go modules for dependency management.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1105
<br>

# v0.23.39
_Released Sep 1, 2020_
## Updated CLI args, config attributes and blocks

- `dependency`


## Description

This fixes dependency config loading for Windows.

## Special thanks

Special thanks to @bltavares for their contribution!

## Related links

https://github.com/gruntwork-io/terragrunt/pull/1320
<br>

# v0.23.38
_Released Aug 28, 2020_
## Updated CLI args, config attributes and blocks

- `dependency`


## Description

A further optimization is made to the dependency fetch optimization routine, where plugin installs are skipped.

## Related links

https://github.com/gruntwork-io/terragrunt/pull/1321
<br>

# v0.23.37
_Released Aug 27, 2020_
## Updated CLI args, config attributes and blocks

- `dependency`


## Description

- You can now benefit from dependency optimization even if you are not managing remote state in `generate` mode.
- You can now disable dependency optimization using a feature flag on the `remote_state` block: `disable_dependency_optimization = true`.

## Related links

https://github.com/gruntwork-io/terragrunt/pull/1315
<br>

# v0.23.36
_Released Aug 26, 2020_
## Updated CLI args, config attributes and blocks

- `dependency`


## Description

This fixes a bug that was introduced in the dependency retrieval optimization, where it was not accounting for IAM role assume configurations.


## Related links

https://github.com/gruntwork-io/terragrunt/pull/1315
<br>

# v0.23.35
_Released Aug 25, 2020_
## Updated CLI args, config attributes and blocks

- `dependency`


## Description

There is now an optimization on `dependency` output fetching if certain conditions are met. See [the updated docs](https://terragrunt.gruntwork.io/docs/reference/config-blocks-and-attributes/#dependency) for more information.


## Related links

https://github.com/gruntwork-io/terragrunt/pull/1311
<br>

# v0.23.34
_Released Aug 24, 2020_
## Updated CLI args, config attributes and blocks

- Makefile

## Description

Starting this release you can use the provided `Makefile` to build the terragrunt binary with `make build`.

## Special thanks

Special thanks to @salewski for their contribution!

## Related links

https://github.com/gruntwork-io/terragrunt/pull/1308
<br>

# v0.23.33
_Released Aug 3, 2020_
## Updated CLI args, config attributes and blocks

- `--terragrunt-debug` [CLI]

## Description

This release introduces a new CLI flag `--terragrunt-debug` which can be used to initiate debug mode for terragrunt. In debug mode, terragrunt will emit a `terragrunt-generated.tfvars.json` file into the terragrunt directory which you can use to inspect what TF variable inputs are being passed to terraform. See [the docs](https://terragrunt.gruntwork.io/docs/features/debugging/) to learn more.

## Special thanks

Special thanks to @eboto for the initial contribution!

## Related links

https://github.com/gruntwork-io/terragrunt/pull/1137
https://github.com/gruntwork-io/terragrunt/pull/1263
<br>

# v0.23.32
_Released Jul 31, 2020_
## Updated CLI args, config attributes and blocks

- `terraform` [block]
- `dependency` [block]

## Description

- You can now use `dependency` block references in subblocks of the `terraform` block with `xxx-all` commands.
- Fix bug where `dependency` blocks ran irrelevant hooks when retrieving outputs.

## Related links

https://github.com/gruntwork-io/terragrunt/pull/1276
https://github.com/gruntwork-io/terragrunt/pull/1275
<br>

# v0.23.31
_Released Jul 1, 2020_
## Updated CLI args, config attributes and blocks

- `prevent_destroy ` [config attribute]

## Description

Fix a bug where if you set `prevent_destroy` in a root `terragrunt.hcl`, it would override any values in child `terragrunt.hcl` files. 

## Special thanks

Thank you to @alexkayabula for the PR!

## Related links

https://github.com/gruntwork-io/terragrunt/pull/1223
<br>

# v0.23.30
_Released Jun 29, 2020_
## Updated CLI args, config attributes and blocks

- `terraform_version_constraint ` [config attribute]

## Description

Terragrunt will now correctly parse and check Terraform version numbers for full releases (e.g., `Terraform 0.12.23`), beta releases (e.g., `Terraform 0.13.0-beta2`), and dev builds (e.g., `Terraform v0.9.5-dev (cad024a5fe131a546936674ef85445215bbc4226+CHANGES)`).

## Special thanks

Thank you to @artemsablin for the PR!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1227
<br>

# v0.23.3
_Released Mar 28, 2020_
## Updated CLI args, config attributes and blocks

- `hclfmt` [command]
- `--terragrunt-hclfmt-file` [CLI Arg] [**NEW**]

## Description

`hclfmt` now supports formatting a single file as opposed to an entire directory tree. Use `--terragrunt-hclfmt-file` to specify the single file you wish to format.

## Special thanks

Thanks to @EDsCODE for the PR!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1104
<br>

# v0.23.29
_Released Jun 25, 2020_
## Updated CLI args, config attributes and blocks

- `--terragrunt-working-dir` [CLI arg]

## Description

You can now set the terragrunt working directory using the environment variable `TERRAGRUNT_WORKING_DIR`.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1235
<br>

# v0.23.28
_Released Jun 23, 2020_
## Updated CLI args, config attributes and blocks

- `plan-all`

## Description

- Terragrunt will no longer incorrectly print `Error with plan:` for `plan-all` commands when there was no error.
- Terragrunt will now log the Terraform version when debug mode is enabled.

## Special thanks

Thank you to @alexkayabula for the contributions!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1219
* https://github.com/gruntwork-io/terragrunt/pull/1229
<br>

# v0.23.27
_Released Jun 14, 2020_
## Updated CLI args, config attributes and blocks

- (none)

## Description

Terragrunt's color output should now work correctly on Windows.

## Special thanks

Thank you to @jereksel for the fix!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1184
<br>

# v0.23.26
_Released Jun 12, 2020_
## Updated CLI args, config attributes and blocks

- main CLI

## Description

Terragrunt now considers terraform json files (e.g., `.tf.json`) as valid terraform code when validating if modules contain Terraform.

## Special thanks

Thank you to @atlaskerr for the fix!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1189
<br>

# v0.23.25
_Released Jun 11, 2020_
## Updated CLI args, config attributes and blocks

- `get_terraform_cli_args` [func]

## Description

This release fixes a bug in `get_terraform_cli_args` function where it would crash when there were no CLI args passed to `terraform`.

## Special thanks

Thank you to @camlow325 for the fix!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1215
<br>

# v0.23.24
_Released Jun 8, 2020_
## Updated CLI args, config attributes and blocks

- `xxx-all` commands (e.g., `apply-all`)

## Description

The `xxx-all` commands will now ignore the Terraform data dir (default `.terraform`) when searching for Terragrunt modules. 

## Special thanks

Thank you to @zachwhaley for the fix!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1208
<br>

# v0.23.23
_Released May 29, 2020_
## Updated CLI args, config attributes and blocks

- `terragrunt_version_constraint` [config attr]

## Description

Fixes a bug where having `terragrunt_version_constraint` in a config causes `terragrunt` to crash when running `xxx-all` commands.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1205
<br>

# v0.23.22
_Released May 28, 2020_
## Updated CLI args, config attributes and blocks

- (none)

## Description

* Fixes a bug in the config comparison function gcsConfigValuesEqual that was deleting Terragrunt specific config values from the source configuration.
* Enables GCS authentication using a fixed token defined in the `GOOGLE_OAUTH_ACCESS_TOKEN` env var.

## Special thanks

Thanks to @Jawshua for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1199
<br>

# v0.23.21
_Released May 28, 2020_
## Updated CLI args, config attributes and blocks

- `get_platform ` [new function]

## Description

Added a new `get_platform()` function you can use in your Terragrunt config to get the name of the current operating system (e.g., darwin, freebsd, linux, or windows).
 
## Special thanks

Thanks to @edgarsandi for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1197
<br>

# v0.23.20
_Released May 23, 2020_
## Updated CLI args, config attributes and blocks

- `terragrunt_version_constraint` [config attr]
- `terraform_version_constraint` [config attr]
- `terraform_binary` [config attr]

## Description

The terragrunt and terraform version checks are now done without parsing the entire configuration.
 

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1196
<br>

# v0.23.2
_Released Mar 10, 2020_
## Updated CLI args, config attributes and blocks

- Config lookup behavior

## Description

Terragrunt will now always look for its config in `terragrunt.hcl` *and* `terragrunt.hcl.json` files. The latter used to be broken in `xxx-all` commands (e.g., `validate-all`).

## Special thanks

Thanks to @rimmington for the PR!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1086
<br>

# v0.23.19
_Released May 23, 2020_
## Updated CLI args, config attributes and blocks

- `terragrunt_version_constraint` [config attr] [**NEW**]

## Description

This release introduces a new config attribute `terragrunt_version_constraint`, which can be used to specify terragrunt versions that the config supports.
 
## Special thanks

Thanks to @jakauppila  for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1192
<br>

# v0.23.18
_Released May 14, 2020_
## Updated CLI args, config attributes and blocks

- `sops_decrypt_file` [func] [**NEW**]

## Description

This release introduces a new function `sops_decrypt_file`, which will decrypt an encrypted file using [SOPS](https://github.com/mozilla/sops).
 
## Special thanks

Thanks to @js-timbirkett for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1180
<br>

# v0.23.17
_Released May 12, 2020_
## Updated CLI args, config attributes and blocks

- (none)

## Description

Fix indentation in the `--help` text of Terragrunt.
 
## Special thanks

Thanks to @jereksel for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1175
<br>

# v0.23.16
_Released May 10, 2020_
## Updated CLI args, config attributes and blocks

- `--terragrunt-parallelism` [new CLI arg]

## Description

This release introduces a new `--terragrunt-parallelism` CLI argument that you can use to limit parallelism when executing `xxx-all` commands (e.g. ,`apply-all`). This is useful if, for example, you are running `apply-all` on a large number of modules, and starting to hit cloud provider rate limits as a result.
 
## Special thanks

Thanks to @mauriciopoppe for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1118
<br>

# v0.23.15
_Released May 8, 2020_
## Updated CLI args, config attributes and blocks

- `get_terraform_cli_args` [function]

## Description

This release introduces a new function (`get_terraform_cli_args`) to introspect terraform args in your terragrunt config. This can be used to conditionally adjust your hooks based on what args are passed in (e.g. `plan --destroy` vs. `plan`).
 
## Special thanks

Thanks to @camlow325 for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1170
<br>

# v0.23.14
_Released May 5, 2020_
## Updated CLI args, config attributes and blocks

- `remote_state` [block]

## Description

The automate S3 bucket created by terragrunt with the `remote_state` block will now properly handle new accounts that do not have the default s3 KMS key created.
 
## Special thanks

Thanks to @jfharden for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1163
<br>

# v0.23.13
_Released Apr 30, 2020_
## Updated CLI args, config attributes and blocks

- `get_env` (func)
- `get_terraform_command` (func) [**NEW**]

## Description

- `get_env` can now be made to force an environment variable to be defined. When used with a single argument, the function will error if the environment variable returns empty string.
- This release introduces a new helper function, `get_terraform_command` which can be used to see what the current executed command by `terragrunt` is.

## Special thanks

Thanks to @hristo-ganekov-sumup  and  @bernardoVale for the contribution!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1154
* https://github.com/gruntwork-io/terragrunt/pull/1160
<br>

# v0.23.12
_Released Apr 27, 2020_
## Updated CLI args, config attributes and blocks

N/A

## Description

Terragrunt will now auto retry on the error "TooManyUpdates error for SSM param".

## Special thanks

Thanks to @ZbigniewZabost  for the fix!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/972
<br>

# v0.23.11
_Released Apr 27, 2020_
## Updated CLI args, config attributes and blocks

- `-var` and `-var-file` [CLI args]

## Description

Terragrunt will now correctly omit `-var` and `-var-file` arguments when you run `destroy` with a plan file.

## Special thanks

Thanks to @FineWolf for the fix!

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1148
<br>

# v0.23.10
_Released Apr 15, 2020_
## Updated CLI args, config attributes and blocks

- `generate` [block]

## Description

`generate` blocks now support an optional attribute `disable_signature`, which, when `true`, will avoid emitting the signature line at the top of the generated file.

## Related links

* https://github.com/gruntwork-io/terragrunt/pull/1141
<br>

# v0.23.1
_Released Mar 9, 2020_
## Updated CLI args, config attributes and blocks

- Internal

## Description

This release fixes a bug where Terragrunt crashes with a `panic` when Terraform exits with a fatal error.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1083
<br>

# v0.23.0
_Released Mar 3, 2020_
## Updated CLI args, config attributes and blocks

- `find_in_parent_folders ` (function) **[BACKWARDS INCOMPATIBLE CHANGE]**

## Description

The `find_in_parent_folders` function will now return an absolute path rather than a relative path. This should make the function less error prone to use, especially in situations where the working directory may change (e.g., due to `.terragrunt-cache`).

## Migration guide

In many config, a common pattern was to use `get_terragrunt_dir` to construct a path with `find_in_parent_folders` to make the relative path consistent. E.g:

```hcl
locals {
  region_vars = file("${get_terragrunt_dir()}/${find_in_parent_folders("region.yaml", "${path_relative_from_include()}/empty.yaml")}")
}
```

With this change, `find_in_parent_folders` now returns the absolute path so you need to update with:

```hcl
locals {
  region_vars = file(find_in_parent_folders("region.yaml", find_in_parent_folders("empty.yaml")))
}
```

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1076
<br>

# v0.22.5
_Released Feb 26, 2020_
## Updated CLI args, config attributes and blocks

- `remote_state` (block)

## Description

Terrarunt will now respect the `external_id` and `session_name` parameters in your S3 backend config.


## Special Thanks

Thanks to @kujon for the contribution!


## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1064
<br>

# v0.22.4
_Released Feb 21, 2020_
## Updated CLI args, config attributes and blocks

- `--terragrunt-include-dir` (CLI arg)
- `--terragrunt-strict-include` (CLI arg) [**NEW**]

## Description

This fixes a bug in the behavior of `--terragrunt-include-dir` where there was an inconsistency in when dependencies that were not directly in the included dirs were included. Now all dependencies of included dirs will consistently be included as stated in the docs.

This release also introduced a new flag `--terragrunt-strict-include` which will ignore dependencies when processing included directories.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1062
<br>

# v0.22.3
_Released Feb 18, 2020_
## Updated config attributes and blocks

- `read_terragrunt_config` (func) [**NEW**]
- `generate` (block)

## Description

This release introduces the `read_terragrunt_config` function, which can be used to load and reference another terragrunt config. See [the function documentation](https://terragrunt.gruntwork.io/docs/reference/built-in-functions/#read_terragrunt_config) for more details.

This release also fixes a bug with the `generate` block where `comment_prefix` was ignored.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1051
<br>

# v0.22.2
_Released Feb 18, 2020_
## Updated config attributes and blocks

- `remote_state` (block)

## Description

This release exposes a configuration to disable checksum computations for validating requests and responses. This can be used to workaround https://github.com/gruntwork-io/terragrunt/issues/1059.

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1060
<br>

# v0.22.1
_Released Feb 18, 2020_
## Updated config attributes and blocks

- `remote_state` (block)

## Description

Allow creating GCS buckets for state storage with the Bucket Policy Only attribute, exposed as the config property `enable_bucket_policy_only`.

This release also includes a documentation fix for the `dependency` block reference.

## Special thanks

Special thanks to @chrissng for their contribution!

## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1034 
- https://github.com/gruntwork-io/terragrunt/pull/1058
<br>

# v0.22.0
_Released Feb 18, 2020_
## Updated config attributes and blocks

- `generate` (block) [**NEW**]
- `remote_state` (block)

## Description

This release introduces the `generate` block and the `generate` attribute to the `remote_state` block. These features can be used to generate code/files in the Terragrunt working directory (where Terragrunt calls out to Terraform) prior to calling `terraform`. See [the updated docs](https://terragrunt.gruntwork.io/docs/features/keep-your-remote-state-configuration-dry/#using-the-generate-property-to-generate-terraform-code-for-managing-remote-state) for more info. Also see [here](https://terragrunt.gruntwork.io/docs/reference/config-blocks-and-attributes/#generate) for documentation on `generate` blocks and [here](https://terragrunt.gruntwork.io/docs/reference/config-blocks-and-attributes/#remote_state) for documentation on the `generate` attribute of the `remote_state` block.


## Related links

- https://github.com/gruntwork-io/terragrunt/pull/1050
<br>

# v0.21.9
_Released Dec 10, 2019_
https://github.com/gruntwork-io/terragrunt/pull/975: `validate` is no longer a command that accepts `-var` and `-var-file` arguments in terraform 0.12.X. Therefore, it is no longer in the list of commands returned by `get_terraform_commands_that_need_vars`.

Special thanks to @ryno75 for their contribution!
<br>

# v0.21.8
_Released Dec 5, 2019_
https://github.com/gruntwork-io/terragrunt/pull/970: Fix bug where `terragrunt console` does not have readline properties and thus the output is garbled.

https://github.com/gruntwork-io/terragrunt/pull/971: Fix bug where `get_aws_account_id` will cause terragrunt to crash if you do not have any credentials in the chain. Instead, terragrunt now fails gracefully with a proper error message.
<br>

# v0.21.7
_Released Dec 4, 2019_
https://github.com/gruntwork-io/terragrunt/pull/969: `RunTerragrunt` is now a public function that can be called programmatically when using terragrunt as a library.

Special thanks to @rvodden for their contribution.
<br>

# v0.21.6
_Released Nov 13, 2019_
https://github.com/gruntwork-io/terragrunt/pull/863: Support JSON format for terragrunt configuration files, as described in [JSON configuration syntax](https://www.terraform.io/docs/configuration/syntax-json.html) for terraform.

Special thanks to @jakauppila for their contribution!
<br>

# v0.21.5
_Released Nov 11, 2019_
https://github.com/gruntwork-io/terragrunt/pull/948: Add additional debug logging for helping debug issues.
<br>

# v0.21.4
_Released Nov 6, 2019_
https://github.com/gruntwork-io/terragrunt/pull/901: You can now set the terragrunt download directory in the `terragrunt.hcl` config using the new `download_dir` attribute.

Special thanks to @ekini for their contribution!
<br>

# v0.21.3
_Released Nov 5, 2019_
https://github.com/gruntwork-io/terragrunt/pull/945: DynamoDB tables generated by terragrunt to initialize state locks are now generated with `PAY_PER_REQUEST` billing mode.

Special thanks to @ThatGerber for their contribution!
<br>

# v0.21.2
_Released Nov 1, 2019_
https://github.com/gruntwork-io/terragrunt/pull/937: Introduce synchronization logic to ensure that for a given dependency config, only have one `terragrunt output` call happening at once. This mitigates the issues https://github.com/gruntwork-io/terragrunt/issues/904 and https://github.com/gruntwork-io/terragrunt/issues/906.
<br>

# v0.21.13
_Released Feb 12, 2020_
https://github.com/gruntwork-io/terragrunt/pull/1042: Fixed a bug in the S3 permissions so they should now work properly for the root user.

Special thanks to @nicomfer for their contribution!
<br>

# v0.21.12
_Released Feb 10, 2020_
https://github.com/gruntwork-io/terragrunt/pull/1030: There is now a special after hook `terragrunt-read-config`, which is triggered immediately after terragrunt loads the config (`terragrunt.hcl`). You can use this to hook important processes that should be run as the first thing in your terragrunt pipeline.

Special thanks to @sleungcy for their contribution!
<br>

# v0.21.11
_Released Jan 14, 2020_
https://github.com/gruntwork-io/terragrunt/pull/1000: Terragrunt IAM role assume functionality will now also set the legacy environment variable `AWS_SECURITY_TOKEN` after assuming the role. This is useful for overriding the credentials set by tools like `aws-vault` when role chaining, as these tools also set the `AWS_SECURITY_TOKEN` which confuses the AWS SDK.

Special thanks to @alexandrst88 for their contribution!
<br>

# v0.21.10
_Released Dec 18, 2019_
#978: This is a bug fix release (fixes #770). After creating the remote state S3 bucket, we need to grant access to the AWS account root user. Without this, the root user cannot use the remote state
bucket. It also adds a `skip_bucket_root_access` option to the remote state config in the event that this behavior is undesirable for some reason.
<br>

# v0.21.1
_Released Oct 25, 2019_
https://github.com/gruntwork-io/terragrunt/pull/928: Fix missing dependencies in the lock file, which breaks certain downstream build processes.

Special thanks to @chenrui333 for their contribution!
<br>

# v0.21.0
_Released Oct 24, 2019_
#917 : Changes `--terragrunt-include-dir` and `--terragrunt-exclude-dir` glob pattern matching to exact matching. Previously comparison was done with `strings.Contains`, which meant that pattern  `**/module-gce-c` also matched `modules-gce-cd`. To exclude `modules-gce-cd`, the pattern now has to exact match, e.g. `**/module-gce-c*`.

Special thanks to @JeanFred for their contribution!
<br>

# v0.20.5
_Released Oct 18, 2019_
https://github.com/gruntwork-io/terragrunt/pull/907: This adds support for the `GOOGLE_CREDENTIALS` environment variable when setting up the gcloud connection to manage the GCS bucket for remote state. Users can either pass in the contents or path to a credentials file.
<br>

# v0.20.4
_Released Oct 14, 2019_
https://github.com/gruntwork-io/terragrunt/pull/909: Bump the version of the Terraform library that we pull in to 0.12.9. This primarily ensures that we have the latest versions of all the [Terraform functions](https://www.terraform.io/docs/configuration/functions.html) that you can use in your `terragrunt.hcl` files.
<br>

# v0.20.3
_Released Oct 11, 2019_
https://github.com/gruntwork-io/terragrunt/pull/871: This release fixes the bucket creation for `remote_state` when using GCS.

Special thanks to @brimstone for their contribution!
<br>

# v0.20.2
_Released Oct 10, 2019_
https://github.com/gruntwork-io/terragrunt/pull/891: You can now use the `--terragrunt-ignore-dependency-order` flag to tell the `xxx-all` commands to ignore dependency order, and apply all modules with as much concurrency as possible. This is mainly useful for the `plan-all` command, where instead of processing dependencies in order, all the plans can be generated in parallel.
<br>

# v0.20.1
_Released Oct 10, 2019_
https://github.com/gruntwork-io/terragrunt/pull/888: Terragrunt now properly takes into account the `TF_DATA_DIR` environment variable.
<br>

# v0.20.0
_Released Oct 8, 2019_
https://github.com/gruntwork-io/terragrunt/pull/898: `include` blocks are now parsed before `locals`. As a consequence, you will **no longer be able to use `locals` to configure `include` blocks**. However, with this change, you will now be able to use all the functions that depended on `include` blocks in your `locals`, such as `get_parent_terragrunt_dir`.
<br>

# v0.2.0
_Released Nov 22, 2016_
- ENHANCEMENT: If you configure s3 as your remote state store, Terraform will now create the S3 bucket for you (prompting you for confirmation first) if that bucket doesn't already exist. It'll also warn you if the bucket doesn't have versioning enabled or if your remote state config doesn't have encryption enabled.
- BUG FIX: Terragrunt will now properly read in `--terragrunt-XXX` flags, even if they are specified _after_ the command (e.g. `terragrunt apply --terragrunt-non-interactive`). This was broken before due to a limitation in the CLI library we were using.
- BUG FIX: The integration test for Terragrunt now creates the S3 bucket for itself rather than using a hard-coded one that only Gruntwork employees can access.

<br>

# v0.19.9
_Released Jul 18, 2019_
https://github.com/gruntwork-io/terragrunt/pull/787: You can now suppress the output of `run_cmd()` (e.g., for sensitive commands) by passing `--terragrunt-quiet` as the first argument to `run_cmd()`.
<br>

# v0.19.8
_Released Jul 2, 2019_
https://github.com/gruntwork-io/terragrunt/pull/771: Ensure `project` and `location` are optional when using the GCS backend. This release also fixes a bug that may lead to Terragrunt creating a GCS bucket in the wrong location.
<br>

# v0.19.7
_Released Jul 1, 2019_
https://github.com/gruntwork-io/terragrunt/pull/761 & https://github.com/gruntwork-io/terragrunt/pull/769: Introduce a new config option for the remote state backend in `terragrunt.hcl` (`disable_init`) which, when set to `"true"`, will skip the backend initialization step.

Special thanks to @stefansedich for their contribution!
<br>

# v0.19.6
_Released Jun 25, 2019_
https://github.com/gruntwork-io/terragrunt/pull/757: Terragrunt now supports remote state using the [gcs](https://www.terraform.io/docs/backends/types/gcs.html) backend including automatically creating a state bucket if it doesn't already exist.
<br>

# v0.19.5
_Released Jun 21, 2019_
https://github.com/gruntwork-io/terragrunt/pull/753: Fix how `include` blocks are parsed so that `path_relative_to_include`, `path_relative_from_include`, and `get_parent_terragrunt_dir` work correctly when used directly in the child `terragrunt.hcl`.
<br>

# v0.19.4
_Released Jun 16, 2019_
https://github.com/gruntwork-io/terragrunt/pull/748: Terragrunt now supports ALL [built-in Terraform functions](https://www.terraform.io/docs/configuration/functions.html)! You can use them to do math (e.g., `ceil`, `min`, `max`), manipulate strings (e.g., `format`, `split`, `substr`), interact with files (`file`, `dirname`), and much more. 
<br>

# v0.19.31
_Released Oct 8, 2019_
#900: Fixes an issue where extraneous quotes were being added to the `aws_get_*` methods
<br>

# v0.19.30
_Released Oct 7, 2019_
#897: This release adds new `get_aws_caller_identity_arn` and `get_aws_caller_identity_user_id` functions
<br>

# v0.19.3
_Released Jun 14, 2019_
https://github.com/gruntwork-io/terragrunt/pull/747: Terragrunt will now clean up `.tf.json` files in addition to `.tf` files in the download dir when fetching code from a `source` URL.
<br>

# v0.19.29
_Released Oct 5, 2019_
https://github.com/gruntwork-io/terragrunt/pull/877: This release updates the behavior of `skip_outputs` to return `mock_outputs` when it is set to `true`.

Special thanks to @barryib for their contribution!
<br>

# v0.19.28
_Released Oct 3, 2019_
#890 Update to AWS SDK 1.25.4 so terragrunt can be used in an EKS container that is assuming an IAM role
<br>

# v0.19.27
_Released Sep 25, 2019_
https://github.com/gruntwork-io/terragrunt/pull/861: You can now pass the `--terragrunt-include-external-dependencies` flag to tell Terragrunt to automatically include all external dependencies without any prompt.
<br>

# v0.19.26
_Released Sep 24, 2019_
https://github.com/gruntwork-io/terragrunt/pull/874: Skip versioning check on S3 bucket when `skip_bucket_versioning` is set to true.
<br>

# v0.19.25
_Released Sep 17, 2019_
https://github.com/gruntwork-io/terragrunt/pull/866: Fixes bug where the dependency logic did not update the terragrunt cache dir. This causes problems if you were using local state, where the state file is stored in the terragrunt cache.
<br>

# v0.19.24
_Released Sep 9, 2019_
https://github.com/gruntwork-io/terragrunt/pull/855: Fix bug where `dependency` block `config_path` that were relative were not resolved from the right directory when detecting cycles in the dependency graph, causing Terragrunt to incorrectly crash.
<br>

# v0.19.23
_Released Sep 4, 2019_
https://github.com/gruntwork-io/terragrunt/pull/850: Implements better error handling for cyclic `dependency` blocks so that you can see when you have a cycle in your `dependency` references.

https://github.com/gruntwork-io/terragrunt/pull/840: Implements better error checking of `dependency` blocks to make sure the target module has already been applied. Additionally, this implements `mock_outputs` which can be used to run `validate` on a module before the dependencies are applied. See [the updated README](https://github.com/gruntwork-io/terragrunt#unapplied-dependency-and-mock-outputs) for more information.
<br>

# v0.19.22
_Released Sep 3, 2019_
https://github.com/gruntwork-io/terragrunt/pull/826: hclfmt now runs before parsing the terragrunt config.
<br>

# v0.19.21
_Released Aug 19, 2019_
https://github.com/gruntwork-io/terragrunt/pull/841: Fix bug where `yamldecode` unnecessarily compacts number strings.
<br>

# v0.19.20
_Released Aug 15, 2019_
https://github.com/gruntwork-io/terragrunt/pull/828: We now have a new block `dependency`, which can be used to read in outputs of other terraform modules managed with terragrunt config. For example, if you wanted to read the outputs of the `vpc` module located at `../vpc` from your current terragrunt config, you can do:

```
dependency "vpc" {
  config_path = "../vpc"
}

inputs = {
  vpc_id = dependency.vpc.outputs.vpc_id
}
```

Take a look at the new section in the [README](https://github.com/gruntwork-io/terragrunt#passing-outputs-between-modules) to learn more.
<br>

# v0.19.2
_Released Jun 12, 2019_
https://github.com/gruntwork-io/terragrunt/pull/741: Fix concurrency issue when running `xxx-all` commands that sometimes caused a panic due to a `fatal error: concurrent map writes` error.
<br>

# v0.19.19
_Released Aug 8, 2019_
#818, #823: If specified use `remote_state.config.credentials` to check if the GCS bucket exists. If it isn't specified then assume the `GOOGLE_APPLICATION_CREDENTIALS` environment variable has been set. Don't throw an error if `remote_state.config.bucket` isn't specified and assume it will be passed through the `extra_arguments` -backend-config variable.

Special thanks to @marko7460 for their contribution!
<br>

# v0.19.18
_Released Aug 8, 2019_
https://github.com/gruntwork-io/terragrunt/pull/751: Added two new settings you can use in your `terragrunt.hcl` configuration: `terraform_binary` lets you specify a custom path to the Terraform binary and `terraform_version_constraint` lets you modify the version constraint Terragrunt enforces when running Terraform. This allows you to use the newest version of Terragrunt with older versions of Terraform. Note, however, that we **only** test Terragrunt with the latest version of Terraform, and provide no guarantees about compatibility with older versions.
<br>

# v0.19.17
_Released Aug 8, 2019_
https://github.com/gruntwork-io/terragrunt/pull/816: Fix a bug where `terragrunt apply` could fail as it tries to delete files that are already deleted.
<br>

# v0.19.16
_Released Aug 2, 2019_
https://github.com/gruntwork-io/terragrunt/pull/807: Introduce a flag `--terragrunt-check` which runs `hclfmt` in check mode. Check mode indicates that the files shouldn't be updated, and the command should instead error out if any files are unformatted.

Special thanks to @thoo5ieb for their contribution!
<br>

# v0.19.15
_Released Aug 2, 2019_
https://github.com/gruntwork-io/terragrunt/pull/813: Update `hclfmt` to ignore the terragrunt cache folders.
<br>

# v0.19.14
_Released Aug 1, 2019_
https://github.com/gruntwork-io/terragrunt/pull/802: Implements `locals`, a new block in the terragrunt config that can be used to bind temporary variables within the config. For example, suppose you had a magic string that you repeat multiple times:

```
inputs = {
  aws_region = "us-east-1"
  s3_endpoint = "com.amazonaws.us-east-1.s3"
}
```

Here, the string `us-east-1` is repeated. You can extract this out into a `local` to DRY:

```
locals {
  region = "us-east-1"
}

inputs = {
  aws_region = local.region
  s3_endpoint = "com.amazonaws.${local.region}.s3"
}
```

You can read more in [the new section of the README that describes locals](https://github.com/gruntwork-io/terragrunt#locals).
<br>

# v0.19.13
_Released Jul 31, 2019_
https://github.com/gruntwork-io/terragrunt/pull/774: Terragrunt now maintains a manifest of the files it copied so it can do a better job of cleaning up those files each time you re-run. 
<br>

# v0.19.12
_Released Jul 29, 2019_
https://github.com/gruntwork-io/terragrunt/pull/801: Introduces a new sub command `hclfmt` which can be used to format your `terragrunt.hcl` files.
<br>

# v0.19.11
_Released Jul 24, 2019_
https://github.com/gruntwork-io/terragrunt/pull/794: Add auto retry for "connection reset by peer" errors. 
<br>

# v0.19.10
_Released Jul 23, 2019_
https://github.com/gruntwork-io/terragrunt/pull/796: Enable public access protection on the S3 bucket used to store state files. This ensures that misconfigurations in terraform and terragrunt can not accidentally enable public access for the state files. You can read more about public access blocking in the [official AWS docs](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html).

**NOTE**: Terragrunt and Terraform do not enable public access ACLs on the bucket or the state files stored in the bucket. Your existing buckets and state files created by previous versions of Terragrunt are already private and you do not need to upgrade to make them private. This feature is an additional protection on top of the existing measures to ensure that you don't accidentally enable public access.

If you wish to enable public access blocking on existing buckets created using previous versions of Terragrunt, follow the instructions on [the official docs to enable the protection through the AWS console](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#console-block-public-access-options).
<br>

# v0.19.1
_Released Jun 11, 2019_
https://github.com/gruntwork-io/terragrunt/pull/739: When the `source` parameter is set to a local file path, copy the files instead of using symlinks and copying. This will hopefully fix issues on Windows. Also, update dependency versions and fix handling of stack traces.
<br>

# v0.19.0
_Released Jun 10, 2019_
https://github.com/gruntwork-io/terragrunt/pull/731: Terragrunt now supports Terraform 0.12!

* This is a **BACKWARDS INCOMPATIBLE** release. 
* Terragrunt 0.19.0 and above will *only* work with Terraform 0.12.0 and above.
* Please see the [migration guide](https://github.com/gruntwork-io/terragrunt/blob/master/_docs/migration_guides/upgrading_to_terragrunt_0.19.x.md) for upgrade instructions!
<br>

# v0.18.7
_Released Jun 4, 2019_
#719 Fixes an issue with symlinks when copying files
<br>

# v0.18.6
_Released May 16, 2019_
https://github.com/gruntwork-io/terragrunt/pull/705 : This implements `get_terraform_commands_that_need_parallelism`, a new interpolation function that you can use to get the list of terraform commands that support the `-parallelism` CLI arg.

Special thanks to @cschroer for their contribution!
<br>

# v0.18.5
_Released May 9, 2019_
https://github.com/gruntwork-io/terragrunt/pull/586: Added a new `terragrunt-info` command you can run to get a JSON dump of Terragrunt settings, including the config path, download dir, working dir, IAM role, etc.
<br>

# v0.18.4
_Released Apr 21, 2019_
https://github.com/gruntwork-io/terragrunt/pull/663: You can now set `skip = true` in your Terragrunt configuration to tell Terragrunt to skip processing a `terraform.tfvars` file. This can be used to temporarily protect modules from changes or to skip over `terraform.tfvars` files that don't define infrastructure by themselves.
<br>

# v0.18.3
_Released Mar 26, 2019_
https://github.com/gruntwork-io/terragrunt/pull/685: You can now use the `--terragrunt-ignore-external-dependencies` flag to tell Terragrunt to ignore any external dependencies (i.e., those outside the current folder path) when running `xxx-all` commands (e.g., `apply-all`).
<br>

# v0.18.2
_Released Mar 8, 2019_
https://github.com/gruntwork-io/terragrunt/pull/670, https://github.com/gruntwork-io/terragrunt/pull/673: Added a new `run_cmd()` helper that you can use to run an arbitrary shell command and use whatever it writes to `stdout` in your Terragrunt config. See the [run_cmd docs](https://github.com/gruntwork-io/terragrunt#run_cmd) for details.
<br>

# v0.18.1
_Released Mar 4, 2019_
https://github.com/gruntwork-io/terragrunt/pull/669: Add checksums to releases.
<br>

# v0.18.0
_Released Feb 11, 2019_
https://github.com/gruntwork-io/terragrunt/pull/645
https://github.com/gruntwork-io/terragrunt/pull/646
https://github.com/gruntwork-io/terragrunt/pull/647
https://github.com/gruntwork-io/terragrunt/pull/650

* Terragrunt now enables server-side encryption for S3 buckets by default. Terraform state may contain secrets, so this is a good security measure to have in place. If you need to disable encryption, set the `skip_bucket_ssencryption` setting to `true`.
* Terragrunt now enables access-logging for S3 buckets by default. This gives you a record of who has been accessing your Terraform state. The logs are stored in the same S3 bucket as the Terraform state. If you need to disable access logging, set the `skip_bucket_accesslogging` setting to `true`.
* Terragrunt now has the ability to enable server-side encryption for DynamoDB. This is not enabled by default since Terraform doesn't store any secrets in DynamoDB (it's only used for hashes and lock info) and enabling encryption takes 1-2 minutes. If you do wish to enable it, set the `enable_lock_table_ssencryption` setting to `true`.
<br>

# v0.17.4
_Released Dec 21, 2018_
https://github.com/gruntwork-io/terragrunt/pull/626: Added support for a `--terragrunt-include-dir` argument that you can use to specify the folders to include for `terragrunt xxx-all` commands (rather than all sub-modules). 
<br>

# v0.17.3
_Released Nov 24, 2018_
https://github.com/gruntwork-io/terragrunt/pull/613: Fix a potential concurrency bug in how Terragrunt handled stdout/stderr that could lead to transient failures with hard-to-understand error messages.
<br>

# v0.17.2
_Released Nov 9, 2018_
https://github.com/gruntwork-io/terragrunt/pull/599: You can now tell Terragrunt to assume an IAM role by setting the `iam_role` parameter in your `terragrunt = { ... }` configuration.
<br>

# v0.17.1
_Released Oct 16, 2018_
https://github.com/gruntwork-io/terragrunt/pull/589: Fix a bug where prompts on `stdin` were not showing up correctly.
<br>

# v0.17.0
_Released Oct 7, 2018_
https://github.com/gruntwork-io/terragrunt/pull/583: Terragrunt will now automatically retry commands if it hits an error known to be transient/intermittent. For example, due to bugs in Terraform, when you run `apply`, you may occasionally get a TLS handshake timeout error, which goes away if you just retry `apply`. Terragrunt will now do this retry automatically for the list of errors in [auto_retry_options.go](https://github.com/gruntwork-io/terragrunt/blob/master/options/auto_retry_options.go). In future releases, we will make the list of errors configurable. You can disable retries with either the `--terragrunt-no-auto-retry` flag or `TERRAGRUNT_AUTO_RETRY` environment variable. See the [auto retry docs](https://github.com/gruntwork-io/terragrunt#auto-retry) for more info.
<br>

# v0.16.9
_Released Sep 15, 2018_
https://github.com/gruntwork-io/terragrunt/pull/567: Add support for `force_path_style` in the S3 config. Add support for skipping S3 bucket versioning via the `skip_bucket_versioning` config.
<br>

# v0.16.8
_Released Aug 30, 2018_
https://github.com/gruntwork-io/terragrunt/pull/553: You can now set environment variables in an `extra_arguments` by specifying key value pairs in the `env_vars` parameter.
<br>

# v0.16.7
_Released Aug 15, 2018_
https://github.com/gruntwork-io/terragrunt/pull/540: There are now two "init" commands you can use in hooks: `init-from-module`, which only executes when downloading remote Terraform configurations based on the `source` parameter, and `init`, which executes for all other `init` invocations (e.g,. to configure backends, download providers, download modules).
<br>

# v0.16.6
_Released Aug 14, 2018_
https://github.com/gruntwork-io/terragrunt/pull/542: When Terragrunt calls `terraform init -from-module=xxx` to download the code for `xxx`, it now sets the `-get=false`, `-get-plugins=false`, and `-backend=false` params. These will all be handled in a later call to `init` instead. This should improve iteration speed when calling Terragrunt commands.
<br>

# v0.16.5
_Released Aug 8, 2018_
https://github.com/gruntwork-io/terragrunt/pull/535: You can now override the Terragrunt download dir using the `TERRAGRUNT_DOWNLOAD` environment variable.
<br>

# v0.16.4
_Released Jul 31, 2018_
https://github.com/gruntwork-io/terragrunt/pull/525: Add `prevent_destroy` flag, which you can use in your Terragrunt configuration to protect a module from anyone running `terragrunt destroy` or `terragrunt destroy-all`.
<br>

# v0.16.3
_Released Jul 14, 2018_
https://github.com/gruntwork-io/terragrunt/pull/523: Fix a bug where Terragrunt would hit an error trying to download Terraform configurations from `source` URLs pointing to the root of a repo.
<br>

# v0.16.2
_Released Jul 11, 2018_
https://github.com/gruntwork-io/terragrunt/pull/519: Properly exclude modules in `xxx-all` commands that show up in `.terragrunt-cache`, a custom download folder specified via `--terragrunt-download-dir`, or in nested subfolders of either of these.
<br>

# v0.16.14
_Released Oct 2, 2018_
https://github.com/gruntwork-io/terragrunt/pull/582: This is a follow-up to `v0.16.13` that fixes a bug where `-var` and `-var-file` were still passed if you called `apply` with a plan file *and* other arguments in between (e.g., `terragrunt apply <other args> <plan file>`).
<br>

# v0.16.13
_Released Oct 1, 2018_
https://github.com/gruntwork-io/terragrunt/pull/577: When you use `extra_arguments`, Terragrunt will no longer pass `-var` or `-var-file` arguments to Terraform when you call `apply` with a plan file.
<br>

# v0.16.12
_Released Sep 30, 2018_
https://github.com/gruntwork-io/terragrunt/pull/579: Fix `prevent_destroy` flag so it works even when configs are inherited from a parent `.tfvars` file.
<br>

# v0.16.11
_Released Sep 27, 2018_
https://github.com/gruntwork-io/terragrunt/pull/563: You can now tell Terragrunt to exclude specific subdirectories when running the `xxx-all` commands (e.g., `apply-all`) by using the `--terragrunt-exclude-dir` flag. This flag supports wildcard expressions and may be specified multiple times.
<br>

# v0.16.10
_Released Sep 16, 2018_
https://github.com/gruntwork-io/terragrunt/pull/569: Terragrunt will now properly respect the `shared_credentials_file` config for S3 backends, using it when creating S3 buckets and DynamoDB tables.
<br>

# v0.16.1
_Released Jul 10, 2018_
https://github.com/gruntwork-io/terragrunt/pull/518: Make sure the `xxx-all` commands (e.g., `apply-all`) don't accidentally try to run Terragrunt in a `.terragrunt-cache` directory.
<br>

# v0.16.0
_Released Jul 8, 2018_
https://github.com/gruntwork-io/terragrunt/pull/516: 

BACKWARDS INCOMPATIBLE CHANGES

1. Terragrunt now downloads remote configurations to the `.terragrunt-cache` folder in the working directory instead of into a temp folder. This makes it easier to understand what Terragrunt is doing, debug issues, and recover state files if something went wrong (e.g., lost Internet connectivity). You can use the new `--terragrunt-download-dir` option to override the download dir.

1. Terragrunt now checks if, after downloading remote configurations, it ended up with no Terraform files, and reports an error accordingly.

1. When running `terraform init` to download remote Terraform configurations, Terragrunt now passes the `-get=false`, `-backend=false`, and `-get-plugins=false` options, as all of these concerns are handled separately from the download. This should speed up local iteration. 
<br>

# v0.15.3
_Released Jul 8, 2018_
https://github.com/gruntwork-io/terragrunt/pull/515: Fix the Terragrunt cache dir on Windows to be in the `$HOME` dir rather than temp.
<br>

# v0.15.2
_Released Jul 2, 2018_
https://github.com/gruntwork-io/terragrunt/pull/506: Fix potential `nil` pointer dereference introduced in v0.15.1. 
<br>

# v0.15.1
_Released Jul 2, 2018_
https://github.com/gruntwork-io/terragrunt/pull/502: Fix the `NeedsInit` method so if you specify `s3_bucket_tags` or `dynamodb_table_tags`, Terragrunt doesn't try to re-run `init` every time.
<br>

# v0.15.0
_Released Jul 1, 2018_
https://github.com/gruntwork-io/terragrunt/pull/497: If you run one of the `xxx-all` commands (e.g, `apply-all`) in a subfolder `/foo`, and the modules in `/foo` have dependencies outside of `/foo`, Terragrunt will prompt you whether the same commands should be executed in those external dependencies too. Terragrunt used to prompt whether those dependencies should be _skipped_, so entering "yes" meant they should be skipped, which is unintuitive. Now Terragrunt prompts whether the same command should apply to those dependencies, and entering "yes" means those dependencies will be modified as well.

https://github.com/gruntwork-io/terragrunt/pull/500: Fix how Terragrunt detects and handles "S3 bucket creation already in progress" errors.
<br>

# v0.14.9
_Released May 2, 2018_
https://github.com/gruntwork-io/terragrunt/pull/468: Improve the way Terragrunt parses `source` URLs so it should be able to handle both URLs with a `//` to indicate a subfolder, and URLs without a `//`, which means to use the root.
<br>

# v0.14.8
_Released Apr 29, 2018_
https://github.com/gruntwork-io/terragrunt/pull/469: Terragrunt will now merge hooks between parent and child configs.
<br>

# v0.14.7
_Released Apr 6, 2018_
https://github.com/gruntwork-io/terragrunt/pull/460: Fix a bug where `after` hooks were not running unless `run_on_error` was set.
<br>

# v0.14.6
_Released Mar 29, 2018_
https://github.com/gruntwork-io/terragrunt/pull/455: Fix a bug where Terragrunt was not preserving Terraform's exit code.
<br>

# v0.14.5
_Released Mar 28, 2018_
https://github.com/gruntwork-io/terragrunt/pull/450: Remove a logging statement left over from development that would spit out the contents of the `TerragruntOptions` object.
<br>

# v0.14.4
_Released Mar 27, 2018_
https://github.com/gruntwork-io/terragrunt/pull/439, https://github.com/gruntwork-io/terragrunt/pull/448: Terragrunt now supports before and after hooks that allow you to execute arbitrary shell commands before and/or after running Terraform! Check out the [before and after hooks docs](https://github.com/gruntwork-io/terragrunt#before-and-after-hooks) for details.
<br>

# v0.14.3
_Released Mar 20, 2018_
https://github.com/gruntwork-io/terragrunt/pull/443: Terragrunt will now automatically run `init` if providers aren't downloaded.
<br>

# v0.14.2
_Released Feb 27, 2018_
https://github.com/gruntwork-io/terragrunt/pull/422, https://github.com/gruntwork-io/terragrunt/pull/426: If you specify a custom `endpoint` property in your S3 backend configuration, Terragrunt will now use that endpoint too. 
<br>

# v0.14.11
_Released Jun 15, 2018_
https://github.com/gruntwork-io/terragrunt/pull/488: If you are using the S3 backend with DynamoDB for locking, and the S3 bucket or DynamoDB table doesn't exist, Terragrunt will now not only automatically create them for you, but can also optionally apply tags to them via the `s3_bucket_tags` and `dynamodb_table_tags` parameters.
<br>

# v0.14.10
_Released May 6, 2018_
https://github.com/gruntwork-io/terragrunt/pull/473: Terragrunt will now properly show help text for Terraform commands (e.g., `terragrunt plan --help`).
<br>

# v0.14.1
_Released Feb 24, 2018_
https://github.com/gruntwork-io/terragrunt/pull/420: Fix out of bounds slice issue when Terragrunt is called with a command that requires multiple arguments, but some of those arguments are missing.
<br>

# v0.14.0
_Released Jan 25, 2018_
https://github.com/gruntwork-io/terragrunt/pull/407: The `apply-all` command now automatically sets the `-auto-approve` parameter so `apply` happens non-interactively with Terraform 0.11.
<br>

# v0.13.9
_Released Oct 17, 2017_
https://github.com/gruntwork-io/terragrunt/pull/321: The `find_in_parent_folders` helper now takes a second optional param that is a fallback value.
<br>

# v0.13.8
_Released Oct 16, 2017_
https://github.com/gruntwork-io/terragrunt/pull/320: The `find_in_parent_folders` helper now accepts an (optional) parameter that tells it the filename to look for.
<br>

# v0.13.7
_Released Oct 5, 2017_
https://github.com/gruntwork-io/terragrunt/pull/309: Terragrunt now properly handles the `dynamodb_table` param, auto-creating the DynamoDB table as necessary.
<br>

# v0.13.6
_Released Oct 5, 2017_
https://github.com/gruntwork-io/terragrunt/pull/308: Terragrunt should now do a better job of checking if S3 backends need initialization, including creating the DynamoDB table automatically, even if the S3 bucket already exists. 
<br>

# v0.13.5
_Released Sep 29, 2017_
https://github.com/gruntwork-io/terragrunt/pull/302: Terragrunt will now exit with an error if you define `remote_state` settings in `terraform.tfvars`, but your Terraform code does not define a corresponding `backend`. 
<br>

# v0.13.4
_Released Sep 29, 2017_
https://github.com/gruntwork-io/terragrunt/pull/301: Terragrunt should now do a better job of checking whether there is source code downloaded into your tmp folder and hopefully avoid the "No Terraform configuration files found in directory" error message.
<br>

# v0.13.3
_Released Sep 18, 2017_
https://github.com/gruntwork-io/terragrunt/pull/277: `plan-all` should now show error messages properly.
<br>

# v0.13.25
_Released Jan 19, 2018_
https://github.com/gruntwork-io/terragrunt/pull/403: Terragrunt will now properly read in state files from `local` backends.
<br>

# v0.13.24
_Released Jan 11, 2018_
https://github.com/gruntwork-io/terragrunt/pull/401: The check for a `backend { ... }` block now also checks `.tf.json` files.
<br>

# v0.13.23
_Released Dec 1, 2017_
https://github.com/gruntwork-io/terragrunt/pull/387: You can now use the environment variable `TERRAGRUNT_SOURCE_UPDATE` instead of the flag `--terragrunt-source-update`.
<br>

# v0.13.22
_Released Nov 27, 2017_
https://github.com/gruntwork-io/terragrunt/pull/383: Fix a crash that could happen with local backends.
<br>

# v0.13.21
_Released Nov 24, 2017_
https://github.com/gruntwork-io/terragrunt/pull/382: The `getAWSAccountId()` helper now assumes the IAM role specified via `--terragrunt-iam-role`, if any.
<br>

# v0.13.20
_Released Nov 13, 2017_
https://github.com/gruntwork-io/terragrunt/pull/366: You can now use the `TF_INPUT` environment variable to set Terragrunt to non-interactive mode. This is the same environment variable [used by Terraform](https://www.terraform.io/docs/configuration/environment-variables.html#tf_input).
<br>

# v0.13.2
_Released Sep 3, 2017_
https://github.com/gruntwork-io/terragrunt/pull/282: Fix a bug where the stdout from “extra” commands executed by Terragrunt (e.g., calling `init`) would pollute the stdout with extra text you wouldn’t want.
<br>

# v0.13.19
_Released Nov 12, 2017_
https://github.com/gruntwork-io/terragrunt/pull/369: Fix a concurrent write bug that would occasionally happen with the `xxx-all` commands (e.g., `apply-all`). Clean up some bugs in Terragrunt's automated tests.
<br>

# v0.13.18
_Released Nov 9, 2017_
https://github.com/gruntwork-io/terragrunt/pull/361: Fix a data race that would happen if you hit `CTRL+C`.
<br>

# v0.13.17
_Released Nov 6, 2017_
https://github.com/gruntwork-io/terragrunt/pull/347: Add missing `--terragrunt-iam-role` param to help text.
<br>

# v0.13.16
_Released Nov 3, 2017_
https://github.com/gruntwork-io/terragrunt/pull/340: Terragrunt will now cache downloaded code in your `HOME` directory instead of tmp. This should hopefully finally fix caching issues!
<br>

# v0.13.15
_Released Oct 27, 2017_
https://github.com/gruntwork-io/terragrunt/pull/331: 

1. Terragrunt was not properly loading recursive dependencies of external modules. E.g., if module A depended on module B, where B was outside of the current folder structure, and B depended on C, Terragrunt would throw an error upon finding C. I added a new failing test case to repro this issue and this PR fixes the failure.

1. While fixing the above, I triggered another issue where `xxx-all` methods would try to create the same S3 bucket in parallel, leading to “operation in progress” style errors that Terragrunt wasn’t handling correctly.
<br>

# v0.13.14
_Released Oct 22, 2017_
https://github.com/gruntwork-io/terragrunt/pull/327: Fix a bug where Terragrunt was silently swallowing parsing errors during the `xxx-all` methods (e.g., `apply-all`, `plan-all`).
<br>

# v0.13.13
_Released Oct 19, 2017_
https://github.com/gruntwork-io/terragrunt/pull/325: The `xxx-all` commands can now use the `--terragrunt-source` parameter. See [Testing multiple modules locally](https://github.com/gruntwork-io/terragrunt#testing-multiple-modules-locally) for details.
<br>

# v0.13.12
_Released Oct 19, 2017_
https://github.com/gruntwork-io/terragrunt/pull/324: You can now use the `validate-all` command to validate all modules in a folder.
<br>

# v0.13.11
_Released Oct 18, 2017_
https://github.com/gruntwork-io/terragrunt/pull/323: You can now tell Terragrunt to assume an IAM role using the `--terragrunt-iam-role` argument or `TERRAGRUNT_IAM_ROLE` environment variable. This offers a convenient way to use Terragrunt/Terraform with multiple AWS accounts. See [Work with multiple AWS accounts](https://github.com/gruntwork-io/terragrunt#work-with-multiple-aws-accounts) for more details.
<br>

# v0.13.10
_Released Oct 17, 2017_
https://github.com/gruntwork-io/terragrunt/pull/322: If you specify the `role_arn` parameter in your S3 backend config, Terragrunt will now automatically assume that IAM role before configuring remote state, creating the S3 bucket, creating the DynamoDB table, etc. This is very useful for multi-account setups where you may need to assume an IAM role in another account.
<br>

# v0.13.1
_Released Aug 31, 2017_
https://github.com/gruntwork-io/terragrunt/pull/278: After the upgrade to support terraform 0.10.x, when Terragrunt ran `init` to download remote configurations, it was running `init` with the wrong working dir.
<br>

# v0.13.0
_Released Aug 15, 2017_
https://github.com/gruntwork-io/terragrunt/pull/261: Terragrunt now supports Terraform 0.10.0. 

Note that this *should* be a largely backwards compatible release that works with both Terraform 0.9.x and 0.10.x, but Terraform 0.10.x has backwards incompatible changes (see the [changelog](https://github.com/hashicorp/terraform/blob/master/CHANGELOG.md)), so bumping the version number of Terragrunt to v0.13.0 to indicate this.

Two other changes:

1. Terragrunt now allows you to explicitly call `init` whenever you want to. Your remote state settings will be appended to the `init` command automatically.
1. Terragrunt still calls `init` for you automatically, also with remote state settings. You can now disable this _Auto Init_ functionality using the `--terragrunt-no-auto-init` flag or `TERRAGRUNT_AUTO_INIT` environment variable. 
<br>

# v0.12.9
_Released Apr 26, 2017_
https://github.com/gruntwork-io/terragrunt/pull/189: Fix a bug where Terragrunt would incorrectly detect a change in the remote state configuration every time you ran it.
<br>

# v0.12.8
_Released Apr 24, 2017_
https://github.com/gruntwork-io/terragrunt/pull/184: Fix path returned by `get_tfvars_dir()` so it works on Windows.
<br>

# v0.12.7
_Released Apr 24, 2017_
https://github.com/gruntwork-io/terragrunt/pull/175: Added three new helper functions: `path_relative_from_include()`, `get_parent_tfvars_dir()`, and `get_aws_account_id()`. See the [helper functions docs](https://github.com/gruntwork-io/terragrunt#helper-functions) for details.
<br>

# v0.12.6
_Released Apr 23, 2017_
https://github.com/gruntwork-io/terragrunt/pull/181: Terragrunt now properly supports using non-string parameters in your `remote_state` config, such as a boolean for the `encrypt` parameter.
<br>

# v0.12.5
_Released Apr 23, 2017_
https://github.com/gruntwork-io/terragrunt/pull/180: Terragrunt will now check the Terraform version you have installed to make sure it's compatible.
<br>

# v0.12.4
_Released Apr 23, 2017_
https://github.com/gruntwork-io/terragrunt/pull/178: Fix how Terragrunt checks if it should run the `init` command.
<br>

# v0.12.3
_Released Apr 23, 2017_
https://github.com/gruntwork-io/terragrunt/pull/177: Fix a bug where Terragrunt was not specifying backend configs when running `terraform init` to download remote Terraform configurations.
<br>

# v0.12.25
_Released Jul 28, 2017_
https://github.com/gruntwork-io/terragrunt/pull/258: The `xxx-all` commands (e.g. `apply-all`) will now ignore `.tfvars` files that don't specify a `source = ...` parameter. 
<br>

# v0.12.24
_Released Jun 18, 2017_
https://github.com/gruntwork-io/terragrunt/pull/238: Fix a number of file path issues on Windows.
<br>

# v0.12.23
_Released Jun 16, 2017_
https://github.com/gruntwork-io/terragrunt/pull/235: Terragrunt will now merge the `extra_arguments` blocks from child `.tfvars` files with those in parent `.tfvars` files. This allows you to define common `extra_arguments` blocks in your root Terragrunt configurations and to define environment or component-specific overrides in the child configurations.
<br>

# v0.12.22
_Released Jun 16, 2017_
https://github.com/gruntwork-io/terragrunt/pull/236: Don't copy hidden files and folders into temp folder.
<br>

# v0.12.21
_Released Jun 13, 2017_
https://github.com/gruntwork-io/terragrunt/pull/232: Make sure Terragrunt always uses the `Logger` (which goes to `stderr`) instead of writing to `stdout`.
<br>

# v0.12.20
_Released Jun 1, 2017_
https://github.com/gruntwork-io/terragrunt/pull/226: Terragrunt will now exit if you try to run the `init` command, since Terragrunt does that automatically for you.
<br>

# v0.12.2
_Released Apr 21, 2017_
https://github.com/gruntwork-io/terragrunt/pull/152: Terragrunt now does a better job of detecting if you're using modules in your code and running `terraform get` when necessary.
<br>

# v0.12.19
_Released May 31, 2017_
https://github.com/gruntwork-io/terragrunt/pull/224: Fix some corner cases with how Terragrunt process interpolations.
<br>

# v0.12.18
_Released May 24, 2017_
https://github.com/gruntwork-io/terragrunt/pull/223: Fix bug in how interpolation functions are processed.
<br>

# v0.12.17
_Released May 23, 2017_
https://github.com/gruntwork-io/terragrunt/pull/205, https://github.com/gruntwork-io/terragrunt/commit/6ab32eabb269e158d9262d412afbf8538385bedb: 

* Terragrunt now supports two new interpolation functions, `get_terraform_commands_that_need_locking` and `get_terraform_commands_that_need_vars`, that are particularly useful with the `extra_args` helper (see the [docs](https://github.com/gruntwork-io/terragrunt#interpolation-syntax) for details). 

* Fixed a bug where Terragrunt was passing `extra_args` to certain commands in the wrong order.
<br>

# v0.12.16
_Released May 4, 2017_
https://github.com/gruntwork-io/terragrunt/pull/206, https://github.com/gruntwork-io/terragrunt/pull/210: Fix a bug in how Terragrunt check's Terraform's version number.

https://github.com/gruntwork-io/terragrunt/pull/209: You can now pass a `--terragrunt-ignore-dependency-errors` flag when calling the `xxx-all` commands to tell Terragrunt to process every module, even if some of its dependencies have errors.
<br>

# v0.12.15
_Released Apr 30, 2017_
https://github.com/gruntwork-io/terragrunt/pull/201, https://github.com/gruntwork-io/terragrunt/pull/203: Fix a bug on Windows where terragrunt was using invalid path names when downloading remote Terraform configurations.
<br>

# v0.12.14
_Released Apr 27, 2017_
https://github.com/gruntwork-io/terragrunt/pull/197: Better handling of exit codes for `***-all` commands.
<br>

# v0.12.13
_Released Apr 27, 2017_
https://github.com/gruntwork-io/terragrunt/pull/195: Terragrunt now supports a `plan-all` command.
<br>

# v0.12.12
_Released Apr 26, 2017_
https://github.com/gruntwork-io/terragrunt/pull/194: The `output-all` command now writes the outputs in the proper order, rather than reverse order.
<br>

# v0.12.11
_Released Apr 26, 2017_
https://github.com/gruntwork-io/terragrunt/pull/185: Add new `required_var_files` and `optional_var_files` helpers that can be used with `extra_arguments` to make it easier to pass `-var-file` arguments to Terraform.
<br>

# v0.12.10
_Released Apr 26, 2017_
https://github.com/gruntwork-io/terragrunt/pull/190: Fix how we execute Terraform commands so stdout isn't dirtied with unrelated output.
<br>

# v0.12.1
_Released Apr 21, 2017_
https://github.com/gruntwork-io/terragrunt/pull/170: Add a `get_tfvars_dir()` helper function that allows you to use relative paths with remote Terraform configurations. Check out the [get_tfvars_dir() docs](https://github.com/gruntwork-io/terragrunt#get_tfvars_dir) for details.
<br>

# v0.12.0
_Released Apr 21, 2017_
BREAKING CHANGE

https://github.com/gruntwork-io/terragrunt/pull/167

This release updates Terragrunt so that it is compatible with Terraform 0.9. This new version of Terraform introduced had some backwards incompatible changes, so you will need to follow the [migration instructions in the README](https://github.com/gruntwork-io/terragrunt#migrating-from-terragrunt-v011x-and-terraform-08x-and-older) to upgrade.
<br>

# v0.11.1
_Released Apr 11, 2017_
- Terragrunt binaries now build with `CGO_ENABLED=0`. This removes the dependency on `libc`, which should enable Terragrunt to run in, for example, an Alpine Linux environment.
<br>

# v0.11.0
_Released Mar 2, 2017_
https://github.com/gruntwork-io/terragrunt/pull/149: 
- The `spin-up` and `tear-down` commands have been renamed to `apply-all` and `destroy-all`. The old command names will still work, but will issue deprecation warnings.
- We've added an `output-all` command that shows all outputs in all subdirectories.

<br>

# v0.10.3
_Released Feb 22, 2017_
- https://github.com/gruntwork-io/terragrunt/pull/141: Terragrunt will now properly wait for the child process (that is, Terraform) to finish before exiting, so if you hit CTRL+C, you don't end up with an extra process hanging around, streaming log output to your console at random. 

<br>

# v0.10.2
_Released Feb 17, 2017_
- https://github.com/gruntwork-io/terragrunt/pull/128: You can now specify `extra_arguments` in your Terragrunt configuration which specify custom arguments to pass to Terraform for specific commands. For example, you can specify custom `-var-file` arguments to pass to Terraform for the `apply` and `plan` commands. For more details, check out the [docs here](https://github.com/gruntwork-io/terragrunt#passing-extra-command-line-arguments-to-terraform). 

<br>

# v0.10.1
_Released Feb 14, 2017_
- BUG FIX: When downloading [remote Terraform configurations](https://github.com/gruntwork-io/terragrunt#remote-terraform-configurations) into a temp folder, Terragrunt would first delete the existing configs in that tmp folder to ensure you had the latest code. Unfortunately, it also deleted the `.tf` files in the `.terraform` folder, so that you would get a "module not found" error. https://github.com/gruntwork-io/terragrunt/pull/126

<br>

# v0.10.0
_Released Feb 9, 2017_
MAJOR CHANGE: Terragrunt now expects its configuration to be defined in a `terraform.tfvars` file instead of a `.terragrunt` file. Check out the [README](https://github.com/gruntwork-io/terragrunt#quick-start) for what the `terraform.tfvars` file should look like and the [migration instructions](https://github.com/gruntwork-io/terragrunt#migrating-from-terragrunt-to-terraformtfvars) for how to move to the new format (it's VERY easy!). For backwards compatibility, the `.terragrunt` format WILL continue to work, but it is now **deprecated**. 

The main reasons for this change are:
1. Instead of introducing a new file format into your Terraform projects, Terragrunt now piggybacks on top of the `.tfvars` format Terraform already uses for defining variables. 
2. We reduce the number of files you need. This is especially valuable with [remote Terraform configurations](https://github.com/gruntwork-io/terragrunt#remote-terraform-configurations), as it means you only need a single terraform.tfvars file to define what code to download and provide the variables for that code.

<br>

# v0.1.5
_Released Nov 22, 2016_
- Terragrunt now compares all settings of the remote configuration between what’s stored in the `.tfstate` file and in `.terragrunt`. If any of those settings don’t match up, it prompts the user whether it should re-run `terraform remote config`.
- Terragrunt has a new `--terragrunt-non-interactive` flag that can be used to disable user prompts. This is useful when running Terragrunt in an automated setting, such as a scripts or automated test.

<br>

# v0.1.4
_Released Nov 15, 2016_
- The `refresh` command now also locks, as it may change the Terraform state file.

<br>

# v0.1.3
_Released Oct 28, 2016_
- The `import` and `remote push` commands now both properly attain a lock before running.

<br>

# v0.1.2
_Released Oct 18, 2016_
- `terragrunt` now supports specifying a custom path to the `.terragrunt` config file using the `--terragrunt-config` option or the `TERRAGRUNT_CONFIG` environment variable.

<br>

# v0.1.1
_Released Oct 3, 2016_
- Update versions of all Go dependencies. We are hopeful this may fix #36. 

<br>

# v0.1.0
_Released Sep 28, 2016_
BREAKING CHANGE

The syntax for the `.terragrunt` file has been modified to use snake_case (which is idiomatic for the HCL language) and to be more consistent between locks and remote state. Please see the [root README](https://github.com/gruntwork-io/terragrunt) for the new syntax.

<br>

# v0.0.9
_Released Jun 20, 2016_
- Terragrunt now detects `CTRL+C` (SIGINT) and will shutdown gracefully and release locks

<br>

# v0.0.8
_Released Jun 13, 2016_
- Fix a panic that would occur when using Terragrunt on a totally new repo with no Terraform state files

<br>

# v0.0.7
_Released Jun 4, 2016_
- Better log output and fixed paths in build script

<br>

# v0.0.6
_Released Jun 3, 2016_
- Only configure remote state if it isn't already configured
- Run `terraform get` automatically if modules haven't been downloaded
- Use IAM username instead of OS username in lock metadata
- Show help text if you run terragrunt with no args or options

<br>

# v0.0.5
_Released Jun 1, 2016_
- Faster gox build

<br>

# v0.0.4
_Released Jun 1, 2016_
First release

<br>

# v0.0.11
_Released Sep 14, 2016_
- Terragrunt now uses `sts.GetCallerIdentity` to fetch a user id instead of `iam.GetUser`. This allows it to work with dynamic credentials from STS (from `GetSessionToken`) and the `AWS_SESSION_TOKEN` environment variable. 

<br>

# v0.0.10
_Released Jul 1, 2016_
- Build terragrunt using module-ci

<br>

