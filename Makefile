SYMBOLS_PATH=pkg/internal/loader/symbols

.PHONY: bin/qviz bin/generated

default: bin/qviz

install:
	cd cmd/qviz && go install

test:
	go test ./...

bin:
	mkdir -p $@

${SYMBOLS_PATH}:
	mkdir $@

# TODO: simplify the including of external packages
generate: ${SYMBOLS_PATH}
	cd ${SYMBOLS_PATH} \
	&& goexports gonum.org/v1/plot \
	&& goexports gonum.org/v1/plot/cmpimg \
	&& goexports gonum.org/v1/plot/palette \
	&& goexports gonum.org/v1/plot/palette/brewer \
	&& goexports gonum.org/v1/plot/palette/moreland \
	&& goexports gonum.org/v1/plot/plotter \
	&& goexports gonum.org/v1/plot/plotutil \
	&& goexports gonum.org/v1/plot/tools/bezier \
	&& goexports gonum.org/v1/plot/vg \
	&& goexports gonum.org/v1/plot/vg/draw \
	&& goexports gonum.org/v1/plot/vg/fonts \
	&& goexports gonum.org/v1/plot/vg/recorder \
	&& goexports gonum.org/v1/plot/vg/vgeps \
	&& goexports gonum.org/v1/plot/vg/vgimg \
	&& goexports gonum.org/v1/plot/vg/vgpdf \
	&& goexports gonum.org/v1/plot/vg/vgsvg \
	&& goexports gonum.org/v1/plot/vg/vgtex

bin/qviz: bin
	cd cmd/qviz && go build -o ../../$@

