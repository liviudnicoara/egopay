# Generate ABI and BIN files in the build directory
build: 
	rm -rf ./internal/app/contracts/build
	solc @chainlink/contracts=/mnt/c/Users/secre/AppData/Roaming/npm/node_modules/@chainlink/contracts --optimize --bin --abi --overwrite -o ./internal/app/contracts/build/ ./internal/app/contracts/*.sol
gen:
	abigen --bin=internal/app/contracts/build/SplitBill.bin --abi=internal/app/contracts/build/SplitBill.abi --pkg=contracts --type=SplitBill --out=internal/app/contracts/splitbill.go

swagger:
	swag init -g ./cmd/main.go -o ./cmd/docs

.PHONY: build gen swagger