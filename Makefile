db: ## Subir o banco localmente.
	docker-compose -f ./build/db/docker-compose.yml up -d

fresh-db: ## Sobe o banco localmente mas recriando tudo do zero
	docker-compose -f ./build/db/docker-compose.yml down -v
	docker-compose -f ./build/db/docker-compose.yml up --build --remove-orphans -d

gqlgen: ## Gera os arquivos para a API GraphQL
	go generate ./config/http/config.go

run: 
	go generate ./cmd/http/http.go

frontend:
	rm -rf services/node_modules/
	yarn --cwd services/frontend/ install
	yarn --cwd services/frontend/ start
	mv yarn.lock services/frontend/
	mv node_modules/ services/frontend/
