package cmd

import (
	"github.com/hedzr/cmdr"
	"github.com/hedzr/go-socketlib/tcp/client"
	"github.com/hedzr/go-socketlib/tcp/server"
)

func socketlib(root cmdr.OptCmd) {
	aCmd := root.NewSubCommand("socketlib", "sl", "socket").
		Description("go-socketlib operations...", "").
		Group("TCP")

	//aCmd.NewSubCommand("server", "svr", "s").
	//	Description("tcp server", "").
	//	Group("2333.TCP").
	//	Action(func(cmd *cmdr.Command, args []string) (err error) {
	//		return
	//	})
	//
	//aCmd.NewSubCommand("client", "c").
	//	Description("tcp client", "").
	//	Group("2333.TCP").
	//	Action(func(cmd *cmdr.Command, args []string) (err error) {
	//		return
	//	})

	server.AttachToCmdr(aCmd)
	client.AttachToCmdr(aCmd)
}
