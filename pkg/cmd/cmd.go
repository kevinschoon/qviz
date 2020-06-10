package cmd

import (
	"fmt"
	"os"

	cli "github.com/jawher/mow.cli"
	"github.com/kevinschoon/qviz/pkg/http"
	"github.com/kevinschoon/qviz/pkg/internal/loader"
)

func Maybe(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func Run(args []string) {
	app := cli.App("qviz", "visualize data")
	app.Spec = "[OPTIONS]"
	app.Command("eval", "evaluate code and generate a plot image", func(cmd *cli.Cmd) {
		opts := loader.DefaultRenderOptions()
		path := cmd.StringArg("PATH", "", "path to a qviz script file")
		cmd.IntOptPtr(&opts.Width, "w width", opts.Width, "image width (inches)")
		cmd.IntOptPtr(&opts.Height, "h height", opts.Height, "image height (inches)")
		cmd.StringOptPtr(&opts.FilePath, "o out", "", "write the plot to this path (defaults to stdout)")
		cmd.StringOptPtr(&opts.FileType, "t type", opts.FileType, "type of file to output [eps,jpg,pdf,png,svg,tiff]")
		cmd.Action = func() {
			ctx, err := loader.LoadPath(*path)
			Maybe(err)
			Maybe(loader.Render(ctx, *opts))
		}
	})
	app.Command("serve", "serve a viz", func(cmd *cli.Cmd) {
		opts := http.DefaultOptions()
		cmd.Action = func() {
			Maybe(http.Serve(opts))
		}
	})
	Maybe(app.Run(args))
}
