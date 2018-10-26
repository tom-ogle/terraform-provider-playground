# terraform-provider-playground

This is a Terraform provider plugin for exploring workarounds for the limitations in Terraform's HCL (Hashicorp Configuration Language).

For example, I used it to test conditionally setting the access_config inline block in the Google Cloud provider resourceComputeInstanceTemplate
to work around the 'should be a list' issue where Terraform can't parse dynamically generated lists and the issue where lists cannot be returned from conditionals.

* Modify the ResourceHCLPlayground function in `playground/resource_hcl_playground.go` to return a &schema.Resource with the appropriate Schema map and Create, Read, Update, Delete functions to test the features of
HCL you are interested in. E.g. You can copy the Schema from a resource you are using.
* Modify the hclUnderTest function in 'playground/resource_hcl_playground_test.go' to test your HCL.
* I recommend using the Golang IDE from Jetbrains to set breakpoints in the Terraform code that parses HCl, as required, so that you can work out
how to make your HCL work.