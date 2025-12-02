.PHONY: run_frontend, build_frontend, run, format, tidy

run_frontend:
	cd vue-frontend && npm run serve
	@echo "run frontend success"

build_frontend:
	cd vue-frontend && npm run build
	@echo "build frontend success"

run:
	go run main.go
	@echo "go run success"

format:
	gofmt -w .
	@echo "go fmt success"

tidy:
	go mod tidy
	@echo "go tidy success"