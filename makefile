mock:
	mockgen -package mock -destination internal/mocks/usecase.go github.com/edgarSucre/moca/internal/controller Usecase

test:
	go test ./... -v -cover

run:
	go run main.go