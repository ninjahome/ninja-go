SHELL=PATH='$(PATH)' /bin/sh

.PHONY: all
all:
	abigen --abi eth-licenses/NinjaChatLicense.abi --pkg contract --type NinjaChatLicense --out eth-licenses.go
	abigen --abi eth-licenses/NinjaToken.abi --pkg contract --type NinjaToken --out ninja-token.go

abi:
	solc --abi --overwrite eth-license.sol -o ./eth-licenses
	solc --abi --overwrite ninja-token.sol -o ./eth-licenses

clean:
#	rm -f *.abi
	rm -f *.go
