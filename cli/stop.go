package cli

import (
	"github.com/dominikbraun/timetrace/core"
	"github.com/dominikbraun/timetrace/out"

	"github.com/spf13/cobra"
)

func stopCommand(t *core.Timetrace) *cobra.Command {
	stop := &cobra.Command{
		Use:   "stop [<MESSAGE>]",
		Short: "Stop tracking your time",
		Run: func(cmd *cobra.Command, args []string) {
			var message string
			if len(args) > 0 {
				message = args[0]
			}

			if err := t.Stop(message); err != nil {
				out.Err("failed to stop tracking: %s", err.Error())
				return
			}

			out.Success("Stopped tracking time")
		},
	}

	return stop
}
