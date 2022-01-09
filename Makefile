.PHONY: deploy

deploy:
	@echo "deploy project"
	docker-compose up -d --build

update:
	@echo "update project services"
	docker-compose pull





