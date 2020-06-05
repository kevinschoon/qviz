# qzvi

`qviz` is a simple utility for generating charts by writing Go scripts. It is sort of a very stripped down
and fast version of Juptyer.

## Usage

`qviz` can be invoked from the commandline by pointing it at a regular Go file.

```go
import (
	"fmt"

	qviz "github.com/kevinschoon/qviz/pkg"
)

func New(*qviz.Options) (*qviz.Viz, error) {
	return &qviz.Viz{
        XYs: [][]float64{
            []float64{1.0, 2.0},
            []float64{2.0, 2.0},
        },
        Labels: []string{"fuu", "bar"},
    }, nil
}
```
