package kube

import (
	"fmt"

	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

func AddFlags(
	flags *pflag.FlagSet,
) *genericclioptions.ConfigFlags {
	configFlags := genericclioptions.NewConfigFlags(true)
	configFlags.AddFlags(flags)
	return configFlags
}

func NewClient(
	configFlags *genericclioptions.ConfigFlags,
) (kubernetes.Interface, error) {
	restConfig, err := configFlags.ToRESTConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get REST config: %w", err)
	}

	client, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	return client, nil
}
