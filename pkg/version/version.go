package version

import "fmt"

var (
	// QViz is the version of this package
	QViz = "UNKNOWN"
	// Plot is the compiled version of gonum.org/v1/plot
	Plot = "UNKNOWN"
	// Gonum is the compiled version of gonum.org/v1/gonum
	Gonum = "UNKNOWN"
	// QFrame is the compiled version of github.com/tobgu/qframe
	QFrame = "UNKNOWN"
)

const versions = `QViz Version: %s
gonum.org/v1/plot: %s
gonum.org/v1/gonum: %s
github.com/tobgu/qframe: %s
`

// String returns a pretty formatted version list
func String() string {
	return fmt.Sprintf(versions, QViz, Plot, Gonum, QFrame)
}
