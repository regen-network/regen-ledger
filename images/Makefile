regen-env: regen-rmi
	docker build --tag regen-ledger/regen-env -f regen-env/Dockerfile \
    		$(shell git rev-parse --show-toplevel)

regen-rmi:
	docker rmi regen-ledger/regen-env 2>/dev/null; true

.PHONY: regen-env regen-rmi
