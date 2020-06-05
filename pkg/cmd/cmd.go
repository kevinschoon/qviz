package cmd

import (
	"fmt"
	"os"

	cli "github.com/jawher/mow.cli"
	qviz "github.com/kevinschoon/qviz/pkg"
)

func Maybe(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
}

func Run(args []string) {
	app := cli.App("qviz", "visualize data")
	var (
		path = app.StringArg("PATH", "", "path to a qviz script file")
	)
	app.Action = func() {
		Maybe(qviz.Read(*path))
	}
	Maybe(app.Run(args))
}
