.PHONY: all kl shen esc generate

all: shen kl

generate: esc
	${GOPATH}/bin/esc -o vm/asset.go -pkg=vm bytecode

esc: ${GOPATH}/bin/esc
	go get -u -v github.com/mjibson/esc

kl:
	go install github.com/tiancaiamao/shen-go/cmd/kl

shen:
	go build -o shen cmd/shen/main.go

test:
	go test github.com/tiancaiamao/shen-go/vm
	go test github.com/tiancaiamao/shen-go/kl
