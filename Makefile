DC_FILE=docker/docker-compose.yml
ETH_TOOLS=${GOPATH}/src/github.com/ethereum/go-ethereum/build/bin

all: solc-compile check-security gen-bind solc-compile-shutdown build

solc-upgrade:
	docker pull ethereum/solc:stable
solc-compile:
	docker-compose -f ${DC_FILE} up && ls -l contracts/gen && docker-compose -f ${DC_FILE} down
solc-compile-shutdown:
	 docker-compose -f ${DC_FILE} down
gen-bind:
	${ETH_TOOLS}/abigen -abi ./contracts/gen/Greeter.abi -bin ./contracts/gen/Greeter.bin -pkg greeter -lang go -out contracts/gen/greeter.go
check-security:
	myth -c `cat contracts/gen/Greeter.bin` -x
build: 
	go build -o bin/dapp .
run:
	./bin/dapp
