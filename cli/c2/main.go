// Copyright © 2020 Hedzr Yeh.

package main

import (
	"github.com/hedzr/cmdr"
	"github.com/hedzr/cmdr-addons/pkg/plugins/trace"
	"github.com/hedzr/go-coaplib/cmd"
	"github.com/hedzr/go-socketlib/tcp/cert"
	"github.com/hedzr/go-socketlib/tcp/client"
	"github.com/hedzr/go-socketlib/tcp/server"
	"github.com/hedzr/log"
	"github.com/hedzr/logex/build"
)

func main() {
	if err := cmdr.Exec(buildRootCmd(),
		// cmdr.WithLogx(logrus.New("debug", false, true)),
		cmdr.WithLogx(build.New(log.NewLoggerConfigWith(true, "sugar", "debug"))),
		trace.WithTraceEnable(true),

		//cmdr.WithUnknownOptionHandler(onUnknownOptionHandler),
		//cmdr.WithUnhandledErrorHandler(onUnhandledErrorHandler),
	); err != nil {
		cmdr.Logger.Fatalf("error: %+v", err)
	}
}

func buildRootCmd() (rootCmd *cmdr.RootCommand) {
	root := cmdr.Root(appName, "1.0.1").
		Copyright(copyright, "hedzr").
		Description(desc, longDesc).
		Examples(examples)
	rootCmd = root.RootCommand()

	socketLib(root)
	return
}

func socketLib(root cmdr.OptCmd) {

	// CoAP

	coapCmd := root.NewSubCommand("coap", "co").
		Description("CoAP server/client operations...", "").
		Group("IoT")

	cmd.AttachToCmdr(coapCmd)

	// TCP/UDP

	tcpCmd := root.NewSubCommand("tcp", "tcp", "socket", "socketlib").
		Description("go-socketlib TCO operations...", "").
		Group("socket-lib")

	server.AttachToCmdr(tcpCmd, server.WithCmdrPort(1983))
	client.AttachToCmdr(tcpCmd, client.WithCmdrPort(1983), client.WithCmdrInteractiveCommand(true))

	udpCmd := root.NewSubCommand("udp", "udp", "UDP", "udplib").
		Description("go-socketlib UDP operations...", "").
		Group("socket-lib")

	server.AttachToCmdr(udpCmd, server.WithCmdrUDPMode(true), server.WithCmdrPort(1984))
	client.AttachToCmdr(udpCmd, client.WithCmdrUDPMode(true), client.WithCmdrPort(1984))

	// Cert

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
