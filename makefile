# DEVELOPMENT
build-dev:
	docker-compose -f docker-compose-dev.yml up --build

up-dev:
	docker-compose -f docker-compose-dev.yml up

remove-dev:
	docker-compose -f docker-compose-dev.yml down --rmi all

build-service:
	@read -p "Masukkan nama service: " service; \
	if [ -z "$$service" ]; then \
		echo "Nama service tidak boleh kosong!"; \
		exit 1; \
	fi; \
	echo "Building $$service..."; \
	docker-compose -f docker-compose-dev.yml up --build $$service

up-service:
	@read -p "Masukkan nama service: " service; \
	if [ -z "$$service" ]; then \
		echo "Nama service tidak boleh kosong!"; \
		exit 1; \
	fi; \
	echo "up $$service..."; \
	docker-compose -f docker-compose-dev.yml up $$service

exec-service:
	@read -p "Masukkan nama service: " service; \
	if [ -z "$$service" ]; then \
		echo "Nama service tidak boleh kosong!"; \
		exit 1; \
	fi; \
	echo "Execute $$service..."; \
	docker-compose -f docker-compose-dev.yml exec $$service sh

stop-service:
	@read -p "Masukkan nama service: " service; \
	if [ -z "$$service" ]; then \
		echo "Nama service tidak boleh kosong!"; \
		exit 1; \
	fi; \
	echo "Stop $$service..."; \
	docker-compose -f docker-compose-dev.yml stop $$service

remove-service:
	@read -p "Masukkan nama service: " service; \
	if [ -z "$$service" ]; then \
		echo "Nama service tidak boleh kosong!"; \
		exit 1; \
	fi; \
	db_service="$$service"_db; \
	echo "Remove $$service..."; \
	docker-compose -f docker-compose-dev.yml stop $$service && \
	docker-compose -f docker-compose-dev.yml stop $$db_service && \
	docker-compose -f docker-compose-dev.yml rm -f $$service && \
	docker-compose -f docker-compose-dev.yml rm -f $$db_service 

migrate-create:
	@read -p "Masukkan nama service: " service; \
	if [ -z "$$service" ]; then \
		echo "Nama service tidak boleh kosong!"; \
		exit 1; \
	fi; \
	db_service="$$service"_db; \
	echo "Migrate $$service..."; \
	docker-compose -f docker-compose-dev.yml exec $$service sh \
	-c "migrate create -ext sql -dir db/migration -seq init_schema";

rebuild-service:
	@read -p "Masukkan nama service: " service; \
	if [ -z "$$service" ]; then \
		echo "Nama service tidak boleh kosong!"; \
		exit 1; \
	fi; \
	db_service="$$service"_db; \
	echo "Rebuild $$service..."; \
	docker-compose -f docker-compose-dev.yml stop $$service && \
	docker-compose -f docker-compose-dev.yml stop $$db_service && \
	docker-compose -f docker-compose-dev.yml rm -f $$service && \
	docker-compose -f docker-compose-dev.yml rm -f $$db_service && \
	rm -rf $${service}/tmp && \
	docker-compose -f docker-compose-dev.yml up --build $$service

migrate-up:
	@read -p "Masukkan nama service: " service; \
	if [ -z "$$service" ]; then \
		echo "Nama service tidak boleh kosong!"; \
		exit 1; \
	fi; \
	echo "Verbose up $$service..."; \
    docker-compose -f docker-compose-dev.yml exec $$service sh -c 'migrate -path db/migration -database "$${DATABASE_URL}" -verbose up'

migrate-down:
	@read -p "Masukkan nama service: " service; \
	if [ -z "$$service" ]; then \
		echo "Nama service tidak boleh kosong!"; \
		exit 1; \
	fi; \
	echo "Verbose down $$service..."; \
    docker-compose -f docker-compose-dev.yml exec $$service sh -c 'migrate -path db/migration -database "$${DATABASE_URL}" -verbose down'

sqlc-init:
	@read -p "Masukkan nama service: " service; \
	if [ -z "$$service" ]; then \
		echo "Nama service tidak boleh kosong!"; \
		exit 1; \
	fi; \
	echo "sqlc init $$service..."; \
	mkdir $${service}/db/query && \
	mkdir $${service}/db/sqlc && \
    docker-compose -f docker-compose-dev.yml exec $$service sh -c 'sqlc version' && \
    docker-compose -f docker-compose-dev.yml exec $$service sh -c 'sqlc init'

sqlc-generate:
	@read -p "Masukkan nama service: " service; \
	if [ -z "$$service" ]; then \
		echo "Nama service tidak boleh kosong!"; \
		exit 1; \
	fi; \
	echo "sqlc generate $$service..."; \
    docker-compose -f docker-compose-dev.yml exec $$service sh -c 'sqlc version' && \
    docker-compose -f docker-compose-dev.yml exec $$service sh -c 'sqlc generate'

sqlc-delete:
	@read -p "Masukkan nama service: " service; \
	if [ -z "$$service" ]; then \
		echo "Nama service tidak boleh kosong!"; \
		exit 1; \
	fi; \
	echo "sqlc delete $$service..."; \
    rm -rf $${service}/db/sqlc/*

install-package:
	@read -p "Masukkan nama service: " service; \
	if [ -z "$$service" ]; then \
		echo "Nama service tidak boleh kosong!"; \
		exit 1; \
	fi; \
	read -p "Masukkan nama package: " package; \
	if [ -z "$$package" ]; then \
		echo "Nama package tidak boleh kosong!"; \
		exit 1; \
	fi; \
	echo "Install package $$package pada service $$service..."; \
    docker-compose -f docker-compose-dev.yml exec $$service sh -c 'go get "$$package"'

go-modtidy:
	@read -p "Masukkan nama service: " service; \
	if [ -z "$$service" ]; then \
		echo "Nama service tidak boleh kosong!"; \
		exit 1; \
	fi; \
	echo "go mod tidy $$service..."; \
    docker-compose -f docker-compose-dev.yml exec $$service sh -c 'go mod tidy'

go-test:
	@read -p "Masukkan nama service: " service; \
	if [ -z "$$service" ]; then \
		echo "Nama service tidak boleh kosong!"; \
		exit 1; \
	fi; \
	echo "go test $$service..."; \
    docker-compose -f docker-compose-dev.yml exec $$service sh -c 'go test -v -cover ./...'

delete-projectservice:
	@read -p "Masukkan nama service: " service; \
	if [ -z "$$service" ]; then \
		echo "Nama service tidak boleh kosong!"; \
		exit 1; \
	fi; \
	echo "Menghapus direktori $$service..."; \
	rm -rf $$service; \
	\
	echo "Menghapus service dan database dari docker-compose-dev.yml..."; \
	sed -i "/^  $$service\_db:/,/^[[:space:]]*$$/{d;}" docker-compose-dev.yml; \
	sed -i "/^  $$service:/,/^[[:space:]]*$$/{d;}" docker-compose-dev.yml; \
	sed -i "/^  $$service\_db:/,/^[[:space:]]*$$/{d;}" docker-compose-depl.yml; \
	sed -i "/^  $$service:/,/^[[:space:]]*$$/{d;}" docker-compose-depl.yml; \
	echo "Service $$service dan database telah dihapus."; \

create-projectservice:
	@read -p "Masukkan nama service: " service; \
	if [ -z "$$service" ]; then \
		echo "Nama service tidak boleh kosong!"; \
		exit 1; \
	fi; \
	\
	cp -r internal/base $$service; \
	db_service="$$service"_db; \
	mv $${service}/cmd/base $${service}/cmd/$${service}; \
	sed -i "s|Ini adalah base service!!!|Ini adalah $${service} service!!!|g" $${service}/cmd/$${service}/main.go; \
	sed -i "s|base|$${service}|g" $${service}/.air.toml; \
	sed -i "s|base|$${service}|g" $${service}/depl.dockerfile; \
	sed -i "s|base|$${service}|g" $${service}/dev.dockerfile; \
	\
	echo "" >> docker-compose-dev.yml; \
	sed -n "3,29p" internal/docker-compose-dev.yml >> docker-compose-dev.yml; \
	sed -i -e "s|base|$${service}|g" -e "s|base_db|$${db_service}|g" -e "s|sule|$${service}|g" docker-compose-dev.yml; \
	\
	echo "" >> docker-compose-depl.yml; \
	sed -n "3,27p" internal/docker-compose-depl.yml >> docker-compose-depl.yml; \
	sed -i -e "s|base|$${service}|g" -e "s|base_db|$${db_service}|g" -e "s|sule|$${service}|g" docker-compose-depl.yml; \
	\
	echo "Service baru telah dibuat, dengan nama service $$service."; \


# DEPLOYMENT
build-depl:
	docker-compose -f docker-compose-dev.yml up --build

up-depl:
	docker-compose -f docker-compose-depl.yml up

remove-depl:
	docker-compose -f docker-compose-depl.yml down --rmi all


.PHONY: build-dev up-dev remove-dev build-service up-service exec-service stop-service remove-service migrate-create rebuild-service migrate-up migrate-down sqlc-init build-depl up-depl remove-depl create-projectservice delete-projectservice