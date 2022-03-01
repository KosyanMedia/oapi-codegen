package server

//go:generate go run github.com/KosyanMedia/oapi-codegen/cmd/oapi-codegen --generate=types,server --package=server -o server.gen.go ../echo-test-schema.yaml
//go:generate go run github.com/matryer/moq -out server_moq.gen.go . ServerInterface
