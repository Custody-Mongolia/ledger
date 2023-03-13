package cmd

import (
	"github.com/formancehq/ledger/pkg/api"
	"github.com/formancehq/stack/libs/go-libs/httpserver"
	app "github.com/formancehq/stack/libs/go-libs/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func NewServe() *cobra.Command {
	return &cobra.Command{
		Use: "serve",
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.New(cmd.OutOrStdout(), resolveOptions(
				viper.GetViper(),
				fx.Invoke(func(lc fx.Lifecycle, h *api.API) {
					lc.Append(httpserver.NewHook(viper.GetString(bindFlag), h))
				}),
			)...).Run(cmd.Context())
		},
	}
}