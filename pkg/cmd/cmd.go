package cmd

import (
	"context"
	"fmt"
	"os"

	cli "github.com/jawher/mow.cli"
	"github.com/kevinschoon/qviz/pkg/internal/runtime"
	"github.com/kevinschoon/qviz/pkg/version"
)

const longDesc = `
 ██████╗ ██╗   ██╗██╗███████╗
██╔═══██╗██║   ██║██║╚══███╔╝
██║   ██║██║   ██║██║  ███╔╝ 
██║▄▄ ██║╚██╗ ██╔╝██║ ███╔╝  
╚██████╔╝ ╚████╔╝ ██║███████╗
 ╚══▀▀═╝   ╚═══╝  ╚═╝╚══════╝

QViz evalutes Go source code to generate Gonum plots and write their output to a file.

%s

Read more about Gonum @ https://www.gonum.org/

Example script:

package main

import qviz "github.com/kevinschoon/qviz/pkg/stdlib"

func main() {
    // qviz.New is a convenience function that returns a
    // new *plot.Plot 
    plt := qviz.New()
    plt.Title.Text = "My New Chart"
    // qviz.Render must be called at the end of your script file and only once,
    // calling it before hand will halt the execution of the program.
    qviz.Render(plt)
}
`

func Maybe(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func Run(args []string) {
	app := cli.App("qviz", "Generate plots by writing Go scripts")
	app.LongDesc = fmt.Sprintf(longDesc, version.String())
	app.Spec = "[OPTIONS] SCRIPT_PATH"
	opts := runtime.DefaultOptions()
	app.StringOptPtr(&opts.FilePath, "o out", "", "file output path (implies headless)")
	app.StringOptPtr(&opts.FileType, "t type", "jpg", "file output type [eps,jpg,pdf,png,svg,tiff]")
	app.IntOptPtr(&opts.Width, "w width", 900, "output width (pixels)")
	app.IntOptPtr(&opts.Height, "h height", 800, "output height (pixels)")
	app.BoolOptPtr(&opts.Headless, "headless", false, "do not render the UI")
	app.BoolOptPtr(&opts.Watch, "watch", true, "watch for changes")
	app.StringArgPtr(&opts.ScriptPath, "SCRIPT_PATH", opts.ScriptPath, "path to a qviz script file")
	app.Action = func() {
		Maybe(runtime.New(*opts).Run(context.Background()))
	}
	Maybe(app.Run(args))
}
