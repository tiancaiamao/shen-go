.PHONY: all kl shen docker test

all: kl cora shen

kl:
	go1.16beta1 install github.com/tiancaiamao/shen-go/cmd/kl

cora:
	go1.16beta1 install github.com/tiancaiamao/shen-go/cmd/cora

shen:
	go1.16beta1 build -o shen github.com/tiancaiamao/shen-go/cmd/shen

docker:
	docker build -t shen-go .
	docker run -i -t --rm -v /tmp:/tmp shen-go \
		/bin/sh -c 'cp -a /usr/local/bin/shen /tmp/'
	cp -a /tmp/shen ./shen

test:
	cd kl; go1.16beta1 test
