#DEV

build-dev:
	docker build -t go-videochat-conference -f containers/images/Dockerfile . && docker uild -t turn -f containers/images/Dockerfile.turn . 

clean-dev:
	docker compose -f containers/composes/dc.dev.yml down

run-dev:
	docker compose -f containers/composes/dc.dev.yml up

.PHONY: build-dev clean-dev run-dev