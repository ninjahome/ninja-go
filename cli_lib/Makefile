BINDIR=bin

.PHONY: mac android ios test

all: mac android ios test

pbs:
	cd chat_msg/ && $(MAKE)


tp:=./bin

test:
	go build  -ldflags '-w -s' -o $(BINDIR)/ctest mac/*.go
mac:
	CGO_CFLAGS=-mmacosx-version-min=10.11 \
	CGO_LDFLAGS=-mmacosx-version-min=10.11 \
	GOARCH=amd64 GOOS=darwin go build  --buildmode=c-archive -o $(BINDIR)/ninja.a mac/*.go
	cp mac/callback.h $(BINDIR)/

android:
	gomobile bind -v -o $(BINDIR)/ninja.aar -target=android/arm64 -ldflags=-s github.com/ninjahome/ninja-go/cli_lib/mobile

ios:
	GO111MODULE=off gomobile bind -v -o $(BINDIR)/chatLib.framework -target=ios -ldflags=-s github.com/ninjahome/ninja-go/cli_lib/mobile

clean:
	gomobile clean
	rm $(BINDIR)/*
