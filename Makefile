.PHONY: all kl docker esc generate generate test

all: kl shen

generate: esc
	${GOPATH}/bin/esc -o runtime/asset.go -pkg=runtime bytecode

esc: ${GOPATH}/bin/esc
	go get -u -v github.com/mjibson/esc

kl:
	go install github.com/tiancaiamao/shen-go/cmd/kl

codegen/codegen.so:
	cd codegen;
	go build -buildmode=plugin

shen:
	go build -o shen github.com/tiancaiamao/shen-go/cmd/shen

docker:
	docker build -t shen-go .
	docker run -i -t --rm -v /tmp:/tmp shen-go \
		/bin/sh -c 'cp -a /usr/local/bin/shen /tmp/'
	cp -a /tmp/shen ./shen

test:
	go test github.com/tiancaiamao/shen-go/runtime
	go test github.com/tiancaiamao/shen-go/kl
