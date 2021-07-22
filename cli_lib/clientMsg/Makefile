SHELL=PATH='$(PATH)' /bin/sh

.PHONY: all
all:
	protoc --go_out=paths=source_relative:./unicast -I./unicast ./unicast/unicast.proto
	protoc --go_out=paths=source_relative:./multicast -I./multicast -I$(GOPATH)/src ./multicast/multicast.proto
