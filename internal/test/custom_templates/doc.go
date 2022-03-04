package custom_templates

//go:generate go run github.com/KosyanMedia/oapi-codegen/cmd/oapi-codegen -templates ./ --generate=types,server --package=custom_templates -o server.gen.go ../echo-test-schema.yaml
