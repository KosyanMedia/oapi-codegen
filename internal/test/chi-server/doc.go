package chi_server

//go:generate go run github.com/KosyanMedia/oapi-codegen/cmd/oapi-codegen --generate=types,chi-server --package=chi_server -o server.gen.go ../test-schema.yaml
//go:generate go run github.com/matryer/moq -out server_moq.gen.go . ServerInterface
