package internal

// This directory contains the OpenAPI 3.0 specification which defines our
// server. The file petstore.gen.go is automatically generated from the schema

// Run oapi-codegen to regenerate the petstore boilerplate
//go:generate go run github.com/KosyanMedia/oapi-codegen/v2/cmd/oapi-codegen --package=petstore --generate types,client -o ../petstore-client.gen.go ../petstore-expanded.yaml
