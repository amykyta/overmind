package main

import (
	"fmt"
	"net"

	"github.com/DarthSim/overmind/utils"

	"gopkg.in/urfave/cli.v1"
)

type cmdQuitHandler struct {
	SocketPath string
}

func (c *cmdQuitHandler) Run(_ *cli.Context) error {
	conn, err := net.Dial("unix", c.SocketPath)
	utils.FatalOnErr(err)

	fmt.Fprintf(conn, "quit")

	return nil
}
