all: regen-env

regen-env: regen-rmi
	docker build --tag regenledger/regen-env -f regen-env/Dockerfile \
    		$(shell git rev-parse --show-toplevel)

regen-rmi:
	docker rmi regenledger/regen-env 2>/dev/null; true

.PHONY: all regen-env regen-rmi