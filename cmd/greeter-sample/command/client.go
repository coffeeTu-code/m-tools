package command

import (
	"fmt"
	"github.com/oklog/oklog/pkg/group"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func NewClientCmd() *cobra.Command {
	var clientCmd = &cobra.Command{
		Use:   "client",
		Short: "client start",
		Run: func(cmd *cobra.Command, args []string) {
			var g group.Group
			{
				// This function just sits and waits for ctrl-C.
				cancelInterrupt := make(chan struct{})
				g.Add(func() error {
					c := make(chan os.Signal, 1)
					signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
					select {
					case sig := <-c:
						return fmt.Errorf("received signal %s", sig)
					case <-cancelInterrupt:
						return nil
					}
				}, func(err error) {
					close(cancelInterrupt)
				})
			}
		},
	}
	return clientCmd
}
