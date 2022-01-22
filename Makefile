.PHONY: deploy stop update

deploy:
	@echo "deploy project"
	docker-compose up -d --build

stop:
	@echo "stop project"
	docker-compose stop

update:
	@echo "update project services"
	docker-compose pull





