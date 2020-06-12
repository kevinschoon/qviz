SYMBOLS_PATH=pkg/internal/loader/symbols
QVIZ_VERSION=$(shell git rev-parse HEAD)
GONUM_VERSION=$(shell grep 'gonum.org/v1/gonum' go.mod | cut -d ' ' -f2)
QFRAME_VERSION=$(shell grep 'github.com/tobgu/qframe' go.mod | cut -d ' ' -f2)
PLOT_VERSION=$(shell grep 'gonum.org/v1/plot' go.mod | cut -d ' ' -f2)

LDFLAGS=\
	-X github.com/kevinschoon/qviz/pkg/version.Gonum=${GONUM_VERSION} \
	-X github.com/kevinschoon/qviz/pkg/version.Plot=${PLOT_VERSION} \
	-X github.com/kevinschoon/qviz/pkg/version.QFrame=${QFRAME_VERSION} \
	-X github.com/kevinschoon/qviz/pkg/version.QViz=${QVIZ_VERSION}

.PHONY: bin/qviz

default: bin/qviz

install:
	cd cmd/qviz \
	&& go install -ldflags '${LDFLAGS}'

test: bin/qviz
	bin/qviz --help

bin:
	mkdir -p $@

bin/qviz: bin
	cd cmd/qviz && \
	go build -ldflags '${LDFLAGS}' \
	-o ../../$@

${SYMBOLS_PATH}:
	mkdir $@

generate: ${SYMBOLS_PATH}
	find ${SYMBOLS_PATH} -type f -name 'go1_*' -exec rm {} \;
	scripts/gen.sh \
		'gonum.org/v1/plot/...' \
		'.*\/internal\/.*' \
		'gonum.org/v1/plot/gob'
	scripts/gen.sh \
		'gonum.org/v1/gonum/...' \
		'.*\/internal\/.*' \
		'gonum.org/v1/gonum/blas/testblas/benchautogen'
	scripts/gen.sh \
		'github.com/tobgu/qframe/...' \
		'.*\/internal\/.*' \
		'github.com/tobgu/qframe/cmd/qfgenerate' \
		'github.com/tobgu/qframe/config'
