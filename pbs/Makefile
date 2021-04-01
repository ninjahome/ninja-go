SHELL=PATH='$(PATH)' /bin/sh

.PHONY: all
all:
	protoc --go_out=paths=source_relative:./contact -I. contact.proto
	protoc --go_out=paths=source_relative:./stream -I. stream.proto
	protoc --go_out=paths=source_relative:./p2p -I. p2p.proto
	protoc --go_out=paths=source_relative:./websocket -I. websocket.proto
	protoc --go_out=plugins=grpc:./cmd --go_opt=paths=source_relative cmdService.proto