package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/pflag"

	"github.com/joseph-cavarretta/kube-plugins/internal/kube"
	"github.com/joseph-cavarretta/kube-plugins/internal/routes"
	"github.com/joseph-cavarretta/kube-plugins/internal/table"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run() error {
	flags := pflag.NewFlagSet("kubectl-routes", pflag.ExitOnError)
	configFlags := kube.AddFlags(flags)

	if err := flags.Parse(os.Args[1:]); err != nil {
		return err
	}

	client, err := kube.NewClient(configFlags)
	if err != nil {
		return err
	}

	ingresses, err := routes.Fetch(context.Background(), client)
	if err != nil {
		return err
	}

	rows := routes.Build(ingresses)
	return table.PrintTable(
		os.Stdout, routes.Headers(), rows,
	)
}
