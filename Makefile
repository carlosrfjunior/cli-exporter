# Makefile

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

.PHONY: help
all: help

build:
	GO111MODULE=on go build -o $(GOBIN)/clis-exporter $(GOBASE)
	# @echo "########## Sending Configuration for the binary file"
	# @cp -f config.json $(GOBIN)/
	@chmod +x $(GOBIN)/*
	@@echo "########## Build finalized" 


help: Makefile
	@echo
	@echo "Usage: make [options]"
	@echo
	@echo "Options:"
	@echo "		build	Create binary file"
	@echo "		Help	"