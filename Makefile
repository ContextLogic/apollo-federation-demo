help: ## show the help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## install required deps/tools for this demo
	@npm install -g @apollo/rover
	@npm install 
	@go get github.com/99designs/gqlgen

graph: ## run rover to compose the supergraph
	@rover supergraph compose --config ./supergraph-config.yaml > supergraph.graphql

gqlgen-accounts: ## run gqlgen generate for accounts service
	@cd services/accounts; gqlgen generate 

accounts: gqlgen-accounts ## start accounts service
	@cd services/accounts; go run server.go

gqlgen-inventory: ## run gqlgen generate for inventory service
	@cd services/inventory; gqlgen generate 

inventory: gqlgen-inventory ## start inventory service
	@cd services/inventory; go run server.go

gqlgen-products: ## run gqlgen generate for products service
	@cd services/products; gqlgen generate 

products: gqlgen-products ## start products service
	@cd services/products; go run server.go 

gqlgen-reviews: ## run gqlgen generate for reviews service
	@cd services/reviews; gqlgen generate 

reviews: gqlgen-reviews ## start reviews service
	@cd services/reviews; go run server.go

gateway: ## start the gateway
	@npm run start-gateway