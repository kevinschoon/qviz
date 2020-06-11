QViz is a utility to quickly and interactively visualize data with by evaluating Go source code at runtime. It works by exposing the [gonum/plot](https://github.com/gonum/plot/) library to a [yaegi](https://github.com/containous/yaegi/) interpreter.

![ss](static/ss.png)

## Usage

```text
Usage: qviz [OPTIONS] SCRIPT_PATH

Generate plots by writing Go scripts

Arguments:
  SCRIPT_PATH     path to a qviz script file

Options:
  -m, --monitor   monitor the script for changes running on each modification
      --width     image width (inches) (default 9)
      --height    image height (inches) (default 9)
  -o, --out       write the plot to this path (default "/dev/stdout")
  -t, --type      type of file to output [eps,jpg,pdf,png,svg,tiff]
```

### Interactive Usage
There are several approaches that can be used to display a plot interactively while editing script files. On Linux
the [feh](https://feh.finalrewind.org/) X11 image viewer works nicely monitoring the output image for changes.
```bash
# Monitor the simple.go script file for changes in real time
qviz --monitor -out /tmp/simple.png examples/simple/simple.go
# Next in a separate pane you can use an image viewer like feh
feh /tmp/simple.png
# now finally open up the script file in your favorite editor
vim examples/simple/simple.go
```

## Writing Scripts

All QViz scripts need to expose a function with the following signature `func QViz(*plot.Plot) error`. Example scripts can be found [here](/examples).

```go
package main

import "gonum.org/v1/plot"

func QViz(plt *plot.Plot) error {
    // your code goes here
    return nil
}
```

## Roadmap

[_] CLI based plotter
[_] allow external user defined packages
[_] first class integration with [qframe](https://github.com/tobgu/qframe)
[_] finish importing the remaining Gonum packages
[_] native UI / code editor...?
