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
	tests := []struct {
		name                     string
		vars                     map[string]interface{}
		expectedName             string
		expectedContextName      string
		expectedLoadbalancerName string
	}{
		{
			name: "simple_one_controller_zero_worker",
			vars: map[string]interface{}{
				"enabled":     true,
				"namespace":   "cl10",
				"environment": "dev",
				"stage":       "test",
			},
			expectedName:             "cl10-dev-test",
			expectedContextName:      "kind-cl10-dev-test",
			expectedLoadbalancerName: "cl10-dev-test-control-plane",
		},
		{
			name: "simple_one_controller_one_worker",
			vars: map[string]interface{}{
				"enabled":     true,
				"namespace":   "cl11",
				"environment": "dev",
				"stage":       "test",
				"nodes": []map[string]interface{}{
					{
						"role": "control-plane",
					},
					{
						"role": "worker",
					},
				},
			},
			expectedName:             "cl11-dev-test",
			expectedContextName:      "kind-cl11-dev-test",
			expectedLoadbalancerName: "cl11-dev-test-control-plane",
		},
		{
			name: "simple_two_controller_two_worker",
			vars: map[string]interface{}{
				"enabled":     true,
				"namespace":   "cl22",
				"environment": "dev",
				"stage":       "test",
				"nodes": []map[string]interface{}{
					{
						"role": "control-plane",
					},
					{
						"role": "control-plane",
					},
					{
						"role": "worker",
					},
					{
						"role": "worker",
					},
				},
			},
			expectedName:             "cl22-dev-test",
			expectedContextName:      "kind-cl22-dev-test",
			expectedLoadbalancerName: "cl22-dev-test-external-load-balancer",
		},
	}

	rootFolder := "../../"
	terraformFolderRelativeToRoot := "examples/complete"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := tt
			t.Parallel()
			tempTestFolder := test_structure.CopyTerraformFolderToTemp(t, rootFolder, terraformFolderRelativeToRoot)

			terraformOptions := &terraform.Options{
				// The path to where our Terraform code is located
				TerraformDir: tempTestFolder,
				Upgrade:      true,
				Vars:         test.vars,
			}

			// At the end of the test, run `terraform destroy` to clean up any resources that were created
			defer cleanup(t, terraformOptions, tempTestFolder)

			// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
			terraform.InitAndApply(t, terraformOptions)

			name := terraform.Output(t, terraformOptions, "name")
			assert.Equal(t, test.expectedName, name)

			ipAddress := terraform.Output(t, terraformOptions, "ipAddress")
			assert.NotEmpty(t, ipAddress)

			kubeconfig := terraform.Output(t, terraformOptions, "kubeconfig")
			assert.NotEmpty(t, kubeconfig)

			kubeconfig_with_ip := terraform.Output(t, terraformOptions, "kubeconfig_with_ip")
			assert.NotEmpty(t, kubeconfig_with_ip)

			context_name := terraform.Output(t, terraformOptions, "context_name")
			assert.Equal(t, test.expectedContextName, context_name)

			loadbalancer_container_name := terraform.Output(t, terraformOptions, "loadbalancer_container_name")
			assert.Equal(t, test.expectedLoadbalancerName, loadbalancer_container_name)

		})
	}
}
