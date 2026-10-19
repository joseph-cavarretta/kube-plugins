package main

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run() error {
	configFlags := genericclioptions.NewConfigFlags(true)
	flags := pflag.NewFlagSet("kubectl-routes", pflag.ExitOnError)
	configFlags.AddFlags(flags)

	if err := flags.Parse(os.Args[1:]); err != nil {
		return err
	}

	restConfig, err := configFlags.ToRESTConfig()
	if err != nil {
		return fmt.Errorf("failed to get REST config: %w", err)
	}

	client, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}

	ingresses, err := client.NetworkingV1().
		Ingresses("").
		List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("failed to list ingresses: %w", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "NAMESPACE\tINGRESS\tHOST")
	for _, ing := range ingresses.Items {
		for _, rule := range ing.Spec.Rules {
			fmt.Fprintf(w, "%s\t%s\t%s\n", ing.Namespace, ing.Name, rule.Host)
		}
	}
	w.Flush()

	return nil
}
