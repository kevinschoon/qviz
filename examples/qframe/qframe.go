package main

import (
	"fmt"
	"strings"

	"github.com/tobgu/qframe"
	"gonum.org/v1/plot"
)

func QViz(plt *plot.Plot) error {
	input := `COL1,COL2
a,1.5
b,2.25
c,3.0`
	f := qframe.ReadCSV(strings.NewReader(input))
	fmt.Println(f)
	return nil
}
