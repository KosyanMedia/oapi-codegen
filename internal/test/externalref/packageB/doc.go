package packageB

//go:generate go run github.com/KosyanMedia/oapi-codegen/v2/cmd/oapi-codegen -generate types,skip-prune,spec --package=packageB -o externalref.gen.go spec.yaml
