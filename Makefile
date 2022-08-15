CONFIG_PATH=.cert
.PHONY: help test compile gencert
.DEFAULT_GOAL := help

init:
	mkdir -p ${CONFIG_PATH}

gencert:
	${HOME}/go/bin/cfssl gencert \
		-initca testutil/ca-csr.json | ${HOME}/go/bin/cfssljson -bare ca
	
	${HOME}/go/bin/cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=testutil/ca-config.json \
		-profile=server \
		testutil/server-csr.json | ${HOME}/go/bin/cfssljson -bare server
	
	${HOME}/go/bin/cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=testutil/ca-config.json \
		-profile=client \
		-cn="root" \
		testutil/client-csr.json | ${HOME}/go/bin/cfssljson -bare root-client

	${HOME}/go/bin/cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=testutil/ca-config.json \
		-profile=client \
		-cn="nobody" \
		testutil/client-csr.json | ${HOME}/go/bin/cfssljson -bare nobody-client

	mv *.pem *.csr ${CONFIG_PATH}

${CONFIG_DIR}/model.conf:
	cp testutil/model.conf ${CONFIG_DIR}/model.conf

${CONFIG_DIR}/policy.csv:
	cp testutil/policy.csv ${CONFIG_DIR}/policy.csv

test: $(CONFIG_DIR)/policy.csv $(CONFIG_DIR)/model.conf
	go test -race -shuffle=on ./...

compile:
	protoc api/v1/*.proto \
	--go_out=. \
	--go-grpc_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative  \
	--proto_path=.

help:
	@grep -E '^[a-zA-Z_-]+:' $(MAKEFILE_LIST)