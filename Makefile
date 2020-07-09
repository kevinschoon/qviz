SYMBOLS_PATH=pkg/internal/runtime/symbols
QVIZ_VERSION=$(shell git describe --tags 2>/dev/null || git rev-parse HEAD)
GONUM_VERSION=$(shell grep 'gonum.org/v1/gonum' go.mod | cut -d ' ' -f2)
QFRAME_VERSION=$(shell grep 'github.com/tobgu/qframe' go.mod | cut -d ' ' -f2)
PLOT_VERSION=$(shell grep 'gonum.org/v1/plot' go.mod | cut -d ' ' -f2)

LINUX_RELEASE=qviz-${QVIZ_VERSION}-linux-amd64
DARWIN_RELEASE=qviz-${QVIZ_VERSION}-darwin-amd64

LDFLAGS=\
	-X github.com/kevinschoon/qviz/pkg/version.Gonum=${GONUM_VERSION} \
	-X github.com/kevinschoon/qviz/pkg/version.Plot=${PLOT_VERSION} \
	-X github.com/kevinschoon/qviz/pkg/version.QFrame=${QFRAME_VERSION} \
	-X github.com/kevinschoon/qviz/pkg/version.QViz=${QVIZ_VERSION}

default: bin/qviz

install:
	cd cmd/qviz \
	&& go install -ldflags '${LDFLAGS}'

release: \
	bin/${LINUX_RELEASE} \
	bin/${LINUX_RELEASE}.sha256 \
	bin/${DARWIN_RELEASE} \
	bin/${DARWIN_RELEASE}.sha256

test: bin/qviz
	bin/qviz --help

bin:
	mkdir -p $@

bin/qviz: bin
	cd cmd/qviz && \
	go build -ldflags '${LDFLAGS}' \
	-o ../../$@

bin/${LINUX_RELEASE}: bin
	cd cmd/qviz \
	&& \
	GOOS=linux \
	GOARCH=amd64 \
	go build -ldflags '${LDFLAGS}' -o ../../$@

bin/${LINUX_RELEASE}.sha256: bin/${LINUX_RELEASE}
	sha256sum $< | sed -e 's/bin\///' > $@

bin/${DARWIN_RELEASE}: bin
	cd cmd/qviz \
	&& \
	GOOS=darwin \
	GOARCH=amd64 \
	go build -ldflags '${LDFLAGS}' -o ../../$@

bin/${DARWIN_RELEASE}.sha256: bin/${DARWIN_RELEASE}
	sha256sum $< | sed -e 's/bin\///' > $@

${SYMBOLS_PATH}:
	mkdir $@

generate: ${SYMBOLS_PATH}
	find ${SYMBOLS_PATH} -type f -name 'go1_*' -exec rm {} \;
	scripts/gen.sh \
		'gonum.org/v1/plot/...' \
		${SYMBOLS_PATH} \
		'.*\/internal\/.*' \
		'gonum.org/v1/plot/gob'
	scripts/gen.sh \
		'gonum.org/v1/gonum/...' \
		${SYMBOLS_PATH} \
		'.*\/internal\/.*' \
		'gonum.org/v1/gonum/blas/testblas/benchautogen'
	scripts/gen.sh \
		'github.com/tobgu/qframe/...' \
		${SYMBOLS_PATH} \
		'.*\/internal\/.*' \
		'github.com/tobgu/qframe/cmd/qfgenerate' \
		'github.com/tobgu/qframe/config'
