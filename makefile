# Generate ABI and BIN files in the build directory
build: 
	rm -rf ./contracts/build
	solc @chainlink/contracts=/mnt/c/Users/secre/AppData/Roaming/npm/node_modules/@chainlink/contracts --optimize --bin --abi --overwrite -o ./contracts/build/ ./contracts/*.sol
gen:
	abigen --bin=contracts/build/SplitBill.bin --abi=contracts/build/SplitBill.abi --pkg=contracts --type=SplitBill --out=contracts/splitbill.go

.PHONY: build gen