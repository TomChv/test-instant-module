package main

import (
	"context"
	"dagger/test-instant/internal/dagger"
)

type TestInstant struct{
		RequiredPaths []string
}

func (t *TestInstant) ModuleRuntime(ctx context.Context, modSource *dagger.ModuleSource, introspectionJSON *dagger.File) (*dagger.Container, error) {
	return dag.Container().From("ttl.sh/dagger/instant-module/typescript/v0@sha256:ce4e9fb0f600d5105917aead850d3c3c40433e3aae75d73d8ad0cb20b2cbf0da"), nil
}

func (t *TestInstant) Codegen(ctx context.Context, modSource *dagger.ModuleSource, introspectionJSON *dagger.File) (*dagger.GeneratedCode, error) {
	return dag.GeneratedCode(dag.Directory()), nil
}
