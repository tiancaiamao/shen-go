.PHONY: all kl shen esc

all: shen kl

generate:
	${GOPATH}/bin/esc -o vm/asset.go -pkg=vm bytecode

esc: ${GOPATH}/bin/esc
	go get github.com/mjibson/esc

kl:
	go install github.com/tiancaiamao/shen-go/cmd/kl

shen:
	go install github.com/tiancaiamao/shen-go/cmd/shen

test:
	go test github.com/tiancaiamao/shen-go/vm
	go test github.com/tiancaiamao/shen-go/kl
