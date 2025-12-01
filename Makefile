.PHONY: run_frontend,run,format,tidy

run_frontend:
	cd vue-frontend && npm run serve

run:
	go run main.go

format:
	gofmt -w .
	@echo "go fmt success"

tidy:
	go mod tidy
	@echo "go tidy success"