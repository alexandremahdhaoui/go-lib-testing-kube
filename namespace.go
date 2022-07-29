package tKube

import (
	"github.com/gruntwork-io/terratest/modules/k8s"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

// CreateNs
//  - Creates a namespace from KubeOpt and Id
//  - Returns a teardown function.
func CreateNs(k KubeTester) func() {
	k8s.CreateNamespace(k.T(), k.KubeOpt(), k.Id())
	// Prepare teardown
	return func() {
		test_structure.RunTestStage(k.T(), "teardown", func() {
			k8s.DeleteNamespace(k.T(), k.KubeOpt(), k.Id())
		})
	}
}
