package history

import (
	"com.fha.gocan/foundation"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewImportHistoryCommand(ctx *foundation.Context) *cobra.Command {
	var sceneName string
	var path string

	cmd := cobra.Command{
		Use:  "import-history",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ui := ctx.Ui
			connection, err := ctx.GetConnection()
			if err != nil {
				return err
			}

			ui.Say("Importing history...")
			h := NewCore(connection)
			if err = h.Import(*ctx, args[0], sceneName, path); err != nil {
				return errors.Wrap(err, "History cannot be imported")
			}

			ui.Ok()
			return nil
		},
	}

	cmd.Flags().StringVarP(&sceneName, "scene", "s", "", "Scene name")
	cmd.Flags().StringVarP(&path, "path", "d", ".", "App directory")
	return &cmd
}
