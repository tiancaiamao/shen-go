.PHONY: all shen esc generate

all: shen

generate: esc
	${GOPATH}/bin/esc -o runtime/asset.go -pkg=runtime bytecode

esc: ${GOPATH}/bin/esc
	go get -u -v github.com/mjibson/esc

shen:
	go build -o shen cmd/shen/main.go

test:
	go test github.com/tiancaiamao/shen-go/runtime
