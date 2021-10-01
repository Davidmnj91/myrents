dockerFile=docker-compose.yml

.PHONY : all

all: docker-start

docker-start:
	docker-compose -f $(dockerFile) up -d --build

docker-stop:
	docker-compose -f $(dockerFile) down

docker-restart:	docker-stop docker-start

server:
	$(MAKE) -C server

app:
	$(MAKE) -C app