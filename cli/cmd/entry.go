// Copyright © 2020 Hedzr Yeh.

package cmd

import (
	"fmt"
	"github.com/hedzr/cmdr"
	"github.com/hedzr/cmdr/examples/demo/svr"
	"github.com/hedzr/cmdr/plugin/daemon"
	"github.com/hedzr/log"
	"github.com/hedzr/logex/build"
	"gopkg.in/hedzr/errors.v2"
	"strings"
)

func Entry() {
	// logrus.SetLevel(logrus.DebugLevel)
	// logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})

	// defer func() {
	// 	fmt.Println("defer caller")
	// 	if err := recover(); err != nil {
	// 		fmt.Printf("recover success. error: %v", err)
	// 	}
	// }()

	if err := cmdr.Exec(buildRootCmd(),

		// To disable internal commands and flags, uncomment the following codes
		// cmdr.WithBuiltinCommands(false, false, false, false, false),
		daemon.WithDaemon(svr.NewDaemon(), nil, nil, nil),

		// cmdr.WithLogx(logrus.New("debug", false, true)),
		cmdr.WithLogx(build.New(log.NewLoggerConfigWith(true, "logrus", "debug"))),

		//// integrate with logex library
		//cmdr.WithLogex(cmdr.DebugLevel),
		//cmdr.WithLogexPrefix("logger"),

		cmdr.WithHelpTabStop(41),

		cmdr.WithWatchMainConfigFileToo(true),
		cmdr.WithNoWatchConfigFiles(false),
		cmdr.WithOptionMergeModifying(func(keyPath string, value, oldVal interface{}) {
			cmdr.Logger.Debugf("%%-> -> %q: %v -> %v", keyPath, oldVal, value)
			if strings.HasSuffix(keyPath, ".mqtt.server.stats.enabled") {
				// mqttlib.FindServer().EnableSysStats(!vxconf.ToBool(value))
			}
			if strings.HasSuffix(keyPath, ".mqtt.server.stats.log.enabled") {
				// mqttlib.FindServer().EnableSysStatsLog(!vxconf.ToBool(value))
			}
		}),

		cmdr.WithUnknownOptionHandler(onUnknownOptionHandler),
		cmdr.WithUnhandledErrorHandler(onUnhandledErrorHandler),

		optAddTraceOption,
		optAddServerExtOption,

		cmdr.WithOnSwitchCharHit(onSwitchCharHit),
		cmdr.WithOnPassThruCharHit(onPassThruCharHit),
	); err != nil {
		cmdr.Logger.Fatalf("error: %+v", err)
	}
}

func onSwitchCharHit(parsed *cmdr.Command, switchChar string, args []string) (err error) {
	if parsed != nil {
		fmt.Printf("the last parsed command is %q - %q\n", parsed.GetTitleNames(), parsed.Description)
	}
	fmt.Printf("SwitchChar FOUND: %v\nremains: %v\n\n", switchChar, args)
	return // cmdr.ErrShouldBeStopException
}

func onPassThruCharHit(parsed *cmdr.Command, switchChar string, args []string) (err error) {
	if parsed != nil {
		fmt.Printf("the last parsed command is %q - %q\n", parsed.GetTitleNames(), parsed.Description)
	}
	fmt.Printf("PassThrough flag FOUND: %v\nremains: %v\n\n", switchChar, args)
	return // ErrShouldBeStopException
}

func onUnhandledErrorHandler(err interface{}) {
	// debug.PrintStack()
	// pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	dumpStacks()
}

func dumpStacks() {
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===\n", errors.DumpStacksAsString(true))
}

func onUnknownOptionHandler(isFlag bool, title string, cmd *cmdr.Command, args []string) (fallbackToDefaultDetector bool) {
	return true
}

var optAddTraceOption, optAddServerExtOption cmdr.ExecOption

func init() {
	// attaches `--trace` to root command
	optAddTraceOption = cmdr.WithXrefBuildingHooks(func(root *cmdr.RootCommand, args []string) {
		cmdr.NewBool(false).
			Titles("trace", "tr").
			Description("enable trace mode for tcp/mqtt send/recv data dump", "").
			AttachToRoot(root)
	}, nil)

	// the following statements show you how to attach an option to a sub-command
	optAddServerExtOption = cmdr.WithXrefBuildingHooks(func(root *cmdr.RootCommand, args []string) {
		serverCmd := cmdr.FindSubCommandRecursive("server", nil)
		serverStartCmd := cmdr.FindSubCommand("start", serverCmd)
		cmdr.NewInt(5100).
			Titles("vnc-server", "vnc").
			Description("start as a vnc server (just a demo)", "").
			Placeholder("PORT").
			AttachTo(cmdr.NewCmdFrom(serverStartCmd))
	}, nil)
}
