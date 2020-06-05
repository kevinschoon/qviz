.PHONY: bin/qviz bin/generated

default: bin/qviz

install:
	cd cmd/qviz && go install

bin:
	mkdir -p $@

bin/qviz: bin
	cd cmd/qviz && go build -o ../../$@

bin/generated:
	mkdir -p $@
	scripts/generate.sh $@
