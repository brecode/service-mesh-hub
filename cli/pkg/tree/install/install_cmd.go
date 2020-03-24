package install

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/google/wire"
	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/installutils/helminstall"
	"github.com/solo-io/go-utils/installutils/helminstall/types"
	"github.com/solo-io/mesh-projects/cli/pkg/cliconstants"
	"github.com/solo-io/mesh-projects/cli/pkg/common"
	common_config "github.com/solo-io/mesh-projects/cli/pkg/common/config"
	"github.com/solo-io/mesh-projects/cli/pkg/common/helmutil"
	"github.com/solo-io/mesh-projects/cli/pkg/common/semver"
	"github.com/solo-io/mesh-projects/cli/pkg/options"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
)

type InstallCommand *cobra.Command

var (
	InstallErr = func(err error) error {
		return eris.Wrap(err, "Error installing Service Mesh Hub")
	}
	InvalidVersionErr = func(version string) error {
		return eris.Errorf(
			"Invalid version: %s. For a list of supported versions, "+
				"see the releases page: https://github.com/solo-io/mesh-projects/releases", version)
	}
	InstallSet = wire.NewSet(
		InstallCmd,
	)
	PreInstallMessage  = "Starting Service Mesh Hub installation...\n"
	PostInstallMessage = "Service Mesh Hub successfully installed!\n"
)

func HelmInstallerProvider(helmClient types.HelmClient, kubeClient kubernetes.Interface) types.Installer {
	return helminstall.NewInstaller(helmClient, kubeClient.CoreV1().Namespaces(), ioutil.Discard)
}

func InstallCmd(opts *options.Options, kubeClientsFactory common.KubeClientsFactory, kubeLoader common_config.KubeLoader, out io.Writer) InstallCommand {
	cmd := &cobra.Command{
		Use:     cliconstants.InstallCommand.Use,
		Short:   cliconstants.InstallCommand.Short,
		PreRunE: validateArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := kubeLoader.GetRestConfigForContext(opts.Root.KubeConfig, opts.Root.KubeContext)
			if err != nil {
				return err
			}
			kubeClients, err := kubeClientsFactory(cfg, opts.Root.WriteNamespace)
			if err != nil {
				return err
			}
			chartUri, err := helmutil.GetChartUri(opts.SmhInstall.HelmChartOverride, opts.SmhInstall.Version)
			if err != nil {
				return InstallErr(err)
			}
			if err := kubeClients.HelmInstaller.Install(&types.InstallerConfig{
				KubeConfig:         opts.Root.KubeConfig,
				KubeContext:        opts.Root.KubeContext,
				DryRun:             opts.SmhInstall.DryRun,
				Verbose:            opts.Root.Verbose,
				InstallNamespace:   opts.Root.WriteNamespace,
				CreateNamespace:    opts.SmhInstall.CreateNamespace,
				ReleaseName:        opts.SmhInstall.HelmReleaseName,
				ReleaseUri:         chartUri,
				ValuesFiles:        opts.SmhInstall.HelmChartValueFileNames,
				PreInstallMessage:  PreInstallMessage,
				PostInstallMessage: PostInstallMessage,
			}); err != nil {
				return InstallErr(err)
			}

			fmt.Fprintf(out, "Service Mesh Hub has been installed to namespace %s\n", opts.Root.WriteNamespace)
			return nil
		},
	}
	options.AddInstallFlags(cmd, opts)
	return cmd
}

func validateArgs(cmd *cobra.Command, _ []string) error {
	// validate version, prefix with 'v' if not already
	version, _ := cmd.Flags().GetString("version")
	if version != "" {
		if !semver.ValidReleaseSemver(version) {
			return InvalidVersionErr(version)
		}
	}
	return nil
}
