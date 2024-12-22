DEFAULT_GOAL := docker-dev-up

.PHONY: docker-dev-up
docker-dev-up:
	docker build . -f ./deployments/Dockerfile --tag aejoy/vk-yourstickers
	docker-compose -f ./deployments/docker-compose.yaml up -d