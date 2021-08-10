SHELL=PATH='$(PATH)' /bin/sh

.PHONY: all
all:
	abigen --abi eth-licenses/NinjaChatLicense.abi --pkg contract --type NinjaChatLicense --out eth-licenses.go

abi:
	solc --abi --overwrite eth-license.sol -o ./eth-licenses

clean:
#	rm -f *.abi
	rm -f *.go
