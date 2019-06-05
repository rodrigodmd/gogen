GOPATH:=$(shell go env GOPATH)

.PHONY: help build install

all: build install ## Build and install the binary in /usr/local/bin

help:             ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build:            ## Creates the binary and place it within script folder
	go build -o gogen
	mv gogen ./scripts/Darwin/

install:          ## Install the created binary in /usr/local/bin
	sudo chmod 777 ./scripts/Darwin/gogen
	sudo rm -fr /usr/local/bin/gogen
	sudo mv ./scripts/Darwin/gogen /usr/local/bin/

