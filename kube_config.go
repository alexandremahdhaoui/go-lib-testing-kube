package tKube

import (
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/stretchr/testify/require"
	tUtils "gitlab.com/alexandre.mahdhaoui/go-lib-testing-utils"
	"testing"
)

func NewKubeConfig(t *testing.T) KubeConfig {
	k := KubeConfig{}

	k.SetConfigPath("")
	k.SetContextName("")
	k.SetId(tUtils.Uuid())
	k.SetT(t)

	return k
}

//----------------------------------------------------------------------------------------------------------------------
//--------------------------------------------------- Functions --------------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------

// Kubectl runs a kubectl command specified in `args` & Returns command's output
func Kubectl(k KubeTester, args ...string) string {
	output, err := k8s.RunKubectlAndGetOutputE(k.T(), k.KubeOpt(), args...)
	require.NoError(k.T(), err)
	return output
}

func NewKubeOpt(kb KubeOptBuilder) *k8s.KubectlOptions {
	return k8s.NewKubectlOptions(kb.ContextName(), kb.ConfigPath(), kb.Id())
}

//----------------------------------------------------------------------------------------------------------------------
//------------------------------------------------------ Struct --------------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------

type KubeConfig struct {

	// Getters
	ConfigPathGetter
	ContextNameGetter
	KubeOptGetter
	tUtils.Identifier
	tUtils.Tester

	// Setters
	ConfigPathSetter
	ContextNameSetter
	IdSetter
	tUtils.TestSetter

	// Fields
	contextName string
	configPath  string
	id          string
	t           *testing.T
}

//----------------------------------------------------------------------------------------------------------------------
//---------------------------------------------------- Interfaces ------------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------

type KubeConfigBuilder interface {
	KubeOptGetter
	tUtils.Identifier
}

type KubeOptBuilder interface {
	ConfigPathGetter
	ContextNameGetter
	tUtils.Identifier
}

type KubeTester interface {
	KubeOptGetter
	tUtils.Identifier
	tUtils.Tester
}

type ConfigPathGetter interface{ ConfigPath() string }
type ContextNameGetter interface{ ContextName() string }
type KubeOptGetter interface{ KubeOpt() *k8s.KubectlOptions }

type ConfigPathSetter interface{ SetConfigPath(string) }
type ContextNameSetter interface{ SetContextName(string) }
type IdSetter interface{ SetId(string) }

//----------------------------------------------------------------------------------------------------------------------
//------------------------------------------------------ Getters -------------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------

// func (k *KubeConfig) Get() {}

func (k *KubeConfig) ConfigPath() string {
	return k.configPath
}
func (k *KubeConfig) ContextName() string {
	return k.contextName
}

func (k *KubeConfig) KubeOpt() *k8s.KubectlOptions {
	return NewKubeOpt(k)
}

func (k *KubeConfig) Id() string {
	return k.id
}

func (k *KubeConfig) T() *testing.T {
	return k.t
}

//----------------------------------------------------------------------------------------------------------------------
//------------------------------------------------------ Setters -------------------------------------------------------
//----------------------------------------------------------------------------------------------------------------------

// func (k *KubeConfig) Set() {}

func (k *KubeConfig) SetConfigPath(s string) {
	k.configPath = s
}

func (k *KubeConfig) SetContextName(s string) {
	k.contextName = s
}

func (k *KubeConfig) SetT(t *testing.T) {
	k.t = t
}
