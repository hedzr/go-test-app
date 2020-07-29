// Copyright © 2020 Hedzr Yeh.

package main

import (
	"github.com/hedzr/cmdr"
	"github.com/hedzr/cmdr-addons/pkg/plugins/trace"
	"github.com/hedzr/go-socketlib/coaplib"
	"github.com/hedzr/go-socketlib/tcp/cert"
	"github.com/hedzr/go-socketlib/tcp/client"
	"github.com/hedzr/go-socketlib/tcp/server"
	"github.com/hedzr/log"
	log2 "log"
)

func main() {
	if err := cmdr.Exec(buildRootCmd(),
		cmdr.WithLogex(cmdr.Level(log.WarnLevel)),
		trace.WithTraceEnable(true),

		//cmdr.WithUnknownOptionHandler(onUnknownOptionHandler),
		//cmdr.WithUnhandledErrorHandler(onUnhandledErrorHandler),
	); err != nil {
		log2.Fatalf("error: %+v", err)
	}
}

func buildRootCmd() (rootCmd *cmdr.RootCommand) {
	root := cmdr.Root(appName, "1.0.1").
		// Header("fluent - test for cmdr - no version - hedzr").
		Copyright(copyright, "hedzr").
		Description(desc, longDesc).
		Examples(examples)
	rootCmd = root.RootCommand()

	socketLib(root)
	return
}

func socketLib(root cmdr.OptCmd) {

	// TCP/UDP

	tcpCmd := root.NewSubCommand("tcp", "tcp", "socket", "socketlib").
		Description("go-socketlib TCO operations...", "").
		Group("TCP")

	server.AttachToCmdr(tcpCmd, server.WithCmdrPort(1983))
	client.AttachToCmdr(tcpCmd, client.WithCmdrPort(1983), client.WithCmdrInteractiveCommand(true))

	udpCmd := root.NewSubCommand("udp", "udp").
		Description("go-socketlib UDP operations...", "").
		Group("UDP")

	server.AttachToCmdr(udpCmd, server.WithCmdrUDPMode(true), server.WithCmdrPort(1984))
	client.AttachToCmdr(udpCmd, client.WithCmdrUDPMode(true), client.WithCmdrPort(1984))

	// Cert

	cert.AttachToCmdr(root)

	// CoAP

	coapCmd := root.NewSubCommand("coap", "co").
		Description("CoAP server/client operations...", "").
		Group("IoT")

	coaplib.AttachToCmdr(coapCmd)
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
