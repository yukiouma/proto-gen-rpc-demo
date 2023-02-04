protoc:
	protoc --codegen_out=. --codegen_opt=paths=source_relative ./api/demo/v1/*.proto

srun:
	go run ./cmd/demo/server/main.go

crun:
	go run ./cmd/demo/client/main.go