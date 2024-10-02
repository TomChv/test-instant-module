package main

import (
	"context"
	"dagger/test-instant/internal/dagger"
)

type TestInstant struct{
		RequiredPaths []string
}

func (t *TestInstant) ModuleRuntime(ctx context.Context, modSource *dagger.ModuleSource, introspectionJSON *dagger.File) (*dagger.Container, error) {
	return dag.Container().From("ttl.sh/dagger/instant-module/v0@sha256:db925277fce2de852a4a2f7abff527f42c067b31e6d1367e5b112a457ea4c881"), nil
}

func (t *TestInstant) Codegen(ctx context.Context, modSource *dagger.ModuleSource, introspectionJSON *dagger.File) (*dagger.GeneratedCode, error) {
	return dag.GeneratedCode(dag.Directory()), nil
}
