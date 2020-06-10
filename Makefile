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

generate: ${SYMBOLS_PATH}
	cd ${SYMBOLS_PATH} \
	&& goexports gonum.org/v1/plot \
	&& goexports gonum.org/v1/plot/plotter

bin/qviz: bin
	cd cmd/qviz && go build -o ../../$@

