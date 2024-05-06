run:
	-docker compose -f ./deploy/local/docker-compose-run.yml down --remove-orphans
	docker compose -f ./deploy/local/docker-compose-run.yml up --build

test:
	-docker compose -f ./deploy/local/docker-compose-test.yml down --remove-orphans
	docker compose -f ./deploy/local/docker-compose-test.yml up --build