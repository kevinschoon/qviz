package cmd

import (
	"fmt"
	"os"

	cli "github.com/jawher/mow.cli"
	"github.com/kevinschoon/qviz/pkg/http"
)

func Maybe(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func Run(args []string) {
	app := cli.App("qviz", "visualize data")
	app.Spec = "[OPTIONS] [PATH]"
	var (
		_ = app.StringArg("PATH", "", "path to a qviz script file")
	)
	app.Action = func() {
		Maybe(http.Serve(http.DefaultOptions()))
	}
	Maybe(app.Run(args))
}
