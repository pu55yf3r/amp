package cluster

import (
	"github.com/appcelerator/amp/cli"
	"github.com/spf13/cobra"
)

// NewCreateCommand returns a new instance of the create command for bootstrapping an cluster.
func NewCreateCommand(c cli.Interface) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create [OPTIONS]",
		Short:   "Create an amp cluster",
		PreRunE: cli.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return create(c, cmd)
		},
	}

	flags := cmd.Flags()
	flags.IntVarP(&opts.workers, "workers", "w", 2, "Initial number of worker nodes")
	flags.IntVarP(&opts.managers, "managers", "m", 3, "Intial number of manager nodes")
	flags.StringVar(&opts.provider, "provider", "local", "Cluster provider")
	flags.StringVar(&opts.name, "name", "", "Cluster Label")
	flags.StringVar(&opts.tag, "tag", "latest", "Specify tag for cluster images (default is 'latest', use 'local' for development)")
	return cmd
}

// Map cli cluster flags to target bootstrap cluster command flags and update the cluster
func create(c cli.Interface, cmd *cobra.Command) error {
	// This is a map from cli cluster flag name to bootstrap script flag name
	m := map[string]string{
		"workers":  "-w",
		"managers": "-m",
		"provider": "-t",
		"name":     "-l",
		"tag":      "-T",
	}

	// TODO: only supporting local cluster management for this release
	// the following ensures that flags are added before the final command arg
	// TODO: refactor reflag to handle this
	args := []string{"hack/deploy"}
	args = reflag(cmd, m, args)
	args = append(args, DefaultLocalClusterID)
	return queryCluster(c, args)
}
