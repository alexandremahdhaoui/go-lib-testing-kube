package tKube

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/k8s"
	tUtils "gitlab.com/alexandre.mahdhaoui/go-lib-testing-utils"
	"strings"
	"time"
)

// TestPodAntiAffinity checks the anti affinity label and tests if it is expectedly working on nodes
func TestPodAntiAffinity(k KubeTester) {
	label := fmt.Sprintf("app.kubernetes.io/instance=%s", k.Id())
	nodeNames := Kubectl(k, "get", "pods", "-l", label,
		"-o=jsonpath='{.items[?(@.status.phase==\"Running\")].spec.nodeName}'")

	if len(strings.Split(nodeNames, " ")) != 5 {
		k.T().Fatal("podAntiAffinity not applied!")
	}
}

// TestServiceEndpoint awaits service's availability, gets the Service Resource & query endpoint with
// tUtils.HttpGetWithRetry() expecting 200
func TestServiceEndpoint(k TlsKubeTester) {
	AwaitResource(k, "service", 5, 5*time.Second)

	service := k8s.GetService(k.T(), k.KubeOpt(), k.Id())
	endpoint := "http://" + k8s.GetServiceEndpoint(k.T(), k.KubeOpt(), service, 80)

	tUtils.HttpGetWithRetry(k, endpoint, 200, 5, 5*time.Second)
}
