package qviz

import "gonum.org/v1/plot"

// Plot returns a new Gonum plot
func New() *plot.Plot

// Render displays the plot output, this function
// may only be called once and must be invoked as the
// last action in a QPlot script
func Render(plt *plot.Plot) error

// Maybe is a convience function that checks the content
// of an error message. If err != nil the execution of the
// program is halted
func Maybe(err error)
