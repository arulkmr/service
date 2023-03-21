all:
	run tidy
run:
	go run ./app/sales-api/main.go
tidy:
	go mod tidy