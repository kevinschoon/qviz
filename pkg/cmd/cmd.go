package cmd

import (
	"fmt"
	"os"

	cli "github.com/jawher/mow.cli"
	"github.com/kevinschoon/qviz/pkg/internal/loader"
)

func Maybe(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func Run(args []string) {
	app := cli.App("qviz", "Generate Gonum plots by writing Go scripts")
	app.LongDesc = `
 ██████╗ ██╗   ██╗██╗███████╗
██╔═══██╗██║   ██║██║╚══███╔╝
██║   ██║██║   ██║██║  ███╔╝ 
██║▄▄ ██║╚██╗ ██╔╝██║ ███╔╝  
╚██████╔╝ ╚████╔╝ ██║███████╗
 ╚══▀▀═╝   ╚═══╝  ╚═╝╚══════╝

QViz evalutes Go scripts that generate a Gonum Plot and write the output to a file.

Read more about Gonum @ https://www.gonum.org/

Example script:

package main

import (
	"gonum.org/v1/plot"
)

func QViz(plot *plot.Plot) error {
	// write your plot code here
	return nil
}

`
	app.Spec = "[OPTIONS] SCRIPT_PATH"
	opts := loader.DefaultOptions()
	app.StringArgPtr(&opts.ScriptPath, "SCRIPT_PATH", opts.ScriptPath, "path to a qviz script file")
	app.BoolOptPtr(&opts.Watch, "m monitor", opts.Watch, "monitor the script for changes running on each modification")
	app.IntOptPtr(&opts.Width, "width", opts.Width, "image width (inches)")
	app.IntOptPtr(&opts.Height, "height", opts.Height, "image height (inches)")
	app.StringOptPtr(&opts.FilePath, "o out", opts.FilePath, "write the plot to this path")
	app.StringOptPtr(&opts.FileType, "t type", opts.FileType, "type of file to output [eps,jpg,pdf,png,svg,tiff]")
	app.Action = func() {
		Maybe(loader.Load(*opts))
	}
	Maybe(app.Run(args))
}
