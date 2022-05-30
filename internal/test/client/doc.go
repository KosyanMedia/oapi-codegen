package client

//go:generate go run github.com/KosyanMedia/oapi-codegen/v2/cmd/oapi-codegen --package=client -o client.gen.go client.yaml
//go:generate go run github.com/KosyanMedia/oapi-codegen/v2/cmd/oapi-codegen --package=no_client_editors --no-req-editors -o no-req-editors/client.gen.go client.yaml
