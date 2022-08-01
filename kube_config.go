package tKube

import (
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/stretchr/testify/require"
	tUtils "gitlab.com/alexandre.mahdhaoui/go-lib-testing-utils"
	"testing"
)

type KubeTester interface {
	KubeOpt() *k8s.KubectlOptions
	tUtils.Identifier
	tUtils.Tester
}

//----------------------------------------------------------------------------------------------------------------------
//------------------------------------------------ Functions -----------------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------

// Kubectl runs a kubectl command specified in `args` & Returns command's output
func Kubectl(k KubeTester, args ...string) string {
	output, err := k8s.RunKubectlAndGetOutputE(k.T(), k.KubeOpt(), args...)
	require.NoError(k.T(), err)
	return output
}

//----------------------------------------------------------------------------------------------------------------------
//----------------------------------------------- KubeOptionsBuilder ---------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------

type KubeOptionsBuilder interface {
	ConfigPath() string
	ContextName() string
	Id() string
}

func KubeOptions(kb KubeOptionsBuilder) *k8s.KubectlOptions {
	return k8s.NewKubectlOptions(kb.ContextName(), kb.ConfigPath(), kb.Id())
}

//----------------------------------------------------------------------------------------------------------------------
//----------------------------------------------- KubeConfigBuilder ----------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------

type KubeConfigBuilder interface {
	Build() KubeConfig
	SetConfigPath(string) KubeConfigBuilder
	SetContextName(string) KubeConfigBuilder
	SetKubeOpt() KubeConfigBuilder
	SetId(string) KubeConfigBuilder
	SetT(t *testing.T) KubeConfigBuilder
}

type kubeConfigBuilder struct {
	KubeConfigBuilder
	kubeConfig kubeConfig
}

func NewKubeConfigBuilder() KubeConfigBuilder {
	return &kubeConfigBuilder{kubeConfig: kubeConfig{}}
}

func (b *kubeConfigBuilder) SetConfigPath(s string) KubeConfigBuilder {
	b.kubeConfig.configPath = s
	return b
}

func (b *kubeConfigBuilder) SetContextName(s string) KubeConfigBuilder {
	b.kubeConfig.contextName = s
	return b
}

func (b *kubeConfigBuilder) SetT(t *testing.T) KubeConfigBuilder {
	b.kubeConfig.t = t
	return b
}

//----------------------------------------------------------------------------------------------------------------------
//---------------------------------------------------- KubeConfig ------------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------

type KubeConfig interface {
	ConfigPath() string
	ContextName() string
	KubeOpt() *k8s.KubectlOptions
	Id() string
	T() *testing.T
}

type kubeConfig struct {
	KubeConfig
	contextName string
	configPath  string
	id          string
	t           *testing.T
}

func NewKubeConfig(t *testing.T) KubeConfig {
	return NewKubeConfigBuilder().
		SetConfigPath("").
		SetContextName("").
		SetId(tUtils.Uuid()).
		SetT(t).
		Build()
}

func (k *kubeConfig) ConfigPath() string {
	return k.configPath
}
func (k *kubeConfig) ContextName() string {
	return k.contextName
}

func (k *kubeConfig) KubeOptions() *k8s.KubectlOptions {
	return KubeOptions(k)
}

func (k *kubeConfig) Id() string {
	return k.id
}

func (k *kubeConfig) T() *testing.T {
	return k.t
}
