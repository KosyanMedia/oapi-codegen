package chi_server

//go:generate go run github.com/KosyanMedia/oapi-codegen/v2/cmd/oapi-codegen --config=cfg.yaml ../test-schema.yaml
//go:generate go run github.com/matryer/moq -out server_moq.gen.go . ServerInterface
