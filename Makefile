.PHONY: bin/qviz bin/generated

default: bin/qviz

install:
	cd cmd/qviz && go install

test:
	go test ./...

bin:
	mkdir -p $@

bin/qviz: bin
	cd cmd/qviz && go build -o ../../$@

bin/generated:
	mkdir -p $@/src/github.com/kevinschoon/qviz/pkg
	cp -v pkg/qviz.go $@/src/github.com/kevinschoon/qviz/pkg/
