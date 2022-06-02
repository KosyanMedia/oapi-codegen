package server

//go:generate go run github.com/KosyanMedia/oapi-codegen/v2/cmd/oapi-codegen --config=cfg.yaml ../echo-test-schema.yaml
//go:generate go run github.com/matryer/moq -out server_moq.gen.go . ServerInterface
//go:generate go run github.com/matryer/moq -out echo_router_moq.gen.go . EchoRouter

type CustomEnumType string
