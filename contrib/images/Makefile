all: regen-env

clean:
	docker rmi -f regenledger/regen-env

regen-env:
	docker build --build-arg UID=$(shell id -u) --build-arg GID=$(shell id -g) --tag regenledger/regen-env regen-env

.PHONY: all regen-env
