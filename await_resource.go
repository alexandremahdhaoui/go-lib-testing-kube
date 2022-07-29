package tKube

import (
	"github.com/gruntwork-io/terratest/modules/k8s"
	"time"
)

// AwaitResource
//  - Waits until specified resource is available
func AwaitResource(k KubeTester, resource string, retries int, sleep time.Duration) {
	switch resource {
	case "ingress":
		k8s.WaitUntilIngressAvailable(k.T(), k.KubeOpt(), k.Id(), retries, sleep)
	case "service":
		k8s.WaitUntilServiceAvailable(k.T(), k.KubeOpt(), k.Id(), retries, sleep)
	}
}
