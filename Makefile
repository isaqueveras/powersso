.PHONY: local down-local clean logs-local

FILES := $(shell docker ps -aq)

local:
	docker-compose -f docker-compose.local.yml up -d --build

down-local:
	docker stop $(FILES)
	docker rm $(FILES)

clean:
	docker system prune -f

logs-local:
	docker logs -f $(FILES)
