.PHONY: all kl shen docker test

all: kl shen

kl:
	go install github.com/tiancaiamao/shen-go/cmd/kl

shen:
	go build -o shen github.com/tiancaiamao/shen-go/cmd/shen

docker:
	docker build -t shen-go .
	docker run -i -t --rm -v /tmp:/tmp shen-go \
		/bin/sh -c 'cp -a /usr/local/bin/shen /tmp/'
	cp -a /tmp/shen ./shen

test:
	cd klambda; go test
