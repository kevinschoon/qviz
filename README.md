# QViz ![QViz](https://github.com/kevinschoon/qviz/workflows/QViz/badge.svg) [![API reference](https://img.shields.io/badge/godoc-reference-5272B4)](https://pkg.go.dev/github.com/kevinschoon/qviz?tab=doc)

QViz is a utility to quickly and interactively visualize data with by evaluating Go source code at runtime. It works by exposing the [gonum/plot](https://github.com/gonum/plot/) library to a [yaegi](https://github.com/containous/yaegi/) interpreter.

![ss](static/ss.png)

## Usage

```text
Usage: qviz [OPTIONS] SCRIPT_PATH

Generate plots by writing Go scripts

Arguments:
  SCRIPT_PATH      path to a qviz script file

Options:
  -o, --out        file output path (implies headless)
  -t, --type       file output type [eps,jpg,pdf,png,svg,tiff] (default "jpg")
  -w, --width      output width (pixels) (default 900)
  -h, --height     output height (pixels) (default 800)
      --headless   do not render the UI
```

### Interactive Usage
QViz has a built in native UI for viewing plots that will start automatically when you run the command.
```bash
# Monitor the simple.go script file for changes in real time
qviz examples/simple/simple.go
# Now open up the script file in your favorite editor
vim examples/simple/simple.go
```

## Writing Scripts

Valid QViz scripts are normal Go files that import the `pkg/stdlib` overlay package. 
Example scripts can be found [here](/examples).

```go
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
```

## Roadmap

- [ ] CLI based plotter
- [ ] allow external user defined packages
- [x] first class integration with [qframe](https://github.com/tobgu/qframe)
- [x] finish importing the remaining Gonum packages
- [x] native UI
