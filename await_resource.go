package tKube

import (
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/testing"
	"time"
)

type resources map[string]func(testing.TestingT, *k8s.KubectlOptions, string, int, time.Duration)

// AwaitResource sleeps until specified resource is available
func AwaitResource(k KubeTester, resource string, retries int, sleep time.Duration) {
	m := resources{
		"ingress": k8s.WaitUntilIngressAvailable,
		"service": k8s.WaitUntilServiceAvailable,
	}
	for K, v := range m {
		if K == resource {
			v(k.T(), k.KubeOpt(), k.Id(), retries, sleep)
			return
		}
	}
	k.T().Fatalf("specified resource is not implemented; got: %s", resource)
}
