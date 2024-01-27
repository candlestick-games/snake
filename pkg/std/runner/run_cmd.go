package runner

import (
	"fmt"
	"os"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/candlestick-games/snake/pkg/std/debugger"
)

func RunCmd(cmd *cobra.Command) {
	cobra.OnInitialize(func() {
		viper.AutomaticEnv()
	})

	if cmd.Version == "" {
		cmd.Version = versionInfo()
	}

	cmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if viper.GetBool("debug") {
			debugger.Enable()
		}
	}

	cmd.Flags().Bool("debug", false, "debug mode")
	MarkHidden(cmd, "debug")
	BindFlag(cmd, "debug")

	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Run cmd: %s", err)
		os.Exit(1)
	}
}

func MarkHidden(cmd *cobra.Command, name string) {
	if err := cmd.Flags().MarkHidden(name); err != nil {
		log.Fatal("Mark hidden flag", "name", name, "error", err)
	}
}

func BindFlag(cmd *cobra.Command, name string) {
	if err := viper.BindPFlag(name, cmd.Flags().Lookup(name)); err != nil {
		log.Fatal("Bind flag", "name", name, "error", err)
	}
}

func versionInfo() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}

	var (
		vcsRevision string
		vcsTime     time.Time
		vcsModified bool
	)

	for _, setting := range info.Settings {
		switch setting.Key {
		case "vcs.revision":
			vcsRevision = setting.Value
		case "vcs.time":
			vcsTime, _ = time.Parse(time.RFC3339, setting.Value)
		case "vcs.modified":
			vcsModified, _ = strconv.ParseBool(setting.Value)
		}
	}

	version := fmt.Sprintf("%s, build with %s", info.Main.Version, info.GoVersion)

	if vcsRevision != "" {
		version += ", revision " + vcsRevision
	}
	if !vcsTime.IsZero() {
		version += ", at " + vcsTime.Local().String()
	}
	if vcsModified {
		version += ", modified"
	}

	return version
}
