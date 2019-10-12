.PHONY: all kl docker esc generate generate test

all: kl shen

generate: esc
	${GOPATH}/bin/esc -o runtime/asset.go -pkg=runtime bytecode

esc: ${GOPATH}/bin/esc
	go get -u -v github.com/mjibson/esc

kl:
	go install github.com/tiancaiamao/shen-go/cmd/kl

shen:
	go build -o shen cmd/shen/main.go

docker:
	docker build -t shen-go .
	docker run -i -t --rm -v /tmp:/tmp shen-go \
		/bin/sh -c 'cp -a /usr/local/bin/shen /tmp/'
	cp -a /tmp/shen ./shen

test:
	go test github.com/tiancaiamao/shen-go/runtime
	go test github.com/tiancaiamao/shen-go/kl
