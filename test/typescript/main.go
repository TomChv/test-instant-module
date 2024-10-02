package main

import (
	"context"
	"dagger/test-instant/internal/dagger"
)

type TestInstant struct{
		RequiredPaths []string
}

func (t *TestInstant) ModuleRuntime(ctx context.Context, modSource *dagger.ModuleSource, introspectionJSON *dagger.File) (*dagger.Container, error) {
	return dag.Container().From("ttl.sh/dagger/instant-module/bun/v0@sha256:7e5cff5318f64bbba76ce65d561f1d94ea890a92654fbde67b47eb728378fa8d"), nil
}

func (t *TestInstant) Codegen(ctx context.Context, modSource *dagger.ModuleSource, introspectionJSON *dagger.File) (*dagger.GeneratedCode, error) {
	return dag.GeneratedCode(dag.Directory()), nil
}
