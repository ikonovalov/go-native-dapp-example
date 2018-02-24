DC_FILE=docker/docker-compose.yml

solc-upgrade:
	docker pull ethereum/solc:stable
solc-compile:
	docker-compose -f ${DC_FILE} up && ls -l contracts/gen && docker-compose -f ${DC_FILE} down
