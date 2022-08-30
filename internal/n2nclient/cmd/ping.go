package cmd

import (
	"github.com/gensliu/n2n-go/internal/n2nclient/analysis"
	"github.com/spf13/cobra"
)

var cmdClientPing = &cobra.Command{
	Use: "ping",
	Short: "测试网络连通和延迟状态",
	// PreRun: initLogger,
	RunE: runClientPing,
}

func init()  {
	rootCmd.AddCommand(cmdClientPing)
}

func runClientPing(cmd *cobra.Command, args []string) error {
	// TODO ctx, ticker, signal
	pinger := analysis.NewPinger()

	pinger.OnSend = func() {
		// TODO OnSend
	}

	pinger.OnRecv = func() {
		// TODO OnRecv
	}

	pinger.OnIdle = func() {
		// TODO OnIdle
	}

	return pinger.Ping()
}