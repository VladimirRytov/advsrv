package main

import (
	"github.com/VladimirRytov/advsrv/internal/front/cli"
)

func main() {
	c := cli.NewCommandLine()
	c.Init()
}
