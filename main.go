package main

import (
	"github.com/spf13/cobra"

	"github.com/candlestick-games/snake/pkg/game"
	"github.com/candlestick-games/snake/pkg/std/runner"
)

func main() {
	cmd := &cobra.Command{
		Use:   "snake",
		Short: "Game",
		Long:  "Snake: Might and Magic",
		Run:   func(_ *cobra.Command, _ []string) { runner.RunGame(&game.Game{}) },
	}

	cmd.Flags().BoolP("window", "w", false, "window mode")
	runner.BindFlag(cmd, "window")

	runner.RunCmd(cmd)
}
