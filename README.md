1. install solc
```shell
sudo add-apt-repository ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install solc
```

2. install protobuf
```shell
curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v24.3/protoc-24.3-linux-x86_64.zip
unzip -d protoc protoc-24.3-linux-x86_64.zip 
cd protoc
mv bin/protoc $GOPATH/bin
rm -rf protoc-24.3-linux-x86_64.zip 
rm -rf protoc/
```

3. install abigen
```shell
go get -u github.com/ethereum/go-ethereum@v1.11.0
cd go/pkg/mod/github.com/ethereum/go-ethereum@v1.13.2/
make devtools
```