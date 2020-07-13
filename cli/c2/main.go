// Copyright © 2020 Hedzr Yeh.

package main

import (
	"github.com/hedzr/cmdr"
	"github.com/hedzr/go-socketlib/tcp/cert"
	"github.com/hedzr/go-socketlib/tcp/client"
	"github.com/hedzr/go-socketlib/tcp/server"
	"log"
)

func main() {
	if err := cmdr.Exec(buildRootCmd(),
		cmdr.WithLogex(cmdr.WarnLevel),
		//cmdr.WithUnknownOptionHandler(onUnknownOptionHandler),
		//cmdr.WithUnhandledErrorHandler(onUnhandledErrorHandler),
	); err != nil {
		log.Fatalf("error: %+v", err)
	}
}

func buildRootCmd() (rootCmd *cmdr.RootCommand) {
	root := cmdr.Root(appName, "1.0.1").
		Header("fluent - test for cmdr - no version - hedzr").
		Description(desc, longDesc).
		Examples(examples)
	rootCmd = root.RootCommand()

	socketlib(root)
	return
}

func socketlib(root cmdr.OptCmd) {
	aCmd := root.NewSubCommand("tcp", "tcp", "socket", "socketlib").
		Description("go-socketlib operations...", "").
		Group("TCP")
	server.AttachToCmdr(aCmd, server.WithPort(1983))
	client.AttachToCmdr(aCmd, client.WithPort(1983))

	cert.AttachToCmdr(root)
}

const (
	appName   = "tcp-tool"
	copyright = "tcp-tool is an effective devops tool"
	desc      = "tcp-tool is an effective devops tool. It make an demo application for `cmdr`."
	longDesc  = "tcp-tool is an effective devops tool. It make an demo application for `cmdr`."
	examples  = `
$ {{.AppName}} --help
  show help screen.
`
)
