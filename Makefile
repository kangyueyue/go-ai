.PHONY: run_frontend, build_frontend, run, format, tidy

run_frontend:
	cd vue-frontend && npm run serve

build_frontend:
	cd vue-frontend && npm run build

run:
	go run main.go

format:
	gofmt -w .
	@echo "go fmt success"

tidy:
	go mod tidy
	@echo "go tidy success"