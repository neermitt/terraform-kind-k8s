package test

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

func cleanup(t *testing.T, terraformOptions *terraform.Options, tempTestFolder string) {
	terraform.Destroy(t, terraformOptions)
	os.RemoveAll(tempTestFolder)
}

// Test the Terraform module in examples/complete using Terratest.
func TestExamplesComplete(t *testing.T) {
	t.Parallel()

	rootFolder := "../../"
	terraformFolderRelativeToRoot := "examples/complete"

	tempTestFolder := test_structure.CopyTerraformFolderToTemp(t, rootFolder, terraformFolderRelativeToRoot)

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: tempTestFolder,
		Upgrade:      true,
		Vars: map[string]interface{}{
			"enabled":     true,
			"namespace":   "eg",
			"environment": "dev",
			"stage":       "test",
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer cleanup(t, terraformOptions, tempTestFolder)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	name := terraform.Output(t, terraformOptions, "name")
	assert.Equal(t, "eg-dev-test", name)

	ipAddress := terraform.Output(t, terraformOptions, "ipAddress")
	assert.NotEmpty(t, ipAddress)

	kubeconfig := terraform.Output(t, terraformOptions, "kubeconfig")
	assert.NotEmpty(t, kubeconfig)

	kubeconfig_with_ip := terraform.Output(t, terraformOptions, "kubeconfig_with_ip")
	assert.NotEmpty(t, kubeconfig_with_ip)

	context_name := terraform.Output(t, terraformOptions, "context_name")
	assert.Equal(t, "kind-eg-dev-test", context_name)

	controller_container_name := terraform.Output(t, terraformOptions, "controller_container_name")
	assert.Equal(t, "eg-dev-test-control-plane", controller_container_name)

}
